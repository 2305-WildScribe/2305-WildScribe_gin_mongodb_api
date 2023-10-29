package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
			log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGOURI")
}

func DatabaseEnvironment() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if os.Getenv("ENV") == "test" {
		return "test"
	}	else if os.Getenv("PROD_ENV") == "production"{
		return "production"
	}
	return "test"
}
