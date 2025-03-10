package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Article struct {
	Author  string
	Content string
}

var articles = []Article{
	{
		Author:  "John Doe",
		Content: "This is a sample article about technology.",
	},
	{
		Author:  "Jane Smith",
		Content: "This article discusses recent advancements in AI.",
	},
	{
		Author:  "Bob Johnson",
		Content: "A brief overview of sustainable energy solutions.",
	},
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "articleId")
	articleId, err := strconv.Atoi(urlId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(articles[articleId].Content))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.Route("/{articleID}", func(r chi.Router) {
		r.Get("/", getArticle)
	})
	http.ListenAndServe(":3000", r)
}
