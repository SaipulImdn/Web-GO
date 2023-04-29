package main

import (
	"fmt"
	"net/http"
)

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello again"))
	})

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)

	fmt.Println("server ready at :9000")
	http.ListenAndServe(":9000", nil)
}
