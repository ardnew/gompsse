# gompsse
##### Go driver for FTDI `LibMPSSE` (Multi-Protocol Synchronous Serial Engine) providing high-level API for I²C and SPI modes with GPIO

## Features
- [x] Designed for and tested with [FT232H](https://www.ftdichip.com/Products/ICs/FT232H.htm)
  - [Adafruit sells a very nice breakout with a bunch of extras](https://www.adafruit.com/product/2264):
    - USB-C and Stemma QT/Qwiic I²C connectors (with a little switch to short the chip's two awkward `SDA` pins!)
    - On-board EEPROM (for storing chip configuration)
    - 5V (`VBUS`) and 3.3V (on-board regulator, up to 500mA draw) outputs
- [x] Includes re-compilable native FTDI drivers for multiple host OS
  - Linux 32-bit (`386`) and 64-bit (`amd64`, `arm64`) - includes Raspberry Pi models 3 and 4
  - macOS (`amd64`)
  - Windows not currently supported

## Drivers
All communication with MPSSE-capable devices is performed with FTDI's open-source driver [`LibMPSSE`](https://www.ftdichip.com/Support/SoftwareExamples/MPSSE.htm). This software however depends on FTDI's proprietary driver [`FTD2XX`](https://www.ftdichip.com/Drivers/D2XX.htm) (based on [`libusb`](https://github.com/libusb/libusb)), which is only available for certain platforms.

Contained in this project are all of the necessary files required to build `LibMPSSE` for each supported OS as well as pre-compiled libraries for your all-important ~~laziness~~convenience (also so that it _Just Works_ with `go get`). A simple GNU Makefile has also been created for each (see: [Building LibMPSSE](#building-libmpsse-optional)). Changes to both `FTD2XX` and `LibMPSSE` were made for general compatibility reasons and should not affect the API.

Under [`native`](native), you will find the headers needed by `gompsse` to communicate with the C libraries, the source code for `LibMPSSE`, and the pre-compiled `FTD2XX` library for each supported platform:

```sh
└── native/
    ├── inc/  # LibMPSSE C APIs and FTD2XX C source code headers needed by cgo
    └── src/  # LibMPSSE C source code, GNU Makefile
        └── `${GOOS}_${GOARCH}`/ # proprietary FTD2XX library
```

#### Building LibMPSSE (optional)
TBD

