package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	log "github.com/sirupsen/logrus"
)

const localfile = "./config.cfg"

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

func (c *Config) ListAllCurrentTasks() {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Index", "Task", "Time"})

	for _, currenttask := range c.Currenttasks {
		t.AppendSeparator()
		t.AppendRow([]interface{}{currenttask.Index, currenttask.Task, currenttask.Time})
	}
	t.SetStyle(table.StyleRounded)
	t.Render()

}

func (c *Config) ListTasksCatWise() {

	var typetasks [][]interface{}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Index", "Task", "Time"})

	//typetask := []interface{}{000, "Current Task", " --"}
	typetasks = append(typetasks, []interface{}{"--", "Current Task", " --"})
	typetasks = append(typetasks, []interface{}{"--", "Completed Task", " --"})
	typetasks = append(typetasks, []interface{}{"--", "Cancelled Task", " --"})

	for idx, typetask := range typetasks {
		t.AppendSeparator()
		t.AppendRow(typetask)
		t.AppendSeparator()

		switch idx {
		case 0:
			for _, tsk := range c.Currenttasks {
				t.AppendRow([]interface{}{tsk.Index, tsk.Task, tsk.Time})
			}
		case 1:
			for _, tsk := range c.Completedtasks {
				t.AppendRow([]interface{}{tsk.Index, tsk.Task, tsk.Time})
			}
		case 2:
			for _, tsk := range c.Cancelledtasks {
				t.AppendRow([]interface{}{tsk.Index, tsk.Task, tsk.Time})
			}
		}

	}
	t.SetStyle(table.StyleRounded)
	t.Render()

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
