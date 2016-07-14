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

	if connectionInfo.sessionError != nil {
		t.Error(connectionInfo.sessionError)
	}

	if connectionInfo.session == nil {
		t.Error("session info is not there")
	}

	if connectionInfo.dbError != nil {
		t.Error(connectionInfo.dbError)
	}

	if connectionInfo.db == nil {
		t.Error("db info is not there")
	}

	connectionInfo.session.Close()

}
