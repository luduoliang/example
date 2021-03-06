package common

import (
	"context"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient   *mongo.Client   // mongo客户端
	MongoDatabase *mongo.Database // mongo数据库
)

//建立mongo连接
func InitMongo(mongodbUrl string, databaseName string) (err error) {
	mongodbUrl = strings.TrimSpace(mongodbUrl)
	databaseName = strings.TrimSpace(databaseName)
	if mongodbUrl == "" {
		err = errors.New("mongodbUrl empty error")
		return
	}
	if databaseName == "" {
		err = errors.New("databaseName empty error")
		return
	}
	// 连接mongo
	MongoClient, err = ConenctMongo(mongodbUrl)
	if err != nil {
		return
	}
	MongoDatabase = MongoClient.Database(databaseName)
	if MongoDatabase == nil {
		err = errors.New("Database select error")
		return
	}
	return nil
}

// 连接mongo数据库
func ConenctMongo(url string) (client *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return
	}
	return
}

//关闭mongo数据库
func CloseMongo(client *mongo.Client) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	if err = client.Disconnect(ctx); err != nil {
		return
	}
	return nil
}

//插入一条记录
func AddOneMongo(tableName string, data interface{}) error {
	if MongoDatabase == nil {
		return errors.New("MongoDatabase is nil")
	}
	mongoCollection := MongoDatabase.Collection(tableName)
	if mongoCollection == nil {
		return errors.New("mongoCollection create error")
	}

	_, err := mongoCollection.InsertOne(context.TODO(), &data)
	return err
}

//查询一条记录
//filter示例：bson.M{"createdat": bson.M{"$gte": start, "$lte": end}}
//filter示例：bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
func FindOneMongo(tableName string, filter bson.M) (interface{}, error) {
	var info interface{}
	if MongoDatabase == nil {
		return nil, errors.New("MongoDatabase is nil")
	}
	mongoCollection := MongoDatabase.Collection(tableName)
	if mongoCollection == nil {
		return nil, errors.New("mongoCollection create error")
	}
	err := mongoCollection.FindOne(context.Background(), filter).Decode(&info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

//修改单条记录
//filter示例：bson.M{"createdat": bson.M{"$gte": start, "$lte": end}}
//filter示例：bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
func UpdateOneMongo(tableName string, filter bson.M, updateData interface{}) error {
	if MongoDatabase == nil {
		return errors.New("MongoDatabase is nil")
	}
	mongoCollection := MongoDatabase.Collection(tableName)
	if mongoCollection == nil {
		return errors.New("mongoCollection create error")
	}

	_, err := mongoCollection.UpdateOne(context.Background(), filter, updateData)
	return err
}

//查多条记录
//filter示例：bson.M{"createdat": bson.M{"$gte": start, "$lte": end}}
//filter示例：bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
//用cursor.All(&list)解析数据，list := []People
func FindMoreItems(tableName string, filter bson.M) (*mongo.Cursor, error) {
	if MongoDatabase == nil {
		return nil, errors.New("MongoDatabase is nil")
	}
	mongoCollection := MongoDatabase.Collection(tableName)
	if mongoCollection == nil {
		return nil, errors.New("mongoCollection create error")
	}
	ctx := context.Background()
	cursor, err := mongoCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}
