package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API de Streaming funcionando 🚀")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/usuarios", listarUsuarios)
	http.HandleFunc("/contenido", listarContenido)
	http.HandleFunc("/suscripciones", listarSuscripciones)
	http.HandleFunc("/reproducir", reproducir)
	http.HandleFunc("/login", login)
	http.HandleFunc("/agregar-usuario", agregarUsuario)
	http.HandleFunc("/agregar-contenido", agregarContenido)

	fmt.Println("Servidor corriendo en http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
