package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const url string = "http://sfbay.craigslist.org/search/eby/apa?bathrooms=2&bedrooms=2&nh=46&nh=47&nh=48&nh=49&nh=62&nh=63&nh=66&s=0&sale_date=-&format=rss"

// The RSS document of lists
type Doc struct {
	XMLName xml.Name `xml:rdf:RDF"`
	Items   []*Item  `xml:"item"`
}

// One posting
type Item struct {
	Title       string       `xml:"title"`
	Link        string       `xml:"link"`
	Description string       `xml:"description"`
	Encs        []*Enclosure `xml:"enclosure"` //http://purl.oclc.org/net/rss_2.0/enc enclosure"`
	Signature   string
}

// Image or other resource
type Enclosure struct {
	Resource string `xml:"resource,attr"`
	Type     string `xml:"type,attr"`
}

func fetch() *Doc {
	r, e := http.Get(url)

	if e != nil {
		log.Fatal(e)
	}

	defer r.Body.Close()

	b, e := ioutil.ReadAll(r.Body)

	if e != nil {
		log.Fatal(e)
	}

	doc := Doc{}

	xml.Unmarshal(b, &doc)

	return &doc
}

func (i *Item) Sig() {
	sha1 := sha1.New().Sum([]byte(i.Title))
	i.Signature = hex.EncodeToString(sha1)
}

func refresh(m *Memory) {
	for {
		doc := fetch()
		m.Update(doc)
		m.Save()

		log.Printf("Refreshed. %d unread items", len(m.UnreadItems))

		time.Sleep(5 * time.Minute)
	}
}
func main() {
	m := CreateMemory(".shacl.json")
	m.Load()

	go refresh(m)
	StartServer(m)
}
