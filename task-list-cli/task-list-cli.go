package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// tasksFileName represents the name of the file where tasks are stored.
const tasksFileName = "tasks.txt"

func main() {
	// Load tasks from file
	tasks := loadTasks()

	// Display initial tasks
	fmt.Println("Task List:")
	displayTasks(tasks)

	// Take user input
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter a command (add/remove/quit): ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "add":
			fmt.Print("Enter task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			tasks = append(tasks, task)
			saveTasks(tasks)
		case "remove":
			fmt.Print("Enter task index to remove: ")
			var index int
			fmt.Scanln(&index)
			if index >= 1 && index <= len(tasks) {
				tasks = removeTask(tasks, index-1)
				saveTasks(tasks)
			} else {
				fmt.Println("Invalid index.")
			}
		case "quit":
			fmt.Println("Quitting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Try again.")
		}

		// Display updated tasks
		fmt.Println("\nUpdated Task List:")
		displayTasks(tasks)
	}
}

// displayTasks prints the tasks to the console.
func displayTasks(tasks []string) {
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

// loadTasks reads tasks from the tasksFileName file and returns them as a slice.
func loadTasks() []string {
	file, err := os.Open(tasksFileName)
	if err != nil {
		// If there's an error opening the file, return an empty slice.
		return []string{}
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read each line from the file and add it to the tasks slice.
		tasks = append(tasks, scanner.Text())
	}
	return tasks
}

// saveTasks writes tasks to the tasksFileName file.
func saveTasks(tasks []string) {
	file, err := os.Create(tasksFileName)
	if err != nil {
		// If there's an error creating the file, print an error message.
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		// Write each task to a new line in the file.
		fmt.Fprintln(file, task)
	}
}

// removeTask removes a task at the specified index from the tasks slice.
func removeTask(tasks []string, index int) []string {
	return append(tasks[:index], tasks[index+1:]...)
}
