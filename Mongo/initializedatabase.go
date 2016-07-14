package Mongo

// InitDatabaseFromConnectionString sets up the database connection
func (connection *ConnectionInfo) InitDatabaseFromConnectionString(connectionString string) {
	connection.InitSessionFromConnectionString(connectionString)
	connection.InitDatabaseFromSession(connection.DatabaseName)
}

// InitDatabaseFromConnection sets session informaiton from a connection string
func (connection *ConnectionInfo) InitDatabaseFromConnection(connectionString string, databaseName string) {
	connection.InitSessionFromConnectionString(connectionString)
	if connection.SessionError != nil {
		return
	}

	connection.Database = connection.Session.DB(databaseName)
}

// InitDatabaseFromSession sets session informaiton from a connection string
func (connection *ConnectionInfo) InitDatabaseFromSession(databaseName string) {
	connection.Database = connection.Session.DB(databaseName)
}
