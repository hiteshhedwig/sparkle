package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kyokomi/emoji"
)

const localfile = ".sparkle"

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
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

type Conversion interface {
	ToCancelledTask() *CancelledTask
	ToCurrentTask() *CurrentTasks
	ToCompletedTask() *CompletedTask
}

func (c *Config) SetProjectName(name string) error {
	fmt.Printf("Setting Project name from : %s ", c.Project)
	c.Project = name
	fmt.Printf("-> %s", name)
	c.SaveConfig()
	return nil
}

func (c Config) DisplayBrief() {
	proj := c.Project
	totallen := len(c.Cancelledtasks) + len(c.Completedtasks) + len(c.Currenttasks)
	currenttasklen := len(c.Currenttasks)
	compllen := len(c.Completedtasks)
	canclen := len(c.Cancelledtasks)
	fmt.Println("         \t\t    SUMMARY 🪶                  ")
	fmt.Printf("        \n\t %s Project Name : %s ", string(colorWhite), proj)
	fmt.Printf("        \n\t %s Total registered task in your project : %d ", string(colorWhite), totallen)
	fmt.Printf("		\n\t %s Total ongoing tasks in your project : %d ", string(colorWhite), currenttasklen)
	fmt.Printf("		\n\t %s Total completed tasks in your project : %d ", string(colorWhite), compllen)
	fmt.Printf("		\n\t %s Total cancelled tasks in your project : %d ", string(colorWhite), canclen)

	if currenttasklen != 0 {
		efficiency := float32(compllen) / float32(currenttasklen+compllen)
		ef := fmt.Sprintf("%.2f", efficiency)
		if efficiency > 0.70 {
			fmt.Printf("		\n\t %s Overall efficiency: %s ", string(colorGreen), ef)
		} else if efficiency < 0.70 && efficiency > 0.40 {
			fmt.Printf("		\n\t %s Overall efficiency: %s ", string(colorYellow), ef)
		} else if efficiency < 0.40 {
			fmt.Printf("		\n\t %s Overall efficiency: %s ", string(colorRed), ef)
		}
	}

}

func (c *Config) UpdateTask(idx int) error {
	var updatedtask string
	idx = idx - 1
	if idx > len(c.Currenttasks) {
		return fmt.Errorf("OverBound")
	}

	fmt.Printf("Current task is : %s \n", c.Currenttasks[idx].Task)
	fmt.Printf("Please type updated task  : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		updatedtask = scanner.Text()
	}
	c.Currenttasks[idx].Task = updatedtask
	c.SaveConfig()
	return nil
}

func (c *Config) PurgeTask(idx int) {
	if idx > len(c.Cancelledtasks) {
		fmt.Printf("%s Index value [%d] is not correct, Please check again in below table \n \n", string(colorRed), idx)
		fmt.Printf("")
		c.ListTasksCatWise()
		return
	}

	c.PurgeforReal(idx)

}

func (c *Config) CompletedTask(idx int) error {
	// idx check!
	if idx > len(c.Currenttasks) {
		fmt.Printf("%s Index value [%d] is not correct, Please check again in below table \n \n", string(colorRed), idx)
		fmt.Printf("")
		c.ListAllCurrentTasks()
		return fmt.Errorf("out of range")
	}

	donetask := c.Currenttasks[idx-1]
	disp := emoji.Sprintf(":pizza:")
	fmt.Println(disp)
	fmt.Printf("%s You have completed the task : %s \n", string(colorGreen), donetask.Task)

	c.Currenttasks = remove(c.Currenttasks, idx-1)
	can := c.ToCompletedTask(donetask)
	c.Completedtasks = append(c.Completedtasks, *can)

	for i := range c.Currenttasks {
		c.Currenttasks[i].Index = i + 1
	}

	c.SaveConfig()
	return nil
}

