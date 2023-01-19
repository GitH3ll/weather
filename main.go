package main

import "weather/internal/weather"

const apiKey = "592ddec5b996fbe95cdd8a649ee3a92b"

func main() {
	service := weather.NewService(apiKey)

	cities := []string{"Minsk", "Moscow", "Berlin", "Tokyo", "Washington", "Tallinn", "Beijing", "Deli", "Rome", "Paris",
		"London", "Baku", "Monaco", "Oslo", "Riga", "Tunis", "Yerevan", "Sofia", "Jerusalem", "Copenhagen", "Cairo",
		"Bangkok", "Minsk", "Moscow", "Berlin", "Tokyo", "Washington", "Tallinn", "Beijing", "Deli", "Rome", "Paris",
		"London", "Baku", "Monaco", "Oslo", "Riga", "Tunis", "Yerevan", "Sofia", "Jerusalem", "Copenhagen", "Cairo",
		"Bangkok"}

	service.PrintTemps(cities)
}
