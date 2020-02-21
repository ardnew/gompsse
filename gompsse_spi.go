package gompsse

// #include "libMPSSE_spi.h"
import "C"

import (
	"fmt"
)

// spiDPinConfig represents the default direction and value for pins associated
// with the lower byte lines of MPSSE, reserved for serial functions SPI/I²C
// (or port "D" on FT232H), but has a few GPIO pins as well.
type spiDPinConfig struct {
	initDir  byte // direction of lines after SPI channel initialization
	initVal  byte // value of lines after SPI channel initialization
	closeDir byte // direction of lines after SPI channel is closed
	closeVal byte // value of lines after SPI channel is closed
}

// Constants controlling the supported SPI transfer options
const (
	spiXferBytes = 0x00000000 // size is provided in bytes
	spiXferBits  = 0x00000001 // size is provided in bits
	spiCSEnable  = 0x00000002 // assert CS before start
	spiCSDisable = 0x00000004 // deassert CS after end
)

// Constants related to board pins when MPSSE operating in SPI mode
const (
	spiClockRateDefault = 12000000 // valid range: 0-30000000 (30 MHz)
	spiLatencyDefault   = 16       // 1-255 USB Hi-Speed, 2-255 USB Full-Speed
)

// Constants defining the available options in the SPI configuration struct.
const (
	// Known SPI operating modes
	//   LIMITATION: libMPSSE only supports mode 0 and mode 2 (CPHA==2).
	SPIMode0       = 0x00000000 // capture on RISE, propagate on FALL
	SPIMode1       = 0x00000001 // capture on FALL, propagate on RISE
	SPIMode2       = 0x00000002 // capture on FALL, propagate on RISE
	SPIMode3       = 0x00000003 // capture on RISE, propagate on FALL
	spiModeMask    = 0x00000003
	spiModeInvalid = 0x000000FF
	spiModeDefault = SPIMode0

	// DPins available for chip-select operation
	spiCSD3      = 0x00000000 // SPI CS on D3
	spiCSD4      = 0x00000004 // SPI CS on D4
	spiCSD5      = 0x00000008 // SPI CS on D5
	spiCSD6      = 0x0000000C // SPI CS on D6
	spiCSD7      = 0x00000010 // SPI CS on D7
	spiCSMask    = 0x0000001C
	spiCSInvalid = 0x000000FF
	spiCSDefault = spiCSD3

	// Other options
	spiCSActiveLow     = 0x00000020 // drive pin low to assert CS
	spiCSActiveHigh    = 0x00000000 // drive pin high to assert CS
	spiCSActiveDefault = spiCSActiveLow
)

// spiCSPin translates a DPin value to its corresponding chip-select mask for
// the SPI configuration struct option.
var spiCSPin = map[DPin]uint32{
	D0: spiCSInvalid,
	D1: spiCSInvalid,
	D2: spiCSInvalid,
	D3: spiCSD3,
	D4: spiCSD4,
	D5: spiCSD5,
	D6: spiCSD6,
	D7: spiCSD7,
}

// spiDPinConfigDefault defines the initial spiDPinConfig value for all pins
// represented by this type.
var spiDPinConfigDefault = [NumDPins]*spiDPinConfig{
	&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D0 SCLK
	&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D1 MOSI
	&spiDPinConfig{initDir: PinIN, initVal: PinLO, closeDir: PinIN, closeVal: PinLO}, // D2 MISO
	&spiDPinConfig{initDir: PinOT, initVal: PinHI, closeDir: PinOT, closeVal: PinHI}, // D3 CS
	&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D4 GPIO
	&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D5 GPIO
	&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D6 GPIO
	&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D7 GPIO
}

// spiConfig holds all of the configuration settings for an SPI channel.
type spiConfig struct {
	clockRate uint32 // in Hertz
	latency   uint8  // in ms
	options   uint32
	dpin      uint32
	reserved  uint16
}

// dpinConfig constructs the 32-bit pin field of the spiConfig struct from the
// provided spiDPinConfig slice cfg for each pin (identified by its index in the
// given slice).
func dpinConfig(cfg [NumDPins]*spiDPinConfig) uint32 {
	var dpin uint32
	for i, c := range cfg {
		dpin |= (uint32(c.initDir) << i) | (uint32(c.initVal) << (8 + i)) |
			(uint32(c.closeDir) << (16 + i)) | (uint32(c.closeVal) << (24 + i))
	}
	return dpin
}

type SPI struct {
	*deviceInfo
	config *spiConfig
}

type spiInfo struct {
}

// func SPIChannels() ([]*spiInfo, error) {
// 	var numChannels C.uint32
// 	if stat := Status(C.SPI_GetNumChannels(&numChannels)); !stat.OK() {
// 		return nil, stat
// 	}

// 	info := make([]C.FT_DEVICE_LIST_INFO_NODE, int(numChannels))
// 	channel := make([]*spiInfo, int(numChannels))

// 	for i := C.uint32(0); i < numChannels; i++ {
// 		if stat := Status(C.SPI_GetChannelInfo(i, &info[i])); !stat.OK() {
// 			continue
// 		}
// 		channel[i] = &spiInfo{
// 			DeviceInfo: NewDeviceInfo(&info[i]),
// 		}
// 	}
// 	return channel, nil
// }

func NewSPI(cs DPin) (*SPI, error) {

	var (
		spiCS uint32
		ok    bool
	)

	if spiCS, ok = spiCSPin[cs]; !ok {
		return nil, fmt.Errorf("invalid CS pin: %d", cs)
	}

	config := &spiConfig{
		clockRate: spiClockRateDefault,
		latency:   spiLatencyDefault,
		options:   spiCSActiveDefault | spiCS | spiModeDefault,
		dpin:      dpinConfig(spiDPinConfigDefault),
		reserved:  0,
	}

	return &SPI{config: config}, nil
}
