package command

import (
	"fmt"
	"io"
	"os"
)

type HelpCommand struct {
	Writer io.Writer
}

func (h *HelpCommand) Execute() error {
	// Если Writer не указан — выводим в стандартный вывод (CLI)
	if h.Writer == nil {
		h.Writer = os.Stdout
	}

	fmt.Fprintln(h.Writer, "\nКоманды CLI:")
	fmt.Fprintln(h.Writer, "  go run ./cmd list             - Показать список задач")
	fmt.Fprintln(h.Writer, `  go run ./cmd add "Задача"     - Добавить новую задачу`)
	fmt.Fprintln(h.Writer, "  go run ./cmd delete ID        - Удалить задачу по ID")
	fmt.Fprintln(h.Writer, "  go run ./cmd done ID          - Отметить задачу выполненной")
	fmt.Fprintln(h.Writer, "  go run ./cmd help             - Показать справку")

	fmt.Fprintln(h.Writer, "API-команды:")
	fmt.Fprintln(h.Writer, "  GET     /todos                - Получить все задачи (JSON)")
	fmt.Fprintln(h.Writer, `  POST    /add {"title":"..."}`+"   - Добавить задачу")
	fmt.Fprintln(h.Writer, "  DELETE  /delete?id=ID         - Удалить задачу по ID")
	fmt.Fprintln(h.Writer, "  PATCH   /done?id=ID           - Отметить задачу выполненной")
	fmt.Fprintln(h.Writer, "  GET     /todos/table          - Табличный вид задач")
	fmt.Fprintln(h.Writer, "  GET     /help                 - Показать справку по API")

	return nil
}
