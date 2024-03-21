package webscraper

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/chromedp/chromedp"
)

var Article_url_root string = "https://www.nature.com/"

type Webcrawler struct {
	Root  string
	Seeds []string
}

// Constructor
func NewWebcrawler(root string) *Webcrawler {
	wc := new(Webcrawler)
	wc.Root = root
	return wc
}

func (w *Webcrawler) Populate_seeds() {
	// This function parses the html of the first link and
	// returns a list of following urls to the seeds param
	var seeds []string
	//Step 1: Get the html
	resp, err := soup.Get(w.Root)
	if err != nil {
		log.Fatal(err)
	}

	chromedp.Evaluate(resp, nil)

	root := soup.HTMLParse(resp)

	fmt.Println(root)

	// 	//Step 2: Parse the html + get links
	href_root_obj := root.FindAll("a")
	for _, r := range href_root_obj {
		for _, attr := range r.Pointer.Attr {
			if attr.Key == "href" {
				if strings.Contains(attr.Val, "articles") {
					full_article_url := path.Join(Article_url_root, attr.Val)
					seeds = append(seeds, full_article_url)
				}
			}
		}

		w.Seeds = seeds
	}

}
