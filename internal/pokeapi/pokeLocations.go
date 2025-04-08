package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jukkapekkaj/pokedex/internal/pokecache"
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

const poke_api_location_area_url = "https://pokeapi.co/api/v2/location-area"
const NO_MORE_RESULTS = "all-locations-read!"

var cache *pokecache.Cache

func GetNextMap(c *Config, _ []string) error {
	if len(c.Next) == 0 {
		c.Next = poke_api_location_area_url
	} else if c.Next == NO_MORE_RESULTS {
		fmt.Println("Printed all results")
		return nil
	}

	body, err := fetchData(c.Next)
	if err != nil {
		return err
	}

	err = parseResults(c, body)
	if err != nil {
		return err
	}

	return nil
}

func GetPrevMap(c *Config, _ []string) error {
	if len(c.Previous) == 0 {
		c.Previous = poke_api_location_area_url
	}

	body, err := fetchData(c.Previous)
	if err != nil {
		return err
	}

	err = parseResults(c, body)
	if err != nil {
		return err
	}

	return nil
}

func fetchData(url string) ([]byte, error) {
	if cache == nil {
		cache = pokecache.NewCache(time.Second * 20)
	}

	var body []byte
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return []byte{}, err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %v and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			return []byte{}, err
		}
		cache.Add(url, body)
	} else {
		fmt.Println("Found item from cache:", url)
	}

	return body, nil
}

func parseResults(c *Config, data []byte) error {
	var fetchedLocations locationResults

	err := json.Unmarshal(data, &fetchedLocations)
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
	} else {
		c.Previous = poke_api_location_area_url
	}
	return nil
}
