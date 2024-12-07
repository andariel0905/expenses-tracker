package models

type ExpenseCategory struct {
	//	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
