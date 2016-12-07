package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Client is an interface of the functions
type Client interface {
	// Get.go
	Find(conn ConnectionInfo, query bson.M, queryOptions QueryOptions) (*mgo.Query, error)
	FindByID(conn ConnectionInfo, documentID string) (*mgo.Query, error)
	FindOne(conn ConnectionInfo, query bson.M) (*mgo.Query, error)
	Count(conn ConnectionInfo, query bson.M) (int, error)
	Distinct(conn ConnectionInfo, query bson.M, propertyName string, results interface{}) (interface{}, error)

	// Modify.go
	Insert(conn ConnectionInfo, item interface{}) error
	Update(conn ConnectionInfo, findQuery bson.M, item interface{}) error
	UpdateByQuery(conn ConnectionInfo, findQuery bson.M, updateQuery bson.M) (*mgo.ChangeInfo, error)
	UpdateByID(conn ConnectionInfo, documentID string, item interface{}) error
	UpdateByIDByQuery(conn ConnectionInfo, documentID string, updateQuery bson.M) error
	UpdateOneAndReturn(conn ConnectionInfo, findQuery bson.M, updateQuery bson.M) (bson.M, *mgo.ChangeInfo, error)
	UpdateByIDAndReturn(conn ConnectionInfo, documentID string, updateQuery bson.M) (bson.M, *mgo.ChangeInfo, error)

	// Delete.go
	DeleteByQuery(conn ConnectionInfo, findQuery bson.M) (*mgo.ChangeInfo, error)
	DeleteByID(conn ConnectionInfo, documentID string) error
}

// ReposClient passed in struct
type ReposClient struct{}

// ConnectionInfo holds the database conneciton information
type ConnectionInfo struct {
	CollectionName string
	Session        *mgo.Session
	SessionError   error
	DatabaseName   string
	Database       *mgo.Database
	DatabaseError  error
}

// App defines the application container
type App struct {
	Client Client
	Conn   ConnectionInfo
}

// currently MGO does not have a good way to see if a session is still active, this is really the only option
func (s *ConnectionInfo) SessionActive() (active bool) {
	defer func() {
		if ierr := recover(); ierr != nil {
			active = false
			return
		}
	}()

	//Ping panics if session is closed. (see mgo.Session.Panic())
	if err := s.Session.Ping(); err != nil {
		active = false
		return
	}
	return true
}

func (s *ConnectionInfo) Close() {
	if s.Session != nil {
		s.Session.Close()
	}

	if s.Database != nil && s.Database.Session != nil {
		s.Database.Session.Close()
	}
}
