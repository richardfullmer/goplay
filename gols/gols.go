package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var listArg bool

func init() {
	flag.BoolVar(&listArg, "l", false, "Show output one line per file")
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
	} else {

		list, err := ioutil.ReadDir(flag.Arg(0))

		if err != nil {
			log.Fatal(err)
		}
		for _, l := range list {
			fmt.Printf("%s ", l.Name())
			if listArg == true {
				fmt.Println()
			}
		}
		fmt.Println()
	}
}
