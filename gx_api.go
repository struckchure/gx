package gx

import (
	"log"

	"github.com/a-h/rest"
)

var Api *rest.API

func ErrorIfNotApi() {
	if Api == nil {
		log.Fatalln("Setup method must be called before defining models")
	}
}
