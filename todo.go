package main

import (
	"time"
	"errors"
	"fmt"
)

type Todo struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time //pointer is used because this value can be null
}

type TodoList []Todo

func (todolist *TodoList) add(title string){
	todo := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*todolist = append(*todolist, todo)
}

func (todolist *TodoList) validateIndex(index int) error{
	if index < 0 || index > len(*todolist) {
		err := errors.New("Inavlid Index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todolist *TodoList) delete(index int) error{
	t:=*todolist	
    if err := t.validateIndex(index); err!=nil{
		return err
	}
	*todolist = append(t[:index], t[index+1:]...)	
	return nil
}