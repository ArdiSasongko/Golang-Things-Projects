package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-cli/todo"
)

type CmdFlags struct {
	Add, Edit      string
	Delete, Toggle int
	List           bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add new todo spesify tile")
	flag.StringVar(&cf.Edit, "edit", "", "Edit title todo spesify index and title")
	flag.IntVar(&cf.Delete, "del", -1, "Delete todo spesify index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle todo spesify index")
	flag.BoolVar(&cf.List, "print", false, "list all todo")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *todo.Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, Invalid format for edit, please user index:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error, invalid index for edit")
			os.Exit(1)
		}

		todos.Edit(index, parts[1])
	case cf.Toggle != -1:
		todos.Toogle(cf.Toggle)
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	default:
		fmt.Println("Invalid command")
	}
}
