package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

type Page struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

var pages map[int]Page = make(map[int]Page)

func FileInit() {
	files, err := ioutil.ReadDir("./articles")
	if err != nil {
		log.Fatal(err)
	}
	// for every article in the articles folder
	for _, file := range files {
		// The following reads the file we are on
		pageFile, err := ioutil.ReadFile("articles/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		// this line makes a new page, called page
		var page Page
		err = json.Unmarshal(pageFile, &page)
		if err != nil {
			log.Fatal(err)
		}
		pageNum, err := strconv.Atoi(file.Name()[:len(file.Name())-len(".json")])
		if err != nil {
			log.Fatal(err)
		}
		pages[pageNum] = page
	}
}

func GetFrontPage() (fpPages []Page) {
	// make a new slice with the capacity to hold 5 , but a length of zero. Lengths and capacity are covered in tour of Go
	fpPages = make([]Page, 0, 5)

	pageNumbers := getPageNumbers()
	// this loop counts backwards from the last element in the page numbers to the fifth to last
	for i := len(pageNumbers) -1; i > len(pageNumbers) -6; i-- {
		// verify that the page actually exists
		if i >= 0 {
			fpPages = append(fpPages, pages[pageNumbers[i]])
		} else {
			// in case it doesn't exist, add and empty page struct which makes everything a nil value
			fpPages = append(fpPages, Page{})
		}
	}
	// no need to specify what we are returning here as we've done that above when we declared the function
	return
}