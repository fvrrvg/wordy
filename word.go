package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getWord() string {
	url := "https://randomword.com/noun"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Failed to get word")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err)
	}

	word := doc.Find("#random_word").Text()

	return word

}

func getRareWord() (string, string) {
	url := "https://randomword.com/"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Failed to get word")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err)
	}

	word := doc.Find("#random_word").Text()
	definition := doc.Find("#random_word_definition").Text()

	return word, definition
}
