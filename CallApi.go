package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	URL          = "https://newsapi.org/v2/top-headlines?country=%s&apiKey=%s"
	NEWS_API_KEY = "NEWS_API_KEY"
)

func GetTopNews(countryCode string) (NewsAPIResponse, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
		return NewsAPIResponse{}, errors.New("Could not load .env file")
	}

	url := formatUrl(countryCode)

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
		return NewsAPIResponse{}, errors.New("Error making get request")
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return NewsAPIResponse{}, errors.New("Error reading response body")
	}

	// Print the response
	jsonData, err := formatResponse(body)
	if err != nil {
		fmt.Println("Error:", err)
		return NewsAPIResponse{}, errors.New("error while parsing JSON response")
	}

	return jsonData, nil
}

func formatUrl(countryCode string) string {
	newsApiKey := os.Getenv(NEWS_API_KEY)
	url := fmt.Sprintf(URL, countryCode, newsApiKey)
	return url
}
