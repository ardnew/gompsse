package main

import (
	"log"

	mpsse "github.com/ardnew/gompsse"
)

func main() {

	for _, desc := range []string{"FT232H", "FT232H-C"} {
		m, err := mpsse.NewMPSSEWithDesc(desc)
		if nil != err {
			log.Printf("NewMPSSEWithDesc(): %+v", err)
			continue
		}
		defer m.Close()
		if err := m.SPI.Init(); nil != err {
			log.Printf("SPI.Init(): %+v", err)
		}
		log.Printf("%s", m)
	}
}
