package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

const url string = "http://sfbay.craigslist.org/search/eby/apa?bathrooms=2&bedrooms=2&nh=46&nh=47&nh=48&nh=49&nh=62&nh=63&nh=66&s=0&sale_date=-&format=rss"

type Doc struct {
	XMLName xml.Name `xml:rdf:RDF"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	Title       string      `xml:"title"`
	Link        string      `xml:"link"`
	Description string      `xml:"description"`
	Encs        []Enclosure `xml:"enclosure"` //http://purl.oclc.org/net/rss_2.0/enc enclosure"`
}

type Enclosure struct {
	Resource string `xml:"resource,attr"`
	Type     string `xml:"type,attr"`
}

func main() {
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

	s := index(&doc)

	fmt.Printf("%s", s.Bytes())
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
				<h4>{{$item.Title}}</h4>
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
