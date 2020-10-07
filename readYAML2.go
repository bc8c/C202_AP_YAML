package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type MenuList struct {
	Coffee   []string `yaml:"COFFEE"`
	Icecream []string `yaml:"ICECREAM"`
}

type CoffeeShop struct {
	ID       int      `yaml:"id"`
	ShopName string   `yaml:"shopname"`
	Clerk    []string `yaml:"clerk"`
	Menu     MenuList `yaml:"menu"`
}

func main() {
	f := &CoffeeShop{}
	source, err := ioutil.ReadFile("configGo2.yaml")
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal([]byte(source), &f)
	if err != nil {
		log.Printf("error: %v", err)
	}
	fmt.Println(f)
	fmt.Println(f.Clerk[2])
	fmt.Println(f.Menu.Icecream)
}
