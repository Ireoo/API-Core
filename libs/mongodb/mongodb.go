package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Ireoo/API-Core/libs/basic"
	"github.com/Ireoo/API-Core/libs/conf"

	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/yaml.v2"
	// "testing" // Remove this line
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
		// log.Printf(`Got config: HOST: %q`, config.Host)
		// log.Printf(`Got config: AUTH: %q`, config.Auth)
		// log.Printf(`Got config: USERNAME: %q`, config.Username)
		// log.Printf(`Got config: PASSWORD: %q`, config.Password)
		// log.Printf(`Got config: TIMEOUT: %s`, config.Timeout*time.Millisecond)
		// log.Printf(`Got config: POOLLIMIT: %d`, config.PoolLimit)
		uri = "mongodb://" + config.Username + ":" + config.Password + "@" + config.Host + "/" + config.Auth
	}
	if config.URI != "" {
		// log.Printf(`Got config: URI: %s`, config.URI)
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

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout*time.Millisecond)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatalf("Create Session: %s\n", err)
	// 	}
	// }()
	return err
}

func connect(db, collection string) *mongo.Collection {

	c := client.Database(db).Collection(collection)
	return c
}

func IsEmpty(db, collection string) bool {
	c := connect(db, collection)

	count, err := c.EstimatedDocumentCount(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(db, collection string, query interface{}) (int64, error) {
	c := connect(db, collection)

	return c.CountDocuments(context.TODO(), query)
}

func Insert(db, collection string, doc interface{}) (*mongo.InsertOneResult, error) {
	c := connect(db, collection)

	return c.InsertOne(context.TODO(), doc)
}

func FindOne(db, collection string, query interface{}, other *conf.Other, result interface{}) error {
	c := connect(db, collection)

	opts := options.FindOne()
	opts.SetSkip(other.Skip)
	if other.Show != nil {
		opts.SetProjection(other.Show)
	}

	return c.FindOne(context.TODO(), query, opts).Decode(result)
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

	Cursor, err := c.Find(context.TODO(), query, opts)
	if err != nil {
		return []bson.M{}, err
	}
	var result []bson.M
	err = Cursor.All(context.TODO(), &result)
	return result, err
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

	Cursor, err := c.Find(context.TODO(), query, opts)
	if err != nil {
		return []bson.M{}, err
	}
	var result []bson.M
	err = Cursor.All(context.TODO(), &result)
	return result, err
}

func Update(db, collection string, where, update interface{}, other *conf.Other) error {
	c := connect(db, collection)

	opts := options.Update()
	opts.SetUpsert(other.Upsert)
	if other.Multi {
		_, err := c.UpdateMany(context.TODO(), where, update, opts)
		return err
	} else {
		_, err := c.UpdateOne(context.TODO(), where, update, opts)
		return err
	}
}

func Remove(db, collection string, selector interface{}) error {
	c := connect(db, collection)

	_, err := c.DeleteOne(context.TODO(), selector)
	return err
}

func RemoveAll(db, collection string, selector interface{}) error {
	c := connect(db, collection)

	_, err := c.DeleteMany(context.TODO(), selector)
	return err
}

func EnsureIndex(db, collection string, selector []string) error {
	c := connect(db, collection)

	index := mongo.IndexModel{
		Keys:    selector,
		Options: options.Index().SetUnique(true),
	}
	_, err := c.Indexes().CreateOne(context.TODO(), index)
	return err
}

func Indexes(db, collection string) []bson.M {
	c := connect(db, collection)

	cur, _ := c.Indexes().List(context.TODO())
	var result []bson.M
	_ = cur.All(context.TODO(), &result)
	return result
}

/**
 * 数据库操作
 */

func connectDB(db string) *mongo.Database {
	d := client.Database(db)
	return d
}

func CollectionNames(db string) ([]string, error) {
	d := connectDB(db)

	opts := options.ListCollections()
	return d.ListCollectionNames(context.TODO(), bson.M{}, opts) //.CollectionNames()
}

// func AddUser(db string, username string, password string, readOnly bool) error {
// 	d := connectDB(db)

// 	err := d //.AddUser(username, password, readOnly)
// 	return err
// }

func DropDatabase(db string) error {
	d := connectDB(db)

	err := d.Drop(context.TODO()) //.DropDatabase()
	return err
}
