package domain

import (
	"log"
	"repo-name/project-name/function-name/internal/domain/model"
	"repo-name/project-name/function-name/internal/port"
)

type ExampleCRUDService interface {
	Create(*model.Example) error
	Update(*model.Example) error
	GetByID(string) (*model.Example, error)
	Delete(string) error
}

type exampleCRUDService struct {
	myExampleClientPort port.ExampleClientPort
}
//debería exponer una interfaz y no una implementación
func NewExampleCRUDService(client port.ExampleClientPort) ExampleCRUDService {
	return &exampleCRUDService{
		myExampleClientPort: client,
	}
}

func (s *exampleCRUDService) Create( resource *model.Example) (err error) {
	log.Logger.Printf("Entering CRUD example create [%s] ", resource.Key)
	return s.myExampleClientPort.Create(resource)
}

func (s *exampleCRUDService) Update( resource *model.Example) (err error) {
	log.Logger.Printf("Entering CRUD example update [%s] ", resource.Key)
	return s.myExampleClientPort.Update(resource)
}

func (s *exampleCRUDService) GetByID( key string) (resource *model.Example, err error) {
	log.Logger.Printf("Entering CRUD example getbyKey [%s] ", key)
	return s.myExampleClientPort.GetByID(key)
}

func (s *exampleCRUDService) Delete( key string) (err error) {
	log.Logger.Printf("Entering CRUD example delete [%s] ", key)
	return s.myExampleClientPort.Delete(key)
}

