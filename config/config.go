package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type CompletedTask struct {
	Index int
	Task  string
	Time  string //time at which it got complete
}

type CurrentTasks struct {
	Index int
	Task  string
	Time  string
}

type CancelledTask struct {
	Index int
	Task  string
	Time  string //time at which it got cancelled
}

type Config struct {
	Project        string
	Currenttasks   []CurrentTasks
	Completedtasks []CompletedTask
	Cancelledtasks []CancelledTask
}

func (c *Config) AddToCurrentTasks(task string) {

	var index int

	switch len(c.Currenttasks) {
	case 0:
		index = 1
	default:
		index = len(c.Currenttasks) + 1
	}

	curr := CurrentTasks{
		Task:  task,
		Index: index,
		Time:  time.Now().Format("2006-01-02 3:4:5 pm"),
	}
	c.Currenttasks = append(c.Currenttasks, curr)
}

const localfile = "./config.cfg"

func (config Config) SaveConfig() {

	configBytes, err := json.MarshalIndent(config, "", "   ")
	if err != nil {
		fmt.Print("Something Wrong happened", err)
	}

	err = ioutil.WriteFile(localfile, configBytes, os.ModePerm)

	if err != nil {
		fmt.Print("Error writing device config!", err)
	}

	fmt.Print("Saved!\n")

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func LoadConfig() (*Config, error) {

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
