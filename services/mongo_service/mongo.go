package mongoservice

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoConnection connects to a mongo instance
func NewMongoConnection() *mongo.Client {
	url := fmt.Sprintf("mongodb+srv://%v:%v@%v/%v?retryWrites=true&w=majority",
		viper.GetString("MONGO.USER"),
		viper.GetString("MONGO.PASS"),
		viper.GetString("MONGO.HOST"),
		viper.GetString("MONGO.DB_NAME"))
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	log.Println("connected to database")

	return client
}
