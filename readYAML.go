package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Engineer struct {
	ID          string             `yaml:"id"`
	Sector      int                `yaml:"sector"`
	Tasks       []string           `yaml:"tasks"`
	DailyHourse []int              `yaml:"dailyHours"`
	Languages   map[string]float32 `yaml:"languages"`
}

func main() {
	f := &Engineer{}
	source, err := ioutil.ReadFile("configGo.yaml")
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal([]byte(source), &f)
	fmt.Println("hello world")
}
