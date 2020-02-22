package gompsse

import "fmt"

type spiXferOption uint32

// Constants controlling the supported SPI transfer options
const (
	spiXferBytes spiXferOption = 0x00000000 // size is provided in bytes
	spiXferBits  spiXferOption = 0x00000001 // size is provided in bits

	spiCSAssert   spiXferOption = 0x00000002 // assert CS before start
	spiCSDeAssert spiXferOption = 0x00000004 // deassert CS after end
)

// Constants related to board pins when MPSSE operating in SPI mode
const (
	spiClockMaximum   = 30000000
	spiClockDefault   = 12000000 // valid range: 0-30000000 (30 MHz)
	spiLatencyDefault = 16       // 1-255 USB Hi-Speed, 2-255 USB Full-Speed
)

type spiOption uint32

// Constants defining the available options in the SPI configuration struct.
const (
	// Known SPI operating modes
	//   LIMITATION: libMPSSE only supports mode 0 and mode 2 (CPHA==2).
	SPIMode0       spiOption = 0x00000000 // capture on RISE, propagate on FALL
	SPIMode1       spiOption = 0x00000001 // capture on FALL, propagate on RISE
	SPIMode2       spiOption = 0x00000002 // capture on FALL, propagate on RISE
	SPIMode3       spiOption = 0x00000003 // capture on RISE, propagate on FALL
	spiModeMask    spiOption = 0x00000003
	spiModeInvalid spiOption = 0x000000FF
	spiModeDefault spiOption = SPIMode0

	// DPins available for chip-select operation
	spiCSD3      spiOption = 0x00000000 // SPI CS on D3
	spiCSD4      spiOption = 0x00000004 // SPI CS on D4
	spiCSD5      spiOption = 0x00000008 // SPI CS on D5
	spiCSD6      spiOption = 0x0000000C // SPI CS on D6
	spiCSD7      spiOption = 0x00000010 // SPI CS on D7
	spiCSMask    spiOption = 0x0000001C
	spiCSInvalid spiOption = 0x000000FF
	spiCSDefault spiOption = spiCSD3

	// Other options
	spiCSActiveLow     spiOption = 0x00000020 // drive pin low to assert CS
	spiCSActiveHigh    spiOption = 0x00000000 // drive pin high to assert CS
	spiCSActiveDefault spiOption = spiCSActiveLow
)

// spiCSPin translates a DPin value to its corresponding chip-select mask for
// the SPI configuration struct option.
var spiCSPin = map[DPin]spiOption{
	D0: spiCSInvalid,
	D1: spiCSInvalid,
	D2: spiCSInvalid,
	D3: spiCSD3,
	D4: spiCSD4,
	D5: spiCSD5,
	D6: spiCSD6,
	D7: spiCSD7,
}

// spiDPinConfig represents the default direction and value for pins associated
// with the lower byte lines of MPSSE, reserved for serial functions SPI/I²C
// (or port "D" on FT232H), but has a few GPIO pins as well.
type spiDPinConfig struct {
	initDir  byte // direction of lines after SPI channel initialization
	initVal  byte // value of lines after SPI channel initialization
	closeDir byte // direction of lines after SPI channel is closed
	closeVal byte // value of lines after SPI channel is closed
}

// spiDPinConfigDefault defines the initial spiDPinConfig value for all pins
// represented by this type. all output pins are configured LOW except for the
// default CS pin (D3) since we also have spiCSActiveLow by default. this means
// we won't activate the default slave line until intended. it also means SCLK
// idles LOW (change initVal to PinHI to idle HIGH).
func spiDPinConfigDefault() uint32 {
	return spiDPin([NumDPins]*spiDPinConfig{
		&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D0 SCLK
		&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D1 MOSI
		&spiDPinConfig{initDir: PinIN, initVal: PinLO, closeDir: PinIN, closeVal: PinLO}, // D2 MISO
		&spiDPinConfig{initDir: PinOT, initVal: PinHI, closeDir: PinOT, closeVal: PinHI}, // D3 CS
		&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D4 GPIO
		&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D5 GPIO
		&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D6 GPIO
		&spiDPinConfig{initDir: PinOT, initVal: PinLO, closeDir: PinOT, closeVal: PinLO}, // D7 GPIO
	})
}

