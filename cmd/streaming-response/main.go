package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
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

var (
	locations = []Geolocation{
		{-97, 37.819929, -122.478255},
		{1899, 39.096849, -120.032351},
		{2619, 37.865101, -119.538329},
		{42, 33.812092, -117.918974},
		{15, 37.77493, -122.419416},
	}

	weather = []Weather{
		{25, 0.5},
		{30, 0.6},
		{35, 0.7},
		{40, 0.8},
	}
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		enc := json.NewEncoder(c.Response())
		for _, l := range locations {

			jsonStr, _ := json.Marshal(l)
			g := XxxStreamResponse{
				Type: "Geolocation",
				Data: jsonStr,
			}
			if err := enc.Encode(g); err != nil {
				return err
			}
			c.Response().Flush()
			time.Sleep(1 * time.Second)
		}

		for _, w := range weather {
			jsonStr, _ := json.Marshal(w)
			g := XxxStreamResponse{
				Type: "Weather",
				Data: jsonStr,
			}

			if err := enc.Encode(g); err != nil {
				return err
			}
			c.Response().Flush()
			time.Sleep(1 * time.Second)
		}
		return nil
	})
	e.Logger.Fatal(e.Start(":1323"))
}
