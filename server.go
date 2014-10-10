package main

import (
	"bytes"
	"fmt"
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

	http.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		m.Refresh()
		http.Redirect(w, r, "/", 302)
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

	.item {
		display: inline-block;
		height: 450px;
		width: 300px;
		padding: 10px;
		vertical-align: top;
	}

	.title {
		height: 100px;
	}

	.image {
		height: 300px;
		width: 300px;
		background: #EEE;
	}
	</style>
	<script type="text/javascript" src="http://code.jquery.com/jquery-2.1.1.min.js"></script>
	<script type="text/javascript">
	$(function(){
		$(".remove").click(function(){
			var i = $(this)
			$.get(i.data('url'))
			i.parents('.item').fadeOut(100)
			return false;
		})
	})
	</script>

	</head>
	<body style="text-align: center">
	<h1>{{len .UnreadItems}} items</h1>
	<div><a href="/refresh">Refresh</a></div>
	{{range .UnreadItems}}
	<div class="item" >
		<div class="image">
			<a href="{{.Link}}" target="_blank">
				{{range .Encs}}
					<img src="{{.Resource}}"/>
				{{end}}
			</a>
		</div>
		<div class="title"><a href="{{.Link}}" target="_blank">{{.Title}}</a></div>
		<div style="padding: 10px 0"><a class="remove" style="border-radius: 4px; padding: 4px 8px; background: #600; color: #FFF; font-size: 10px" data-url="/remove?id={{.Signature}}" href="#"><b>Remove</b></a></div>
	</div>
	{{end}}
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
