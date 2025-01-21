package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type WeatherData struct {
	Elevation float32 `json:"elevation"`
	GenerationTime_ms float64 `json:"generationtime_ms"`
	Hourly struct {
		Temperature_2m []float32 `json:"temperature_2m"`
		Time []string `json:"time"`
	} `json:"hourly"`
	HourlyUnits struct {
		Temperature_2m string `json:"temperature_2m"`
		Time string `json:"time"`
	} `json:"hourly_units"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone string `json:"timezone"`
	TimezoneAbbreviation string `json:"timezone_abbreviation"`
	UtcOffsetSeconds int `json:"utc_offset_seconds"`
}

// loose coordinates of the science museum of minnesota
// 44.94, -93.10
// make a type for this response
func getWeatherData() (weatherData WeatherData) {
	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=44.94&longitude=-93.10&hourly=temperature_2m&models=gfs_seamless")
	Check(err)
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	Check(err)
	return
}

// need to break this into more functions
func main() {
	weatherData := getWeatherData()
	
	// Open the database connection
	db, err := sql.Open("sqlite3", "./service1.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Database connection opened successfully.")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS weather_data (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			elevation REAL,
			generation_time_ms REAL,
			temperature_2m REAL,
			time TEXT,
			latitude REAL,
			longitude REAL,
			timezone TEXT,
			timezone_abbreviation TEXT,
			utc_offset_seconds INTEGER
		)
	`)
	Check(err)

	// delete table
	// _, err = db.Exec(`
	// 	DROP TABLE IF EXISTS weather_data
	// `)
	// if err != nil {
	// 	log.Fatalf("Error deleting table: %v\n", err)
	// }
	// fmt.Println("Table deleted successfully.")

	// this doesnt feel right 🤔
	// insert data 
	_, err = db.Exec("INSERT INTO weather_data (elevation, generation_time_ms, temperature_2m, time, latitude, longitude, timezone, timezone_abbreviation, utc_offset_seconds) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", 
	weatherData.Elevation, 
	weatherData.GenerationTime_ms, 
	weatherData.Hourly.Temperature_2m[0], 
	weatherData.Hourly.Time[0], 
	weatherData.Latitude, 
	weatherData.Longitude, 
	weatherData.Timezone, 
	weatherData.TimezoneAbbreviation, 
	weatherData.UtcOffsetSeconds)
	if err != nil {
		log.Fatalf("Error inserting data: %v\n", err)
	}

	// query rows from table
	rows, err := db.Query("SELECT * FROM weather_data")
	if err != nil {
		log.Fatalf("Error querying data: %v\n", err)
	}
	defer rows.Close()

	// read and print rows from table to terminal
	for rows.Next() {
		var id int // sqlite automatically adds an ID column
		var elevation float32
		var generationTimeMs float64
		var temperature2m float32
		var time string
		var latitude float64
		var longitude float64
		var timezone string
		var timezoneAbbreviation string
		var utcOffsetSeconds int
		err := rows.Scan(&id, &elevation, &generationTimeMs, &temperature2m, &time, &latitude, &longitude, &timezone, &timezoneAbbreviation, &utcOffsetSeconds)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Elevation: %v, Generation Time (ms): %v, Temperature (2m): %v, Time: %s, Latitude: %v, Longitude: %v, Timezone: %s, Timezone Abbreviation: %s, UTC Offset Seconds: %d\n", 
			id, elevation, generationTimeMs, temperature2m, time, latitude, longitude, timezone, timezoneAbbreviation, utcOffsetSeconds)
	}
}
