package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

const port = "9000"

// EchoServer returns the query params to the requestor, in JSON format.
func EchoServer(w http.ResponseWriter, req *http.Request) {

	result, err := url.ParseQuery(req.URL.RawQuery)

	if err != nil {
		log.Fatal("Could not parse query params: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(result)

	if err != nil {
		log.Fatal("Could not marshal into JSON: ", err)
	}

	w.Write(j)

}

func main() {
	http.HandleFunc("/", EchoServer)
	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
