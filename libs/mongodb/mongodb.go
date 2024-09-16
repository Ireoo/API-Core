package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Ireoo/API-Core/libs/basic"
	"github.com/Ireoo/API-Core/libs/conf"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/yaml.v2"
)

var (
	uri    string
	client *mongo.Client
)

func New(command_uri string) error {
	exist, _ := basic.PathExists("./api-core.conf")
	if !exist {
		_ = os.WriteFile("./api-core.conf", []byte(`# This is a api-core config file.
# This is demo.

Host: "127.0.0.1:27017"
Auth: "admin"
Username: "root"
Password: "meiyoumima"
Timeout: 20000
PoolLimit: 4096
#URI: "127.0.0.1:27017" # 如果不通过此URI直接连接，请删掉或者注释掉`), 0666)

		fmt.Println(`api-core.conf is not exists.`)
		fmt.Println("You can use \033[1mvim ./api-core.conf\033[0m to change MongoDB connect.")
		os.Exit(0)
	}

	config := new(conf.MongoDB)
	configYaml, err := os.ReadFile("./api-core.conf")
	if err != nil {
		log.Fatalf("ReadFile: %v", err)
	}

	err = yaml.Unmarshal(configYaml, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	if config.Host != "" {
		uri = "mongodb://" + config.Username + ":" + config.Password + "@" + config.Host + "/" + config.Auth
	}
	if config.URI != "" {
		uri = config.URI
	}
	if os.Getenv("MONGODB_URI") != "" {
		uri = os.Getenv("MONGODB_URI")
	}
	if command_uri != "" {
		uri = command_uri
	}

	if uri == "" {
		log.Fatalln(`Can't get MongoDB connect uri!`)
	} else {
		log.Printf(`Got MongoDB Connect URI: %s`, uri)
	}

	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout*time.Millisecond)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
	return err
}

func connect(db, collection string) *mongo.Collection {
	return client.Database(db).Collection(collection)
}

func IsEmpty(db, collection string) bool {
	c := connect(db, collection)

	count, err := c.EstimatedDocumentCount(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(db, collection string, query interface{}) (int64, error) {
	c := connect(db, collection)

	return c.CountDocuments(context.Background(), query)
}

func Insert(db, collection string, doc interface{}) (*mongo.InsertOneResult, error) {
	c := connect(db, collection)

	return c.InsertOne(context.Background(), doc)
}

func FindOne(db, collection string, query interface{}, other *conf.Other, result interface{}) error {
	c := connect(db, collection)

	opts := options.FindOne()
	opts.SetSkip(other.Skip)
	if other.Show != nil {
		opts.SetProjection(other.Show)
	}

	return c.FindOne(context.Background(), query, opts).Decode(result)
}

func FindAll(db, collection string, query interface{}, other *conf.Other) ([]bson.M, error) {
	c := connect(db, collection)

	opts := options.Find()
	opts.SetSkip(other.Skip)
	opts.SetLimit(other.Limit)
	if other.SortFormat != nil {
		opts.SetSort(other.SortFormat)
	}
	if other.Show != nil {
		opts.SetProjection(other.Show)
	}

	Cursor, err := c.Find(context.Background(), query, opts)
	if err != nil {
		return nil, err
	}
	var result []bson.M
	if err = Cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func FindPage(db, collection string, other *conf.Other, query interface{}) ([]bson.M, error) {
	c := connect(db, collection)

	opts := options.Find()
	opts.SetSkip(other.Skip)
	opts.SetLimit(other.Limit)
	if other.SortFormat != nil {
		opts.SetSort(other.SortFormat)
	}
	if other.Show != nil {
		opts.SetProjection(other.Show)
	}

	Cursor, err := c.Find(context.Background(), query, opts)
	if err != nil {
		return nil, err
	}
	var result []bson.M
	if err = Cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func Update(db, collection string, where, update interface{}, other *conf.Other) error {
	c := connect(db, collection)

	opts := options.Update()
	opts.SetUpsert(other.Upsert)
	if other.Multi {
		_, err := c.UpdateMany(context.Background(), where, update, opts)
		return err
	} else {
		_, err := c.UpdateOne(context.Background(), where, update, opts)
		return err
	}
}

func Remove(db, collection string, selector interface{}) error {
	c := connect(db, collection)

	_, err := c.DeleteOne(context.Background(), selector)
	return err
}

func RemoveAll(db, collection string, selector interface{}) error {
	c := connect(db, collection)

	_, err := c.DeleteMany(context.Background(), selector)
	return err
}

func EnsureIndex(db, collection string, selector []string) error {
	c := connect(db, collection)

	index := mongo.IndexModel{
		Keys:    selector,
		Options: options.Index().SetUnique(true),
	}
	_, err := c.Indexes().CreateOne(context.Background(), index)
	return err
}

func Indexes(db, collection string) []bson.M {
	c := connect(db, collection)

	cur, _ := c.Indexes().List(context.Background())
	var result []bson.M
	_ = cur.All(context.Background(), &result)
	return result
}

func connectDB(db string) *mongo.Database {
	return client.Database(db)
}

func CollectionNames(db string) ([]string, error) {
	d := connectDB(db)

	opts := options.ListCollections()
	return d.ListCollectionNames(context.Background(), bson.M{}, opts)
}

func DropDatabase(db string) error {
	d := connectDB(db)

	return d.Drop(context.Background())
}
