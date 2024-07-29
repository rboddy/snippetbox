package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// a hello message as the response body

func home(w http.ResponseWriter, r *http.Request) {
  // use the header add method to add a 'Server: Go' header to the response
  // header map. The first parameter is the header name, and the second
  // parameter is the header value. This must be done before any calls
  // to write or WriteHeader

  w.Header().Add("Server", "Go")

	w.Write([]byte("Hello from Snippetbox"))
}

// add a snippetview handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

  // we can use fmt.Fprintf instead of w.Write because it still satisfies the
  // io.Writer interface
  fmt.Fprintf(w, "Display a specific snippet with ID %d", id)

}

// add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {

  // use the w.WriteHeader() method to modify the status code that is sent back
  w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet..."))
}

func main() {
	// use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/"
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // add the {id} wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)
  mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// print a log message to say that the server is starting
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
