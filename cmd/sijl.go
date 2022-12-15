package main

import (
	"log"

	"github.com/CSC354/sijl/internal/serv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	serv.StartSijlServer()
}
