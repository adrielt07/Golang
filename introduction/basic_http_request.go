package main

import (
	"fmt"
	"net/http"
)

type myHandler struct {
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("<h1>Hello %v</h1>", mh.greeting)))
}

func main() {
	http.Handle("/", &myHandler{greeting: "World"})

	http.ListenAndServe(":8000", nil)
}
