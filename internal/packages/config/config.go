package config

var Port string
var Mongo *MongoConfig
var AuthenticationHeader string

func Init() {
	Port = ":8000"
	AuthenticationHeader = "Authorization"
	Mongo = &MongoConfig{
		URI: "mongodb://localhost:27017",
	}
}
