package hof

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func CarsIndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response, err := getAllCarsJson()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
}

func CarHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	carIndex, err := strconv.Atoi(p[0].Value)
	if err != nil {
		log.Fatal("CarHandler unable to find car (%v) by index\n", p[0].Value)
	}
	response, err := getThisCarJson(carIndex)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
}

func getAllCarsJson() ([]byte, error) {
	return json.MarshalIndent(Payload{CarsDB}, "", "  ")
}

func getThisCarJson(carIndex int) ([]byte, error) {
	return json.MarshalIndent(CarsDB[carIndex], "", "  ")
}

// retrieve performs a HTTP Get request for the rss feed and decodes the results.
func GetThisCar(carIndex int) (*IndexedCar, error) {

	thisCarJson, err := getThisCarJson(carIndex)
	if err != nil {
		panic(err)
	}

	// Decode the json into our struct type.
	// We don't need to check for errors, the caller can do this.

	var thisIndexedCar IndexedCar
	if err := json.Unmarshal(thisCarJson, &thisIndexedCar); err != nil {
		panic(err)
	}
	//log.Printf("carIndex: %v\n", carIndex)
	//log.Printf("thisCarJson: %v\n", thisCarJson)
	//log.Printf("thisIndexedCar: %v\n", thisIndexedCar)
	return &thisIndexedCar, err
}
