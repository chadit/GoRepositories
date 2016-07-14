package Mongo

import "testing"

var (
	databaseConnectionString = "mongodb://127.0.0.1/GoTest/?w=majority;journal=true;maxPoolSize=1000"
	sessionTestCollection    = "SessionTestings"
)

func TestCreateToken(t *testing.T) {
	connectionInfo := new(ConnectionInfo)
	collection, collectionError := connectionInfo.InitCollectionAndDatabaseFromConnectionString(databaseConnectionString, sessionTestCollection)
	if collectionError != nil {
		t.Error(collectionError)
	}
	if collection == nil {
		t.Error("failed to initialize collection")
	}

	if connectionInfo.SessionError != nil {
		t.Error(connectionInfo.SessionError)
	}

	if connectionInfo.Session == nil {
		t.Error("session info is not there")
	}

	if connectionInfo.DatabaseError != nil {
		t.Error(connectionInfo.DatabaseError)
	}

	if connectionInfo.Database == nil {
		t.Error("db info is not there")
	}

	connectionInfo.Session.Close()

}
