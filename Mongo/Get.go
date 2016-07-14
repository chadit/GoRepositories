package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Find - return all documents that match the query
func (connection *ConnectionInfo) Find(collectionName string, query bson.M, queryOptions QueryOptions) (*mgo.Query, error) {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return nil, collectionError
	}
	if queryOptions.Sort == "" {
		return collection.Find(query).Skip(queryOptions.Skip).Limit(queryOptions.Limit).Select(queryOptions.Projection), nil
	}

	return collection.Find(query).Sort(queryOptions.Sort).Skip(queryOptions.Skip).Limit(queryOptions.Limit).Select(queryOptions.Projection), nil
}

// FindByID returns the document by it's bson id
func (connection *ConnectionInfo) FindByID(collectionName, documentID string) (*mgo.Query, error) {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return nil, collectionError
	}
	return collection.FindId(documentID), nil
}

// FindOne returns the first document found
func (connection *ConnectionInfo) FindOne(collectionName string, query bson.M) (*mgo.Query, error) {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return nil, collectionError
	}
	return collection.Find(query), nil
}

// Count returns the count from a query
func (connection *ConnectionInfo) Count(collectionName string, query bson.M) (int, error) {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return 0, collectionError
	}
	return collection.Find(query).Count()
}

// Distinct returns list of unique results
func (connection *ConnectionInfo) Distinct(collectionName string, query bson.M, propertyName string, results interface{}) (interface{}, error) {
	collection, collectionError := connection.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		return nil, collectionError
	}
	err := collection.Find(query).Distinct(propertyName, &results)
	return results, err
}
