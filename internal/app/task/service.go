package task

type Service struct {
	todos  []Todo
	nextID int
}

func NewService(initialTodos []Todo) *Service {
	// Присваиваем ID задачам, если они есть
	for i := range initialTodos {
		initialTodos[i].ID = i + 1
	}
	return &Service{
		todos:  initialTodos,
		nextID: len(initialTodos) + 1,
	}
}

func (s *Service) GetAll() []Todo {
	return s.todos
}

func (s *Service) Add(title string) Todo {
	todo := Todo{
		ID:     s.nextID,
		Title:  title,
		Status: "active",
	}
	s.todos = append(s.todos, todo)
	s.nextID++
	return todo
}
