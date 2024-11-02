package handlers

import (
	"html/template"
	"net/http"
)

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("static/html/index.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)

	}
}
