package initializers

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	PORT string
	USER string
	PASS string
	NAME string
}

func ConDb() (*sql.DB, error) {
	host := "localhost"
	port := 5432
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	/**
	  defining the connection string &&
	  connecting to the database
	*/
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error creating database connection: %w", err)
	}

	/**
	  We use the Ping() method on the *sql.DB object to test the connection to the database server.
	  If the connection is successful,
	  Ping() returns nil, indicating no error.
	*/
	if error := db.Ping(); error != nil {
		db.Close()
		log.Fatal(error)
	}
	return db, nil
}

func CreateOrgTable() error {
	db, err := ConDb()
	if err != nil {
		return err
	}
	defer db.Close()

	//-- open the SQL file
	file, err := os.Open("db/pq/org/CreateOrgTable.sql")
	if err != nil {
		return fmt.Errorf("error opening the sql file: %w", err)
	}
	defer file.Close()

	//-- read the file contents from the given file
	cnt, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading the sql file: %w", err)
	}

	//-- preapare the query for execution && execute the query
	stmt, err := db.Prepare(string(cnt))
	if err != nil {
		return fmt.Errorf("error preaparing the query: %w", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("error executing the query: %w", err)
	}

	return nil
}

func CreateUserTable() error {
	db, err := ConDb()
	if err != nil {
		return err
	}
	defer db.Close()

	//-- open the SQL file
	file, err := os.Open("db/pq/user/CreateUserTable.sql")
	if err != nil {
		return fmt.Errorf("error opening the sql file: %w", err)
	}
	defer file.Close()

	//-- read the file contents from the given file
	cnt, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading the sql file: %w", err)
	}

	//-- preapare the query for execution && execute the query
	stmt, err := db.Prepare(string(cnt))
	if err != nil {
		return fmt.Errorf("error preaparing the query: %w", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("error executing the query: %w", err)
	}

	return nil
}