package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExpenseCategory struct {
	Id primitive.ObjectID	`json:"id" bson:"_id"`
}
