package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	port := "8081"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})
	r.Get("/api/getExample", getHandler)
	r.Get("/api/zincGetExample", getZincBasic)
	r.Post("/api/getZincSearch", getZincSearch)
	r.Post("/nameExample", nameExample)
	r.Get("/postexample", postHandler)
	r.Post("/postzinc", getZincBasic)

	r.Mount("/posts", postsResource{}.Routes())

	log.Fatal(http.ListenAndServe(":"+port, r))

}
