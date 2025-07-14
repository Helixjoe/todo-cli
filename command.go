package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add string
	Del int
	toggle int
	Edit string
	List bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add,"add","","Add a todo item")
	flag.StringVar(&cf.Edit,"edit","","Edit a todo item")
	flag.IntVar(&cf.Del,"del",-1,"Delete a todo item")
	flag.IntVar(&cf.toggle,"toggle",-1,"Toggle a todo item")
	flag.BoolVar(&cf.List,"list",false,"List all todo")
	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute(todos *TodoList){
	switch{
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":" ,2)		
		if len(parts) !=2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err!=nil{
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index,parts[1])
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.toggle != -1:
		todos.toggle(cf.toggle)
	default:
		fmt.Println("Invalid Command")
	}
}