package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
	errs "github.com/ofs/alpha-dataset-api/apierrors"
	"github.com/pkg/errors"
)

// GetCourse retrieves a single course resource
func (cfg *Config) GetCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, contextServiceName, datasetAPI)
	defer drainBody(ctx, r)

	vars := mux.Vars(r)
	institution := vars["institution"]
	courseID := vars["course"]
	mode := vars["mode"]
	logData := log.Data{"institution": institution, "course_id": courseID, "mode": mode}

	log.InfoCtx(ctx, "GetCourse handler: attempting to get course resource", logData)

	// Get course details from datastore
	course, err := cfg.DataStore.GetCourse(ctx, institution, courseID, mode)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "GetCourse handler: failed to get course resource"), logData)

		errorValues := make(map[string](string))
		errorValues["institution"] = institution
		errorValues["course"] = courseID
		errorValues["mode"] = mode

		datastoreError(ctx, w, errs.New(err, 0, errorValues))
		return
	}

	b, err := json.Marshal(course)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "GetCourse handler: failed to marshal course resource into bytes"), logData)
		http.Error(w, errs.ErrInternalServer.Error(), http.StatusInternalServerError)
		return
	}

	log.InfoCtx(ctx, "GetCourse handler: successfully got course resource", logData)
	writeBody(ctx, w, b)
}
