package main

import (
	"log"

	mpsse "github.com/ardnew/gompsse"
)

func main() {

	if _, err := mpsse.NewMPSSE(); nil != err {
		log.Printf("err: %+v", err)
	}

}
