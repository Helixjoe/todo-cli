package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//todo-directory check
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	todoFound := false
	for _, file := range files {
		if file.Name() == "todos"{
			todoFound = true
			fmt.Println("todos directory found")
		}
	}
	if !todoFound {
		fmt.Println("todos directory not found")
		os.Mkdir("todos",0750)
		fmt.Println("todos directory created")
	}

	//todolists check
	files, err = os.ReadDir("todos")
	if err != nil {
		log.Fatal(err)
	}
	
	if len(files) == 0 {
		fmt.Println("0 todolists found")
		var newTodoName string
		fmt.Println("Enter name for new todolist:")
		fmt.Scan(&newTodoName)
		err := os.WriteFile("todos/"+newTodoName + ".json", []byte("{}"), 0650)
		if err != nil {
			fmt.Println("Failed to create file:", err)
		}			
		} else {
		fmt.Println("Todolists found are:")
		for _,file := range files {
			fmt.Println(file.Name())
		}
	}
	// todo := TodoList{}
	// storage := NewStorage[TodoList]("todo.json")
	// storage.Load(&todo)
	// commands := NewCmdFlags()
	// commands.Execute(&todo)
	// storage.Save(todo)
}