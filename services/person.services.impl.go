package services

import (
	"context"
	"errors"

	"github.com/Hussein-miracle/hng-go-task-2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	UserCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(UserCollection *mongo.Collection, ctx context.Context) PersonService {
	return &UserServiceImpl{
		UserCollection: UserCollection,
		ctx:            ctx,
	}
}

func (UserServiceImplPointer *UserServiceImpl) CreateUser(user *models.Person) error {
	_, err := UserServiceImplPointer.UserCollection.InsertOne(UserServiceImplPointer.ctx, user)

	return err
}

func (UserServiceImpl *UserServiceImpl) GetUser(user_id string) (*models.Person, error) {
	var user *models.Person

	var query primitive.D

	userId, err := primitive.ObjectIDFromHex(user_id)

	if err != nil {
		query = bson.D{bson.E{Key: "name", Value: user_id}}
	} else {
		query = bson.D{bson.E{Key: "_id", Value: userId}}
	}

	err = UserServiceImpl.UserCollection.FindOne(UserServiceImpl.ctx, query).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (UserServiceImpl *UserServiceImpl) DeleteUser(user_id string) error {
	var query primitive.D

	userId, err := primitive.ObjectIDFromHex(user_id)

	if err != nil {
		query = bson.D{bson.E{Key: "name", Value: user_id}}
	} else {
		query = bson.D{bson.E{Key: "_id", Value: userId}}
	}

	result, err := UserServiceImpl.UserCollection.DeleteOne(UserServiceImpl.ctx, query)

	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}

	return nil
}

func (UserServiceImpl *UserServiceImpl) UpdateUser(user *models.Person, user_id string) error {
	var query primitive.D

	userId, err := primitive.ObjectIDFromHex(user_id)

	if err != nil {
		query = bson.D{bson.E{Key: "name", Value: user.Name}}
	} else {
		query = bson.D{bson.E{Key: "_id", Value: userId}}
	}

	updateInter := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{
		Key: "name", Value: user.Name,
	}}}}

	result, err := UserServiceImpl.UserCollection.UpdateOne(UserServiceImpl.ctx, query, updateInter)

	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}

	return nil
}
