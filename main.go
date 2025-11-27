package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// "net/http" package provides HTTP client and server implementations.
// "io" package provides basic interfaces to I/O primitives.
// "os" package provides a platform-independent interface to operating system functionality.
// "errors" package implements functions to manipulate errors.

// in this method signature, we are using http.ResponseWriter and *http.Request as parameters.
// http.ResponseWriter is an interface that allows us to write HTTP response data. w is the variable name (alias) for the ResponseWriter parameter.
// *http.Request is a struct that represents an HTTP request received by a server or to be sent by a client. r is the variable name (alias) for the Request parameter handles the incoming HTTP request.
// this method handles requests to the root endpoint /.
// ResponseWriter is an interface meant to be passed as an object, whereas Request is a struct type, so it is passed as a pointer to avoid copying the entire struct.
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request for /")
	io.WriteString(w, "Hello World\n")
	// this line writes "Hello World" to the response writer w, which sends it back to the client that made the request.
}

// this method handles requests to the /hello endpoint.
// w is the variable name (alias) for the ResponseWriter parameter.
// r is the variable name (alias) for the Request parameter handles the incoming HTTP request.
// this method writes "Hello from /hello endpoint" to the response writer w, which sends it back to the client that made the request.
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request for /hello")
	io.WriteString(w, "Hello from /hello endpoint\n")
}

func main() {
	// handleFunc associates the specified path with the handler function.
	// when a request is made to the specified path, the corresponding handler function is invoked to process the request.
	// here we are associating the root endpoint / with the getRoot function and /hello endpoint with getHello function.
	// when a request is made to /, getRoot function will be called.
	// when a request is made to /hello, getHello function will be called.
	// these functions will handle the requests and send responses back to the client.
	// it is like using urlpatterns in Django framework.
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	// this will open a web server on port 8080 and listen for incoming HTTP requests.
	err := http.ListenAndServe(":8080", nil)
	// ListenAndServe starts an HTTP server with a given address and handler. it is a blocking call, meaning it will run indefinitely until the program is terminated or an error occurs.

	// ErrServerClosed is returned when the server is told to shut down or close.
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("HTTP server has been shutdown")
	} else if err != nil {
		fmt.Printf("Error starting HTTP server: %v\n", err)
		os.Exit(1)
		// Exit function terminates the program with the given exit code. A non-zero exit code indicates an error.
	}

	// MULTIPLEXING REQUEST HANDLERS USING SERVEMUX
	// A ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.
	// Here, we create a new ServeMux instance and register our handler functions with it.
	// Finally, we start the HTTP server with the ServeMux as the handler.
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	err = http.ListenAndServe(":8080", mux)
}
