package todo

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Complete    bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	newTodo := Todo{
		Title:       title,
		Complete:    false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, newTodo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		return err
	}

	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Toogle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Complete

	if !isCompleted {
		completeTime := time.Now()
		t[index].CompletedAt = &completeTime
	}

	t[index].Complete = !isCompleted

	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Index", "Title", "Completed", "Created At", "Completed At")
	for i, t := range *todos {
		completed := "x"
		completedAt := ""

		if t.Complete {
			completed = "o"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(i), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
