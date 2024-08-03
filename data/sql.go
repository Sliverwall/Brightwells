package data

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func SQL_query(query string) [][]interface{} {
	// Connect to database
	db, err := sql.Open("sqlite", "data/brightwells.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	// Query the "place_entity" table
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error querying table:", err)
		return nil
	}
	defer rows.Close()

	// Get column names to handle the columns dynamically
	columns, err := rows.Columns()
	if err != nil {
		log.Println("Error getting columns:", err)
		return nil
	}

	// Create a list to hold all entries
	allEntries := make([][]interface{}, 0)

	// Iterate through the rows
	for rows.Next() {
		// Create a slice to hold each row's values
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan the row into the value pointers
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil
		}

		// Append the row to the list of all entries
		allEntries = append(allEntries, values)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil
	}

	// Return final list
	return allEntries
}

// SQL_exec executes a non-select SQL query such as INSERT, UPDATE, or DELETE.
func SQL_exec(query string) (sql.Result, error) {
	// Connect to the database
	db, err := sql.Open("sqlite", "data/brightwells.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	// Execute the query
	result, err := db.Exec(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}

	return result, nil
}
