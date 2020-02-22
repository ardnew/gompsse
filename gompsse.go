package gompsse

import (
	"fmt"
	"strings"
)

type MPSSE struct {
	info *deviceInfo
	mode Mode
	I2C  *I2C
	SPI  *SPI
	GPIO *GPIO
}

func (m *MPSSE) String() string {
	return fmt.Sprintf("{ Info: %s, Mode: %s, I2C: %+v, SPI: %+v, GPIO: %+v }",
		m.info, m.mode, m.I2C, m.SPI, m.GPIO)
}

func NewMPSSE() (*MPSSE, error) {
	return NewMPSSEWithMask(nil) // first device found
}

func NewMPSSEWithIndex(index uint) (*MPSSE, error) {
	s := fmt.Sprintf("%d", index)
	return NewMPSSEWithMask(&OpenMask{Index: s})
}

func NewMPSSEWithVIDPID(vid uint16, pid uint16) (*MPSSE, error) {
	v := fmt.Sprintf("%04x", vid)
	p := fmt.Sprintf("%04x", pid)
	return NewMPSSEWithMask(&OpenMask{VID: v, PID: p})
}

func NewMPSSEWithSerial(serial string) (*MPSSE, error) {
	return NewMPSSEWithMask(&OpenMask{Serial: serial})
}

func NewMPSSEWithDesc(desc string) (*MPSSE, error) {
	return NewMPSSEWithMask(&OpenMask{Desc: desc})
}

func NewMPSSEWithMask(mask *OpenMask) (*MPSSE, error) {
	m := &MPSSE{info: nil, mode: ModeNone, I2C: nil, SPI: nil}
	if err := m.openDevice(mask); nil != err {
		return nil, err
	}
	m.I2C = &I2C{device: m, config: i2cConfigDefault()}
	m.SPI = &SPI{device: m, config: spiConfigDefault()}
	m.GPIO = &GPIO{device: m, config: gpioConfigDefault()}
	if err := m.GPIO.Init(); nil != err {
		return nil, err
	}
	return m, nil
}

type OpenMask struct {
	Index  string
	VID    string
	PID    string
	Serial string
	Desc   string
}

func (m *MPSSE) openDevice(mask *OpenMask) error {

	var (
		dev []*deviceInfo
		sel *deviceInfo
		err error
	)

	if dev, err = devices(); nil != err {
		return err
	}

	for _, d := range dev {
		if nil == mask {
			sel = d
			break
		}
		if "" != mask.Index {
			if mask.Index != fmt.Sprintf("%d", d.index) {
				continue
			}
		}
		if "" != mask.VID {
			ms := strings.ToLower(mask.VID)
			dx := fmt.Sprintf("%x", d.vid)
			dz := fmt.Sprintf("%04x", d.vid)
			if (ms != dx) && (ms != ("0x" + dx)) &&
				(ms != dz) && (ms != ("0x" + dz)) &&
				(ms != fmt.Sprintf("%d", d.vid)) {
				continue
			}
		}
		if "" != mask.PID {
			ms := strings.ToLower(mask.PID)
			dx := fmt.Sprintf("%x", d.pid)
			dz := fmt.Sprintf("%04x", d.pid)
			if (ms != dx) && (ms != ("0x" + dx)) &&
				(ms != dz) && (ms != ("0x" + dz)) &&
				(ms != fmt.Sprintf("%d", d.pid)) {
				continue
			}
		}
		if "" != mask.Serial {
			if strings.ToLower(mask.Serial) != strings.ToLower(d.serial) {
				continue
			}
		}
		if "" != mask.Desc {
			if strings.ToLower(mask.Desc) != strings.ToLower(d.desc) {
				continue
			}
		}
		sel = d
		break
	}

	if nil == sel {
		return SDeviceNotFound
	}

	if err = sel.open(); nil != err {
		return err
	}
	m.info = sel
	return nil
}

func (m *MPSSE) Close() error {
	if nil != m.info {
		return m.info.close()
	}
	m.mode = ModeNone
	return nil
}

