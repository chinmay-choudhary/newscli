package main

import "fmt"

func CommandTop(countryCode string) error {
	newsData, err := GetTopNews(countryCode)
	if err != nil {
		fmt.Println("An error occured while fetching the news can you check the country code and try again")
	}

	fmt.Println()
	for i, articles := range newsData.Articles {
		if i == 5 {
			break
		}
		fmt.Println("Title: ", articles.Title)
		fmt.Println("Link to articl: ", articles.URL)
		fmt.Println("Article Summary: ", articles.Description)
		fmt.Println()
	}
	return nil
}
