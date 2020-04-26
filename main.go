package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Nome: "Camiseta", Descricao: "camiseta do verão", Preco: 123.0, Quantidade: 100},
		{Nome: "Notebook", Descricao: "acer bolado", Preco: 1000.01, Quantidade: 10},
		{Nome: "Tenis Nike SB", Descricao: "Nike bom pra sk8", Preco: 250.05, Quantidade: 50},
	}
	templates.ExecuteTemplate(w, "Index", products)
}
