package main

// if you want to export the data names then the letters should be capital
import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Breed struct {
	ID     int
	NAME   string
	FAMILY string
	ORIGIN string
}

type Dog struct {
	ID    int
	NAME  string
	BREED *Breed
}

// type Breed struct {
// 	ID     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Family string `json:"family"`
// 	Origin string `json:"origin"`
// }

// type Dog struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Breed *Breed `json:"breed"`
// }

func SendData(w http.ResponseWriter, r *http.Request) {
	breed := Breed{
		ID:     101,
		NAME:   "Lhasa Apso",
		FAMILY: "Canidae",
		ORIGIN: "Tibet",
	}
	dog := Dog{
		ID:    1,
		NAME:  "Lucy",
		BREED: &breed,
	}
	// setting the response type to the json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Change this to your frontend URL
	w.Header().Set("Access-Control-Allow-Methods", "GET")                  // Adjust as needed

	// encoding the dog struct to the json and send it as a response
	err := json.NewEncoder(w).Encode(dog)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
func main() {
	// creating a new instance for a router using mux v 1.8.0
	r := mux.NewRouter()
	r.HandleFunc("/data", SendData)

	http.Handle("/", r)
	http.ListenAndServe(":3001", nil)
}
