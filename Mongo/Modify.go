package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Insert a new document
func (c ReposClient) Insert(conn ConnectionInfo, item interface{}) error {
	var (
		col *mgo.Collection
		err error
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return err
	}
	return col.Insert(item)
}

// Update update an existing document, similar to find and replace document
func (c ReposClient) Update(conn ConnectionInfo, findQuery bson.M, item interface{}) error {
	var (
		col *mgo.Collection
		err error
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return err
	}
	return col.Update(findQuery, item)
}

// UpdateByQuery updates an existing document by an update query
func (c ReposClient) UpdateByQuery(conn ConnectionInfo, findQuery bson.M, updateQuery bson.M) (*mgo.ChangeInfo, error) {
	var (
		col *mgo.Collection
		err error
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return nil, err
	}
	return col.UpdateAll(findQuery, updateQuery)
}

// UpdateByID updates an existing document with a matching BSON ID
func (c ReposClient) UpdateByID(conn ConnectionInfo, documentID string, item interface{}) error {
	var (
		col *mgo.Collection
		err error
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return err
	}
	return col.UpdateId(documentID, item)
}

// UpdateByIDByQuery updates an existing document with a query with a matching BSON ID
func (c ReposClient) UpdateByIDByQuery(conn ConnectionInfo, documentID string, updateQuery bson.M) error {
	var (
		col *mgo.Collection
		err error
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return err
	}
	return col.UpdateId(documentID, updateQuery)
}

// UpdateOneAndReturn updates and existing document and returns those documents
func (c ReposClient) UpdateOneAndReturn(conn ConnectionInfo, findQuery bson.M, updateQuery bson.M) (bson.M, *mgo.ChangeInfo, error) {
	var (
		col *mgo.Collection
		err error
		m   bson.M
	)

	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return m, nil, err
	}
	change := mgo.Change{
		Update:    updateQuery,
		ReturnNew: true,
	}
	info, err := col.Find(findQuery).Apply(change, &m)
	return m, info, err
}

// UpdateByIDAndReturn updates and existing documents and returns those documents with a matching BSON ID
func (c ReposClient) UpdateByIDAndReturn(conn ConnectionInfo, documentID string, updateQuery bson.M) (bson.M, *mgo.ChangeInfo, error) {
	var (
		col *mgo.Collection
		err error
		m   bson.M
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return m, nil, err
	}
	change := mgo.Change{
		Update:    updateQuery,
		ReturnNew: true,
	}
	info, err := col.FindId(documentID).Apply(change, &m)
	return m, info, err
}
