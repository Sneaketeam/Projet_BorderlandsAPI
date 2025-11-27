package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ==========================================
// 1. STRUCTURES DE DONNÉES
// ==========================================

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
	IsFavorite   bool
}

type PageData struct {
	Weapons             []Weapon
	CurrentCategory     string
	CurrentRarity       string
	CurrentName         string
	CurrentManufacturer string
	IsLoggedIn          bool
	Username            string
	FlashMessage        string
	FlashType           string
	CurrentPage         int
	TotalPages          int
	NextPage            int
	PrevPage            int
}

type LoginData struct {
	ErrorMessage string
	IsError      bool
}

// ==========================================
// 2. GESTION BASE DE DONNÉES
// ==========================================

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"

	// TES INFOS ALWAYSDATA
	dbUser := "443067"
	dbPass := "giogio220706"
	dbHost := "mysql-borderlandsapi.alwaysdata.net"
	dbName := "borderlandsapi_database"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getUserID(db *sql.DB, username string) (int, error) {
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&id)
	return id, err
}

// ==========================================
// 3. PAGES PRINCIPALES
// ==========================================

func IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Session
	var isLoggedIn bool
	var username string
	cookie, err := r.Cookie("session_token")
	if err == nil {
		isLoggedIn = true
		username = cookie.Value
	}

	// Messages Flash
	msgCode := r.URL.Query().Get("msg")
	favName := r.URL.Query().Get("fav_name")
	var flashMsg, flashType string

	if msgCode == "nologin" {
		flashMsg = "⚠️ Connecte-toi pour gérer tes favoris !"
		flashType = "error"
	}
	if msgCode == "added" {
		if favName != "" {
			flashMsg = "★ " + favName + " ajoutée aux Favoris !"
		} else {
			flashMsg = "★ Arme ajoutée aux Favoris !"
		}
		flashType = "success"
	}
	if msgCode == "removed" {
		flashMsg = "🗑️ Arme retirée."
		flashType = "error"
	}

	db := dbConn()
	defer db.Close()

	// Filtres & Pagination
	cat := r.URL.Query().Get("category")
	rar := r.URL.Query().Get("rarity")
	man := r.URL.Query().Get("manufacturer")
	nam := r.URL.Query().Get("name")

	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	whereClause := "1=1"
	var args []interface{}

	if cat != "" {
		whereClause += " AND category = ?"
		args = append(args, cat)
	}
	if rar != "" {
		whereClause += " AND rarity = ?"
		args = append(args, rar)
	}
	if man != "" {
		whereClause += " AND manufacturer = ?"
		args = append(args, man)
	}
	if nam != "" {
		whereClause += " AND name LIKE ?"
		args = append(args, "%"+nam+"%")
	}

	// Total Pages
	var totalItems int
	countQuery := "SELECT COUNT(*) FROM weapons WHERE " + whereClause
	db.QueryRow(countQuery, args...).Scan(&totalItems)
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	// Requête Armes
	query := "SELECT * FROM weapons WHERE " + whereClause + " ORDER BY id ASC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)
	rows, _ := db.Query(query, args...)
	defer rows.Close()

	var weapons []Weapon
	var userID int
	if isLoggedIn {
		userID, _ = getUserID(db, username)
	}

	for rows.Next() {
		var wpn Weapon
		rows.Scan(&wpn.ID, &wpn.Category, &wpn.Name, &wpn.Manufacturer, &wpn.Rarity, &wpn.FlavorText, &wpn.Details, &wpn.Source, &wpn.ImageURL)

		if isLoggedIn {
			var count int
			db.QueryRow("SELECT COUNT(*) FROM favorites WHERE user_id = ? AND weapon_id = ?", userID, wpn.ID).Scan(&count)
			if count > 0 {
				wpn.IsFavorite = true
			}
		}
		weapons = append(weapons, wpn)
	}

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, PageData{
		Weapons: weapons, CurrentCategory: cat, CurrentRarity: rar, CurrentManufacturer: man, CurrentName: nam,
		IsLoggedIn: isLoggedIn, Username: username, FlashMessage: flashMsg, FlashType: flashType,
		CurrentPage: page, TotalPages: totalPages,
		NextPage: page + 1, PrevPage: page - 1,
	})
}

func WeaponPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	id := r.URL.Query().Get("id")
	db := dbConn()
	defer db.Close()

	var wpn Weapon
	err := db.QueryRow("SELECT * FROM weapons WHERE id = ?", id).Scan(&wpn.ID, &wpn.Category, &wpn.Name, &wpn.Manufacturer, &wpn.Rarity, &wpn.FlavorText, &wpn.Details, &wpn.Source, &wpn.ImageURL)
	if err != nil {
		http.Error(w, "Introuvable", 404)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err == nil {
		userID, _ := getUserID(db, cookie.Value)
		var count int
		db.QueryRow("SELECT COUNT(*) FROM favorites WHERE user_id = ? AND weapon_id = ?", userID, wpn.ID).Scan(&count)
		if count > 0 {
			wpn.IsFavorite = true
		}
	}
	t, _ := template.ParseFiles("templates/weapon.html")
	t.Execute(w, wpn)
}

