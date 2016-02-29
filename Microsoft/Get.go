package Microsoft

import (
	"database/sql"
	"log"
)

// GetByQuery will execute a query against the MsSQL database
func GetByQuery(connectionString string, query string) (*sql.Rows, error) {
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
		return nil, err
	}
	defer db.Close()
	return db.Query(query)
}
