package main

import (
	"encoding/json"
	"net/http"
)

type Suscripcion struct {
	ID     int     `json:"id"`
	Plan   string  `json:"plan"`
	Precio float64 `json:"precio"`
}

var suscripciones = []Suscripcion{
	{ID: 1, Plan: "Básico", Precio: 5.99},
	{ID: 2, Plan: "Premium", Precio: 12.99},
}

func listarSuscripciones(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(suscripciones)
}

func reproducir(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Reproduciendo contenido 🎬",
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Inicio de sesión exitoso 🔐",
	})
}
