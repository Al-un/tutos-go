package intermediate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type pokemonSpecies struct {
	Name string `json:"name"`
}

type pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species pokemonSpecies `json:"pokemon_species"`
}

type pokeResp struct {
	Name     string    `json:"name"`
	Pokemons []pokemon `json:"pokemon_entries"`
}

// RunRestClient https://tutorialedge.net/golang/consuming-restful-api-with-go/
func RunRestClient() {
	resp, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(data))

	var responseObject pokeResp
	json.Unmarshal(data, &responseObject)
	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemons))

	for i := 0; i < len(responseObject.Pokemons); i++ {
		fmt.Println(responseObject.Pokemons[i].Species.Name)
	}
}
