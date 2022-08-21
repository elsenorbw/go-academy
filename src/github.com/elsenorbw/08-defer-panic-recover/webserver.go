package main

import "net/http"

func theHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go!"))
}

func main() {

	// Configure a handler
	http.HandleFunc("/", theHandler)

	// Startup the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

}
