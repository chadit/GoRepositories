package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Insert a new document
func Insert(collection *mgo.Collection, item interface{}) error {
	return collection.Insert(item)
}

// Update update an existing document, similar to find and replace document
func Update(collection *mgo.Collection, findQuery bson.M, item interface{}) error {
	return collection.Update(findQuery, item)
}

// UpdateByQuery updates an existing document by an update query
func UpdateByQuery(collection *mgo.Collection, findQuery bson.M, updateQuery bson.M) (*mgo.ChangeInfo, error) {
	return collection.UpdateAll(findQuery, updateQuery)
}

// UpdateByID updates an existing document with a matching BSON ID
func UpdateByID(collection *mgo.Collection, documentID string, item interface{}) error {
	return collection.UpdateId(documentID, item)
}

// UpdateByIDByQuery updates an existing document with a query with a matching BSON ID
func UpdateByIDByQuery(collection *mgo.Collection, documentID string, updateQuery bson.M) error {
	return collection.UpdateId(documentID, updateQuery)
}

// UpdateOneAndReturn updates and existing document and returns those documents
func UpdateOneAndReturn(collection *mgo.Collection, findQuery bson.M, updateQuery bson.M) (bson.M, *mgo.ChangeInfo, error) {
	var m bson.M
	change := mgo.Change{
		Update:    updateQuery,
		ReturnNew: true,
	}
	info, err := collection.Find(findQuery).Apply(change, &m)
	return m, info, err
}

// UpdateByIDAndReturn updates and existing documents and returns those documents with a matching BSON ID
func UpdateByIDAndReturn(collection *mgo.Collection, documentID string, updateQuery bson.M) (bson.M, *mgo.ChangeInfo, error) {
	var m bson.M
	change := mgo.Change{
		Update:    updateQuery,
		ReturnNew: true,
	}
	info, err := collection.FindId(documentID).Apply(change, &m)
	return m, info, err
}
