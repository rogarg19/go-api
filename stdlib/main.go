package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", HandleRoot)

	http.ListenAndServe(":8080", mux)
}

func HandleRoot(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello"))
}

func GetAlbums(res http.ResponseWriter, req *http.Request) {
}
