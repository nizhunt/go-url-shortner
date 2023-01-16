package main

import (
	"github.com/nizhunt/urlShortner/model"
	"github.com/nizhunt/urlShortner/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}