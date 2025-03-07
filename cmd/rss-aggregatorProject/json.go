package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Server Error:", msg)
	}
	type errResponse struct {
		//Below is called a 'struct field tag' or simply 'tag'
		Error string `json:"error"` //The field Error should be serialized as "error" instead of "Error" .
	}
	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	//Marshal refers to serialization—converting a Go object (struct, map, etc.) into a JSON (or other format) string.
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to marshal JSON: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json") //To add Headers
	w.WriteHeader(code)                                // TO add status codes
	w.Write(data)                                      //To add response data
}
