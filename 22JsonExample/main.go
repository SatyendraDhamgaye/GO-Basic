package main

import (
	"encoding/json"
	"fmt"
)

type shop struct {
	Name     string `json:"shopName"` //to provide diffrent name in json
	Owner    string `json:"ownerName"`
	Rating   int
	Password string   `json:"-"`              //to remove any feild entierly form the json
	Items    []string `json:"tags,omitempty"` //omiempty to remove empty feilds
}

func EncodeJson() {

	market := []shop{
		{"Dev1", "Saty", 4, "ffff1111", []string{"graphic-card", "ram"}},
		{"Dev2", "kappya", 2, "ffff1111", []string{"graphic-card", "ram"}},
		{"Dev3", "dattya", 3, "ffff1111", nil},
	}

	jsonData, err := json.MarshalIndent(market, "", "\t")
	catchError(err)

	fmt.Printf("%s\n", jsonData)
}

func main() {
	EncodeJson()
	DecodeJson()
}

func catchError(err error) {
	if err != nil {
		panic(err)
	}
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"shopName": "Dev1",
			"ownerName": "Saty",
			"Rating": 4,
			"tags": ["graphic-card", "ram"]
	}
	`)

	var marketfromweb shop

	checkJson := json.Valid(jsonDataFromWeb)
	if checkJson {
		fmt.Println("JSON is Valide")
		json.Unmarshal(jsonDataFromWeb, &marketfromweb)
		fmt.Printf("%#v\n", marketfromweb)
	}

	var jsoninmap map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &jsoninmap)
	for k, v := range jsoninmap {
		fmt.Printf("Key is %v and Value is %v\n", k, v)
	}
}
