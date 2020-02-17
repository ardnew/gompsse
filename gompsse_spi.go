package gompsse

// #include "libMPSSE_spi.h"
import "C"

// Constants controlling the supported SPI transfer options
const (
	SPITransferOptionsSizeInBytes       = 0x00000000 // size is provided in bytes
	SPITransferOptionsSizeInBits        = 0x00000001 // size is provided in bits
	SPITransferOptionsChipselectEnable  = 0x00000002 // assert CS before start
	SPITransferOptionsChipselectDisable = 0x00000004 // deassert CS after end
)

const SPIClockRateDefault = 100000 // valid range: 0-30000000 (30 MHz)
const SPILatencyTimer = value      // 1-255 USB Hi-Speed, 2-255 USB Full-Speed

// Constants defining the supported SPI modes.
//   NOTE: the libMPSSE engine only supports mode 0 and mode 2 (CPHA==2).
const (
	SPIConfigOptionModeMask = 0x00000003
	SPIConfigOptionMode0    = 0x00000000 // capture on RISE, propagate on FALL
	SPIConfigOptionMode1    = 0x00000001 // capture on FALL, propagate on RISE
	SPIConfigOptionMode2    = 0x00000002 // capture on FALL, propagate on RISE
	SPIConfigOptionMode3    = 0x00000003 // capture on RISE, propagate on FALL

	SPIConfigOptionModeDefault = SPIConfigOptionMode0
)

// Constants defining the supported chip-select pins and options
const (
	SPIConfigOptionCSMask  = 0x0000001C
	SPIConfigOptionCSDBUS3 = 0x00000000 // CS on D3
	SPIConfigOptionCSDBUS4 = 0x00000004 // CS on D4
	SPIConfigOptionCSDBUS5 = 0x00000008 // CS on D5
	SPIConfigOptionCSDBUS6 = 0x0000000C // CS on D6
	SPIConfigOptionCSDBUS7 = 0x00000010 // CS on D7

	SPIConfigOptionsCSDefault = SPIConfigOptionCSDBUS3

	SPIConfigOptionCSActivelow = 0x00000020 // drive pin low to assert CS
)

// Constants related to board pins when MPSSE operating in SPI mode
const (
	DBUSNumPins = 8
	CBUSNumPins = 8
)

// SPIPinConfig represents the default direction and value for pins associated
// with the lower byte lines of MPSSE, reserved for serial functions SPI/I²C
// (or port "D" on FT232H)
type SPIPinConfig struct {
	InitDir  byte // direction of lines after SPI channel initialization
	InitVal  byte // value of lines after SPI channel initialization
	CloseDir byte // direction of lines after SPI channel is closed
	CloseVal byte // value of lines after SPI channel is closed
}

// DefaultSPIPinConfig defines the initial SPIPinConfig value for all pins
// represented by this type.
var DefaultSPIPinConfig = [DBUSNumPins]*SPIPinConfig{
	&SPIPinConfig{InitDir: PIN_OT, InitVal: PIN_LO, CloseDir: PIN_OT, CloseVal: PIN_LO}, // D0 SCLK
	&SPIPinConfig{InitDir: PIN_OT, InitVal: PIN_LO, CloseDir: PIN_OT, CloseVal: PIN_LO}, // D1 MOSI
	&SPIPinConfig{InitDir: PIN_IN, InitVal: PIN_LO, CloseDir: PIN_IN, CloseVal: PIN_LO}, // D2 MISO
	&SPIPinConfig{InitDir: PIN_OT, InitVal: PIN_HI, CloseDir: PIN_OT, CloseVal: PIN_HI}, // D3 CS
	&SPIPinConfig{InitDir: PIN_OT, InitVal: PIN_LO, CloseDir: PIN_OT, CloseVal: PIN_LO}, // D4 GPIO
	&SPIPinConfig{InitDir: PIN_OT, InitVal: PIN_LO, CloseDir: PIN_OT, CloseVal: PIN_LO}, // D5 GPIO
	&SPIPinConfig{InitDir: PIN_OT, InitVal: PIN_LO, CloseDir: PIN_OT, CloseVal: PIN_LO}, // D6 GPIO
	&SPIPinConfig{InitDir: PIN_OT, InitVal: PIN_LO, CloseDir: PIN_OT, CloseVal: PIN_LO}, // D7 GPIO
}

// SPIConfig holds all of the configuration settings for an SPI channel.
type SPIConfig struct {
	ClockRate uint32 // in Hertz
	Latency   uint8  // in ms
	Options   uint32
	Pin       uint32
	Reserved  uint16
}

// SetPin constructs the 32-bit Pin field of the SPIConfig struct from the
// provided SPIPinConfig slice cfg for each pin (identified by its index in the
// given slice).
func (sc *SPIConfig) SetPin(cfg []*SPIPinConfig) {

	sc.Pin = 0
	for i, c := range cfg {
		sc.Pin |= (uint32(c.InitDir) << i) | (uint32(c.InitVal) << (8 + i)) |
			(uint32(c.CloseDir) << (16 + i)) | (uint32(c.CloseVal) << (24 + i))
	}
}

type SPI struct {
	Config SPIConfig
}
