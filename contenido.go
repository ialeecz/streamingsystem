package main

import (
	"encoding/json"
	"net/http"
)

type Contenido struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Tipo   string `json:"tipo"`
}

var contenidos = []Contenido{
	{ID: 1, Titulo: "Curso de Go", Tipo: "Educativo"},
	{ID: 2, Titulo: "Dexter", Tipo: "Serie"},
}

func listarContenido(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contenidos)
}

func agregarContenido(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var nuevo Contenido
	json.NewDecoder(r.Body).Decode(&nuevo)

	nuevo.ID = len(contenidos) + 1
	contenidos = append(contenidos, nuevo)

	json.NewEncoder(w).Encode(nuevo)
}
