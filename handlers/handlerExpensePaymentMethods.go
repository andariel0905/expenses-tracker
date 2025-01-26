package handlers

import (
	"fmt"

	"github.com/andariel0905/expenses-tracker/db"
	"github.com/andariel0905/expenses-tracker/global"
	"github.com/andariel0905/expenses-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PostPaymentMethod(newPaymentMethodName string) {
	collection := db.GetMongoDBCollection(global.Client, "expensePaymentMethods")

	newPaymentMethod := models.ExpenseCategory{Name: newPaymentMethodName}

	insertResult, err := collection.InsertOne(global.Context, newPaymentMethod)

	if err != nil {
		panic(fmt.Sprintf("Doc insert issue %s", err))
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func GetPaymentMethods() []bson.M {
	collection := db.GetMongoDBCollection(global.Client, "expensePaymentMethods")

	findOptions := options.Find()

	cursor, err := collection.Find(global.Context, bson.D{}, findOptions)
	if err != nil {
		panic(fmt.Sprintf("Error while obtaining documents: %s", err))
	}
	defer cursor.Close(global.Context)

	var results []bson.M

	for cursor.Next(global.Context) {
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

func SetPaymentMethod(currentName string, newName string) {
	collection := db.GetMongoDBCollection(global.Client, "expensePaymentMethods")

	filter := bson.M{"name": currentName}
	updateInterface := bson.M{
		"$set": bson.M{"name": newName},
	}

	insertResult, err := collection.UpdateOne(global.Context, filter, updateInterface)

	if err != nil {
		panic(fmt.Sprintf("Doc insert issue %s", err))
	}

	if insertResult.MatchedCount == 0 {
		fmt.Println("Operation didn't match any document")
	} else {
		fmt.Println("Document's name updated from " + currentName + " to " + newName)
	}
}

func DeletePaymentMethod(expenseCategoryName string) {
	collection := db.GetMongoDBCollection(global.Client, "expensePaymentMethods")

	filter := bson.M{"name": expenseCategoryName}

	_, err := collection.DeleteOne(global.Context, filter)

	if err != nil {
		panic(fmt.Sprintf("Doc delete issue %s", err))
	}

	fmt.Println("Document '" + expenseCategoryName + "' deleted successfully")
}
