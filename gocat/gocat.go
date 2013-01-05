package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bytes)
}
