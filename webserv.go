package main

import (
    "net/http"
    "log"

    "github.com/julienschmidt/httprouter"
)


func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("HIT -- Index")
	http.ServeFile(w,r, "client/index.html")
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("HIT -- Home")
	http.ServeFile(w,r, "client/home.html")
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("HIT -- About")
	http.ServeFile(w,r, "client/about.html")
}


func main() {
    router := httprouter.New()
	
    router.GET("/", index)

	router.GET("/home", home)
	router.GET("/about", about)

	router.ServeFiles("/static/*filepath", http.Dir("client/static"))
	
	log.Println("-- Starting Server --")
    log.Fatal(http.ListenAndServe(":3000", router))
}

