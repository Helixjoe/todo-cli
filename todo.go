package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
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

func (todolist *TodoList) toggle(index int) error{
	t:=*todolist
	if err := t.validateIndex(index); err!=nil {
		return err
	}
	isCompleted := t[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	t[index].Completed = !isCompleted	
	return nil
}

func (todolist *TodoList) edit(index int,title string) error {
	t := *todolist
	if err := t.validateIndex(index); err!=nil {
		return err
	}
	t[index].Title = title
	return nil	
}

func (todolist *TodoList) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#","Title","Completed","CreatedAt","CompletedAt")
	for index,t := range *todolist{
		completed := "❌"
		completedAt := ""
		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()	
}