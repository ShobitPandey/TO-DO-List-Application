package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID     int
	Name   string
	Status bool
}

var tasks []Task
var nextID int = 1

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- To-Do List Application ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			addTask(reader)
		case "2":
			viewTasks()
		case "3":
			markTaskDone(reader)
		case "4":
			deleteTask(reader)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again later.")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Task name cannot be blanked.")
		return
	}

	task := Task{
		ID:     nextID,
		Name:   name,
		Status: false,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Task added successfully!")
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("\n--- Task List ---")
	for _, task := range tasks {
		status := "Pending"
		if task.Status {
			status = "Done"
		}
		fmt.Printf("ID: %d | Name: %s | Status: %s\n", task.ID, task.Name, status)
	}
}

func markTaskDone(reader *bufio.Reader) {
	fmt.Print("Enter task ID to mark as complete: ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = true
			fmt.Println("Task marked as complete!")
			return
		}
	}

	fmt.Println("Task not found.")
}

func deleteTask(reader *bufio.Reader) {
	fmt.Print("Enter task ID to delete: ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return
		}
	}

	fmt.Println("Task not found.")
}
