package main

import (
	"service"
)

func main() {
	m := server.NewServer()
	m.Run(":8080")
}
