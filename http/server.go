package http

import (
	"context"
	"github.com/bandgren/classified-ads/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var srv *http.Server

// Start starts the http server
func Start() {
	createServer()
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

// Shutdown shuts down the http server
func Shutdown(ctx context.Context) (done chan struct{}) {
	done = make(chan struct{})
	go func() {
		defer close(done)
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("couldnt shutdown server error [%s]\n", err)
		}
	}()
	return
}

func createServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/ads", controllers.GetAds).Methods("GET")
	router.HandleFunc("/api/ads", controllers.CreateAd).Queries().Methods("POST")
	router.HandleFunc("/api/ads/{id}", controllers.DeleteAd).Methods("DELETE")
	srv = &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

}
