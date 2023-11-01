package main

import (
  "net/http"
  "ctf_api/handlers"

  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
)

//Setup with chi for routing
func main() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/", handlers.GetUserHandler)
  http.ListenAndServe(":3000", r)
}

