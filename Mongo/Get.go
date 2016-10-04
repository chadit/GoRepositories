package Mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Find - return all documents that match the query
func (c ReposClient) Find(conn ConnectionInfo, query bson.M, queryOptions QueryOptions) (*mgo.Query, error) {
	var (
		col *mgo.Collection
		err error
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return nil, err
	}
	if queryOptions.Sort == "" {
		return col.Find(query).Skip(queryOptions.Skip).Limit(queryOptions.Limit).Select(queryOptions.Projection), nil
	}

	return col.Find(query).Sort(queryOptions.Sort).Skip(queryOptions.Skip).Limit(queryOptions.Limit).Select(queryOptions.Projection), nil
}

// FindByID returns the document by it's bson id
func (c ReposClient) FindByID(conn ConnectionInfo, documentID string) (*mgo.Query, error) {
	var (
		col *mgo.Collection
		err error
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return nil, err
	}
	return col.FindId(documentID), nil
}

// FindOne returns the first document found
func (c ReposClient) FindOne(conn ConnectionInfo, query bson.M) (*mgo.Query, error) {
	var (
		col *mgo.Collection
		err error
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return nil, err
	}
	return col.Find(query), nil
}

// Count returns the count from a query
func (c ReposClient) Count(conn ConnectionInfo, query bson.M) (int, error) {
	var (
		col *mgo.Collection
		err error
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return 0, err
	}
	return col.Find(query).Count()
}

// Distinct returns list of unique results
func (c ReposClient) Distinct(conn ConnectionInfo, query bson.M, propertyName string, results interface{}) (interface{}, error) {
	var (
		col *mgo.Collection
		err error
	)
	if col, err = conn.InitCollectionFromDatabase(conn.CollectionName); err != nil {
		return nil, err
	}
	err = col.Find(query).Distinct(propertyName, &results)
	return results, err
}
