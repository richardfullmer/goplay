package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()

	Commander("git", "clone", "git@github.com:opensoft/simple-serializer", "opensoft/simple-serializer")
	os.Chdir("opensoft/simple-serializer")
	Commander("composer", "-v", "install", "--dev")
	Commander("php", "vendor/bin/phpunit")
	os.Chdir("../..")
	os.RemoveAll("opensoft")

}

// A command runner executes commands 
func Commander(name string, arg ...string) {

	cmd := exec.Command(name, arg...)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Print(err)
	}

	if err := cmd.Start(); err != nil {
		log.Print(err)
	}

	io.Copy(os.Stdout, stdout)

	if err := cmd.Wait(); err != nil {
		log.Print(err)

	}

}
