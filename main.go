package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type embedStruct struct {
	Title       string
	Description string
	Color       string
	Image       string
	Redirect    string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

	fmt.Println(query)

	var embedInfo embedStruct
	if title, ok := query["title"]; ok {
		embedInfo.Title = title[0]
	}

	if desc, ok := query["desc"]; ok {
		embedInfo.Description = desc[0]
	}

	if color, ok := query["color"]; ok {
		embedInfo.Color = color[0]
	}

	if image, ok := query["image"]; ok {
		embedInfo.Image = image[0]
	}

	if url, ok := query["url"]; ok {
		embedInfo.Redirect = url[0]
	}

	tmpl, _ := template.ParseFiles("public/embed.html")
	tmpl.Execute(w, embedInfo)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/e", embed)
    http.HandleFunc("/oembed", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`{"type": "photo", "author_name": "Bees"}`))
    })

	http.ListenAndServe(":8080", nil)
}