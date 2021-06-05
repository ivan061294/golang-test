package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JoseEvangelistaCucho/mod/controller"
	"github.com/JoseEvangelistaCucho/mod/utils"
	"github.com/gorilla/mux"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, DELETE, PUT, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		if r.Method == http.MethodOptions {
			w.Header().Set("Content-Type", "application/json")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.Handle("/api/v1/PostArray", http.HandlerFunc(controller.EntitiesArray)).Methods("POST")

	spa := utils.SpaHandler{StaticPath: "public", IndexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	port := os.Getenv("PORT")
	fmt.Printf(port)
	if port == "" {
		port = "8080"
	}
	fmt.Printf(port)
	srv := &http.Server{
		Handler:      CORS(router),
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
