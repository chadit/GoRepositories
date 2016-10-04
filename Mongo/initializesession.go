package Mongo

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

// InitSessionFromConnectionString get the session information for a conneciton
func (connection *ConnectionInfo) InitSessionFromConnectionString(connectionString string) {
	dialInformation, sessionMode, err := GetDialInformation(connectionString)
	if err != nil {
		connection.SessionError = err
		return
	}
	connection.DatabaseName = dialInformation.Database
	connection.InitSessionFromDialInfo(dialInformation, sessionMode)
}

// InitSessionFromDialInfo get the session information for a conneciton
func (connection *ConnectionInfo) InitSessionFromDialInfo(dialInfo *mgo.DialInfo, sessionMode mgo.Mode) {
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		connection.SessionError = err
		return
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(sessionMode, true)
	connection.Session = session
	connection.SessionError = nil
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
