package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DeleteByQuery all existing document by find query
func DeleteByQuery(collection *mgo.Collection, findQuery bson.M) (*mgo.ChangeInfo, error) {
	return collection.RemoveAll(findQuery)
}

// DeleteByID all existing document by find query
func DeleteByID(collection *mgo.Collection, documentID string) error {
	return collection.RemoveId(documentID)
}
