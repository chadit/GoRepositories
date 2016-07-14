package Mongo

import "testing"

func TestInsert(t *testing.T) {
	connectionInfo := new(ConnectionInfo)
	connectionInfo.InitDatabaseFromConnectionString(databaseConnectionString)
	defer connectionInfo.session.Close()

	newTenant := NewTenant("test")
	insertError := connectionInfo.Insert(sessionTestCollection, newTenant)
	if insertError != nil {
		t.Error(insertError)
	}

}
