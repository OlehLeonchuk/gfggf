package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
		return
	}

	fmt.Println("Received data: ", string(body))
}

func main() {
	http.HandleFunc("/", handlePost)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
