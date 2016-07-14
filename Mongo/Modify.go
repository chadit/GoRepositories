package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Insert a new document
func (connection *ConnectionInfo) Insert(collectionName string, item interface{}) error {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return collectionError
	}
	return collection.Insert(item)
}

// Update update an existing document, similar to find and replace document
func (connection *ConnectionInfo) Update(collectionName string, findQuery bson.M, item interface{}) error {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return collectionError
	}
	return collection.Update(findQuery, item)
}

// UpdateByQuery updates an existing document by an update query
func (connection *ConnectionInfo) UpdateByQuery(collectionName string, findQuery bson.M, updateQuery bson.M) (*mgo.ChangeInfo, error) {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return nil, collectionError
	}
	return collection.UpdateAll(findQuery, updateQuery)
}

// UpdateByID updates an existing document with a matching BSON ID
func (connection *ConnectionInfo) UpdateByID(collectionName string, documentID string, item interface{}) error {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return collectionError
	}
	return collection.UpdateId(documentID, item)
}

// UpdateByIDByQuery updates an existing document with a query with a matching BSON ID
func (connection *ConnectionInfo) UpdateByIDByQuery(collectionName string, documentID string, updateQuery bson.M) error {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return collectionError
	}
	return collection.UpdateId(documentID, updateQuery)
}

// UpdateOneAndReturn updates and existing document and returns those documents
func (connection *ConnectionInfo) UpdateOneAndReturn(collectionName string, findQuery bson.M, updateQuery bson.M) (bson.M, *mgo.ChangeInfo, error) {
	var m bson.M
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return m, nil, collectionError
	}
	change := mgo.Change{
		Update:    updateQuery,
		ReturnNew: true,
	}
	info, err := collection.Find(findQuery).Apply(change, &m)
	return m, info, err
}

// UpdateByIDAndReturn updates and existing documents and returns those documents with a matching BSON ID
func (connection *ConnectionInfo) UpdateByIDAndReturn(collectionName string, documentID string, updateQuery bson.M) (bson.M, *mgo.ChangeInfo, error) {
	var m bson.M
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return m, nil, collectionError
	}
	change := mgo.Change{
		Update:    updateQuery,
		ReturnNew: true,
	}
	info, err := collection.FindId(documentID).Apply(change, &m)
	return m, info, err
}
