package mongo

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/Ireoo/API-Core/libs/basic"
	"github.com/Ireoo/API-Core/libs/conf"

	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
)

var (
	globalS *mgo.Session
	// isInit  bool
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
		dialInfo := &mgo.DialInfo{
			Addrs:     []string{config.Host},
			Timeout:   config.Timeout * time.Millisecond,
			Source:    config.Auth,
			Username:  config.Username,
			Password:  config.Password,
			PoolLimit: config.PoolLimit,
			Mechanism: "LDAP",
		}
		globalS, err = mgo.DialWithInfo(dialInfo)
	} else {
		log.Printf(`Got config: URI: %s`, config.URI)
		globalS, err = mgo.Dial(config.URI)
	}

	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
	// isInit = true
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	// if !isInit {
	// 	_init()
	// }
	ms := globalS.Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

func IsEmpty(db, collection string) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(db, collection string, query interface{}) (int, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Count()
}

func Insert(db, collection string, doc interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Insert(doc)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
}

func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func Update(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Update(selector, update)
}

func Upsert(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

func UpdateAll(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

func Remove(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

func EnsureIndex(db, collection string, selector []string) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	index := mgo.Index{
		Key:        selector,
		Unique:     true,
		DropDups:   true,
		Background: true, // See notes.
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	return err
}

func Indexes(db, collection string) ([]mgo.Index, error) {
	ms, c := connect(db, collection)
	defer ms.Clone()

	return c.Indexes()
}

/**
 * 数据库操作
 */

func connectDB(db string) (*mgo.Session, *mgo.Database) {
	ms := globalS.Copy()
	d := ms.DB(db)
	ms.SetMode(mgo.Monotonic, true)
	return ms, d
}

func CollectionNames(db string) ([]string, error) {
	ms, d := connectDB(db)
	defer ms.Clone()

	return d.CollectionNames()
}

func AddUser(db string, username string, password string, readOnly bool) error {
	ms, d := connectDB(db)
	defer ms.Clone()

	err := d.AddUser(username, password, readOnly)
	return err
}

func DropDatabase(db string) error {
	ms, d := connectDB(db)
	defer ms.Clone()

	err := d.DropDatabase()
	return err
}
