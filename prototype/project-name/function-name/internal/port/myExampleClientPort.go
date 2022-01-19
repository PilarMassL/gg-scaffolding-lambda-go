package port

import "repo-name/project-name/function-name/internal/domain/model"

// ExampleClientPort describes the example port repository.
type ExampleClientPort interface {
	Create(*model.Example) error
	Update(*model.Example) error
	GetByID(string) (*model.Example, error)
	Delete(string) error
}