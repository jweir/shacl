package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"unicode/utf8"
)

func urls() []string {
	return []string{
		"http://sfbay.craigslist.org/search/eby/apa?bathrooms=2&bedrooms=2&nh=46&nh=47&nh=48&nh=49&nh=59&nh=60&nh=61&nh=62&nh=63&nh=66&s=0&sale_date=-&format=rss",
		"http://sfbay.craigslist.org/search/eby/apa?bathrooms=2&bedrooms=2&maxAsk=4000&minAsk=1800&nh=112&format=rss",
		"http://sfbay.craigslist.org/search/sfc/apa?bathrooms=2&bedrooms=2&maxAsk=3500&minAsk=1800&format=rss",
		"http://sfbay.craigslist.org/search/pen/apa?bathrooms=2&bedrooms=2&maxAsk=3500&minAsk=1800&nh=73&nh=75&nh=82&nh=89&format=rss",
		// "http://sfbay.craigslist.org/search/nby/apa?bathrooms=2&bedrooms=2&maxAsk=3500&minAsk=1800&nh=102&nh=104&nh=105&format=rss",
		// "http://sfbay.craigslist.org/search/eby/apa?bathrooms=2&bedrooms=2&maxAsk=4000&minAsk=1800&nh=52&nh=57&nh=69&format=rss",

		"http://sfbay.craigslist.org/search/eby/apa?bathrooms=1&bedrooms=3&nh=46&nh=47&nh=48&nh=49&nh=59&nh=60&nh=61&nh=62&nh=63&nh=66&s=0&sale_date=-&format=rss",
		"http://sfbay.craigslist.org/search/eby/apa?bathrooms=1&bedrooms=3&nh=46&nh=47&nh=48&nh=49&nh=59&nh=60&nh=61&nh=62&nh=63&nh=66&s=0&sale_date=-&format=rss",
		"http://sfbay.craigslist.org/search/eby/apa?bathrooms=1&bedrooms=3&maxAsk=4000&minAsk=1800&nh=112&format=rss",
		"http://sfbay.craigslist.org/search/sfc/apa?bathrooms=1&bedrooms=3&maxAsk=3500&minAsk=1800&format=rss",
		"http://sfbay.craigslist.org/search/pen/apa?bathrooms=1&bedrooms=3&maxAsk=3500&minAsk=1800&nh=73&nh=75&nh=82&nh=89&format=rss",
		// "http://sfbay.craigslist.org/search/nby/apa?bathrooms=1&bedrooms=3&maxAsk=3500&minAsk=1800&nh=102&nh=104&nh=105&format=rss",
		// "http://sfbay.craigslist.org/search/eby/apa?bathrooms=1&bedrooms=3&maxAsk=4000&minAsk=1800&nh=52&nh=57&nh=69&format=rss",
	}
}

// The RSS document of lists
type Doc struct {
	XMLName xml.Name `xml:"RDF"`
	Items   []*Item  `xml:"item"`
}

// One posting
type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	// Description string       `xml:"description"`
	Encs      []*Enclosure `xml:"enclosure"` //http://purl.oclc.org/net/rss_2.0/enc enclosure"`
	Signature string
}

// Image or other resource
type Enclosure struct {
	Resource string `xml:"resource,attr"`
	Type     string `xml:"type,attr"`
}

func fetch(url string) (*Doc, error) {
	r, e := http.Get(url)

	if e != nil {
		log.Printf("Error: %s", e)
		return &Doc{}, e
	}

	defer r.Body.Close()

	b, e := ioutil.ReadAll(r.Body)

	if e != nil {
		log.Printf("Error: %s", e)
		return &Doc{}, e
	}

	return parse(b), nil
}

func parse(b []byte) *Doc {
	doc := Doc{}

	s := fmt.Sprintf("%s", b)

	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}

		s = string(v)
	}

	xml.Unmarshal([]byte(s), &doc)

	return &doc
}

func (i *Item) Sig() {
	sha1 := sha1.New().Sum([]byte(i.Title))
	i.Signature = hex.EncodeToString(sha1)
}

func loop(m *Memory) {
	for {
		m.Refresh()
		time.Sleep(5 * time.Minute)
	}
}
func main() {
	m := CreateMemory(".shacl.json")
	m.Load()

	go loop(m)
	StartServer(m)
}
