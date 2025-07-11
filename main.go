package main

func main() {
	todo := TodoList{}
	todo.add("Buy Groceries")
	todo.add("Go to the Gym")
	todo.toggle(0)
	todo.print()
}