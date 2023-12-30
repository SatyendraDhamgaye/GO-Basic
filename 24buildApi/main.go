package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Shop struct {
	Item_Name   string       `json:"itemname"`
	Item_Price  int          `json:"itemprice"`
	Item_Desc   string       `json:"itemdescription"`
	Manufaturer *Manufaturer `json:"itemmanufaturer"`
}

type Manufaturer struct {
	CompanyName string `json:"companyname"`
	website     string
}

var shop []Shop

func (s *Shop) IsEmpty() bool {
	return s.Item_Name == "" && s.Item_Desc == ""
}

func main() {

}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hi this is API Building first page with nothing after '/' :)</h1>"))
}

func getShopDetails(w http.ResponseWriter, r http.Request) {
	fmt.Println("Getting all the data from shop Database")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shop)
}
