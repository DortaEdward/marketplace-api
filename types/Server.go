package types

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	addr    string
	storage string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	ResponseWithJson(w, 200, struct{}{})
}

func (s *Server) Run() {
	// Middlewares
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthcheck", handleHealthCheck)

	r.Mount("/v1", v1Router)

	fmt.Printf("Listening on PORT: %s\n", s.addr)
	log.Fatal(
		http.ListenAndServe(s.addr, r),
	)
}
