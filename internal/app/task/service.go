package task

type Service struct {
	todos []Todo
}

func NewService(initialTodos []Todo) *Service {
	return &Service{todos: initialTodos}
}

func (s *Service) GetAll() []Todo {
	return s.todos
}

func (s *Service) Add(todo Todo) {
	s.todos = append(s.todos, todo)
}