// Types representing individual port pins.
type (
	DPin byte // pin on MPSSE low-byte lines (port "D" on FT232H)
	CPin byte // pin on MPSSE high-byte lines (port "C" on FT232H)
)

// Constants related to GPIO pin configuration
const (
	PinLO byte = 0 // pin value clear
	PinHI byte = 1 // pin value set
	PinIN byte = 0 // pin direction input
	PinOT byte = 1 // pin direction output

	NumDPins = 8 // number of MPSSE low-byte line pins
	NumCPins = 8 // number of MPSSE high-byte line pins
)

// Constants defining the available board pins on MPSSE low-byte lines
const (
	D0 DPin = 1 << iota
	D1
	D2
	D3
	D4
	D5
	D6
	D7
)

// Constants defining the available board pins on MPSSE high-byte lines
const (
	C0 CPin = 1 << iota
	C1
	C2
	C3
	C4
	C5
	C6
	C7
)

type deviceInfo struct {
	index     int
	isOpen    bool
	isHiSpeed bool
	chip      Chip
	vid       uint32
	pid       uint32
	locID     uint32
	serial    string
	desc      string
	handle    Handle
}

func (dev *deviceInfo) String() string {
	return fmt.Sprintf("%d:{ Open = %t, HiSpeed = %t, Chip = \"%s\" (0x%02X), "+
		"VID = 0x%04X, PID = 0x%04X, Location = %04X, "+
		"Serial = \"%s\", Desc = \"%s\", Handle = %p }",
		dev.index, dev.isOpen, dev.isHiSpeed, dev.chip, uint32(dev.chip),
		dev.vid, dev.pid, dev.locID, dev.serial, dev.desc, dev.handle)
}

func (dev *deviceInfo) open() error {
	if ce := dev.close(); nil != ce {
		return ce
	}
	if oe := _FT_Open(dev); nil != oe {
		return oe
	}
	dev.isOpen = true
	return nil
}

func (dev *deviceInfo) close() error {
	if !dev.isOpen {
		return nil
	}
	if ce := _FT_Close(dev); nil != ce {
		return ce
	}
	dev.isOpen = false
	return nil
}

func devices() ([]*deviceInfo, error) {

	n, ce := _FT_CreateDeviceInfoList()
	if nil != ce {
		return nil, ce
	}

	if 0 == n {
		return []*deviceInfo{}, nil
	}

	info, de := _FT_GetDeviceInfoList(n)
	if nil != de {
		return nil, de
	}

	return info, nil
}

type gpioConfig struct {
	dir uint8
	val uint8
}

func gpioConfigDefault() *gpioConfig {
	return &gpioConfig{
		dir: 0xFF, // each bit set, all pins OUTPUT by default
		val: 0x00, // each bit clear, all pins LOW by default
	}
}

type GPIO struct {
	device *MPSSE
	config *gpioConfig
}

func (gpio *GPIO) Init() error {
	return gpio.Write(gpio.config.dir, gpio.config.val)
}

func (gpio *GPIO) Write(dir uint8, val uint8) error {

	val &= dir // only set output bits

	if err := _FT_WriteGPIO(gpio, dir, val); nil != err {
		return err
	}

	gpio.config.dir = dir
	gpio.config.val = val

	return nil
}

func (gpio *GPIO) Read() (uint8, error) {

	val, err := _FT_ReadGPIO(gpio)
	if nil != err {
		return 0, err
	}

	gpio.config.val = val

	return val, nil
}

func (gpio *GPIO) Set(pin CPin, val bool) error {

	dir := gpio.config.dir | uint8(pin)
	set := gpio.config.val

	if val {
		set |= uint8(pin)
	} else {
		set &= ^uint8(pin)
	}

	return gpio.Write(dir, set)
}

func (gpio *GPIO) Get(pin CPin) (bool, error) {

	set, err := gpio.Read()
	if nil != err {
		return false, err
	}

	return (set & uint8(pin)) > 0, nil
}
