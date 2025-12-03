package main

import (
	"log"

	"shikposh-backend/cmd/command"
)

func main() {
	command.Execute()
}

func init() {
	log.SetFlags(log.Lshortfile)
}
