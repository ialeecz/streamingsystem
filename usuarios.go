package main

import (
	"encoding/json"
	"net/http"
)

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios = []Usuario{
	{ID: 1, Nombre: "Carlos", Email: "carlos@email.com"},
	{ID: 2, Nombre: "Ana", Email: "ana@email.com"},
}

func listarUsuarios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(usuarios)
}

func agregarUsuario(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var nuevo Usuario
	json.NewDecoder(r.Body).Decode(&nuevo)

	nuevo.ID = len(usuarios) + 1
	usuarios = append(usuarios, nuevo)

	json.NewEncoder(w).Encode(nuevo)
}
