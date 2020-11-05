package cycapi

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoConn struct {
	Conn *mongo.Client
}

// CreateMongoConnection Creates a DB Connection with the Database specified on the url
func CreateMongoConnection(username, password, url string) (*MongoConn, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s@%s", "mongodb+srv", username, password, url)))
	if err != nil {
		return nil, err
	}

	connection := MongoConn{
		Conn: client,
	}

	return &connection, err
}

// CloseConnection Closes a DB Connection
func (db *MongoConn) CloseConnection() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return db.Conn.Disconnect(ctx)
}
