package main

import "todo-cli/todo"

func main() {
	todos := todo.Todos{}
	storage := NewStorage[todo.Todos]("./result/todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}
