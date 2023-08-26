package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	files, err := os.ReadDir("client")

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, e := range files {
		fmt.Println(e)
	}

	http.HandleFunc("/", root)
	http.ListenAndServe(":3000", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- ROOT --")

	http.ServeFile(w, r, "client/index.html")
}
