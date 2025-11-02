package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Carro struct {
	Name string
	Age  int
}

func main() {
	r := chi.NewRouter()
	r.Get("/{username}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "Jorge")
		w.Write([]byte(param))
	})

	r.Post("/create_email", func(w http.ResponseWriter, r *http.Request) {

		var car Carro
		render.DecodeJSON(r.Body, &car)
		fmt.Println(car)

		render.JSON(w, r, car)
	})
	http.ListenAndServe(":3000", r)

}
