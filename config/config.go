package config

import (
	"github.com/goccy/go-yaml"
	"io/ioutil"
	"log"
)

type ConfT struct {
	Db       map[string]interface{}
	Security map[string]interface{}
}

var data []byte

func Config() ConfT {
	if len(data) == 0 {
		readData, err := ioutil.ReadFile("config/app.yaml")
		if err != nil {
			panic(err)
		}
		data = readData
	}

	m := ConfT{}
	err := yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return m
}
