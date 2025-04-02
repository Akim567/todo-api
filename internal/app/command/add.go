package command

import (
	"fmt"
	"todo-pet/internal/app/task"
)

type AddCommand struct {
	Service *task.Service
	Title   string
}

func (c *AddCommand) Execute() error {
	todo := c.Service.Add(c.Title)
	fmt.Printf("Задача добавлена: [%d] %s\n", todo.ID, todo.Title)
	return nil
}
