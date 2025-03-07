package main

import "net/http"

func handle_error(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 200, "Something Went Wrong")
}
