package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Hussein-miracle/hng-go-task-2/controllers"
	"github.com/Hussein-miracle/hng-go-task-2/routes"
	"github.com/Hussein-miracle/hng-go-task-2/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	routerEngine   *gin.Engine
	userservice    services.PersonService
	usercontroller controllers.PersonController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.Please add")
	}

	ctx = context.TODO()

	mongoconnection := options.Client().ApplyURI(uri)
	mongoclient, err = mongo.Connect(ctx, mongoconnection)

	if err != nil {
		log.Fatal(err)
	}
	// defer func() {
	// 	if err = mongoclient.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	if err := mongoclient.Database("hng-go-task-2").RunCommand(ctx, bson.D{bson.E{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	fmt.Println("connection established")

	usercollection = mongoclient.Database("hng-go-task-2").Collection("person")
	fmt.Println("collection retrieved")

	userservice = services.NewUserService(usercollection, ctx)
	fmt.Println("services retrieved")
	usercontroller = controllers.NewUserController(userservice)
	fmt.Println("controllers retrieved")
	routerEngine = gin.Default()
	routes.RegisterUserRoutes(routerEngine, &usercontroller)
}

func main() {
	defer mongoclient.Disconnect(ctx)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := routerEngine.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
