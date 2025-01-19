package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	dataDir      = ".data/todo_json"
	dataFileName = "data.json"

	todoStatus = 0
	inProgress = 1
	done       = 2
)

var (
	dataPath string
	err      error
	taskList []task
)

type task struct {
	Id         int
	Status     int
	Content    string
	CreateAt   string
	ModifiedAt string
}

func newTask(content string) task {
	return task{
		Id:         len(taskList),
		Status:     todoStatus,
		Content:    content,
		CreateAt:   time.Now().String(),
		ModifiedAt: time.Now().String(),
	}
}

func init() {
	dataPath, err = os.UserHomeDir()
	if err != nil {
		panic("Can't find HOME directory.")
	}

	dataPath = mustDataFile(filepath.Join(dataPath, dataDir), dataFileName)
	getTaskList()
}

func mustDataFile(path, file string) string {
	if _, err = os.Stat(path); err != nil {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			panic(err)
		}
	}
	return filepath.Join(path, file)
}

func mustFile(file *os.File, err error) *os.File {
	if err != nil {
		panic(err)
	}
	return file
}

func getTaskList() {
	reader := json.NewDecoder(mustFile(os.Open(dataPath)))
	reader.Decode(&taskList)
}

func saveTasks() {
	f := mustFile(os.OpenFile(dataPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm))
	defer f.Sync()
	writer := json.NewEncoder(f)
	writer.Encode(taskList)
}

func AddTasks(content string) {
	defer saveTasks()
	add := newTask(content)
	taskList = append(taskList, add)
}

func DeleteTask(id int) {
	defer saveTasks()
	if id < 0 || id >= len(taskList) {
		panic(fmt.Errorf("invalid task id"))
	}
	for i := id + 1; i < len(taskList); i++ {
		taskList[i].Id--
	}
	taskList = append(taskList[:id], taskList[id+1:]...)
}

func UpdateTask(id int, content string) {
	defer saveTasks()
	if id < 0 || id >= len(taskList) {
		panic(fmt.Errorf("invalid task id"))
	}
	taskList[id].Content = content
	taskList[id].ModifiedAt = time.Now().String()
}

func ListTasks(mark string) {
	status := -1
	switch mark {
	case "todo":
		status = 0
	case "inprogress":
		status = 1
	case "done":
		status = 2
	case "all":
		status = -1
	default:
		panic(fmt.Errorf("unknown flag %s", mark))
	}
	for _, task := range taskList {
		if status == -1 {
			fmt.Println(task.Id, " ", task.Content)
		} else if task.Status == status {
			fmt.Println(task.Id, " ", task.Content)
		}
	}
}

func MarkTasks(id int, mark string) {
	defer saveTasks()
	if id < 0 || id >= len(taskList) {
		panic(fmt.Errorf("invalid task id"))
	}

	status := 0
	switch mark {
	case "inprogress":
		status = 1
	case "done":
		status = 2
	}
	taskList[id].Status = status
}
