package mongo

import (
	"API-Core.go/libs/basic"
	"API-Core.go/libs/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"

	"github.com/globalsign/mgo"
)

var globalS *mgo.Session

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
PoolLimit: 4096`), 0666)
		log.Fatal(`api-core.conf is not exists.`)
	}

	config := new(conf.MongoDB)
	configYaml, err := ioutil.ReadFile("./api-core.conf")

	err = yaml.Unmarshal(configYaml, config)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	if config.Host == "" && config.URI == "" {
		log.Fatalln(config)
	}

	if config.URI == "" {
		log.Printf(`Got config: HOST: %s`, config.Host)
		log.Printf(`Got config: AUTH: %s`, config.Auth)
		log.Printf(`Got config: USERNAME: %s`, config.Username)
		log.Printf(`Got config: PASSWORD: %s`, config.Password)
		log.Printf(`Got config: TIMEOUT: %s`, config.Timeout*time.Millisecond)
		log.Printf(`Got config: POOLLIMIT: %d`, config.PoolLimit)
		dialInfo := &mgo.DialInfo{
			Addrs:     []string{config.Host},
			Timeout:   config.Timeout * time.Millisecond,
			Source:    config.Auth,
			Username:  config.Username,
			Password:  config.Password,
			PoolLimit: config.PoolLimit,
		}
		globalS, err = mgo.DialWithInfo(dialInfo)
	} else {
		log.Printf(`Got config: URI: %s`, config.URI)
		globalS, err = mgo.Dial(config.URI)
	}

	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
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
