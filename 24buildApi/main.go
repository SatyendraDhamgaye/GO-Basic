package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Shop struct {
	Item_Name   string       `json:"itemname"`
	Item_Price  int          `json:"itemprice"`
	Item_Desc   string       `json:"itemdescription"`
	Manufaturer *Manufaturer `json:"itemmanufaturer"`
	Shop_Name   string       `json:"shopname"`
	Shop_Id     string       `json:"shopid"`
}

type Manufaturer struct {
	CompanyName string `json:"companyname"`
	website     string
}

var shop []Shop

func (s *Shop) IsEmpty() bool {
	return s.Item_Name == ""
}

func main() {

}

// controller starts
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hi this is API Building first page with nothing after '/' :)</h1>"))
}

func getShopDetails(w http.ResponseWriter, r http.Request) {
	fmt.Println("Getting all the data from shop Database")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shop)
}

func getOneShopData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting values from shop")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range shop {
		if course.Shop_Name == params["sname"] {
			json.NewEncoder(w).Encode(shop)
			return
		}
	}
	json.NewEncoder(w).Encode("No data found, please enter correct data")
	return
}

func createOneShop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one Shop")
	w.Header().Set("Content-Type", "applciation/json")

	//if body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is empty, send some data")
	}

	//if json data is {}
	var shops Shop
	_ = json.NewDecoder(r.Body).Decode(&shop)
	if shops.IsEmpty() {
		json.NewEncoder(w).Encode("json is empty, send some data")
		return
	}

	//creating random unique shopID integer
	rand.Seed(time.Now().Unix())
	shops.Shop_Id = strconv.Itoa(rand.Intn(99))
	shops = append(shops, shop)
	json.NewEncoder(w).Encode("data is given")
	return
}

//controller ends
