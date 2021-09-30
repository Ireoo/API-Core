package mongo

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Ireoo/API-Core/libs/basic"
	"github.com/Ireoo/API-Core/libs/conf"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

var (
	uri    string
	client *mongo.Client
)

func init() {
	exist, _ := basic.PathExists("./api-core.conf")
	if !exist {
		_ = ioutil.WriteFile("./api-core.conf", []byte(`# This is a api-core config file.
# This is demo.

Host: "127.0.0.1:27017"
Auth: "admin"
Username: "root"
Password: "meiyoumima"
Timeout: 200
PoolLimit: 4096
#URI: "127.0.0.1:27017" # 如果不通过此URI直接连接，请删掉或者注释掉`), 0666)

		log.Println(`api-core.conf is not exists.`)
		log.Fatal("You can use \033[1mvim ./api-core.conf\033[0m to change MongoDB connect.")
	}

	config := new(conf.MongoDB)
	configYaml, err := ioutil.ReadFile("./api-core.conf")
	if err != nil {
		log.Fatalf("ReadFile: %v", err)
	}

	err = yaml.Unmarshal(configYaml, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Print(err)
		if config.Host == "" && config.URI == "" {
			log.Fatalln(config)
		}

		if config.URI == "" {
			log.Printf(`Got config: HOST: %q`, config.Host)
			log.Printf(`Got config: AUTH: %q`, config.Auth)
			log.Printf(`Got config: USERNAME: %q`, config.Username)
			log.Printf(`Got config: PASSWORD: %q`, config.Password)
			log.Printf(`Got config: TIMEOUT: %s`, config.Timeout*time.Millisecond)
			log.Printf(`Got config: POOLLIMIT: %d`, config.PoolLimit)
			uri = "mongodb://" + config.Username + ":" + config.Password + "@" + config.Host + "/" + config.Auth
		} else {
			log.Printf(`Got config: URI: %s`, config.URI)
			uri = config.URI
		}
	} else {
		uri = os.Getenv("MONGODB_URI")
		log.Printf(`Got config: URI: %s`, uri)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Create Session: %s\n", err)
		}
	}()
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

func FindOne(db, collection string, query, selector, result interface{}) error {
	c := connect(db, collection)

	return c.FindOne(context.TODO(), query).Decode(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	c := connect(db, collection)

	result, err := c.Find(context.TODO(), query)
	return err
}

func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	c := connect(db, collection)

	// return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
	result, err := c.Find(context.TODO(), query)
	return err
}

func Update(db, collection string, selector, update interface{}) error {
	c := connect(db, collection)

	_, err := c.UpdateOne(context.TODO(), selector, update)
	return err
}

func Upsert(db, collection string, selector, update interface{}) error {
	c := connect(db, collection)

	opts := options.Update().SetUpsert(true)
	_, err := c.UpdateOne(context.TODO(), selector, update, opts)
	return err
}

func UpdateAll(db, collection string, selector, update interface{}) error {
	c := connect(db, collection)

	_, err := c.UpdateMany(context.TODO(), selector, update)
	return err
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

func Indexes(db, collection string) *mongo.Cursor {
	c := connect(db, collection)

	cur, _ := c.Indexes().List(context.TODO())
	return cur
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

	return d.ListCollectionNames(context.TODO(), bson.M{}) //.CollectionNames()
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
