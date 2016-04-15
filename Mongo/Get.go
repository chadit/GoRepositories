package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// // Find - return all documents that match the query
// func Find(collection *mgo.Collection, query bson.M, queryOptions QueryOptions) ([]bson.M, error) {
// 	var m []bson.M
// 	if queryOptions.Sort == "" {
// 		err := collection.Find(query).Skip(queryOptions.Skip).Limit(queryOptions.Limit).Select(queryOptions.Projection).All(&m)
// 		return m, err
// 	}
//
// 	err := collection.Find(query).Sort(queryOptions.Sort).Skip(queryOptions.Skip).Limit(queryOptions.Limit).Select(queryOptions.Projection).All(&m)
// 	return m, err
// }

// Find - return all documents that match the query
func Find(collection *mgo.Collection, query bson.M, queryOptions QueryOptions) *mgo.Query {
	//	var m []bson.M
	if queryOptions.Sort == "" {
		return collection.Find(query).Skip(queryOptions.Skip).Limit(queryOptions.Limit).Select(queryOptions.Projection)
		//	return m, err
	}

	return collection.Find(query).Sort(queryOptions.Sort).Skip(queryOptions.Skip).Limit(queryOptions.Limit).Select(queryOptions.Projection)
	//	return m, err
}

// // FindByID returns the document by it's bson id
// func FindByID(collection *mgo.Collection, documentID string) (bson.M, error) {
// 	var m bson.M
// 	err := collection.FindId(documentID).One(&m)
// 	return m, err
// }

// FindByID returns the document by it's bson id
func FindByID(collection *mgo.Collection, documentID string) *mgo.Query {
	//	var m bson.M
	return collection.FindId(documentID)
	//	return m, err
}

// // FindOne returns the first document found
// func FindOne(collection *mgo.Collection, query bson.M) (bson.M, error) {
// 	var m bson.M
// 	err := collection.Find(query).One(&m)
// 	return m, err
// }

// FindOne returns the first document found
func FindOne(collection *mgo.Collection, query bson.M) *mgo.Query {
	//	var m bson.M
	return collection.Find(query)
	//return m, err
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
