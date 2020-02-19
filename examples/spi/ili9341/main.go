package main

import (
	"log"

	mpsse "github.com/ardnew/gompsse"
)

func main() {

	var (
		spi []*mpsse.SPIChannelInfo
		err error
	)
	if spi, err = mpsse.SPIChannels(); nil != err {
		log.Fatalf("SPIChannels(): %s", err)
	}

	for i, d := range spi {
		log.Printf("%d: %+v", i, d)
	}

	log.Printf("exiting %d", len(spi))
}
