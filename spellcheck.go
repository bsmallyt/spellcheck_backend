package main

import (
	"encoding/json"
	"net/http"
)

func spellcheck(w http.ResponseWriter, r *http.Request) {
    // Check that GET
    if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

    // Create response data
	words := []string{"apple", "dog", "ben", "job", "run"}

	// Convert response to JSON and write it
	if err := json.NewEncoder(w).Encode(words); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

    switch val := r.URL.Query().Get("val"); val {
    case "spellcheck":
        spellcheck(w, r)
    default:
        http.Error(w, `{"error": "Invalid or missing 'val' parameter"}`, http.StatusBadRequest)
    }

}

func main() {
	http.HandleFunc("/", handler)

	port := ":8090"
	println("Server running on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
