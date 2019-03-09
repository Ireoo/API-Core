package conf

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoDB struct {
	Host      string        `yaml:"Host"`
	Auth      string        `yaml:"Auth"`
	Username  string        `yaml:"Username"`
	Password  string        `yaml:"Password"`
	Timeout   time.Duration `yaml:"Timeout"`
	PoolLimit int           `yaml:"PoolLimit"`
	URI       string        `yaml:"URI"`
}

type Input struct {
	Where bson.M `json:"where"`
	Data  bson.M `json:"data"`
	Other struct {
		page  int
		limit int
	} `json:"other"`
	Table string `json:"table"`
	Mode  string `json:"mode"`
	Auth  string `json:"auth"`
}

type AppInfo struct {
	Id bson.ObjectId `bson:"_id"`
}
