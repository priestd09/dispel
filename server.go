package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.ServeFile(w, req, "static/index.html")
}

func main() {
	// open image DB
	imgDB := newImageDB()

	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/images", imgDB.imageSearchHandler)
	router.GET("/images/show/:img", imgDB.imageShowHandler)

	router.ServeFiles("/static/*filepath", http.Dir("static"))

	log.Println("Listening...")
	http.ListenAndServe(":3000", router)
}
