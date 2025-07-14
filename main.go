package main

func main() {
	todo := TodoList{}
	storage := NewStorage[TodoList]("todo.json")
	storage.Load(&todo)
	commands := NewCmdFlags()
	commands.Execute(&todo)
	storage.Save(todo)
}