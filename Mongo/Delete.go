package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DeleteByQuery all existing document by find query
func (c ReposClient) DeleteByQuery(conn ConnectionInfo, findQuery bson.M) (*mgo.ChangeInfo, error) {
	var (
		col *mgo.Collection
		err error
	)

	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return nil, err
	}
	return col.RemoveAll(findQuery)
}

// DeleteByID all existing document by find query
func (c ReposClient) DeleteByID(conn ConnectionInfo, documentID string) error {
	var (
		col *mgo.Collection
		err error
	)

	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return err
	}
	return col.RemoveId(documentID)
}
