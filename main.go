package main

import (
	"fmt"
)

func main() {
	conf, err := LoadConfig(".")

	if err != nil {
		fmt.Println("Error occured while loading config", err)
	}

	StartGin(conf)
}
