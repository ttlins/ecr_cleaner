package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/mitchellh/go-homedir"
	"github.com/titolins/ecr_cleaner/cmd"
)

type loadErr struct {
	fileName string
	err      error
}

func mustGetEnvFiles() []string {
	home, err := homedir.Dir()
	if err != nil {
		log.Printf("Failed to locate home dir: %s", err)
	}

	return []string{
		".env",
		fmt.Sprintf("%s/.ecr_cleaner", home),
	}
}

func loadEnv() {
	envFiles := mustGetEnvFiles()

	var loadErrs []loadErr
	for _, f := range envFiles {
		if err := godotenv.Load(f); err != nil {
			loadErrs = append(loadErrs, loadErr{f, err})
		}
	}

	if len(loadErrs) == len(envFiles) {
		log.Println("Couldn't load any env files - this will fail if aws variables were not set manually")
		for _, loadErr := range loadErrs {
			log.Printf("Error loading .env file %q: %s", loadErr.fileName, loadErr.err.Error())
		}
	}

}

func main() {
	loadEnv()

	cmd.Execute()
}
