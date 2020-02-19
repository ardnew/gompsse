package gompsse

// #cgo darwin,amd64 LDFLAGS: -framework CoreFoundation -framework IOKit
// #cgo  CFLAGS: -I${SRCDIR}/native/inc
// #cgo LDFLAGS: -lMPSSE -lftd2xx -ldl
// #include "ftd2xx.h"
// #include "stdlib.h"
import "C"

type Handle C.FT_HANDLE
type Status C.FT_STATUS

// Constants related to device status
const (
	SOK                      Status = C.FT_OK
	SInvalidHandle           Status = C.FT_INVALID_HANDLE
	SDeviceNotFound          Status = C.FT_DEVICE_NOT_FOUND
	SDeviceNotOpened         Status = C.FT_DEVICE_NOT_OPENED
	SIOError                 Status = C.FT_IO_ERROR
	SInsufficientResources   Status = C.FT_INSUFFICIENT_RESOURCES
	SInvalidParameter        Status = C.FT_INVALID_PARAMETER
	SInvalidBaudRate         Status = C.FT_INVALID_BAUD_RATE
	SDeviceNotOpenedForErase Status = C.FT_DEVICE_NOT_OPENED_FOR_ERASE
	SDeviceNotOpenedForWrite Status = C.FT_DEVICE_NOT_OPENED_FOR_WRITE
	SFailedToWriteDevice     Status = C.FT_FAILED_TO_WRITE_DEVICE
	SEEPROMReadFailed        Status = C.FT_EEPROM_READ_FAILED
	SEEPROMWriteFailed       Status = C.FT_EEPROM_WRITE_FAILED
	SEEPROMEraseFailed       Status = C.FT_EEPROM_ERASE_FAILED
	SEEPROMNotPresent        Status = C.FT_EEPROM_NOT_PRESENT
	SEEPROMNotProgrammed     Status = C.FT_EEPROM_NOT_PROGRAMMED
	SInvalidArgs             Status = C.FT_INVALID_ARGS
	SNotSupported            Status = C.FT_NOT_SUPPORTED
	SOtherError              Status = C.FT_OTHER_ERROR
	SDeviceListNotReady      Status = C.FT_DEVICE_LIST_NOT_READY
)

func (s Status) OK() bool {
	return SOK == s
}

func (s Status) Error() string {
	switch s {
	case SOK:
		return "OK"
	case SInvalidHandle:
		return "invalid handle"
	case SDeviceNotFound:
		return "device not found"
	case SDeviceNotOpened:
		return "device not opened"
	case SIOError:
		return "IO error"
	case SInsufficientResources:
		return "insufficient resources"
	case SInvalidParameter:
		return "invalid parameter"
	case SInvalidBaudRate:
		return "invalid baud rate"
	case SDeviceNotOpenedForErase:
		return "device not opened for erase"
	case SDeviceNotOpenedForWrite:
		return "device not opened for write"
	case SFailedToWriteDevice:
		return "failed to write device"
	case SEEPROMReadFailed:
		return "EEPROM read failed"
	case SEEPROMWriteFailed:
		return "EEPROM write failed"
	case SEEPROMEraseFailed:
		return "EEPROM erase failed"
	case SEEPROMNotPresent:
		return "EEPROM not present"
	case SEEPROMNotProgrammed:
		return "EEPROM not programmed"
	case SInvalidArgs:
		return "invalid args"
	case SNotSupported:
		return "not supported"
	case SOtherError:
		return "other error"
	case SDeviceListNotReady:
		return "device list not ready"
	default:
		return "unknown error"
	}
}

type DeviceInfo struct {
	IsOpen      bool
	IsHiSpeed   bool
	Type        uint32
	ID          uint32
	LocID       uint32
	SerialNo    string
	Description string
	Handle      Handle
}

func NewDeviceInfo(info *C.FT_DEVICE_LIST_INFO_NODE) *DeviceInfo {
	return &DeviceInfo{
		IsOpen:      1 == (info.Flags & 0x01),
		IsHiSpeed:   2 == (info.Flags & 0x02),
		Type:        uint32(info.Type),
		ID:          uint32(info.ID),
		LocID:       uint32(info.LocId),
		SerialNo:    C.GoString(&info.SerialNumber[0]),
		Description: C.GoString(&info.Description[0]),
		Handle:      Handle(info.ftHandle),
	}
}
