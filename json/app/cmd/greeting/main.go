package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func main() {

	router := chi.NewRouter()
	router.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		var body Person
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "text/plain")
			_, err := w.Write([]byte(`{"message": "pong"}`))
			if err != nil {
				// Handle error here. For example, you might want to log it and return.
				fmt.Println("Error writing response")
				return
			}
			return
		}
		err := json.NewEncoder(w).Encode("Hello " + body.Name + " " + body.LastName)
		if err != nil {
			// Handle error here. For example, you might want to log it and return.
			fmt.Println("Error writing response")
			return
		}
	})

	http.ListenAndServe(":8080", router)
}
