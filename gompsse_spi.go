package gompsse

// #include "libMPSSE_i2c.h"
import "C"

// Constants controlling the various SPI communication options
const (
	// transferOptions-Bit0: If this bit is 0 then it means that the transfer size
	// provided is in bytes
	SPITransferOptionsSizeInBytes = 0x00000000
	// transferOptions-Bit0: If this bit is 1 then it means that the transfer size
	// provided is in bytes
	SPITransferOptionsSizeInBits = 0x00000001
	// transferOptions-Bit1: if BIT1 is 1 then CHIP_SELECT line will be enabled at
	// start of transfer
	SPITransferOptionsChipselectEnable = 0x00000002
	// transferOptions-Bit2: if BIT2 is 1 then CHIP_SELECT line will be disabled
	// at end of transfer
	SPITransferOptionsChipselectDisable = 0x00000004
)

// Constants defining the supported SPI modes
const (
	SPIConfigOptionModeMask = 0x00000003
	SPIConfigOptionMode0    = 0x00000000
	SPIConfigOptionMode1    = 0x00000001
	SPIConfigOptionMode2    = 0x00000002
	SPIConfigOptionMode3    = 0x00000003
)

// Constants defining the supported chip-select pins and options
const (
	SPIConfigOptionCSMask  = 0x0000001C // 111 00
	SPIConfigOptionCSDBUS3 = 0x00000000 // 000 00
	SPIConfigOptionCSDBUS4 = 0x00000004 // 001 00
	SPIConfigOptionCSDBUS5 = 0x00000008 // 010 00
	SPIConfigOptionCSDBUS6 = 0x0000000C // 011 00
	SPIConfigOptionCSDBUS7 = 0x00000010 // 100 00

	SPIConfigOptionCSActivelow = 0x00000020
)

// SPIConfig holds all of the configuration settings for an SPI channel.
type SPIConfig struct {
	ClockRate uint32
	Latency   uint8
	// This member provides a way to enable/disable features specific to the
	// protocol that are implemented in the chip
	//  Bits 1-0   (CPOL/CPHA)
	//    00 - MODE0 - data captured on rising edge, propagated on falling
	//    01 - MODE1 - data captured on falling edge, propagated on rising
	//    10 - MODE2 - data captured on falling edge, propagated on rising
	//    11 - MODE3 - data captured on rising edge, propagated on falling
	//  Bits 4-2   (chip-select)
	//    000 - A/B/C/D_DBUS3
	//    001 - A/B/C/D_DBUS4
	//    010 - A/B/C/D_DBUS5
	//    011 - A/B/C/D_DBUS6
	//    100 - A/B/C/D_DBUS7
	//  Bit  5     (chip-select is active high if this bit is 0)
	//  Bits 6-31  (reserved)
	Options uint32
	//  Bits  7- 0 (Initial direction of the pins)
	//  Bits 15- 8 (Initial values of the pins)
	//  Bits 23-16 (Final direction of the pins)
	//  Bits 31-24 (Final values of the pins)
	Pin      uint32
	Reserved uint16
}

type SPI struct {
	Config SPIConfig
}
