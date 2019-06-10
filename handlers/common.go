package handlers

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	errs "github.com/ofs/alpha-dataset-api/apierrors"
	"github.com/ofs/alpha-dataset-api/models"
	"github.com/ofs/alpha-dataset-api/store"
	"github.com/pkg/errors"
)

// Config represents the structure which holds information for the datasetAPI handlers
type Config struct {
	DataStore store.Storer
	Host      string
}

const (
	// ContextServiceName represents the context key for the name of the service
	contextServiceName = contextKey("service")
	// DatasetAPI represents the alias name for the dataset API
	datasetAPI = "dataset"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

// datastoreError handles errors returned from datastore
func datastoreError(ctx context.Context, w http.ResponseWriter, err error) {
	log.ErrorCtx(ctx, err, nil)

	if err.Error() == errs.ErrCourseNotFound.Error() {
		err.(*errs.ErrorObject).Message = "resource not found"
		err.(*errs.ErrorObject).Code = http.StatusNotFound

		Error(ctx, w, err)
		return
	}

	http.Error(w, errs.ErrInternalServer.Error(), http.StatusInternalServerError)
	return
}

// Error handles transforms an error into structured error message
func Error(ctx context.Context, w http.ResponseWriter, err error) {
	errorResponse := &models.ErrorResponse{
		Errors: []*models.ErrorObject{
			&models.ErrorObject{
				Error:       err.(*errs.ErrorObject).Error(),
				ErrorValues: err.(*errs.ErrorObject).Values(),
			},
		},
	}

	ErrorResponse(ctx, w, err.(*errs.ErrorObject).Status(), errorResponse)
}

// ErrorResponse sets the structured error message in the http response body
func ErrorResponse(ctx context.Context, w http.ResponseWriter, status int, errorResponse *models.ErrorResponse) {
	b, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, errs.ErrInternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := w.Write(b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "failed to write error response body"), nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// drainBody drains the body of the given of the given HTTP request.
func drainBody(ctx context.Context, r *http.Request) {
	if r.Body == nil {
		return
	}

	_, err := io.Copy(ioutil.Discard, r.Body)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "error draining request body"), nil)
	}

	err = r.Body.Close()
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "error closing request body"), nil)
	}
}

func writeBody(ctx context.Context, w http.ResponseWriter, b []byte) {
	w.Header().Set("Content-Type", "application/json")

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if _, err := w.Write(b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "failed to write response body"), nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
