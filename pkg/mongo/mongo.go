package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/ClanceyLu/echo-api/conf"

	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect 连接 mongodb 并且返回 mongo.Client
func Connect() *m.Client {
	var (
		mongoConf = conf.Conf.Sub("mongo")
		host      = mongoConf.GetString("host")
		user      = mongoConf.GetString("user")
		password  = mongoConf.GetString("password")
	)

	connectURL := fmt.Sprintf("mongodb://%s:%s@%s", user, password, host)
	client, err := m.NewClient(options.Client().ApplyURI(connectURL))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return client
}
