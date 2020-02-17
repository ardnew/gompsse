package gompsse

// #cgo darwin,amd64 LDFLAGS: -framework CoreFoundation -framework IOKit
// #cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/native/darwin_amd64/lib
// #cgo darwin,amd64  CFLAGS: -I${SRCDIR}/native/darwin_amd64/inc
// #cgo linux,386    LDFLAGS: -L${SRCDIR}/native/linux_386/lib
// #cgo linux,386     CFLAGS: -I${SRCDIR}/native/linux_386/inc
// #cgo linux,amd64  LDFLAGS: -L${SRCDIR}/native/linux_amd64/lib
// #cgo linux,amd64   CFLAGS: -I${SRCDIR}/native/linux_amd64/inc
// #cgo linux,arm64  LDFLAGS: -L${SRCDIR}/native/linux_arm64/lib
// #cgo linux,arm64   CFLAGS: -I${SRCDIR}/native/linux_arm64/inc
// #cgo              LDFLAGS: -lMPSSE
// #include "ftd2xx.h"
import "C"

type FTHandle C.FT_HANDLE
type FTStatus C.FT_STATUS

type FTError int

// Constants related to device status
const (
	FTOK                      FTError = C.FT_OK
	FTInvalidHandle           FTError = C.FT_INVALID_HANDLE
	FTDeviceNotFound          FTError = C.FT_DEVICE_NOT_FOUND
	FTDeviceNotOpened         FTError = C.FT_DEVICE_NOT_OPENED
	FTIOError                 FTError = C.FT_IO_ERROR
	FTInsufficientResources   FTError = C.FT_INSUFFICIENT_RESOURCES
	FTInvalidParameter        FTError = C.FT_INVALID_PARAMETER
	FTInvalidBaudRate         FTError = C.FT_INVALID_BAUD_RATE
	FTDeviceNotOpenedForErase FTError = C.FT_DEVICE_NOT_OPENED_FOR_ERASE
	FTDeviceNotOpenedForWrite FTError = C.FT_DEVICE_NOT_OPENED_FOR_WRITE
	FTFailedToWriteDevice     FTError = C.FT_FAILED_TO_WRITE_DEVICE
	FTEEPROMReadFailed        FTError = C.FT_EEPROM_READ_FAILED
	FTEEPROMWriteFailed       FTError = C.FT_EEPROM_WRITE_FAILED
	FTEEPROMEraseFailed       FTError = C.FT_EEPROM_ERASE_FAILED
	FTEEPROMNotPresent        FTError = C.FT_EEPROM_NOT_PRESENT
	FTEEPROMNotProgrammed     FTError = C.FT_EEPROM_NOT_PROGRAMMED
	FTInvalidArgs             FTError = C.FT_INVALID_ARGS
	FTNotSupported            FTError = C.FT_NOT_SUPPORTED
	FTOtherError              FTError = C.FT_OTHER_ERROR
	FTDeviceListNotReady      FTError = C.FT_DEVICE_LIST_NOT_READY
)

// Constants related to GPIO pin configuration
const (
	PIN_LO byte = 0 // pin value clear
	PIN_HI byte = 1 // pin value set
	PIN_IN byte = 0 // pin direction input
	PIN_OT byte = 1 // pin direction output
)

type MPSSE struct {
	I2C *I2C
	SPI *SPI
}

func NewMPSSE() *MPSSE {
	return &MPSSE{I2C: nil, SPI: nil}
}
