package main

func main() {
	todo := TodoList{}
	storage := NewStorage[TodoList]("todo.json")
	storage.Load(&todo)
	todo.add("Buy Groceries")
	todo.add("Go to the Gym")
	todo.toggle(0)
	todo.print()
	storage.Save(todo)
}