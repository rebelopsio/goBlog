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
		log.fatal(err)
	}
	// for every article in the articles folder
	for _, file := range files {
		// The following reads the file we are on
		pageFile, err := ioutil.ReadFile("articles/" + file.Name())
		if err != nil {
			log.fatal(err)
		}
		// this line makes a new page, called page
		var page Page
		err := json.Unmarshal(pageFile, &page)
		if err != nil {
			log.fatal(err)
		}
		pageNum, err := strconv.Atoi(file.name()[:len(file.Name())-len(".json")])
		if err != nil {
			log.fatal(err)
		}
		pages[pageNum] = page
	}
}