package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/go-resty/resty/v2"
)

type XxxStreamResponse struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type Geolocation struct {
	Altitude  float64 `json:"Altitude"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}

type Weather struct {
	Temperature float64 `json:"Temperature"`
	Humidity    float64 `json:"Humidity"`
}

func main() {
	client := resty.New()

	resp, err := client.R().
		SetDoNotParseResponse(true).
		Get("http://localhost:1323")

	if err != nil {
		log.Fatalf("Error while making request: %v", err)
	}
	defer resp.RawBody().Close()

	decoder := json.NewDecoder(resp.RawBody())
	for {
		var msg XxxStreamResponse
		if err := decoder.Decode(&msg); err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error while decoding JSON: %v", err)
		}

		switch msg.Type {
		case "Geolocation":
			var loc Geolocation
			if err := json.Unmarshal(msg.Data, &loc); err != nil {
				log.Printf("Error decoding Geolocation: %v", err)
				continue
			}
			fmt.Printf("Received Geolocation: %+v\n", loc)
		case "Weather":
			var wth Weather
			if err := json.Unmarshal(msg.Data, &wth); err != nil {
				log.Printf("Error decoding Weather: %v", err)
				continue
			}
			fmt.Printf("Received Weather: %+v\n", wth)
		default:
			fmt.Printf("Unknown type: %s\n", msg.Type)
		}
	}
}
