package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Person struct {
	ID   primitive.ObjectID `bson:"_id"`
	NAME string             `bson:"name"`
}

func main() {
	// creating the router
	var router = CreateServer()
	// creating the client
	client := Dbclient()
	fmt.Println("Connected to MongoDB!")
	// connecting to the database and collection
	collection := ConnectToDatabase("gotest", "users", client)
	fmt.Println("Collection instance created!")

	// creating the routes
	router.GET("/users", func(c *gin.Context) {
		var users []Person
		cursor, _ := collection.Find(context.Background(), bson.M{})
		for cursor.Next(context.Background()) {
			var person Person
			if err := cursor.Decode(&person); err != nil {
				log.Fatal(err)
			}
			fmt.Println(person)
			users = append(users, person)
		}
		c.JSON(200, users)
		return
	})
	router.POST("/users", func(c *gin.Context) {
		username := c.PostForm("name")
		one, err := collection.InsertOne(context.Background(), bson.M{
			"name": username,
		})
		if err != nil {
			return
		}
		fmt.Println(one)
	})
	router.PUT("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		one, err := collection.UpdateOne(context.Background(), bson.M{
			"name": name,
		}, bson.M{
			"$set": bson.M{
				"name": "bro",
			},
		})
		if err != nil {
			return
		}
		fmt.Println(one)
	})
	router.DELETE("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		_, err := collection.DeleteOne(context.Background(), bson.M{
			"name": name,
		})
		if err != nil {
			return
		}
		c.String(200, "Deleted")
	})
	// running the server
	RunServer(router)

	// closing the connection
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())
}
