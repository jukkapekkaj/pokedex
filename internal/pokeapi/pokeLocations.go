package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Config struct {
	Next     string
	Previous string
}

type result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationResults struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []result `json:"results"`
}

const poke_api_location_area_url = "https://pokeapi.co/api/v2/location-area/"
const NO_MORE_RESULTS = "all-locations-read!"

func GetNextMap(c *Config) error {
	if len(c.Next) == 0 {
		c.Next = poke_api_location_area_url
	} else if c.Next == NO_MORE_RESULTS {
		fmt.Println("Printed all results")
		return nil
	}
	res, err := http.Get(c.Next)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %v and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	//fmt.Println(body)

	var fetchedLocations locationResults

	err = json.Unmarshal(body, &fetchedLocations)
	if err != nil {
		return err
	}

	for _, r := range fetchedLocations.Results {
		fmt.Println(r.Name)
	}

	if len(fetchedLocations.Next) > 0 {
		c.Next = fetchedLocations.Next
	} else {
		c.Next = NO_MORE_RESULTS
	}

	if len(fetchedLocations.Previous) > 0 {
		c.Previous = fetchedLocations.Previous
	}

	return nil
}

func GetPrevMap(c *Config) error {
	if len(c.Previous) == 0 {
		c.Previous = poke_api_location_area_url
	}
	res, err := http.Get(c.Previous)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %v and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	//fmt.Println(body)

	var fetchedLocations locationResults

	err = json.Unmarshal(body, &fetchedLocations)
	if err != nil {
		return err
	}

	for _, r := range fetchedLocations.Results {
		fmt.Println(r.Name)
	}

	if len(fetchedLocations.Next) > 0 {
		c.Next = fetchedLocations.Next
	}

	if len(fetchedLocations.Previous) > 0 {
		c.Previous = fetchedLocations.Previous
	} else {
		c.Previous = poke_api_location_area_url
	}

	return nil
}
