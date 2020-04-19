package main

//Importamos modulo de red para implementar el servidor
import (
	"net/http"
)

func main() {

	// routers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	// startthe server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello world"))

}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Contact Page"))

}