func (c *Config) ToCompletedTask(com CurrentTasks) *CompletedTask {

	return &CompletedTask{
		Index: com.Index,
		Task:  com.Task,
		Time:  time.Now().Format("2006-01-02 3:4:5 pm"),
	}

}

func (c *Config) ToCancelledTask(com CurrentTasks) *CancelledTask {

	return &CancelledTask{
		Index: com.Index,
		Task:  com.Task,
		Time:  time.Now().Format("2006-01-02 3:4:5 pm"),
	}

}

func (c *Config) Cancel(idx int) {

	var userin string

	dumptask := c.Currenttasks[idx-1]
	fmt.Println("Are you sure you want to cancel this task? ", dumptask.Task)
	fmt.Print("Type yes/no to confirm cancellation : ")
	fmt.Scanf("%s", &userin)

	switch strings.ToLower(userin) {
	case "no":
		fmt.Println("Aborting")
		return
	case "yes":
		fmt.Println("ok to cancel this task")
		c.CancelforReal(idx)
	default:
		fmt.Println("Aborting")
		return
	}

}

func (c *Config) PurgeforReal(idx int) {

	dumptask := c.Cancelledtasks[idx-1]
	fmt.Println("Purging : ", dumptask.Task)
	c.Cancelledtasks = removefromCancel(c.Cancelledtasks, idx-1)
	c.SaveConfig()
}

func (c *Config) CancelforReal(idx int) {

	dumptask := c.Currenttasks[idx-1]
	c.Currenttasks = remove(c.Currenttasks, idx-1)
	fmt.Println(dumptask)
	can := c.ToCancelledTask(dumptask)
	c.Cancelledtasks = append(c.Cancelledtasks, *can)

	for i := range c.Currenttasks {
		c.Currenttasks[i].Index = i + 1
	}

	c.SaveConfig()
}

func remove(slice []CurrentTasks, s int) []CurrentTasks {
	return append(slice[:s], slice[s+1:]...)
}

func removefromCancel(slice []CancelledTask, s int) []CancelledTask {
	return append(slice[:s], slice[s+1:]...)
}

func (c Config) CheckCurrentTask() {
	if len(c.Currenttasks) == 0 && len(c.Cancelledtasks) == 0 && len(c.Completedtasks) == 0 {
		fmt.Println("No Task Found! Start around adding them, try :")
		fmt.Println("     sparkle add --help")
	}
}

func (c Config) ListAllCurrentTasks() {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Index", "Task", "Time"})
	c.CheckCurrentTask()
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

	typetasks = append(typetasks, []interface{}{"******", "Current Task", "****************"})
	typetasks = append(typetasks, []interface{}{"******", "Completed Task", "****************"})
	typetasks = append(typetasks, []interface{}{"******", "Cancelled Task", "****************"})

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

	encrypted := config.SaveJwtConf()

	file, err := os.Create(localfile)
	if err != nil {
		fmt.Println("Some error while saving", err)
	}

	file.WriteString(encrypted)

}

func (config Config) LoadFromTokenFile(tokenString string) (*Config, error) {

	var hmacSampleSecret []byte
	var conf Config

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		maptoconfig(&conf, claims["config"])

	} else {
		fmt.Println(err)
	}
	return &conf, nil
}

func (config Config) SaveJwtConf() string {
	var hmacSampleSecret []byte
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"config": config,
	})

	tokenString, _ := token.SignedString(hmacSampleSecret)

	return tokenString
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func LoadConfig(project string) (*Config, error) {

	var config Config

	if fileExists(localfile) {
		configFile, err := os.Open(localfile)
		if err != nil {
			fmt.Errorf("Error opening device file %v", err)
			return nil, err
		}
		defer configFile.Close()
		configBytes, _ := ioutil.ReadAll(configFile)
		tokenloaded := string(configBytes)
		conf, _ := config.LoadFromTokenFile(tokenloaded)
		return conf, nil
	}

	return &config, nil

}

func maptoconfig(temp *Config, config interface{}) {
	temp.LoadFromMap(config)
}
