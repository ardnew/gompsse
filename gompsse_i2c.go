package gompsse

// #include "libMPSSE_i2c.h"
import "C"

// Constants controlling the various I2C communication options
const (
	// Generate start condition before transmitting
	i2cTransferOptionsStartBit = 0x00000001

	// Generate stop condition before transmitting
	i2cTransferOptionsStopBit = 0x00000002

	// Continue transmitting data in bulk without caring about Ack or nAck from
	// device if this bit is not set. If this bit is set then stop transitting the
	// data in the buffer when the device nAcks
	i2cTransferOptionsBreakOnNACK = 0x00000004

	// libMPSSE-I2C generates an ACKs for every byte read. Some I2C slaves require
	// the I2C master to generate a nACK for the last data byte read. Setting this
	// bit enables working with such I2C slaves
	i2cTransferOptionsNACKLastByte = 0x00000008

	// no address phase, no USB interframe delays
	i2cTransferOptionsFastTransferBytes = 0x00000010
	i2cTransferOptionsFastTransferBits  = 0x00000020
	i2cTransferOptionsFastTransfer      = 0x00000030

	// if I2CTransferOptionsFastTransfer is set then setting this bit would mean
	// that the address field should be ignored. The address is either a part of
	// the data or this is a special I2C frame that doesn't require an address
	i2cTransferOptionsNoAddress = 0x00000040

	i2cCmdGetdeviceidRD = 0xF9
	i2cCmdGetdeviceidWR = 0xF8

	i2cGiveACK  = 1
	i2cGiveNACK = 0
)

const (
	// 3-phase clocking is enabled by default. Setting this bit in ConfigOptions
	// will disable it
	i2c3PhaseClockingEnable  = 0x00000000
	i2c3PhaseClockingDisable = 0x00000001
	i2c3PhaseClockingDefault = i2c3PhaseClockingEnable

	// The I2C master should actually drive the SDA line only when the output is
	// LOW. It should tristate the SDA line when the output should be HIGH.
	i2cDriveOnlyZeroDisable = 0x00000000
	i2cDriveOnlyZeroEnable  = 0x00000002
	i2cDriveOnlyZeroDefault = i2cDriveOnlyZeroEnable
)

// I2CClockRate holds one of the supported I2C clock rate constants
type I2CClockRate uint32

// Constants defining the supported I2C clock rates
const (
	i2cClockStandardMode  I2CClockRate = 100000  // 100 kb/sec
	i2cClockFastMode      I2CClockRate = 400000  // 400 kb/sec
	i2cClockFastModePlus  I2CClockRate = 1000000 // 1000 kb/sec
	i2cClockHighSpeedMode I2CClockRate = 3400000 // 3.4 Mb/sec
	i2cClockDefault       I2CClockRate = i2cClockStandardMode
)

const (
	i2cLatencyDefault = 16
)

// i2cConfig holds all of the configuration settings for an I2C channel
type i2cConfig struct {
	clockRate I2CClockRate
	latency   uint8
	options   uint32
}

func i2cConfigDefault() *i2cConfig {
	return &i2cConfig{
		clockRate: i2cClockDefault,
		latency:   i2cLatencyDefault,
		options:   i2cDriveOnlyZeroDefault | i2c3PhaseClockingDefault,
	}
}

// I2C holds an active I2C channel through which all communication is performed.
type I2C struct {
	device *MPSSE
	config *i2cConfig
}
