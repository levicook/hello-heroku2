package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world")
}

func main() {
	http.HandleFunc("/", helloWorld)

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
