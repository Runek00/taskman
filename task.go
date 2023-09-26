package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var nextTaskId int = 0
var taskStore map[int]Task = make(map[int]Task)

type Priority int8

const (
	Low    Priority = 1
	Medium Priority = 2
	High   Priority = 3
)

type Task struct {
	Id        int
	Title     string
	Completed bool
	Priority  Priority
}

func createTask(title string) Task {
	out := Task{nextTaskId, title, false, 1}
	nextTaskId++
	return out
}

func addTask(title string) Task {
	task := createTask(title)
	taskStore[task.Id] = task
	return task
}

func showTasks() {
	var tasks []Task = make([]Task, 0, len(taskStore))
	for _, v := range taskStore {
		tasks = append(tasks, v)
	}
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].Priority > tasks[j].Priority
	})

	for i := 0; i < len(tasks); i++ {
		t := tasks[i]
		fmt.Printf("\nPriority: %v	ID: %v	Title: %v	Completed: %v", t.Priority, t.Id, t.Title, t.Completed)
	}
	fmt.Println()
}

func completeTask(id int) {
	t, ok := taskStore[id]
	if !ok {
		fmt.Print("No such task")
		return
	}
	t.Completed = true
	taskStore[id] = t
}

func deleteTask(id int) {
	_, ok := taskStore[id]
	if !ok {
		fmt.Print("No such task")
		return
	}
	delete(taskStore, id)
}

func interactive() {
	load()
MainLoop:
	for {
		fmt.Print("\nWhat do you want to do?\n(add) add task\n(del) delete task\n(com) complete task\n(show) show tasks\n(set) set priority\n(exit) exit\n")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		switch strings.TrimSpace(input) {
		case "add":
			addTask(getTitle())
		case "del":
			deleteTask(getId())
		case "com":
			completeTask(getId())
		case "show":
			showTasks()
		case "set":
			setPriority(getId())
		case "exit":
			fmt.Print("BAIII!\n")
			break MainLoop
		default:
			fmt.Print("No such option\n")
		}
	}
	save()
}

func setPriority(id int) {
	t := taskStore[id]
	fmt.Printf("\nThe old priority is %v.\n1-Low, 2-Medium, 3-High\nWhat is the new priority?\n", t.Priority)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	pr, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(err)
		return
	}
	switch pr {
	case 1:
		t.Priority = Low
	case 2:
		t.Priority = Medium
	case 3:
		t.Priority = High
	}
	taskStore[id] = t
}

func getTitle() string {
	fmt.Print("\nWhat is the title?\n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(input)
}

func getId() int {
	fmt.Print("\nWhat is the id?\n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return -1
	}
	inn, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return inn
}

func save() {
	marshaled, err := json.Marshal(taskStore)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile("store.json", marshaled, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func load() {
	marshaled, err := os.ReadFile("store.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(marshaled, &taskStore)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k := range taskStore {
		if k >= nextTaskId {
			nextTaskId = k + 1
		}
	}

}

func main() {
	// // For testing
	// addTask("q")
	// addTask("w")
	// addTask("e")
	// addTask("r")
	// addTask("t")
	// addTask("y")
	completeTask((3))
	showTasks()
	interactive()
}
