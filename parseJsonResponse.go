package main

import (
	"encoding/json"
	"fmt"
)

type NewsAPIResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      Source  `json:"source"`
	Author      string  `json:"author"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         string  `json:"url"`
	URLToImage  string  `json:"urlToImage"`
	PublishedAt string  `json:"publishedAt"`
	Content     *string `json:"content"`
}

type Source struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
}

func formatResponse(unformattedJson []byte) (NewsAPIResponse, error) {
	var response NewsAPIResponse
	err := json.Unmarshal(unformattedJson, &response)
	if err != nil {
		return NewsAPIResponse{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return response, nil
}
