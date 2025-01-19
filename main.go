package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// loose coordinates of the science museum of minnesota
// 44.94, -93.10
func main() {
	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=44.94&longitude=-93.10&hourly=temperature_2m&models=gfs_seamless")
	Check(err)
	defer resp.Body.Close()
	var val map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&val)
	Check(err)
	fmt.Println(val["hourly_units"])
	// Open the database connection
	// db, err := sql.Open("sqlite3", "./example.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// fmt.Println("Database connection opened successfully.")
	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS users (
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		name TEXT NOT NULL,
	// 		age INTEGER
	// 	)
	// `)
	// _, err = db.Exec(`
	// 	DROP TABLE IF EXISTS users
	// `)
	// if err != nil {
	// 	log.Fatalf("Error deleting table: %v\n", err)
	// }
	// fmt.Println("Table deleted successfully.")

	// // Insert data
	// _, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "bob", 100)
	// if err != nil {
	// 	log.Fatalf("Error inserting data: %v\n", err)
	// }

	// rows, err := db.Query("SELECT id, name, age FROM users")
	// if err != nil {
	// 	log.Fatalf("Error querying data: %v\n", err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	var age int
	// 	err := rows.Scan(&id, &name, &age)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	// }
}
