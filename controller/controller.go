package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Weapon struct {
	ID           int
	Category     string
	Name         string
	Manufacturer string
	Rarity       string
	FlavorText   string
	Details      string
	Source       string
	ImageURL     string
}

type PageData struct {
	Weapons         []Weapon
	CurrentCategory string
	CurrentRarity   string
	CurrentName     string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "borderlands_db"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUser, dbPass, dbName)

	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		// On garde le panic ici car sans BDD, le site ne sert à rien
		panic(err.Error())
	}
	return db
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	// Gestion propre des erreurs 404 pour les fichiers inexistants (favicon, robots.txt...)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	db := dbConn()
	defer db.Close()

	categoryFilter := r.URL.Query().Get("category")
	rarityFilter := r.URL.Query().Get("rarity")
	nameFilter := r.URL.Query().Get("name")

	query := "SELECT * FROM weapons WHERE 1=1"
	var args []interface{}

	if categoryFilter != "" {
		query += " AND category = ?"
		args = append(args, categoryFilter)
	}
	if rarityFilter != "" {
		query += " AND rarity = ?"
		args = append(args, rarityFilter)
	}
	if nameFilter != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+nameFilter+"%")
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		// On affiche l'erreur dans le navigateur, pas dans le terminal
		http.Error(w, "Erreur BDD", 500)
		return
	}
	defer rows.Close()

	var weapons []Weapon
	for rows.Next() {
		var wpn Weapon
		err = rows.Scan(&wpn.ID, &wpn.Category, &wpn.Name, &wpn.Manufacturer, &wpn.Rarity, &wpn.FlavorText, &wpn.Details, &wpn.Source, &wpn.ImageURL)
		if err != nil {
			continue
		}
		weapons = append(weapons, wpn)
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Erreur Template", 500)
		return
	}

	data := PageData{
		Weapons:         weapons,
		CurrentCategory: categoryFilter,
		CurrentRarity:   rarityFilter,
		CurrentName:     nameFilter,
	}

	t.Execute(w, data)
}

func GetWeapons(w http.ResponseWriter, r *http.Request) {
	// Cette fonction API ne sert plus au site, on la laisse vide ou on peut la supprimer.
}
