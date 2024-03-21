package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "strings"

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

    //-- creating the data_table_of_users if NOT EXISTS
    data, err := os.ReadFile("db/pq/user/CreateUser.sql")
    if err != nil {
        fmt.Println("Error reading sql file", err)
    }

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
    
    fmt.Println("Connected to the db.........")

    /**
        We use the Ping() method on the *sql.DB object to test the connection to the database server.
        If the connection is successful,
        Ping() returns nil, indicating no error.
    */
    if error := db.Ping(); error != nil {
        db.Close()
        log.Fatal(error)
    }

    //-- executing the sql statements
    tx, err := db.Begin()
    if err != nil {
        fmt.Println("Error starting transaction", err)
        return nil, err
    }

    sqlStatements := string(data)
    for _, stmt := range strings.Split(sqlStatements, ";") {
        _, err := tx.Exec(stmt)
        if err != nil {
            fmt.Println("Error executing statement: ", err)
            tx.Rollback()
            return nil, err
        }
    }
	if err := tx.Commit(); err != nil {
        fmt.Println("Error committing transaction: ", err)
        return nil, err
    }

    //-- Print inserted data fromt the table &&
    //-- specific query to retrieve inserted data
    rows, err := db.Query("SELECT * FROM root.users;")
    if err != nil {
        fmt.Println("Error fetching inserted data:", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name, address, email string

        if err := rows.Scan(&id, &name, &address, &email); err != nil {
            fmt.Println("Error scanning data:", err)
            continue
        }

        fmt.Printf("ID: %d, Name: %s, Address: %s, Email: %s, \n", id, name, address, email)
    }

    fmt.Println("User data created and retrieved successfully!")

    return db, nil
}