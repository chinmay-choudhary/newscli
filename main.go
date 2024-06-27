package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter country for which you want the top news")
	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()

	userInput := reader.Text()

	words := strings.Fields(userInput)

	countryCode := words[0]
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
}
