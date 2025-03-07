package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load("/home/adi/Projects/rss-aggregator/.env")
	fmt.Println("Hello World")
	portString := os.Getenv("PORT")
	if portString == "" {
		fmt.Println("No PORT env variable found")
		log.Fatal("PORT is not found in environment")
	}
	fmt.Println(portString)

	router := chi.NewRouter()
	//Cors Handler

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK! Hello Go"))
	})

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)
	v1Router.Get("/healthz", handle_readiness) //to check health
	v1Router.Get("/err", handle_error)
	fmt.Println("Server running at " + portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
