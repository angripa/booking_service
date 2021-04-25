package models

type Config struct {
	Port       int    `json:"port"`
	DBHost     string `json:"dbHost"`
	DBPort     int    `json:"dbPort"`
	DBName     string `json:"dbName"`
	DBDriver   string `json:"dbDriver"`
	DBUsername string `json:"dbUsername"`
	DBPassword string `json:"dbPassword"`
}
