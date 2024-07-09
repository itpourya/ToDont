package main

import (
	"github.com/itpourya/ToDont/todo"
	"log"
)

func main() {
	todolist := todo.Item{}
	_, err := todolist.Load()
	_, err = todolist.CompleteTask(1)
	_, err = todolist.Save()

	if err != "" {
		log.Fatalln(err)
	}
}
