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
	Title     string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func(todos *Todos) add(Title string){
	todo := Todo{
		Title: Title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func(todos *Todos) ValidateIndex(index int) error{
	if index < 0 || index >= len(*todos){
		err := errors.New("index out of range")
		fmt.Println(err)
		return err
	}

	return nil
}

func(todos *Todos) Delete(index int) error{
	t := *todos

	if err := t.ValidateIndex(index); err != nil{
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func(todos *Todos) Toggle(index int) error{
	t := *todos

	if err := t.ValidateIndex(index); err != nil{
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted{
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted
	return nil
}

func(todos *Todos) Edit(index int, newTitle string) error{
	t := *todos
	if err := t.ValidateIndex(index); err != nil{
		return err
	}

	t[index].Title = newTitle
	return nil
}

func(todos *Todos) Print() {
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("Index", "Title", "Completed", "Created At", "Completed At")
	
	for index, t := range *todos{
		completed := "❌"
		completedAt := ""

		if t.Completed{
			completed ="✅"
			if t.CompletedAt != nil{
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}