package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	timer := time.NewTicker(3 * time.Second)
	defer timer.Stop()

	for range timer.C {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		waterStatus := "aman"
		if water >= 6 && water <= 8 {
			waterStatus = "siaga"
		} else if water > 8 {
			waterStatus = "waspada"
		}

		windStatus := "aman"
		if wind >= 7 && wind <= 15 {
			windStatus = "siaga"
		} else if wind > 15 {
			windStatus = "bahaya"
		}

		weather := Weather{
			Water: water,
			Wind:  wind,
		}

		requestJson, err := json.Marshal(weather)

		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body))
		log.Println("status water : ", waterStatus)
		log.Println("status wind : ", windStatus)
	}

}
