package config

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Port     int
	MongoUrl string
}

func (receiver *Config) InitMongoUrl() error {
	mongoUrl := os.Getenv("MONGO_URL")
	if len(mongoUrl) > 0 {
		receiver.MongoUrl = mongoUrl
		return nil
	}

	data, err := ioutil.ReadFile("./config/.mongourl")
	if err != nil {
		return fmt.Errorf("error while reading mongourl \n\t%v", err)
	}
	receiver.MongoUrl = string(data)
	err = os.Setenv("MONGO_URL", string(data))
	if err != nil {
		return fmt.Errorf("error while setting os variable MONGO_URL\n\t%v", err)
	}
	return nil
}
