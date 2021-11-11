package main

import (
	"fmt"
	"sotater/copy-opgg-go/pkg/server"
)

func main() {
	fmt.Println("main fun started")
	server.GetServer().Run()
}
