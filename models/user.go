package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	State   string `json:"state"`
	City    string `json:"city"`
	Pincode int    `json:"pincode"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name"`
	Address   `json:"adress"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
