target  := libMPSSE
install := /usr/local/lib

os   := linux
#os   := darwin
arch := amd64
#arch := 386
#arch := arm64

ftd2xx_name        = ftd2xx
linux-ftd2xx_vers  = 1.4.8
darwin-ftd2xx_vers = 1.4.16
ftd2xx_root        = $(os)-$(arch)/lib$(ftd2xx_name)-$(os)-$(arch)-$($(os)-ftd2xx_vers)

includes = $(addprefix -I,$(ftd2xx_root) . ../inc)
libpaths = $(addprefix -L,$(ftd2xx_root))
libnames = $(addprefix -l,$(ftd2xx_name))

i2c_src  = ftdi_i2c.c
spi_src  = ftdi_spi.c
sources  = ftdi_infra.c ftdi_mid.c $(i2c_src) $(spi_src)
objects  = $(sources:.c=.o)
linux-outputs  = $(target).a $(target).so $(ftd2xx_root)/lib$(ftd2xx_name).a $(ftd2xx_root)/lib$(ftd2xx_name).so.$($(os)-ftd2xx_vers)
darwin-outputs = $(target).a $(target).dylib $(ftd2xx_root)/lib$(ftd2xx_name).a $(ftd2xx_root)/lib$(ftd2xx_name).$($(os)-ftd2xx_vers).dylib

#defines = -DINFRA_debug_ENABLE
defines  =
optimize = -Og
debug    = -g
warnings = -Wall
linux-cflags   = -fPIC
linux-ldflags  = -fPIC -shared
darwin-cflags  =
darwin-ldflags = -dynamiclib

cross   =
objdump = $(cross)objdump
CC      = $(cross)gcc
AR      = $(cross)ar
LD      = $(cross)ld
CFLAGS  = $(debug) $(optimize) $(warnings) $(defines) $(includes) $($(os)-cflags)
LDFLAGS = $($(os)-ldflags) $(libpaths) $(libnames)

span      := $(shell perl -e 'print "-"x80')
buildtime := $(shell perl -e 'printf "\tBuild date:\t%s", scalar localtime()')

.PHONY: all linux-all darwin-all install linux-install darwin-install clean

all: $(os)-all

linux-all: $(target)
	@printf "%s\n%s\n%s\n" "$(span)" "$(buildtime)" "$(span)" | tee "$(target).a.info" "$(target).so.info"
	$(objdump) -fpt $(target).a >> $(target).a.info
	$(objdump) -fpt $(target).so >> $(target).so.info

darwin-all: $(target)
	@printf "%s\n%s\n%s\n" "$(span)" "$(buildtime)" "$(span)" | tee "$(target).a.info" "$(target).dylib.info"
	$(objdump) -macho -dylibs-used -archive-headers -private-header -t -indirect-symbols -g $(target).a >> $(target).a.info
	$(objdump) -macho -dylibs-used -archive-headers -private-header -t -indirect-symbols -g $(target).dylib >> $(target).dylib.info

install: $(os)-install
	@echo "$(span)"
	@echo "libraries installed to: $(install)"
	@printf "\t%s\n" $($(os)-outputs)
	@echo "$(span)"

linux-install: $($(os)-outputs)
	@ln -sf lib$(ftd2xx_name).so.$($(os)-ftd2xx_vers) $(install)/lib$(ftd2xx_name).so

darwin-install: $($(os)-outputs)
	@ln -sf lib$(ftd2xx_name).$($(os)-ftd2xx_vers).dylib $(install)/lib$(ftd2xx_name).dylib

$($(os)-outputs): all
	@test -d $(install) || mkdir -p $(install)
	cp $@ $(install)

$(target): $(os)-$(target)
	$(AR) -rcs $(target).a $(objects)

linux-$(target): $(objects)
	$(CC) -o $(target).so $(objects) $(LDFLAGS)

darwin-$(target): $(objects)
	$(CC) -o $(target).dylib $(objects) $(LDFLAGS)

$(objects): %.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<

clean:
	$(RM) -f *.o *.a *.so *.dylib *.info