// ==========================================
// 4. GESTION FAVORIS
// ==========================================

func FavoritesPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/auth?error=nologin", http.StatusSeeOther)
		return
	}

	db := dbConn()
	defer db.Close()
	query := `SELECT w.id, w.category, w.name, w.manufacturer, w.rarity, w.flavor_text, w.details, w.source, w.image_url FROM weapons w JOIN favorites f ON w.id = f.weapon_id JOIN users u ON f.user_id = u.id WHERE u.username = ?`
	rows, _ := db.Query(query, cookie.Value)
	defer rows.Close()

	var weapons []Weapon
	for rows.Next() {
		var wpn Weapon
		rows.Scan(&wpn.ID, &wpn.Category, &wpn.Name, &wpn.Manufacturer, &wpn.Rarity, &wpn.FlavorText, &wpn.Details, &wpn.Source, &wpn.ImageURL)
		weapons = append(weapons, wpn)
	}
	t, _ := template.ParseFiles("templates/favorites.html")
	t.Execute(w, PageData{Weapons: weapons, IsLoggedIn: true, Username: cookie.Value})
}

func AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/?msg=nologin", http.StatusSeeOther)
		return
	}
	db := dbConn()
	defer db.Close()
	weaponID := r.URL.Query().Get("id")
	userID, _ := getUserID(db, cookie.Value)

	var weaponName string
	db.QueryRow("SELECT name FROM weapons WHERE id = ?", weaponID).Scan(&weaponName)

	db.Exec("INSERT IGNORE INTO favorites (user_id, weapon_id) VALUES (?, ?)", userID, weaponID)

	// Redirection avec le nom pour l'affichage
	http.Redirect(w, r, "/?msg=added&fav_name="+url.QueryEscape(weaponName), http.StatusSeeOther)
}

func RemoveFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	db := dbConn()
	defer db.Close()
	userID, _ := getUserID(db, cookie.Value)
	weaponID := r.URL.Query().Get("id")
	db.Exec("DELETE FROM favorites WHERE user_id = ? AND weapon_id = ?", userID, weaponID)
	http.Redirect(w, r, "/?msg=removed", http.StatusSeeOther)
}

// API AJAX (Au cas où tu voudrais réutiliser le JS plus tard)
func ToggleFavoriteAPI(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Non connecté", http.StatusUnauthorized)
		return
	}
	db := dbConn()
	defer db.Close()

	weaponID := r.URL.Query().Get("id")
	userID, err := getUserID(db, cookie.Value)
	if err != nil {
		http.Error(w, "Error", 500)
		return
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM favorites WHERE user_id = ? AND weapon_id = ?", userID, weaponID).Scan(&count)
	var name string
	db.QueryRow("SELECT name FROM weapons WHERE id = ?", weaponID).Scan(&name)

	w.Header().Set("Content-Type", "application/json")
	if count > 0 {
		db.Exec("DELETE FROM favorites WHERE user_id = ? AND weapon_id = ?", userID, weaponID)
		fmt.Fprintf(w, `{"status": "removed", "name": "%s"}`, name)
	} else {
		db.Exec("INSERT INTO favorites (user_id, weapon_id) VALUES (?, ?)", userID, weaponID)
		fmt.Fprintf(w, `{"status": "added", "name": "%s"}`, name)
	}
}

// ==========================================
// 5. AUTHENTIFICATION
// ==========================================

func LoginPage(w http.ResponseWriter, r *http.Request) {
	errCode := r.URL.Query().Get("error")
	var msg string
	var isError bool
	if errCode == "exists" {
		msg = "Ce pseudo est déjà pris !"
		isError = true
	}
	if errCode == "wrong" {
		msg = "Identifiants incorrects !"
		isError = true
	}
	t, _ := template.ParseFiles("templates/login.html")
	t.Execute(w, LoginData{ErrorMessage: msg, IsError: isError})
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	db := dbConn()
	defer db.Close()
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		http.Redirect(w, r, "/auth?error=exists", http.StatusSeeOther)
		return
	}
	createCookie(w, username)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	db := dbConn()
	defer db.Close()
	var dbPass string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&dbPass)
	if err != nil || dbPass != password {
		http.Redirect(w, r, "/auth?error=wrong", http.StatusSeeOther)
		return
	}
	createCookie(w, username)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Anti-Cache agressif pour la déconnexion
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Expires: time.Unix(0, 0), Path: "/"})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func createCookie(w http.ResponseWriter, username string) {
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: username, Expires: time.Now().Add(24 * time.Hour), Path: "/"})
}

func GetWeapons(w http.ResponseWriter, r *http.Request) {}
