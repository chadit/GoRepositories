package Mongo

import (
	"errors"

	"gopkg.in/mgo.v2"
)

// InitCollectionFromDatabase - initialize collection from mgo database
func (connection *ConnectionInfo) InitCollectionFromDatabase(collectionName string) (*mgo.Collection, error) {
	return connection.Database.C(collectionName), nil
}

// InitCollectionFromSession - initialize collection from mgo session
func (connection *ConnectionInfo) InitCollectionFromSession(databaseName string, collectionName string) (*mgo.Collection, error) {
	connection.Database = connection.Session.DB(databaseName)
	return connection.Database.C(collectionName), nil
}

// InitCollectionFromConnectionString - initialize collection from a connection string and passin database
func (connection *ConnectionInfo) InitCollectionFromConnectionString(connectionString string, databaseName string, collectionName string) (*mgo.Collection, error) {
	return connection.initializeCollection(databaseName, connectionString, collectionName)
	// db, err := connection.InitDatabaseFromConnection(connectionString, databaseName)
	// if err != nil {
	// 	return new(mgo.Collection), err
	// }
	// return db.C(collectionName), nil
}

// InitCollectionAndDatabaseFromConnectionString - initialize collection from a connection string
func (connection *ConnectionInfo) InitCollectionAndDatabaseFromConnectionString(connectionString string, collectionName string) (*mgo.Collection, error) {
	if collectionName == "" {
		return nil, errors.New("collectionName cannot be empty")
	}

	if connectionString == "" {
		return nil, errors.New("connectionString cannot be empty")
	}
	return connection.initializeCollection(connection.DatabaseName, connectionString, collectionName)
}

func (connection *ConnectionInfo) initializeCollection(databaseName, connectionString, collectionName string) (*mgo.Collection, error) {
	if databaseName == "" {
		dialInformation, _, err := GetDialInformation(connectionString)
		if err != nil {
			return nil, err
		}
		connection.DatabaseName = dialInformation.Database
	}

	connection.InitDatabaseFromConnection(connectionString, connection.DatabaseName)
	if connection.SessionError != nil {
		return nil, connection.SessionError
	}

	if connection.DatabaseError != nil {
		return nil, connection.DatabaseError
	}
	return connection.Database.C(collectionName), nil
}
