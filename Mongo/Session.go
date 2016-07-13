package Mongo

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

// ConnectionInfo holds the database conneciton information
type ConnectionInfo struct {
	session      *mgo.Session
	sessionError error
	databaseName string
	db           *mgo.Database
	dbError      error
}

// InitCollectionFromDatabase - initialize collection from mgo database
func (connection *ConnectionInfo) InitCollectionFromDatabase(collectionName string) (*mgo.Collection, error) {
	return connection.db.C(collectionName), nil
}

// InitCollectionFromSession - initialize collection from mgo session
func (connection *ConnectionInfo) InitCollectionFromSession(databaseName string, collectionName string) (*mgo.Collection, error) {
	db := connection.session.DB(databaseName)
	connection.db = db
	return db.C(collectionName), nil
}

// InitCollectionFromConnectionString - initialize collection from a connection string and passin database
func (connection *ConnectionInfo) InitCollectionFromConnectionString(connectionString string, databaseName string, collectionName string) (*mgo.Collection, error) {
	return connection.initializeCollection(databaseName, connectionString, collectionName)
	// db, err := connection.InitDatabaseFromConnection(connectionString, databaseName)
	// if err != nil {
	// 	return new(mgo.Collection), err
	// }
	// return db.C(collectionName), nil
}

// InitCollectionAndDatabaseFromConnectionString - initialize collection from a connection string
func (connection *ConnectionInfo) InitCollectionAndDatabaseFromConnectionString(connectionString string, collectionName string) (*mgo.Collection, error) {
	if collectionName == "" {
		return nil, errors.New("collectionName cannot be empty")
	}

	if connectionString == "" {
		return nil, errors.New("connectionString cannot be empty")
	}
	return connection.initializeCollection(connection.databaseName, connectionString, collectionName)
	// if connection.databaseName == "" {
	// 	dialInformation, _, err := GetDialInformation(connectionString)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	connection.databaseName = dialInformation.Database
	// }
	//
	// connection.InitDatabaseFromConnection(connectionString, connection.databaseName)
	// if connection.sessionError != nil {
	// 	return nil, connection.sessionError
	// }
	//
	// if connection.dbError != nil {
	// 	return nil, connection.dbError
	// }
	// return connection.db.C(collectionName), nil
}

func (connection *ConnectionInfo) initializeCollection(databaseName, connectionString, collectionName string) (*mgo.Collection, error) {
	if databaseName == "" {
		dialInformation, _, err := GetDialInformation(connectionString)
		if err != nil {
			return nil, err
		}
		connection.databaseName = dialInformation.Database
	}

	connection.InitDatabaseFromConnection(connectionString, connection.databaseName)
	if connection.sessionError != nil {
		return nil, connection.sessionError
	}

	if connection.dbError != nil {
		return nil, connection.dbError
	}
	return connection.db.C(collectionName), nil
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

// InitSessionFromConnectionString get the session information for a conneciton
func (connection *ConnectionInfo) InitSessionFromConnectionString(connectionString string) {
	dialInformation, sessionMode, err := GetDialInformation(connectionString)
	if err != nil {
		connection.sessionError = err
		return
	}
	connection.databaseName = dialInformation.Database
	connection.InitSessionFromDialInfo(dialInformation, sessionMode)
}

// InitSessionFromDialInfo get the session information for a conneciton
func (connection *ConnectionInfo) InitSessionFromDialInfo(dialInfo *mgo.DialInfo, sessionMode mgo.Mode) {
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		connection.sessionError = err
		return
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(sessionMode, true)
	connection.session = session
	connection.sessionError = nil
}

// GetDialInformation get the dial information
func GetDialInformation(connectionString string) (*mgo.DialInfo, mgo.Mode, error) {
	info := &mgo.DialInfo{
		Addrs:    []string{},
		Timeout:  60 * time.Second,
		Database: "",
		Username: "",
		Password: "",
	}

	result := getConnectionStringItems(connectionString)
	//fmt.Println(len(result))
	if len(result) < 2 {
		return info, 2, errors.New("could not parse connecitonString")
	}

	host := strings.Split(result[0], ",")
	databaseName := result[1]
	info.Addrs = host
	info.Database = databaseName
	sessionMode := mgo.Primary
	// get options
	if len(result) == 3 {
		resultOptions := getOptionStringItems(result[2])
		for i := range resultOptions {
			switch {
			case strings.Contains(resultOptions[i], "replicaSet"):
				optionValue := getValueFromPairString(resultOptions[i])
				if optionValue != "" {
					info.ReplicaSetName = optionValue
				}
			case strings.Contains(resultOptions[i], "maxPoolSize"):
				optionValue := getValueFromPairInt(resultOptions[i])
				info.PoolLimit = optionValue
			case strings.Contains(resultOptions[i], "readPreference"):
				optionValue := getValueFromPairString(resultOptions[i])
				sessionMode = getReadPreference(optionValue)

			} // end switch

		} // end for loop
	}
	return info, sessionMode, nil
}

func getConnectionStringItems(connectionString string) []string {
	connectionString = strings.Replace(connectionString, "mongodb://", "", 1)
	return strings.Split(connectionString, "/")
}

func getOptionStringItems(optionString string) []string {
	optionString = strings.Replace(optionString, "?", "", 1)
	return strings.Split(optionString, ";")
}

func getValueFromPairString(optionString string) string {
	keyValue := strings.Split(optionString, "=")
	if len(keyValue) == 2 {
		return keyValue[1]
	}
	return ""
}

func getValueFromPairInt(optionString string) int {
	keyValue := strings.Split(optionString, "=")
	if len(keyValue) == 2 {
		value, err := strconv.Atoi(keyValue[1])
		if err != nil {
			return 100
		}

		return value
	}
	return 100
}

func getReadPreference(optionString string) mgo.Mode {
	// Primary            Mode = 2 // Default mode. All operations read from the current replica set primary.
	// PrimaryPreferred   Mode = 3 // Read from the primary if available. Read from the secondary otherwise.
	// Secondary          Mode = 4 // Read from one of the nearest secondary members of the replica set.
	// SecondaryPreferred Mode = 5 // Read from one of the nearest secondaries if available. Read from primary otherwise.
	// Nearest            Mode = 6 // Read from one of the nearest members, irrespective of it being primary or secondary.
	//
	// // Read preference modes are specific to mgo:
	// Eventual  Mode = 0 // Same as Nearest, but may change servers between reads.
	// Monotonic Mode = 1 // Same as SecondaryPreferred before first write. Same as Primary after first write.
	// Strong    Mode = 2 // Same as Primary.
	switch optionString {
	case "Primary":
		return mgo.Primary

	}
	return mgo.Primary
}
