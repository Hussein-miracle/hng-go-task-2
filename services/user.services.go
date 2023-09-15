package services

import "github.com/Hussein-miracle/hng-go-task-2/models"

type UserService interface {
	CreateUser(*models.User) error
}
