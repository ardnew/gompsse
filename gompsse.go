package gompsse

// unfortunately it isn't possible to use ${GOOS}/${GOARCH} in the #cgo
// directives, so we have to duplicate for each platform.

// #cgo darwin,amd64 LDFLAGS: -framework CoreFoundation -framework IOKit
// #cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/darwin_amd64/lib
// #cgo darwin,amd64  CFLAGS: -I${SRCDIR}/darwin_amd64/inc
// #cgo linux,386    LDFLAGS: -L${SRCDIR}/linux_386/lib
// #cgo linux,386     CFLAGS: -I${SRCDIR}/linux_386/inc
// #cgo linux,amd64  LDFLAGS: -L${SRCDIR}/linux_amd64/lib
// #cgo linux,amd64   CFLAGS: -I${SRCDIR}/linux_amd64/inc
// #cgo linux,arm64  LDFLAGS: -L${SRCDIR}/linux_arm64/lib
// #cgo linux,arm64   CFLAGS: -I${SRCDIR}/linux_arm64/inc
// #cgo              LDFLAGS: -lMPSSE
// #include "libMPSSE_spi.h"
// #include "libMPSSE_i2c.h"
import "C"
