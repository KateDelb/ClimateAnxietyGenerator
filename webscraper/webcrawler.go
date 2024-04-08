package webscraper

import (
	"log"
	"path"
	"strings"
	"sync"

	"github.com/anaskhan96/soup"
	"github.com/chromedp/chromedp"
	"github.com/forestgiant/sliceutil"
)

var Article_url_root string = "https://www.nature.com/"

type Webcrawler struct {
	Todo_urls []string
	Done_urls []string
}

// Constructor
func NewWebcrawler(root string) *Webcrawler {
	wc := new(Webcrawler)
	wc.Todo_urls = append(wc.Todo_urls, root)
	return wc
}

func (w *Webcrawler) Populate_seeds(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	// This function parses the html of the first link and
	// returns a list of following urls to the seeds param
	var seeds []string
	//Step 1: Get the html
	if !sliceutil.Contains(w.Done_urls, url) {
		resp, err := soup.Get(url)
		if err != nil {
			log.Println(err)
		}

		chromedp.Evaluate(resp, nil)

		root := soup.HTMLParse(resp)

		// fmt.Println(root)

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

			w.Todo_urls = append(w.Todo_urls, seeds...)
			w.Done_urls = append(w.Done_urls, url)
		}
	}

}
