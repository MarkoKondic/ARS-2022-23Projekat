package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	server := postServer{
		dataConfig: make(map[string]*Config),
		data:       make(map[string]*ConfigGroup),
	}
	router.HandleFunc("/config/", server.createPostHandler).Methods("POST")
	router.HandleFunc("/configs/", server.getAllHandler).Methods("GET")
	router.HandleFunc("/config/{id}/", server.getPostHandler).Methods("GET")
	router.HandleFunc("/config/{id}/", server.delPostHandler).Methods("DELETE")

	router.HandleFunc("/configGroup/", server.createConfigGroupHandler).Methods("POST")
	router.HandleFunc("/configGroups/", server.getAllConfigGroupHandlers).Methods("GET")
	router.HandleFunc("/configGroup/{id}/", server.getConfigGroupHandler).Methods("GET")
	router.HandleFunc("/configGroup/{id}/", server.delConfigGroupHandler).Methods("DELETE")
	router.HandleFunc("/configGroup/{config}/{configGroup}/", server.addConfigInConfigGroup).Methods("GET")

	// start server
	srv := &http.Server{Addr: "0.0.0.0:8000", Handler: router}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}