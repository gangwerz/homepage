package main

import (
	"fmt"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	http.HandleFunc("/", root)

	fmt.Println("Started Server on Port 3000")
	http.ListenAndServe(":3000", nil)
}

// --- HANDLERS ---

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Route Hit:  ROOT")

	http.ServeFile(w, r, "client/index.html")
}
