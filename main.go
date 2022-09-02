package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	postStrapi()

}

func GetStrapi() {
	fmt.Println("Getting data...")
	res, err := http.Get("http://localhost:1337/api/restaurants")
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
}

func postStrapi() {
	// value := make(map[string]interface{})

	// data := map[string]string{
	// 	"name":        "Nwanyi Igbo",
	// 	"description": "This is a very nice place to eat navite sup",
	// }

	// value["data"] = data

	file, _ := os.Open("new.json")

	// data := make(map[string])

	// responseBody := bytes.NewBuffer(file)

	response, err := http.Post("http://localhost:1337/api/restaurants", "application/json", file)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
}
