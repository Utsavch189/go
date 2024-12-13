package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://pokeapi.co/api/v2/pokemon?limit=151"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // caller's responsibility to close the connection
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	string_data := string(databytes)

	fmt.Printf("Data type for string_data is : %T\n ", string_data)
	//fmt.Println("string_data : ", string_data) // But the json response data in string now

	// Parse into json
	type Pokemon struct {
		Name string `json:"name"` // Maps JSON key "name" to Go field "Name"
		URL  string `json:"url"`  // Maps JSON key "url" to Go field "URL"
	}

	type PokemonResponse struct {
		Count    int       `json:"count"`    // Maps JSON key "count" to Go field "Count"
		Next     string    `json:"next"`     // Maps JSON key "next" to Go field "Next"
		Previous string    `json:"previous"` // Maps JSON key "previous" to Go field "Previous"
		Results  []Pokemon `json:"results"`  // Maps JSON key "results" to Go field "Results"
	}

	// Parse the JSON data
	var pokemonResponse PokemonResponse
	err1 := json.Unmarshal(databytes, &pokemonResponse)
	if err1 != nil {
		fmt.Println("Error parsing JSON:", err1)
		return
	}
	// Print the parsed data
	fmt.Printf("Total Pokémon Count: %d\n", pokemonResponse.Count)
	fmt.Printf("Next Page URL: %s\n", pokemonResponse.Next)
	fmt.Println("Pokémon List:")
	for idx, p := range pokemonResponse.Results {
		fmt.Printf("%d. %s (%s)\n", idx+1, p.Name, p.URL)
	}

	// Re-encode the JSON (if needed)
	json_data, err := json.Marshal(pokemonResponse)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data type for json_data is : %T\n", json_data)
	fmt.Println("json_data : ", string(json_data))
}
