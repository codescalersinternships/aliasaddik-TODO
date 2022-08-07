package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aliasaddik/todo-project/controllers"
	"github.com/aliasaddik/todo-project/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	handles     services.Handles
	controller  controllers.Controller
	ctx         context.Context
	taskc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://aliasaddik:12345@mongo:27017/?authSource=admin")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	taskc = mongoclient.Database("taskdb").Collection("tasks")
	handles = services.NewHandle(taskc, ctx)
	controller = controllers.New(handles)
	server = gin.New()
	server.Use(cors.New(cors.Config{
		AllowHeaders: []string{"Content-Type", "Access-Control-Allow-Origin"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "DELETE", "GET", "POST", "PATCH"},
	}))
	server.Static("/swaggerui/", "cmd/api/swaggerui")

}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/")
	controller.RegisterRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}
