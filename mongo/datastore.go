package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	errs "github.com/ofs/alpha-dataset-api/apierrors"
	"github.com/ofs/alpha-dataset-api/models"
)

// Mongo represents a simplistic MongoDB configuration.
type Mongo struct {
	Collection     string
	Database       string
	lastPingTime   time.Time
	lastPingResult error
	Session        *mgo.Session
	URI            string
}

// Init creates a new mgo.Session with a strong consistency and a write mode of "majority".
func (m *Mongo) Init() (session *mgo.Session, err error) {
	if session != nil {
		return nil, errors.New("session already exists")
	}

	if session, err = mgo.Dial(m.URI); err != nil {
		return nil, err
	}

	session.EnsureSafe(&mgo.Safe{WMode: "majority"})
	session.SetMode(mgo.Strong, true)
	return session, nil
}

// GetCourse retrieves a single course resource from datastore
func (m *Mongo) GetCourse(ctx context.Context, institutionID, courseID, mode string) (*models.Course, error) {
	s := m.Session.Copy()
	defer s.Close()

	selector := bson.M{
		"institution.public_ukprn": institutionID,
		"kis_course_id":            courseID,
		"mode.code":                mode,
	}

	var course models.Course
	err := s.DB(m.Database).C(m.Collection).Find(selector).One(&course)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, errs.ErrCourseNotFound
		}
		return nil, err
	}

	return &course, nil
}
