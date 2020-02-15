package gompsse

// #include "libMPSSE_i2c.h"
import "C"

// Constants controlling the various I2C communication options
const (
	// Generate start condition before transmitting
	I2CTransferOptionsStartBit = 0x00000001

	// Generate stop condition before transmitting
	I2CTransferOptionsStopBit = 0x00000002

	// Continue transmitting data in bulk without caring about Ack or nAck from
	// device if this bit is not set. If this bit is set then stop transitting the
	// data in the buffer when the device nAcks
	I2CTransferOptionsBreakOnNACK = 0x00000004

	// libMPSSE-I2C generates an ACKs for every byte read. Some I2C slaves require
	// the I2C master to generate a nACK for the last data byte read. Setting this
	// bit enables working with such I2C slaves
	I2CTransferOptionsNACKLastByte = 0x00000008

	// no address phase, no USB interframe delays
	I2CTransferOptionsFastTransferBytes = 0x00000010
	I2CTransferOptionsFastTransferBits  = 0x00000020
	I2CTransferOptionsFastTransfer      = 0x00000030

	// if I2CTransferOptionsFastTransfer is set then setting this bit would mean
	// that the address field should be ignored. The address is either a part of
	// the data or this is a special I2C frame that doesn't require an address
	I2CTransferOptionsNoAddress = 0x00000040

	I2CCmdGetdeviceidRD = 0xF9
	I2CCmdGetdeviceidWR = 0xF8

	I2CGiveACK  = 1
	I2CGiveNACK = 0

	// 3-phase clocking is enabled by default. Setting this bit in ConfigOptions
	// will disable it
	I2CDisable3phaseClocking = 0x0001

	// The I2C master should actually drive the SDA line only when the output is
	// LOW. It should be tristate the SDA line when the output should be high.
	// This tristating the SDA line during output HIGH is supported only in FT232H
	// chip. This feature is called DriveOnlyZero feature and is enabled when the
	// following bit is set in the options parameter in function I2C_Init
	I2CEnableDriveOnlyZero = 0x0002
)

// I2CClockRate holds one of the supported I2C clock rate constants
type I2CClockRate uint32

// Constants defining the supported I2C clock rates
const (
	I2CClockStandardMode  I2CClockRate = 100000  // 100 kb/sec
	I2CClockFastMode      I2CClockRate = 400000  // 400 kb/sec
	I2CClockFastModePlus  I2CClockRate = 1000000 // 1000 kb/sec
	I2CClockHighSpeedMode I2CClockRate = 3400000 // 3.4 Mb/sec
)

// I2CConfig holds all of the configuration settings for an I2C channel
type I2CConfig struct {
	ClockRate I2CClockRate
	Latency   uint8
	Options   uint32
}

// I2C holds an active I2C channel through which all communication is performed.
type I2C struct {
	Config I2CConfig
}
