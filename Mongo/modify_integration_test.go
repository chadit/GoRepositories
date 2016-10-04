package Mongo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestInsertAndDelete(t *testing.T) {
	c := ReposClient{}
	conn := ConnectionInfo{}
	conn.CollectionName = sessionTestCollection + GetNewBsonIDString()
	conn.InitDatabaseFromConnectionString(databaseConnectionString)
	defer conn.Session.Close()

	newTenant := NewTenant("test")
	insertError := c.Insert(conn, newTenant)
	if insertError != nil {
		t.Error(insertError)
	}
	var m bson.M
	changeInfo, deleteError := c.DeleteByQuery(conn, m)
	if deleteError != nil {
		t.Error(deleteError)
	}

	if changeInfo.Removed != 1 {
		t.Error("failed to delete records")
	}

	collection, collectionError := conn.InitCollectionFromDatabase(conn.CollectionName)
	if collectionError != nil {
		t.Error(collectionError)
	}

	collection.DropCollection()

}
