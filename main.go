package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ONSdigital/go-ns/log"
	"github.com/ofs/alpha-dataset-api/config"
	"github.com/ofs/alpha-dataset-api/mongo"
	"github.com/ofs/alpha-dataset-api/service"
	"github.com/ofs/alpha-dataset-api/store"
)

// DatasetAPIStore is a wrapper which embeds a Mongo stuct which satisfies the store.Storer interface.
type DatasetAPIStore struct {
	*mongo.Mongo
}

func main() {
	log.Namespace = "alpha-dataset-api"

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Get()
	if err != nil {
		log.ErrorC("errored getting configuration", err, log.Data{"config": cfg})
		os.Exit(1)
	}

	log.Info("configuration on startup", log.Data{"config": cfg})

	mongodb := &mongo.Mongo{
		Collection: cfg.MongoConfig.Collection,
		Database:   cfg.MongoConfig.Database,
		URI:        cfg.MongoConfig.BindAddr,
	}

	// connect to mongo
	session, err := mongodb.Init()
	if err != nil {
		log.ErrorC("failed to initialise mongo", err, nil)
		os.Exit(1)
	}

	mongodb.Session = session

	store := store.DataStore{Backend: DatasetAPIStore{Mongo: mongodb}}

	apiErrors := make(chan error, 1)

	service.CreateDatasetAPI(*cfg, store.Backend, apiErrors)

	// Gracefully shutdown the application closing any open resources.
	gracefulShutdown := func() {
		start := time.Now()
		log.Info("shutdown with timeout in seconds", log.Data{"timeout": cfg.GracefulShutdownTimeout})
		ctx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)

		// stop any incoming requests before closing any outbound connections
		service.Close(ctx)

		// TODO close connection to database

		log.Info("shutdown complete", log.Data{"shutdown_duration": time.Since(start)})

		cancel()
		os.Exit(1)
	}

	for {
		select {
		case err := <-apiErrors:
			log.Info("api error received", log.Data{"api_error": err})
			gracefulShutdown()
		case signal := <-signals:
			log.Info("os signal received", log.Data{"os_signal": signal})
			gracefulShutdown()
		}
	}
}