// spiDPin constructs the 32-bit field pin of the spiConfig struct from the
// provided spiDPinConfig slice cfg for each pin (identified by its index in the
// given slice).
func spiDPin(cfg [NumDPins]*spiDPinConfig) uint32 {
	var pin uint32
	for i, c := range cfg {
		pin |= (uint32(c.initDir) << i) | (uint32(c.initVal) << (8 + i)) |
			(uint32(c.closeDir) << (16 + i)) | (uint32(c.closeVal) << (24 + i))
	}
	return pin
}

// spiConfig holds all of the configuration settings for an SPI channel.
type spiConfig struct {
	clockRate uint32 // in Hertz
	latency   uint8  // in ms
	options   spiOption
	pin       uint32 // port D pins ("low byte lines of MPSSE")
	reserved  uint16
}

func spiConfigDefault() *spiConfig {
	return &spiConfig{
		clockRate: spiClockDefault,
		latency:   spiLatencyDefault,
		options:   spiCSActiveDefault | spiCSDefault | spiModeDefault,
		pin:       spiDPinConfigDefault(),
		reserved:  0,
	}
}

type SPI struct {
	device *MPSSE
	config *spiConfig
}

func (spi *SPI) ChangeCS(cs DPin) error {

	var csOpt spiOption

	if csOpt, ok := spiCSPin[cs]; !ok || (spiCSInvalid == csOpt) {
		return fmt.Errorf("invalid CS pin: %d", cs)
	}

	spi.config.options &= ^(spiCSMask)
	spi.config.options |= csOpt

	return nil
}

func (spi *SPI) SetOptions(cs DPin, activeLow bool, mode byte) error {

	var (
		activeOpt spiOption
		modeOpt   spiOption
	)

	if activeLow {
		activeOpt = spiCSActiveLow
	} else {
		activeOpt = spiCSActiveHigh
	}

	if spiOption(mode) > spiModeMask {
		return fmt.Errorf("invalid SPI mode: Mode %d", mode)
	} else {
		modeOpt = spiOption(mode)
	}

	spi.config.options = activeOpt | modeOpt
	return spi.ChangeCS(cs)
}

func (spi *SPI) SetConfig(clock uint32, latency byte, cs DPin, activeLow bool, mode byte) error {

	if 0 == clock {
		spi.config.clockRate = spiClockDefault
	} else {
		if clock <= spiClockMaximum {
			spi.config.clockRate = clock
		} else {
			return fmt.Errorf("invalid clock rate: %d", clock)
		}
	}

	if 0 == latency {
		spi.config.latency = spiLatencyDefault
	} else {
		spi.config.latency = latency
	}

	return spi.SetOptions(cs, activeLow, mode)
}

func (spi *SPI) Init() error {

	if err := _SPI_InitChannel(spi); nil != err {
		return err
	}

	spi.device.mode = ModeSPI

	return spi.device.GPIO.Init() // reset GPIO
}

func (spi *SPI) Write(data []uint8, start bool, stop bool) (uint32, error) {
	opt := spiXferBytes
	if start {
		opt |= spiCSAssert
	}
	if stop {
		opt |= spiCSDeAssert
	}
	return _SPI_Write(spi, data, opt)
}

func (spi *SPI) WriteWith(cs CPin, data []uint8, start bool, stop bool) (uint32, error) {

	assert := 0 == uint32(spiCSActiveLow&spi.config.options)

	if start {
		if err := spi.device.GPIO.Set(cs, assert); nil != err {
			return 0, err
		}
	}
	if stop {
		defer func() { _ = spi.device.GPIO.Set(cs, !assert) }()
	}

	return _SPI_Write(spi, data, spiXferBytes)
}
