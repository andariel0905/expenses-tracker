package global

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client
var Context context.Context

func SetGlobalVariables(mainClient *mongo.Client, mainCxt context.Context) {
	Client = mainClient
	Context = mainCxt
}
