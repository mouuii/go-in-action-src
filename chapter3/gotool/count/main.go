package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/goinaction/code/chapter3/words"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("should pass at least one arg")
	}
	filename := os.Args[1]
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(contents)
	count := words.CountWords(text)
	fmt.Printf("There are %d word in this file", count)

}
