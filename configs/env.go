package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    if os.Getenv("PROD_ENV") == "production" {
				return os.Getenv("MONGOURI")
    }	else	{
			err := godotenv.Load()
			if err != nil {
					log.Fatal("Error loading .env file")
			}
			return os.Getenv("MONGOURI")
		}
}
