package conf

import (
	"time"

	"gopkg.in/mgo.v2/bson"
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
	Where string `json:"where" form:"where" query:"where"`
	Data  string `json:"data" form:"data" query:"data"`
	Other struct {
		Page  int
		Limit int
	} `json:"other" form:"other" query:"other"`
	Table string `json:"table" form:"table" query:"table"`
	Mode  string `json:"mode" form:"mode" query:"mode"`
	Auth  string `json:"auth" form:"auth" query:"auth"`
}

type AppInfo struct {
	Id bson.ObjectId `bson:"_id"`
}
