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
	// Если Writer не указан — используем os.Stdout по умолчанию
	if h.Writer == nil {
		h.Writer = os.Stdout
	}

	fmt.Fprintln(h.Writer, "Доступные команды:\n")

	// CLI-команды
	fmt.Fprintln(h.Writer, " ================= CLI Команды ================= ")
	fmt.Fprintln(h.Writer, "  list                      - Показать все задачи")
	fmt.Fprintln(h.Writer, "  add \"название\"           - Добавить новую задачу")
	fmt.Fprintln(h.Writer, "  delete ID                - Удалить задачу по ID")
	fmt.Fprintln(h.Writer, "  done ID                  - Отметить задачу выполненной")
	fmt.Fprintln(h.Writer, "  help                     - Показать это сообщение\n")

	// API-команды
	fmt.Fprintln(h.Writer, " ================= API Команды ================= ")
	fmt.Fprintln(h.Writer, "  GET    /todos                 - Получить все задачи (в JSON)")
	fmt.Fprintln(h.Writer, "  POST   /add                   - Добавить задачу (JSON {\"title\": \"...\"})")
	fmt.Fprintln(h.Writer, "  DELETE /delete?id=ID         - Удалить задачу по ID")
	fmt.Fprintln(h.Writer, "  PATCH  /done?id=ID           - Отметить задачу выполненной")
	fmt.Fprintln(h.Writer, "  GET    /todos/table          - Получить задачи в виде таблицы")

	return nil
}
