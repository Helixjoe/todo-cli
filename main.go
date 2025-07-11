package main

import "fmt"

func main() {
	todo := TodoList{}
	todo.add("Buy Groceries")
	todo.add("Go to the Gym")
	fmt.Printf("%+v\n\n",todo)
	todo.delete(0)
	fmt.Printf("%+v\n\n",todo)
}