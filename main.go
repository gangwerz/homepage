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

	fmt.Println("Available Client Files")

	for _, e := range files {
		fmt.Println(e)
	}

	http.HandleFunc("/", root)

	fmt.Println("Starting Server on port 3000...")
	http.ListenAndServe(":3000", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- ROOT --")

	http.ServeFile(w, r, "client/index.html")
}
