package main
 
import (
	"github.com/EdmundMartin/democrawl"
	"github.com/PuerkitoBio/goquery"
)
 
type DummyParser struct {
}
 
func (d DummyParser) ParsePage(doc *goquery.Document) democrawl.ScrapeResult {
	data := democrawl.ScrapeResult{}
	data.Title = doc.Find("title").First().Text()
	data.H1 = doc.Find("h1").First().Text()
	return democrawl.ScrapeResult{}
}
 
func main() {
	d := DummyParser{}
	democrawl.Crawl("https://www.theguardian.com/uk", d, 10)
}