package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type randomNameResponse struct {
	Name string `json:"name"`
}

type randomJokeResponse struct {
	Type  string `json:"type"`
	Value joke   `json:"value"`
}

type joke struct {
	Joke string `json:"joke"`
}

const randomJokeNameConstant = "John Doe"
const errorConstant = "Error: "

func main() {
	randomName := fetchRandomName("")
	randomJoke := fetchRandomJoke("")
	jokingRandomly := strings.Replace(randomJoke, randomJokeNameConstant, randomName, 1)
	log.Println(jokingRandomly)
}

func fetchRandomName(extraString string) string {
	var response randomNameResponse

	resp, err := http.Get("https://api.namefake.com/")
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
	log.Println(response.Name + extraString)
	return response.Name + extraString
}

func fetchRandomJoke(extraLink string) string {
	var response randomJokeResponse

	resp, err := http.Get("http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]" + extraLink)
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
