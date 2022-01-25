package repository

// Repository represent the repositories
type Repository interface {
	Close()
	FindByID(id string) (*TodoListModel, error)
	Find() ([]*TodoListModel, error)
	Create(user *TodoListModel) error
	Update(user *TodoListModel) error
	Delete(id string) error
}

// TodoListModel represent the user model
type TodoListModel struct {
	Id         string `json:"Id"`
	Text       string `json:"Text"`
	IsSelected string `json:"IsSelected"`
}

type MockRepository struct {
}

func (mockRepository MockRepository) Close()                                     {}
func (mockRepository MockRepository) FindByID(id string) (*TodoListModel, error) { return nil, nil }
func (mockRepository MockRepository) Find() ([]*TodoListModel, error)            { return nil, nil }
func (mockRepository MockRepository) Create(todolist *TodoListModel) error       { return nil }
func (mockRepository MockRepository) Update(todolist *TodoListModel) error       { return nil }
func (mockRepository MockRepository) Delete(id string) error                     { return nil }

func GetMockRepository() Repository {
	return MockRepository{}
}
