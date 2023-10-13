package main

import "net/http"

func main() {
	// Criando um Multiplexer
	// Ele é um componente no qual adicionamos nossas rotas
	mux := http.NewServeMux()

	// Adicionando uma nova rota
	mux.HandleFunc("/", HomeHandler)

	// Outra forma de criar uma rota (Utilizando Structs)
	// Dessa forma, ele fica mais customizavel
	mux.Handle("/blog", blog{title: "Blog do Erick"})

	// Adicionando o nosso Mux ao Servidor
	http.ListenAndServe(":8080", mux)

	// Criar outra Mux
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Erick"))
	})
	http.ListenAndServe(":8081", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blog struct {
	title string
}

// Adicionando um método a Struct blog
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
