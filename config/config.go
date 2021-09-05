package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

type CompletedTask struct {
	index int
	task  string
	time  string //time at which it got complete
}

type CurrentTasks struct {
	index int
	task  string
	time  string //time at which it was created
}

type CancelledTask struct {
	index int
	task  string
	time  string //time at which it got cancelled
}

type Config struct {
	Project        string
	Currenttasks   []CurrentTasks
	Completedtasks []CompletedTask
	Cancelledtasks []CancelledTask
}

func LoadConfig() (*Config, error) {

	const localfile = "./config.cfg"
	var config Config

	if !fileExists(localfile) {
		fmt.Printf("Setting up config file! \n")
		config.Project = "TO setup!"
		configBytes, err := json.MarshalIndent(config, "", "   ")
		if err != nil {
			fmt.Print("Something Wrong happened", err)
			return nil, err
		}

		err = ioutil.WriteFile(localfile, configBytes, os.ModePerm)

		if err != nil {
			fmt.Print("Error writing device config!", err)
			return nil, err
		}
	} else {
		configFile, err := os.Open(localfile)
		if err != nil {
			log.Error("Error opening device file", err)
			return nil, err
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer configFile.Close()
		configBytes, _ := ioutil.ReadAll(configFile)
		err = json.Unmarshal(configBytes, &config)

		if err != nil {
			log.Error("11Error loading  config!", err)
			return nil, err
		}
	}

	return &config, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
