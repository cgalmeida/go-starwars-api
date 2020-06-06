package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-rest-api-master/helper"
	"github.com/go-rest-api-master/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getPlanets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Planet array
	var planets []models.Planet

	//Connection mongoDB with helper class
	collection := helper.ConnectDB()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var planet models.Planet
		// & character returns the memory address of the following variable.
		err := cur.Decode(&planet) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		planets = append(planets, planet)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(planets) // encode similar to serialize process.
}

func getPlanet(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var planet models.Planet
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&planet)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(planet)
}

// Get single planet byName
func getPlanetByName(w http.ResponseWriter, r *http.Request) { // set header.
	w.Header().Set("Content-Type", "application/json")

	var planet models.Planet
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	name, _ := params["name"]
	fmt.Println(name)
	collection := helper.ConnectDB()

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"name": name}
	err := collection.FindOne(context.TODO(), filter).Decode(&planet)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(planet)

}

func createPlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var planet models.Planet
	//initializes planet info:

	planet.Movies = strconv.Itoa(0) //TODO: Get movies info

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&planet)

	//get planet infos from https://swapi.dev/:
	fmt.Println(planet.Name)
	planet = getPlanetInfo(planet.Name)

	// connect db
	collection := helper.ConnectDB()

	// insert our planet model.
	result, err := collection.InsertOne(context.TODO(), planet)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getPlanetInfo(planetName string) models.Planet {

	// A Planet Struct to map every planet to.
	type PlanetData struct {
		Name    string   `json:"name"`
		Climate string   `json:"climate"`
		Terrain string   `json:"terrain"`
		Movie   []string `json:"films"`
	}

	//A Response struct to map the Entire Response
	type Response struct {
		Name       string       `json:"name"`
		PlanetData []PlanetData `json:"results"`
	}

	var planet models.Planet

	planetName = strings.ReplaceAll(planetName, " ", "%20")
	searchURL := fmt.Sprint("https://swapi.dev/api/planets/?search=", planetName)
	response, err := http.Get(searchURL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.PlanetData); i++ {
		fmt.Println(responseObject.PlanetData[i].Name)
		fmt.Println(responseObject.PlanetData[i].Climate)
		fmt.Println(responseObject.PlanetData[i].Terrain)
		fmt.Println(len(responseObject.PlanetData[i].Movie))

		planet.Name = responseObject.PlanetData[i].Name
		planet.Climate = responseObject.PlanetData[i].Climate
		planet.Terrain = responseObject.PlanetData[i].Terrain
		planet.Movies = strconv.Itoa(len(responseObject.PlanetData[i].Movie))
	}

	//fmt.Println(responseObject.Planet)
	return planet

}

func updatePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var planet models.Planet

	collection := helper.ConnectDB()

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&planet)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"name", planet.Name},
			{"climate", planet.Climate},
			{"terrain", planet.Terrain},
			{"movies", planet.Movies},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&planet)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	planet.ID = id

	json.NewEncoder(w).Encode(planet)
}

func deletePlanet(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

// var client *mongo.Client

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/planets", getPlanets).Methods("GET")
	r.HandleFunc("/api/planets/{id}", getPlanet).Methods("GET")
	r.HandleFunc("/api/planets/search/{name}", getPlanetByName).Methods("GET")
	r.HandleFunc("/api/planets", createPlanet).Methods("POST")
	r.HandleFunc("/api/planets/{id}", updatePlanet).Methods("PUT")
	r.HandleFunc("/api/planets/{id}", deletePlanet).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
