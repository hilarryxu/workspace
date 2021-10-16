// Package main provides ...
package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "默认name", "帮助信息")
	flag.StringVar(&name, "n", "默认name", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
}
