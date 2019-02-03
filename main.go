package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"./src/parser"
)

func readFile(path string) string {
	d, _ := ioutil.ReadFile(path)

	return string(d)
}

func absPath(s string) string {
	dir, err := filepath.Abs(s)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func printChildToken(x *parser.Token) {
	fmt.Println(x.ChildTokens[2].ChildTokens[0].ChildTokens[1])
}

func main() {

	file := absPath("./code.fjs")

	d := readFile(file)

	parser.Parse(file, d)

	// fmt.Println(tr)
}
