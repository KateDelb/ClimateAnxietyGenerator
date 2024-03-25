package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"sync"
	"time"
	"webscraper/webscraper"
)

var Source_url string = "https://www.nature.com/nclimate/"

func main() {
	// crawler := webscraper.NewWebcrawler("https://www.sciencedirect.com/journal/environmental-research")
	start := time.Now()
	var wg sync.WaitGroup
	crawler := webscraper.NewWebcrawler(Source_url)
	for _, url := range crawler.Todo_urls {
		go crawler.Populate_seeds(url, &wg)
		wg.Add(1)
	}
	wg.Wait() // Wait for all workers to finish
	fmt.Println("All workers finished execution, len of todo list", len(crawler.Todo_urls))
	elapsed := time.Since(start)
	log.Printf("Fetching the urls took %s", elapsed)

	// 	//Dyncamically insert the random url so the button redirects to it.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Define the data to be passed to the template
		rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
		random_url := crawler.Todo_urls[rand.Intn(len(crawler.Todo_urls))]
		Data := random_url
		fmt.Printf("URL: %s\n", Data)

		// Parse the HTML template
		tmpl, err := template.ParseFiles(filepath.Join("templates", "test.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Render the template with the data and write it to the response
		err = tmpl.Execute(w, Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	// Start the web server
	http.ListenAndServe(":8000", nil)
}
