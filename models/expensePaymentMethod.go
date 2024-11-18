package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExpensePaymentMethod struct {
	Id primitive.ObjectID	`json:"id" bson:"_id"`
}