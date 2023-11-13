package utils

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file
func LoadEnv() {
	projectName := regexp.MustCompile(`^(.` + "*" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	err := godotenv.Load(string(rootPath) + `/.env`)

	requiredVars := []string{}

	missingVars := []string{}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			missingVars = append(missingVars, v)
		}
	}

	if len(missingVars) > 0 {
		log.Default().Println("Missing required environment variables: " + strings.Join(missingVars, ", "))
		return
	}

	if err != nil {
		log.Default().Println(".env file not found; using env vars if found")
		return
	}
}
