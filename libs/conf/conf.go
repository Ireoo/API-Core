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
	Where struct {
		_id bson.ObjectId `json:"_id" form:"_id" query:"_id"`
	} `json:"where" form:"where" query:"where"`
	Data struct {
		_id bson.ObjectId `json:"_id" form:"_id" query:"_id"`
	} `json:"data" form:"data" query:"data"`
	// Other string `json:"other" form:"other" query:"other"`
	Other struct {
		Page     int    `json:"page" form:"page" query:"page"`
		Limit    int    `json:"limit" form:"limit" query:"limit"`
		Distinct bson.M `json:"distinct" form:"distinct" query:"distinct"`
	} `json:"other" form:"other" query:"other"`
	Table string `json:"table" form:"table" query:"table"`
	Mode  string `json:"mode" form:"mode" query:"mode"`
	Auth  string `json:"auth" form:"auth" query:"auth"`
	App   string `json:"app" form:"app" query:"app"`
}

type AppInfo struct {
	Id bson.ObjectId `bson:"_id"`
}

type Result struct {
	Success bool
	Data    interface{}
}
