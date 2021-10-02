package conf

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	Where bson.M `json:"where,omitempty" form:"where,omitempty" query:"where,omitempty"`
	Data  bson.M `json:"data,omitempty" form:"data,omitempty" query:"data,omitempty"`
	// Other string `json:"other" form:"other" query:"other"`
	Other struct {
		Page     int    `json:"page,omitempty" form:"page,omitempty" query:"page,omitempty"`
		Limit    int    `json:"limit,omitempty" form:"limit,omitempty" query:"limit,omitempty"`
		Distinct bson.M `json:"distinct" form:"distinct" query:"distinct"`
	} `json:"other,omitempty" form:"other,omitempty" query:"other,omitempty"`
	Table string `json:"table,omitempty" form:"table,omitempty" query:"table,omitempty"`
	Mode  string `json:"mode,omitempty" form:"mode,omitempty" query:"mode,omitempty"`
	Auth  string `json:"auth,omitempty" form:"auth,omitempty" query:"auth,omitempty"`
	App   string `json:"app,omitempty" form:"app,omitempty" query:"app,omitempty"`
}

type AppInfo struct {
	Id string `bson:"_id"`
}

type Result struct {
	Success bool
	Data    interface{}
}
