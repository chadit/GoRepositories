package Mongo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestInsertAndDelete(t *testing.T) {
	connectionInfo := new(ConnectionInfo)
	collectionName := sessionTestCollection + GetNewBsonIDString()
	connectionInfo.InitDatabaseFromConnectionString(databaseConnectionString)
	defer connectionInfo.session.Close()

	newTenant := NewTenant("test")
	insertError := connectionInfo.Insert(collectionName, newTenant)
	if insertError != nil {
		t.Error(insertError)
	}
	var m bson.M
	changeInfo, deleteError := connectionInfo.DeleteByQuery(collectionName, m)
	if deleteError != nil {
		t.Error(deleteError)
	}

	if changeInfo.Removed != 1 {
		t.Error("failed to delete records")
	}

	collection, collectionError := connectionInfo.InitCollectionFromDatabase(collectionName)
	if collectionError != nil {
		t.Error(collectionError)
	}

	collection.DropCollection()

}
