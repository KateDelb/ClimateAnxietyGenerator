package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"
	"webscraper/webscraper"
)

var Source_url string = "https://www.nature.com/nclimate/" //DEFINE global variables like this

func main() {
	// crawler := webscraper.NewWebcrawler("https://www.sciencedirect.com/journal/environmental-research")
	crawler := webscraper.NewWebcrawler(Source_url)
	crawler.Populate_seeds() // TO CALL FROM ANOTHER PACKAGE THE FUNCTION HAS TO START WITH A CAPITAL
	fmt.Println(crawler.Seeds)

	//Dyncamically insert the random url so the button redirects to it.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Define the data to be passed to the template
		rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
		random_url := crawler.Seeds[rand.Intn(len(crawler.Seeds))]
		fmt.Println(random_url)
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
