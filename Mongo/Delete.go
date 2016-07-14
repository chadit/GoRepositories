package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DeleteByQuery all existing document by find query
func (connection *ConnectionInfo) DeleteByQuery(collectionName string, findQuery bson.M) (*mgo.ChangeInfo, error) {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return nil, collectionError
	}
	return collection.RemoveAll(findQuery)
}

// DeleteByID all existing document by find query
func (connection *ConnectionInfo) DeleteByID(collectionName, documentID string) error {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return collectionError
	}
	return collection.RemoveId(documentID)
}
