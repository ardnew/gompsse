package gompsse

type MPSSE struct {
	I2C *I2C
	SPI *SPI
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

func NewMPSSE() *MPSSE {
	return &MPSSE{I2C: nil, SPI: nil}
}
