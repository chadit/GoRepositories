package Microsoft

import (
	"database/sql"
	"log"
	"time"
	//_ "github.com/denisenkom/go-mssqldb"
)

// GetByQuery will execute a query against the MsSQL database
func GetByQuery(connectionString string, query string) (*sql.Rows, error) {
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
		return nil, err
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(20 * time.Second)
	defer db.Close()
	return db.Query(query)
}
