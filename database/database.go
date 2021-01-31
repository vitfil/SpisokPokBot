package database

import (
	"context"

	logger "est-server/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbName string
	client *mongo.Client
)

func Init() {
}

func Database() *mongo.Database {
	return client.Database(dbName)
}

func Collection(name string) *mongo.Collection {
	return client.Database(dbName).Collection(name)
}

func Connect(connString, name string) error {
	var err error

	dbName = name

	if connString[:10] != "mongodb://" {
		connString = "mongodb://" + connString
	}
	opt := options.Client().
		ApplyURI(connString)

	logger.WithFields(logger.Fields{
		"ConnectionString": connString,
		"DatabaseName":     dbName,
	}).Debug("Try to connect database")

	client, err = mongo.NewClient(opt)
	if err != nil {
		logger.WithFields(logger.Fields{
			"Error": err,
		}).Fatal("Connect to database failed")
		return err
	}

	err = client.Connect(context.Background())
	if err != nil {
		logger.WithFields(logger.Fields{
			"Error": err,
		}).Fatal("Connect to database failed")
		return err
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		logger.WithFields(logger.Fields{
			"Error": err,
		}).Fatal("Connect to database failed")
		return err
	}

	return nil
}

func Disconnect() {

	if err := client.Disconnect(context.Background()); err != nil {
		logger.WithFields(logger.Fields{
			"Error": err,
		}).Error("Disconnect from database")
	} else {
		logger.Debug("Disconnect from database")
	}
}
