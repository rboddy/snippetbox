package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
  addr := flag.String("addr", ":4000", "HTTP Network Address")

  flag.Parse()

  mux := http.NewServeMux()

  fileServer := http.FileServer(http.Dir("./ui/static/"))

  mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

  mux.HandleFunc("GET /{$}", home)
  mux.HandleFunc("GET /snippet/view/{id}", snippetView)
  mux.HandleFunc("GET /snippet/view/create", snippetCreate)
  mux.HandleFunc("POST /snippet/view/create", snippetCreatePost)

  log.Printf("starting server on %s", *addr)

  err := http.ListenAndServe(*addr, mux)
  log.Fatal(err)
}
