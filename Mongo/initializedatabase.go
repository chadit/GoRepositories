package Mongo

// InitDatabaseFromConnectionString sets up the database connection
func (connection *ConnectionInfo) InitDatabaseFromConnectionString(connectionString string) {
	connection.InitSessionFromConnectionString(connectionString)
	connection.InitDatabaseFromSession(connection.databaseName)
}

// InitDatabaseFromConnection sets session informaiton from a connection string
func (connection *ConnectionInfo) InitDatabaseFromConnection(connectionString string, databaseName string) {
	connection.InitSessionFromConnectionString(connectionString)
	if connection.sessionError != nil {
		return
	}

	connection.db = connection.session.DB(databaseName)
}

// InitDatabaseFromSession sets session informaiton from a connection string
func (connection *ConnectionInfo) InitDatabaseFromSession(databaseName string) {
	connection.db = connection.session.DB(databaseName)
}
