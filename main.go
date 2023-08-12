package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://gmg:password@localhost:27017/"),
	)

	defer func() {
		cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("mongodb disconnect error : %v", err)
		}
	}()

	if err != nil {
		log.Fatalf("connection error :%v", err)
		return
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping mongodb error :%v", err)
		return
	}
	fmt.Println("ping success")

	db := mongoClient.Database("demo")
	sampleCollection := db.Collection("sample")
	sampleCollection.Drop(ctx)

	insertedDocument := bson.M{
		"name":       "cookie",
		"content":    "test content",
		"bank_money": 1000,
		"create_at":  time.Now(),
	}
	insertedRes, err := sampleCollection.InsertOne(ctx, insertedDocument)

	if err != nil {
		log.Fatalf("inserted error: %v", err)
		return
	}

	log.Printf("inserted ID is: %v", insertedRes.InsertedID)

	// QUERY ALL DATA
	fmt.Println("== qeury all data ==")
	cursor, err := sampleCollection.Find(ctx, options.Find())
	if err != nil {
		log.Fatalf("find collection err: %v", err)
		return
	}

	var queryResult []bson.M
	if err := cursor.All(ctx, &queryResult); err != nil {
		log.Fatalf("query mongodb error: %v", err)
		return
	}

	for _, doc := range queryResult {
		fmt.Println(doc)
	}

	updateManyFilter := bson.D{
		bson.E{
			Key:   "name",
			Value: "cookie",
		},
	}

	updateSet := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{
					Key:   "bank_money",
					Value: 2000,
				},
			},
		},
	}

	updateManyResult, err := sampleCollection.UpdateMany(ctx, updateManyFilter, updateSet)
	if err != nil {
		log.Fatalf("update error: %v", err)
		return
	}

	fmt.Println("=== update modified count")
	fmt.Println(updateManyResult.ModifiedCount)

}
