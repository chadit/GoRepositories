package Microsoft

import (
	"database/sql"

	"github.com/corneldamian/golog"
)

// PingMsServer uses the connection string in the config file to verify it is available
func PingMsServer(connectionString string) (bool, error) {
	log := golog.GetLogger("log")
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Error("Open connection failed:", err.Error())
		return false, err
	}
	defer db.Close()

	// Is the database running?
	pingError := db.Ping()
	if err != nil {
		log.Error("MicrosoftSql Ping : " + pingError.Error() + " connectionstring=\"" + connectionString + "\"")
		return false, err
	}
	return true, nil
}

// DoesDatabaseExist checks if the database "strDBName" exists on the MSSQL database engine.
func DoesDatabaseExist(connectionString string, strDBName string) (bool, error) {
	log := golog.GetLogger("log")
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Error("Open connection failed:", err.Error())
		return false, err
	}
	defer db.Close()

	return checkDB(db, strDBName)
}

// CheckDB checks if the database "strDBName" exists on the MSSQL database engine.
func checkDB(db *sql.DB, strDBName string) (bool, error) {

	// Does the database exist?
	result, err := db.Query("SELECT db_id('" + strDBName + "')")
	defer result.Close()
	if err != nil {
		return false, err
	}

	for result.Next() {
		var s sql.NullString
		err := result.Scan(&s)
		if err != nil {
			return false, err
		}

		// Check result
		if s.Valid {
			return true, nil
		}
		return false, nil
	}

	// This return() should never be hit...
	return false, err
}