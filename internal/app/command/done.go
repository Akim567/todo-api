package command

import (
	"fmt"
	"todo-pet/internal/app/task"
)

type DoneCommand struct {
	Service *task.Service
	ID      int
}

func (h *DoneCommand) Execute() error {
	ok := h.Service.Done(h.ID)

	if !ok {
		fmt.Printf("Задача с ID %d не найдена\n", h.ID)
	} else {
		fmt.Printf("Задача с ID %d отмечена выполненной\n", h.ID)
	}
	return nil
}
