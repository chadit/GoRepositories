package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Find - return all documents that match the query
func Find(collection *mgo.Collection, query bson.M, queryOptions QueryOptions) *mgo.Query {
	if queryOptions.Sort == "" {
		return collection.Find(query).Skip(queryOptions.Skip).Limit(queryOptions.Limit).Select(queryOptions.Projection)
	}

	return collection.Find(query).Sort(queryOptions.Sort).Skip(queryOptions.Skip).Limit(queryOptions.Limit).Select(queryOptions.Projection)
}

// FindByID returns the document by it's bson id
func FindByID(collection *mgo.Collection, documentID string) *mgo.Query {
	return collection.FindId(documentID)
}

// FindOne returns the first document found
func FindOne(collection *mgo.Collection, query bson.M) *mgo.Query {
	return collection.Find(query)
}

// Count returns the count from a query
func Count(collection *mgo.Collection, query bson.M) (int, error) {
	return collection.Find(query).Count()
}

// Distinct returns list of unique results
func Distinct(collection *mgo.Collection, query bson.M, propertyName string, results interface{}) (interface{}, error) {
	err := collection.Find(query).Distinct(propertyName, &results)
	return results, err
}
