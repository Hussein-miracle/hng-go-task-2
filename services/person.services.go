package services

import (
	"github.com/Hussein-miracle/hng-go-task-2/models"
)

type PersonService interface {
	CreateUser(*models.Person) error
	GetUser(user_id string) (*models.Person, error)
	DeleteUser(user_id string) error
	UpdateUser(*models.Person, string) error
}
