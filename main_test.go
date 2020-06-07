package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/planets", getPlanets).Methods("GET")
	return router
}

//Get All Planets 
func TestGetPlanets(t *testing.T) {
	//fmt.Println("Test Planets")
	request, _ := http.NewRequest("GET", "/api/planets", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}

//Get Single Planet By ID
func TestGetPlanet(t *testing.T) {
	//fmt.Println("Test Planet")
	router := mux.NewRouter()
	router.HandleFunc("/api/planets/{id}", getPlanets).Methods("GET")

	//Existing Planet
	request, _ := http.NewRequest("GET", "/api/planets/5edbc537496dddb68a9136b3", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK response is expected")

}

//Get Single Planet By non Existing ID
func TestGetNonExistingPlanet(t *testing.T) {
	//fmt.Println("Test Planet")
	router := mux.NewRouter()
	router.HandleFunc("/api/planets/{id}", getPlanet).Methods("GET")

	//Non Existing Planet
	request, _ := http.NewRequest("GET", "/api/planets/00000", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code, "OK response is expected")

}

//Create Planet
func TestCreatePlanet(t *testing.T) {
	//fmt.Println("Test CreatePlanet")
	router := mux.NewRouter()
	router.HandleFunc("/api/planets", createPlanet).Methods("POST")

	var jsonStr = []byte(`{
		"name": "Alderaan",
		"climate": "tropical",
		"terrain": "jungle rainforests",
		"movies": "1"
	 }`)

	//Non Existing Planet
	request, _ := http.NewRequest("POST", "/api/planets", bytes.NewBuffer(jsonStr))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK response is expected")

}

//Create Non StarWars Planet
func TestCreateNonExistingPlanet(t *testing.T) {
	//fmt.Println("Test CreatePlanet")
	router := mux.NewRouter()
	router.HandleFunc("/api/planets", createPlanet).Methods("POST")

	var jsonStr = []byte(`{
		"name": "Alda",
		"climate": "tropical",
		"terrain": "jungle rainforests",
		"movies": "1"
	 }`)

	//Non Existing Planet
	request, _ := http.NewRequest("POST", "/api/planets", bytes.NewBuffer(jsonStr))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code, "OK response is expected")

}

//Delete Non Existing Planet
func TestDeleteNonExistingPlanet(t *testing.T) {
	//fmt.Println("Test CreatePlanet")
	router := mux.NewRouter()
	router.HandleFunc("/api/planets/{id}", deletePlanet).Methods("DELETE")

	//Non Existing Planet
	request, _ := http.NewRequest("DELETE", "/api/planets/00000", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK response is expected")

}

//Delete Existing Planet
func TestDeleteExistingPlanet(t *testing.T) {
	//fmt.Println("Test CreatePlanet")
	router := mux.NewRouter()
	router.HandleFunc("/api/planets/{id}", deletePlanet).Methods("DELETE")

	//Non Existing Planet
	request, _ := http.NewRequest("DELETE", "/api/planets/5edd447c14709871249c308e", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK response is expected")

}

//Update Existing Planet
func TestUpdateExistingPlanet(t *testing.T) {
	//fmt.Println("Test CreatePlanet")
	router := mux.NewRouter()
	router.HandleFunc("/api/planets/{id}", updatePlanet).Methods("PUT")

	var jsonStr = []byte(`{
		"name": "Alderaan",
		"climate": "tropical UPDATED",
		"terrain": "jungle rainforests",
		"movies": "1"
	 }`)

	//Non Existing Planet
	request, _ := http.NewRequest("PUT", "/api/planets/5edbcf8e4f0ad983dcf039c8", bytes.NewBuffer(jsonStr))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK response is expected")

}

//Update Non Existing Planet
func TestUpdateNonExistingPlanet(t *testing.T) {
	//fmt.Println("Test UpdatePlanet")
	router := mux.NewRouter()
	router.HandleFunc("/api/planets/{id}", updatePlanet).Methods("PUT")

	var jsonStr = []byte(`{
		"name": "Alderaan",
		"climate": "tropical UPDATED",
		"terrain": "jungle rainforests",
		"movies": "1"
	 }`)

	//Non Existing Planet
	request, _ := http.NewRequest("PUT", "/api/planets/00000", bytes.NewBuffer(jsonStr))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code, "OK response is expected")

}

//Try to Get Non Existing Planet By Name
func TestGetNonExistingPlanetByName(t *testing.T) {
	//fmt.Println("Test CreatePlanet")
	router := mux.NewRouter()
	router.HandleFunc("/api/planets/search/{name}", getPlanetByName).Methods("GET")

	//Non Existing Planet
	request, _ := http.NewRequest("GET", "/api/planets/search/Terra", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code, "OK response is expected")

}

//Get Existing Planet By Name
func TestGetExistingPlanetByName(t *testing.T) {
	//fmt.Println("Test CreatePlanet")
	router := mux.NewRouter()
	router.HandleFunc("/api/planets/search/{name}", getPlanetByName).Methods("GET")

	//Non Existing Planet
	request, _ := http.NewRequest("GET", "/api/planets/search/Alderaan", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK response is expected")

}
