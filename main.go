package main

// The following implements the main Go
// package starting up the WOOT server

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/woot-crdt/handlers"
)

const (
	// PORT is the WOOT
	// server port
	PORT = "8080"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := handlers.Router()

	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowCredentials: true,
	// })
	// handler := c.Handler(r)

	// headersOk := gohandlers.AllowedHeaders([]string{"X-Requested-With"})
	// originsOk := gohandlers.AllowedOrigins([]string{"*"})
	// methodsOk := gohandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("started WOOT node server")
	r.Use(accessControlMiddleware)
	http.ListenAndServe(":"+PORT, r)

	// http.ListenAndServe(":"+PORT, gohandlers.CORS(headersOk, originsOk, methodsOk)(r))
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")

		if r.Method == "OPTIONS" {
			// return ok
			return
		}

		next.ServeHTTP(w, r)
	})
}
