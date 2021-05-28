package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/ben-eh/cocktailTime/ingredient"
	"github.com/ben-eh/cocktailTime/measurement"
	"github.com/gorilla/mux"
)

type RecipeNecessities struct {
	Measurements []measurement.Measurement
	Ingredients  []ingredient.Ingredient
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	htmlPage, err := ioutil.ReadFile("templates/index.html")
	if err != nil {
		log.Fatal("Could not read index.html")
	}

	fmt.Fprintf(w, string(htmlPage))
}

func addCocktailHandler(w http.ResponseWriter, r *http.Request) {
	var allIngredients []ingredient.Ingredient
	allIngredients = ingredient.GetAllIngredients()
	var allMeasurements []measurement.Measurement
	allMeasurements = measurement.GetAllMeasurements()

	data := RecipeNecessities{
		Measurements: allMeasurements,
		Ingredients:  allIngredients,
	}
	log.Println(data)
	log.Println("pause")

	t, _ := template.ParseFiles("templates/addCocktail.html")
	t.Execute(w, data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/addCocktail", addCocktailHandler)
	// r.HandleFunc("/saveEntry", saveEntryHandler)
	// r.HandleFunc("/showEntry/{entry_id}", showEntryHandler)
	// r.HandleFunc("/editEntry/{entry_id}", editEntryHandler)
	// r.HandleFunc("/updateEntry/{entry_id}", updateEntryHandler)
	// r.HandleFunc("/deleteEntry/{entry_id}", deleteEntryHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
