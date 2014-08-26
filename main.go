package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
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
	defer r.Body.Close()

	if e != nil {
		log.Fatal(e)
	}

	b, e := ioutil.ReadAll(r.Body)

	if e != nil {
		log.Fatal(e)
	}

	doc := Doc{}

	xml.Unmarshal(b, &doc)

	return &doc
}

func main() {
	// get latest doc
	doc := fetch()

	// set the hashes
	doc.Setup()

	// filter out existing
	// print the results

	s := index(doc)

	fmt.Printf("%s", s.Bytes())
}

func (doc *Doc) Setup() {
	for _, i := range doc.Items {
		i.Sig()
	}
}

func (i *Item) Sig() {
	sha1 := sha1.New().Sum([]byte(i.Title))
	i.Signature = hex.EncodeToString(sha1)
}

func index(doc *Doc) bytes.Buffer {
	var b bytes.Buffer

	s := `
<doctype html>
<html>
	<body>
		<h1>{{len .Items}} items</h1>
		{{range $i, $item := .Items}}
			<div>
				<h4>{{$item.Signature}} {{$item.Title}}</h4>
				{{range $item.Encs}}
					<img src="{{.Resource}}"/>
				{{end}}
			</div>
			<hr/>
		{{end}}
	</body>
</html>
	`

	tmpl := template.Must(template.New("test").Parse(s))

	e := tmpl.Execute(&b, doc)

	if e != nil {
		log.Fatal(e)
	}

	return b

}

func temp() string {
	return `
	`
}
