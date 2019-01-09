package store

import (
	"context"

	"github.com/ofs/alpha-dataset-api/models"
)

// DataStore provides a datastore.Storer interface used to store, retrieve, remove or update datasets
type DataStore struct {
	Backend Storer
}

// Storer represents basic data access via Get, Remove and Upsert methods.
type Storer interface {
	GetCourse(ctx context.Context, institutionID, courseID, mode string) (*models.Course, error)
}
