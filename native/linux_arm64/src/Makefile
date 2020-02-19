target  = libMPSSE
install = ../lib

goos   = linux
goarch = arm64

ftd2xx_name = ftd2xx
ftd2xx_vers = 1.4.8
ftd2xx_root = lib$(ftd2xx_name)-$(goos)-$(goarch)-$(ftd2xx_vers)

includes = $(addprefix -I,$(ftd2xx_root) .)
libpaths = $(addprefix -L,$(ftd2xx_root))
libnames = $(addprefix -l,$(ftd2xx_name))

i2c_src  = ftdi_i2c.c
spi_src  = ftdi_spi.c
sources  = ftdi_infra.c ftdi_mid.c $(i2c_src) $(spi_src)
objects  = $(sources:.c=.o)
outputs  = $(target).a $(target).a.info $(ftd2xx_root)/lib$(ftd2xx_name).a

#defines = -DINFRA_debug_ENABLE
defines  =
optimize = -Og
debug    = -g
warnings = -Wall
cflags   = -fno-PIC
ldflags  = -fno-PIC -static

cross   =
objdump = $(cross)objdump
CC      = $(cross)gcc
AR      = $(cross)ar
LD      = $(cross)ld
CFLAGS  = $(debug) $(optimize) $(warnings) $(defines) $(includes) $(cflags)
LDFLAGS = $(ldflags) $(libpaths) $(libnames)

span := $(shell perl -e 'print "-"x80')
buildtime := $(shell perl -e 'printf "\tBuild date:\t%s", scalar localtime()')

.PHONY: all install clean

all: $(target)
	@echo "$(span)"
	@echo "$(buildtime)"
	@echo "$(span)"
	$(shell echo "$(span)"       > $(target).a.info)
	$(shell echo "$(buildtime)" >> $(target).a.info)
	$(shell echo "$(span)"      >> $(target).a.info)
	$(shell $(objdump) -fpt $(target).a >> $(target).a.info)

install: $(outputs)
	@echo "$(span)"
	@echo "libraries installed to \"$(install)/\":"
	@echo "	$(target).a"
	@echo "	lib$(ftd2xx_name).a"
	@echo "$(span)"

clean:
	$(RM) -f *.o *.a *.info

$(outputs): all
	@test -d $(install) || mkdir -p $(install)
	cp $@ $(install)

$(target): $(objects)
	$(AR) -rcs $(target).a $(objects)

$(objects): %.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<
