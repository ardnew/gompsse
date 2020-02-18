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

func (e FTError) Error() string {
	switch e {
	case FTOK:
		return "OK"
	case FTInvalidHandle:
		return "invalid handle"
	case FTDeviceNotFound:
		return "device not found"
	case FTDeviceNotOpened:
		return "device not opened"
	case FTIOError:
		return "IO error"
	case FTInsufficientResources:
		return "insufficient resources"
	case FTInvalidParameter:
		return "invalid parameter"
	case FTInvalidBaudRate:
		return "invalid baud rate"
	case FTDeviceNotOpenedForErase:
		return "device not opened for erase"
	case FTDeviceNotOpenedForWrite:
		return "device not opened for write"
	case FTFailedToWriteDevice:
		return "failed to write device"
	case FTEEPROMReadFailed:
		return "EEPROM read failed"
	case FTEEPROMWriteFailed:
		return "EEPROM write failed"
	case FTEEPROMEraseFailed:
		return "EEPROM erase failed"
	case FTEEPROMNotPresent:
		return "EEPROM not present"
	case FTEEPROMNotProgrammed:
		return "EEPROM not programmed"
	case FTInvalidArgs:
		return "invalid args"
	case FTNotSupported:
		return "not supported"
	case FTOtherError:
		return "other error"
	case FTDeviceListNotReady:
		return "device list not ready"
	default:
		return "unknown error"
	}
}
