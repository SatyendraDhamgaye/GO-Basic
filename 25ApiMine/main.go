package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Game struct {
	Game_id        string     `json:"id"`
	Game_name      string     `json:"name"`
	Game_price     float64    `json:"price"`
	Game_meta_data *Meta_data `json:"meta"`
	//can also remove * and it'll work but it copies it, using * will use reference of it rather that copyinf it
}

type Meta_data struct {
	Studio_name  string `json:"studio"`
	Release_data string `json:"release-date"`
}

// creating temp DB
var Games []Game

// helper methods here :-
func (g Game) isEmpty() bool {
	return g.Game_name == "" && g.Game_price == 0 && g.Game_price < 0
}

func main() {
	fmt.Println("gaming list")
	r := mux.NewRouter()

	//seeding, putting values inside the fake DB
	Games = append(Games, Game{Game_id: "1", Game_name: "RE9",
		Game_price: 3999.99, Game_meta_data: &Meta_data{Studio_name: "CAPCOM", Release_data: "TBD"}})
	Games = append(Games, Game{Game_id: "2", Game_name: "GTAVI",
		Game_price: 4999.99, Game_meta_data: &Meta_data{Studio_name: "ROCKSTAR", Release_data: "TBD"}})
	Games = append(Games, Game{Game_id: "3", Game_name: "ERDLC",
		Game_price: 1999.99, Game_meta_data: &Meta_data{Studio_name: "FROMSOFT", Release_data: "TBD"}})

	//router
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/games", getGameList).Methods("GET")
	r.HandleFunc("/game/{id}", getOneGame).Methods("GET")
	r.HandleFunc("/game", createOneGame).Methods("POST")
	r.HandleFunc("/game/{id}", updateOneGame).Methods("PUT")
	r.HandleFunc("/game/{id}", deleteOneGame).Methods("DELETE")

	//web localhost listner
	log.Fatal(http.ListenAndServe(":8000", r))
}

//controller here :-

// homepage route
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to home page</h1>"))
}

//getGameList route

func getGameList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all the data from the Database")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Games)
}

// getOneGame route
func getOneGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting only one game data using GameId")
	w.Header().Set("Content-Type", "application/json")

	//getting data from user
	params := mux.Vars(r)
	//to get one game
	//loop through games, find matching id, return the response

	for _, value := range Games {
		if value.Game_id == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	json.NewEncoder(w).Encode("No Game data found for the given ID")
	return
}

//for user to create a new game and add it to Games slice

func createOneGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a new Game and adding it database")
	w.Header().Set("Content-Type", "application/json")

	//if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// if data is -{}
	var gameList Game
	_ = json.NewDecoder(r.Body).Decode(&gameList)
	if gameList.isEmpty() {
		json.NewEncoder(w).Encode("Please send some data, No data inside the json")
		return
	}

	//if title is duplicate
	for _, value := range Games {
		if gameList.Game_name == value.Game_name {
			json.NewEncoder(w).Encode("The Game name provided is duplication please change the name")
			return
		}
	}

	rand.Seed(time.Now().Unix())
	gameList.Game_id = strconv.Itoa(rand.Intn(100))
	Games = append(Games, gameList)
	json.NewEncoder(w).Encode(gameList)
	return
}

// for user to update one of the game using gameid
func updateOneGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one of the game using Id")
	w.Header().Set("Content-Type", "application/json")

	//getting Id from the user
	params := mux.Vars(r)

	//loop to find the given id, remove it, add new data with the given id
	for index, value := range Games {
		if value.Game_id == params["id"] {
			Games = append(Games[:index], Games[index+1:]...)
			var gameList Game
			_ = json.NewDecoder(r.Body).Decode(&gameList)
			gameList.Game_id = params["id"]
			Games = append(Games, gameList)
			json.NewEncoder(w).Encode(Games)
			return
		}
	}

	//if no id is provided
	if params["id"] == "0" {
		json.NewEncoder(w).Encode("Please send a id lol")
		return
	}
}

//delete data of the given ID

func deleteOneGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleteing one game data of the given ID")
	w.Header().Set("Content-Type", "application/json")

	//to get data from the user
	params := mux.Vars(r)

	//loop to find given ID, delete
	for index, value := range Games {
		if value.Game_id == params["id"] {
			Games = append(Games[:index], Games[index+1:]...)
			json.NewEncoder(w).Encode("The data for given id is removed")
			json.NewEncoder(w).Encode(Games)
			break
		}
	}
}
