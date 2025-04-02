package command

import (
	"fmt"
	"todo-pet/internal/app/task"
)

type DeleteCommand struct {
	Service *task.Service
	ID      int
}

func (h *DeleteCommand) Execute() error {
	ok := h.Service.DeleteById(h.ID)

	if !ok {
		fmt.Printf("Задача с ID %d не найдена\n", h.ID)
	} else {
		fmt.Printf("Задача с ID %d удалена\n", h.ID)
	}
	return nil
}
