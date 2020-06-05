package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Planet struct (Model)
type Planet struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init planets var as a slice Planet struct
var planets []Planet

// Get all planets
func getPlanets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(planets)
}

// Get single planet
func getPlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through planets and find one with the id from the params
	for _, item := range planets {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Planet{})
}

// Get single planet
func getPlanetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("name....")
	params := mux.Vars(r) // Gets params
	// Loop through planets and find one with the id from the params
	for _, item := range planets {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Planet{})
}

// Add new planet
func createPlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var planet Planet
	_ = json.NewDecoder(r.Body).Decode(&planet)
	planet.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	planets = append(planets, planet)
	json.NewEncoder(w).Encode(planet)
}

// Update planet
func updatePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range planets {
		if item.ID == params["id"] {
			planets = append(planets[:index], planets[index+1:]...)
			var planet Planet
			_ = json.NewDecoder(r.Body).Decode(&planet)
			planet.ID = params["id"]
			planets = append(planets, planet)
			json.NewEncoder(w).Encode(planet)
			return
		}
	}
}

// Delete planet
func deletePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range planets {
		if item.ID == params["id"] {
			planets = append(planets[:index], planets[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(planets)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	planets = append(planets, Planet{ID: "1", Name: "Alderaan", Climate: "Planet One", Terrain: "arrid"})
	planets = append(planets, Planet{ID: "2", Name: "454555", Climate: "Planet Two", Terrain: "tropical"})

	// Route handles & endpoints
	r.HandleFunc("/planets", getPlanets).Methods("GET")
	r.HandleFunc("/planets/{id}", getPlanet).Methods("GET")
	r.HandleFunc("/planets/search/{name}", getPlanetByName).Methods("GET")
	r.HandleFunc("/planets", createPlanet).Methods("POST")
	r.HandleFunc("/planets/{id}", updatePlanet).Methods("PUT")
	r.HandleFunc("/planets/{id}", deletePlanet).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Request sample
// {
// 	"name": "Alderaan",
//  "climate": "temperate",
//  "terrain": "grasslands, mountains"
// }
