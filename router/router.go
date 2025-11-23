package router

import (
	"BorderlandsAPI/controller"
	"fmt"
	"net/http"
)

func InitServer() {
	// 1. API
	http.HandleFunc("/api/weapons", controller.GetWeapons)

	// 2. Images
	imgServer := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", imgServer))

	// 3. CSS
	cssServer := http.FileServer(http.Dir("./stylecss"))
	http.Handle("/stylecss/", http.StripPrefix("/stylecss/", cssServer))

	// 4. Site Web
	http.HandleFunc("/", controller.IndexPage)

	// --- MESSAGE DE DÉMARRAGE PROPRE ---
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
