package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Service struct {
	apiKey string
	client http.Client
}

func NewService(apiKey string) *Service {
	return &Service{
		apiKey: apiKey,
		client: http.Client{},
	}
}

type respBody struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func (s Service) PrintTemps(cities []string) {
	start := time.Now()

	var wg = &sync.WaitGroup{}

	for i := range cities {
		wg.Add(1)
		go s.getTemp(cities[i], wg)
	}

	wg.Wait()

	fmt.Printf("Time took: %s\n", time.Since(start))
}

func (s Service) getTemp(city string, wg *sync.WaitGroup) {
	const requestURL = "http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s"

	resp, err := s.client.Get(fmt.Sprintf(requestURL, s.apiKey, city))
	if err != nil {
		log.Printf("request to openweathermap failed: %s\n", err.Error())
	}

	var data respBody
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Printf("failed to decode response body: %s\n", err.Error())
	}

	fmt.Printf("City: %s, Temp: %f\n", city, kelvinToCelsius(data.Main.Temp))
	wg.Done()
}

func kelvinToCelsius(temp float64) float64 {
	const kelvinConstant = 273
	return temp - kelvinConstant
}
