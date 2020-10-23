package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	port := 5000
	fmt.Println("Joking Randomly >> starting...")

	http.HandleFunc("/", jokingRandomlyHandler)
	fmt.Printf("Joking Randomly >> serving on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

	fmt.Println("Joking Randomly >> exited...")
}

func jokingRandomlyHandler(w http.ResponseWriter, r *http.Request) {
	randomName := fetchRandomName(randomNameURL)
	randomJoke := fetchRandomJoke(randomJokeURL)
	jokingRandomly := strings.Replace(randomJoke, randomJokeNameConstant, randomName, 1)
	log.Println(jokingRandomly)
	response := jokingRandomlyResponse{Message: jokingRandomly}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

func fetchRandomName(URL string) string {
	var response randomNameResponse

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err.Error())
		return errorConstant + err.Error()
	}

	decoder := json.NewDecoder(resp.Body)
	decodeErr := decoder.Decode(&response)
	if decodeErr != nil {
		fmt.Println(decodeErr.Error())
		return errorConstant + decodeErr.Error()
	}
	log.Println(response.Name)
	return response.Name
}

func fetchRandomJoke(URL string) string {
	var response randomJokeResponse

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err.Error())
		return errorConstant + err.Error()
	}

	decoder := json.NewDecoder(resp.Body)
	decodeErr := decoder.Decode(&response)
	if decodeErr != nil {
		fmt.Println(decodeErr.Error())
		return errorConstant + decodeErr.Error()
	}
	log.Println(response.Value.Joke)
	return response.Value.Joke
}
