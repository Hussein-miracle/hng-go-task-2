package services

import (
	"context"

	"github.com/Hussein-miracle/hng-go-task-2/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	UserCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(UserCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		UserCollection: UserCollection,
		ctx:            ctx,
	}
}

func (UserServiceImplPointer *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := UserServiceImplPointer.UserCollection.InsertOne(UserServiceImplPointer.ctx, user)

	return err
}
