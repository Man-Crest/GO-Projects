package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

var Movies []Movie

func main() {
	r := mux.NewRouter()

	Movies = append(Movies, Movie{
		ID:       "1",
		Isbn:     "001",
		Title:    "speed",
		Director: &Director{"jack", "sparrow"},
	})

	Movies = append(Movies, Movie{
		ID:       "2",
		Isbn:     "002",
		Title:    "Ironman",
		Director: &Director{"Tony", "Lee"},
	})

	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreatMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Println("server is running of port :8000")
	http.ListenAndServe(":8000", r)
}

func GetMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "applicatoin/json")
	json.NewEncoder(w).Encode(Movies)
}
func GetMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "applicatoin/json")
	params := mux.Vars(r)
	for _, val := range Movies {
		if params["id"] == val.ID {
			json.NewEncoder(w).Encode(val)
			return
		}
	}
}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "applicatoin/json")
	params := mux.Vars(r)

	fmt.Println(params["id"])
	for i, value := range Movies {
		fmt.Println(value.ID)
		if params["id"] == value.ID {
			fmt.Println(value)
			Movies = append(Movies[:i], Movies[i+1:]...)
			json.NewEncoder(w).Encode(Movies)
			break
		}
	}
}

func CreatMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(1000))
	Movies = append(Movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func UpdateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatoin/json")
	params := mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	fmt.Println(params["id"], movie)

	for i, val := range Movies {
		if params["id"] == val.ID {
			movie.ID = params["id"]
			Movies = append(Movies[:i], Movies[i+1:]...)
			Movies = append(Movies, movie)
			fmt.Println(movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
		break
	}
}
