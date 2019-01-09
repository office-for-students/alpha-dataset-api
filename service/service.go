package service

import (
	"context"

	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/server"
	"github.com/gorilla/mux"
	"github.com/ofs/alpha-dataset-api/config"
	"github.com/ofs/alpha-dataset-api/handlers"
	"github.com/ofs/alpha-dataset-api/store"
)

var (
	httpServer   *server.Server
	serverErrors chan error
)

// DatasetAPI manages stored data
type DatasetAPI struct {
	DataStore store.Storer
	Router    *mux.Router
}

// CreateDatasetAPI manages all the routes configured to API
func CreateDatasetAPI(cfg config.Configuration, dataStore store.Storer, errorChan chan error) {
	router := mux.NewRouter()
	Routes(cfg, router, dataStore)

	httpServer = server.New(cfg.BindAddr, router)

	// Disable this here to allow main to manage graceful shutdown of the entire app.
	httpServer.HandleOSSignals = false

	go func() {
		log.Info("Starting dataset API...", nil)
		if err := httpServer.ListenAndServe(); err != nil {
			log.ErrorC("dataset API http server returned error", err, nil)
			errorChan <- err
		}
	}()
}

// Routes represents a list of endpoints that exist with this api
func Routes(cfg config.Configuration, router *mux.Router, dataStore store.Storer) *DatasetAPI {

	api := DatasetAPI{
		DataStore: dataStore,
		Router:    router,
	}

	host := cfg.Host + cfg.BindAddr

	routeConfig := handlers.Config{DataStore: dataStore, Host: host}
	// api.Router.HandleFunc("/institutions/{institution}", routeConfig.GetInstitutions).Methods("GET")
	// api.Router.HandleFunc("/institutions/{institution}/courses", routeConfig.GetCourses).Methods("GET")
	api.Router.HandleFunc("/institutions/{institution_id}/courses/{course_id}/modes/{mode}", routeConfig.GetCourse).Methods("GET")
	return &api
}

// Close represents the graceful shutting down of the http server
func Close(ctx context.Context) error {
	if err := httpServer.Shutdown(ctx); err != nil {
		return err
	}

	log.Info("graceful shutdown of http server complete", nil)
	return nil
}
