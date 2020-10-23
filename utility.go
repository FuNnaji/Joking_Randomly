package main

type jokingRandomlyResponse struct {
	Message string `json:"message"`
}

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
const randomNameURL = "https://api.namefake.com/"
const randomJokeURL = "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]"
