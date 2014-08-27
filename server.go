package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func loadPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func StartServer(m *Memory) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := index(m)
		w.Header().Set("Content-Type", "text/html")
		w.Write(b.Bytes())
	})

	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query()["id"][0]
		item := m.UnreadItems[id]
		m.Remove(item)
		http.Redirect(w, r, "/", 302)
	})

	http.HandleFunc("/find", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query()["id"][0]
		item := m.UnreadItems[id]

		log.Printf("%s\n%s\n", id, item.Link)

		resp, e := http.Get(item.Link)

		defer resp.Body.Close()

		if e != nil {
			log.Fatal(e)
		}

		b, e := ioutil.ReadAll(resp.Body)

		if e != nil {
			log.Fatal(e)
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(m *Memory) bytes.Buffer {
	var b bytes.Buffer

	s := `<doctype html>
	<html>
	<head>
	<link href='http://fonts.googleapis.com/css?family=Open+Sans:400,700' rel='stylesheet' type='text/css'>
	<style>
	body {
		font-family: 'Open Sans', sans-serif;
	}

	a {
		text-decoration: none;
		color: #333;
	}
	</style>
	</head>
	<body style="text-align: center">
	<div style="width: 20%; margin: 20px auto; text-align: left; display: inline-block">
	<h1>{{len .UnreadItems}} items</h1>
	{{range .UnreadItems}}
	<div>
	<div><a href="/find?id={{.Signature}}" target="cl">{{.Title}}</a></div>
	<div style="padding: 10px 0"><a style="border-radius: 4px; padding: 4px 8px; background: #600; color: #FFF; font-size: 10px" href="/remove?id={{.Signature}}"><b>Remove</b></a></div>
		<a href="/find?id={{.Signature}}" target="cl">
			{{range .Encs}}
				<img src="{{.Resource}}"/>
			{{end}}
		</a>
	</div>
	<hr/>
	{{end}}
	</div>
	<div style="width: 78%; display: inline-block; vertical-align: top">
	<iframe src="" id="cl" name="cl" style="width: 100%; border: 0; height: 1000px"></iframe>
	</div>
	</body>
	</html>
	`

	tmpl := template.Must(template.New("test").Parse(s))

	e := tmpl.Execute(&b, m)

	if e != nil {
		log.Fatal(e)
	}

	return b
}
