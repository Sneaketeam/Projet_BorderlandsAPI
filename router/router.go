package router

import (
	"BorderlandsAPI/controller"
	"fmt"
	"net/http"
)

func InitServer() {
	// 1. API & Static
	http.HandleFunc("/api/weapons", controller.GetWeapons)

	// --- NOUVEAU : API POUR LE JAVASCRIPT (FAVORIS SANS RECHARGEMENT) ---
	http.HandleFunc("/api/fav/toggle", controller.ToggleFavoriteAPI)

	// 2. Images
	imgServer := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", imgServer))

	// 3. CSS
	cssServer := http.FileServer(http.Dir("./stylecss"))
	http.Handle("/stylecss/", http.StripPrefix("/stylecss/", cssServer))

	// 4. Pages du Site
	http.HandleFunc("/", controller.IndexPage)
	http.HandleFunc("/weapon", controller.WeaponPage)

	// 5. Favoris (Pages classiques)
	http.HandleFunc("/favorites", controller.FavoritesPage)                // Voir la liste
	http.HandleFunc("/favorites/add", controller.AddFavoriteHandler)       // Action classique (redirection)
	http.HandleFunc("/favorites/remove", controller.RemoveFavoriteHandler) // Action classique (redirection)

	// 6. Authentification
	http.HandleFunc("/auth", controller.LoginPage)
	http.HandleFunc("/signup", controller.SignupHandler)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)

	// --- MESSAGE DE DÉMARRAGE ---
	fmt.Println("")
	fmt.Println("          ★ GIOVANNI STREET ARMOURY 2 - ONLINE ★")
	fmt.Println("          --------------------------------------")
	fmt.Println("          🌍 Site accessible : http://localhost:8000")
	fmt.Println("")

	// Lancement
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("❌ Erreur critique :", err)
	}
}
