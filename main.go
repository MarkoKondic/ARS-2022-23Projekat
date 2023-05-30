// Post API

//    Title: Post API

//    Schemes: http
//    Version: 0.0.1
//    BasePath: /

//    Produces:
//      - application/json

// swagger:meta
package main

import (
	"context"
	ps "github.com/MihajloJankovic/Alati/Dao"
	pss "github.com/MihajloJankovic/Alati/Dao2"
	pm "github.com/MihajloJankovic/Alati/prometheus"
	"github.com/go-openapi/runtime/middleware"
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
	Dao, err := ps.New()
	if err != nil {
		log.Fatal(err)
	}
	Dao2, err := pss.New()
	if err != nil {
		log.Fatal(err)
	}

	server := postServer{Dao: Dao, Dao2: Dao2, keys: make(map[string]string), keys2: make(map[string]string)}
	router.HandleFunc("/config/", pm.CountCreateConfig(server.createPostHandler)).Methods("POST")
	router.HandleFunc("/configs/", pm.CountGetAllConfig(server.getAllHandler)).Methods("GET")
	router.HandleFunc("/config/{id}/{version}/", pm.CountGetConfig(server.getPostHandler)).Methods("GET")
	router.HandleFunc("/config/{id}/{version}/", pm.CountDelConfig(server.delPostHandler)).Methods("DELETE")
	router.HandleFunc("/config/{id}/{version}/{labels}", pm.CountGetConfigByLabels(server.getPostByLabel)).Methods("GET")

	router.HandleFunc("/configGroup/", pm.CountCreateGroup(server.createConfigGroupHandler)).Methods("POST")
	router.HandleFunc("/configGroups/", pm.CountGetAllGroup(server.getAllConfigGroupHandlers)).Methods("GET")
	router.HandleFunc("/configGroup/{id}/", pm.CountGetGroup(server.getConfigGroupHandler)).Methods("GET")
	router.HandleFunc("/configGroup/{id}/", pm.CountDelGroup(server.delConfigGroupHandler)).Methods("DELETE")
	router.HandleFunc("/configGroup/{id}/{version}/{configGroup}/", pm.CountAppendGroup(server.addConfigInConfigGroup)).Methods("GET")

	router.HandleFunc("/swagger.yaml", server.swaggerHandler).Methods("GET")

	router.Path("/metrics").Handler(pm.MetricsHandler())

	//show Traces UI on http://localhost:16686
	//show Metrics UI on show UI on localhost:9090

	// SwaggerUI
	optionsDevelopers := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	developerDocumentationHandler := middleware.SwaggerUI(optionsDevelopers, nil)
	router.Handle("/docs", developerDocumentationHandler)

	// optionsShared := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	//sharedDocumentationHandler := middleware.Redoc(optionsShared, nil)
	// router.Handle("/docs", sharedDocumentationHandler)

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
