package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ClassDef struct {
	ID          int                `yaml:"id"`
	Classnumber int                `yaml:"classnumber"`
	Contents    []string           `yaml:"contents"`
	Assignment  []string           `yaml:"assignment"`
	Difficulty  map[string]float32 `yaml:"difficulty"`
}

func main() {
	f := &ClassDef{}
	source, err := ioutil.ReadFile("configGo.yaml")
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal([]byte(source), &f)
	if err != nil {
		log.Printf("error: %v", err)
	}
	fmt.Println(f)
	fmt.Println(f.Difficulty["SS"])
}
