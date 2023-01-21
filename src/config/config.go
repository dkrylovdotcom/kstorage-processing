package config

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"log"
)

type Config struct {
	volumesCacheDirName string `yaml:"volumesCacheDirName"`
}

func Load(configPath string) {
	yfile, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]Config)
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(data["config"])

	for k, v := range data {
		fmt.Printf("%s -> %d\n", k, v)
	}
}

