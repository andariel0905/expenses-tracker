package handlers

import (
	"context"
	"fmt"

	"github.com/andariel0905/expenses-tracker/db"
	"github.com/andariel0905/expenses-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PostExpenseCategory(client *mongo.Client, cxt context.Context, newExpenseCategoryName string) {
	collection := db.GetMongoDBCollection(client, "expenseCategories")

	newExpenseCategory := models.ExpenseCategory{Name: newExpenseCategoryName}

	insertResult, err := collection.InsertOne(cxt, newExpenseCategory)

	if err != nil {
		panic(fmt.Sprintf("Doc insert issue %s", err))
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func GetExpenseCategories(client *mongo.Client, cxt context.Context) []bson.M {
	collection := db.GetMongoDBCollection(client, "expenseCategories")

	findOptions := options.Find()

	cursor, err := collection.Find(cxt, bson.D{}, findOptions)
	if err != nil {
		panic(fmt.Sprintf("Error while obtaining documents: %s", err))
	}
	defer cursor.Close(cxt)

	var results []bson.M

	for cursor.Next(cxt) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			panic(fmt.Sprintf("Error while decoding document: %s", err))
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		panic(fmt.Sprintf("Error while iterating the cursor: %s", err))
	}

	return results
}
