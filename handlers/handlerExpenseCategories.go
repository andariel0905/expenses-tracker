package handlers

import (
	"context"
	"fmt"

	"github.com/andariel0905/expenses-tracker/db"
	"github.com/andariel0905/expenses-tracker/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func postExpenseCategory(client *mongo.Client, cxt context.Context, newExpenseCategoryName string) {
	collection := db.GetMongoDBCollection(client, "expenseCategories")

	newExpenseCategory := models.ExpenseCategory{newExpenseCategoryName}

	insertResult, err := collection.InsertOne(cxt, newExpenseCategory)

	if err != nil {
		panic(fmt.Sprintf("Doc insert issue %s", err))
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
