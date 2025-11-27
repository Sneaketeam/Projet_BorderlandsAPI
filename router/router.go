package router

import (
	"BorderlandsAPI/controller" // Vérifie le nom du module
	"fmt"
	"net/http"
)

func InitServer() {
	// Pages
	http.HandleFunc("/", controller.IndexPage)
	http.HandleFunc("/weapon", controller.WeaponPage)
	http.HandleFunc("/favorites", controller.FavoritesPage)

	// Actions
	http.HandleFunc("/favorites/add", controller.AddFavoriteHandler)
	http.HandleFunc("/favorites/remove", controller.RemoveFavoriteHandler)
	http.HandleFunc("/api/fav/toggle", controller.ToggleFavoriteAPI) // Pour l'AJAX

	// Auth
	http.HandleFunc("/auth", controller.LoginPage)
	http.HandleFunc("/signup", controller.SignupHandler)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)

	// Static
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.Handle("/stylecss/", http.StripPrefix("/stylecss/", http.FileServer(http.Dir("./stylecss"))))

	// API factice pour compatibilité
	http.HandleFunc("/api/weapons", controller.GetWeapons)

	fmt.Println("✅ SERVEUR STABLE Lancé : http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
