package todo

import (
	"encoding/json"
	"os"
	"time"
)

type ItemUtils interface {
	AddTask(name string) (bool, string)
	CompleteTask(index int) (bool, string)
	Save() (bool, string)
	Load() (bool, string)
	Delete(index int) (bool, string)
}

type Item struct {
	Name        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

var Todos []Item

const jsonfile = ".todo.json"

func (t *Item) AddTask(name string) (bool, string) {
	todo := &Item{
		Name:        name,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	Todos = append(Todos, *todo)

	return true, ""
}

func (t *Item) CompleteTask(index int) (bool, string) {
	if index-1 < 0 {
		return false, "index is lower than 0"
	}

	Todos[index-1].Done = true
	Todos[index-1].CompletedAt = time.Now()

	return true, ""
}

func (t *Item) Save() (bool, string) {
	jsonData, err := json.Marshal(&Todos)
	if err != nil {
		return false, ""
	}

	err = os.WriteFile(jsonfile, jsonData, 0666)
	if err != nil {
		return false, ""
	}

	return true, ""
}

func (t *Item) Load() (bool, string) {
	file, err := os.ReadFile(jsonfile)
	if err != nil {
		return false, ""
	}

	err = json.Unmarshal(file, &Todos)
	if err != nil {
		return false, ""
	}

	return true, ""
}

func (t *Item) Delete(index int) (bool, string) {
	Todos = append(Todos[:index-1], Todos[index:]...)

	return true, ""
}
