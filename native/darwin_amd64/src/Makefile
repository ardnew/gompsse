target  = libMPSSE
install = /usr/lib

goos   = darwin
goarch = amd64

ftd2xx_name = ftd2xx
ftd2xx_vers = 1.4.16
ftd2xx_root = lib$(ftd2xx_name)-$(goos)-$(goarch)-$(ftd2xx_vers)

includes = $(addprefix -I,$(ftd2xx_root) .)
libpaths = $(addprefix -L,$(ftd2xx_root))
libnames = $(addprefix -l,$(ftd2xx_name))

i2c_src  = ftdi_i2c.c
spi_src  = ftdi_spi.c
sources  = ftdi_infra.c ftdi_mid.c $(i2c_src) $(spi_src)
objects  = $(sources:.c=.o)
outputs  = $(target).a $(target).dylib $(ftd2xx_root)/lib$(ftd2xx_name).a $(ftd2xx_root)/lib$(ftd2xx_name).$(ftd2xx_vers).dylib

#defines = -DINFRA_debug_ENABLE
defines  =
optimize = -Og
debug    = -g
warnings = -Wall
cflags   =
ldflags  = -dynamiclib

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
	@echo "	$(target).a	- static library"
	@echo "	$(target).dylib	- shared object library"
	@echo "$(span)"
	@# generate the library info text for static library
	$(shell echo "$(span)"       > $(target).a.info)
	$(shell echo "$(buildtime)" >> $(target).a.info)
	$(shell echo "$(span)"      >> $(target).a.info)
	$(shell $(objdump) -macho -dylibs-used -archive-headers -private-header -t -indirect-symbols -g $(target).a >> $(target).a.info)
	@# generate the library info text for dynamic library
	$(shell echo "$(span)"       > $(target).dylib.info)
	$(shell echo "$(buildtime)" >> $(target).dylib.info)
	$(shell echo "$(span)"      >> $(target).dylib.info)
	$(shell $(objdump) -macho -dylibs-used -archive-headers -private-header -t -indirect-symbols -g $(target).dylib >> $(target).dylib.info)

install: $(outputs)
	@ln -sf lib$(ftd2xx_name).$(ftd2xx_vers).dylib $(install)/lib$(ftd2xx_name).dylib
	@echo "$(span)"
	@echo "libraries installed to \"$(install)/\":"
	@echo "	$(target).a"
	@echo "	$(target).dylib"
	@echo "	lib$(ftd2xx_name).a"
	@echo "	lib$(ftd2xx_name).$(ftd2xx_vers).dylib"
	@echo "	lib$(ftd2xx_name).dylib -> lib$(ftd2xx_name).$(ftd2xx_vers).dylib"
	@echo "$(span)"

clean:
	$(RM) -f *.o *.a *.dylib *.info

$(outputs): all
	@test -d $(install) || mkdir -p $(install)
	cp $@ $(install)

$(target): $(objects)
	$(CC) -o $(target).dylib $(objects) $(LDFLAGS)
	$(AR) -rcs $(target).a $(objects)

$(objects): %.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<
