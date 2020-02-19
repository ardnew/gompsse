target  = libMPSSE
install = /usr/local/lib

#os   = linux
os   := darwin
arch := amd64
#arch = 386
#arch = arm64

ftd2xx_name = ftd2xx
linux-ftd2xx_vers = 1.4.8
darwin-ftd2xx_vers = 1.4.16
ftd2xx_root = ../$(os)-$(arch)/lib$(ftd2xx_name)-$(os)-$(arch)-$($(os)-ftd2xx_vers)

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
darwin-cflags  =
linux-ldflags  = -shared
darwin-ldflags = -dynamiclib

cross   =
objdump = $(cross)objdump
CC      = $(cross)gcc
AR      = $(cross)ar
LD      = $(cross)ld
CFLAGS  = $(debug) $(optimize) $(warnings) $(defines) $(includes) $($(os)-cflags)
LDFLAGS = $($(os)-ldflags) $(libpaths) $(libnames)

span := $(shell perl -e 'print "-"x80')
buildtime := $(shell perl -e 'printf "\tBuild date:\t%s", scalar localtime()')

.PHONY: all linux-all darwin-all install linux-install darwin-install clean linux-clean darwin-clean

all: $(os)-all

linux-all: $(target)
	@echo "$(span)"
	@echo "$(buildtime)"
	@echo "$(span)"
	@echo "	$(target).a	- static library"
	@echo "	$(target).so	- shared object library"
	@echo "$(span)"
	@# generate the library info text for static library
	$(shell echo "$(span)"       > $(target).a.info)
	$(shell echo "$(buildtime)" >> $(target).a.info)
	$(shell echo "$(span)"      >> $(target).a.info)
	$(shell $(objdump) -fpt $(target).a >> $(target).a.info)
	@# generate the library info text for dynamic library
	$(shell echo "$(span)"       > $(target).so.info)
	$(shell echo "$(buildtime)" >> $(target).so.info)
	$(shell echo "$(span)"      >> $(target).so.info)
	$(shell $(objdump) -fpt $(target).so >> $(target).so.info)

darwin-all: $(target)
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

install: $(os)-install

linux-install: $(objects)
	@ln -sf lib$(ftd2xx_name).so.$($(os)-ftd2xx_vers) $(install)/lib$(ftd2xx_name).so
	@echo "$(span)"
	@echo "libraries installed to \"$(install)/\":"
	@echo "	$(target).a"
	@echo "	$(target).so"
	@echo "	lib$(ftd2xx_name).a"
	@echo "	lib$(ftd2xx_name).so.$($(os)-ftd2xx_vers)"
	@echo "	lib$(ftd2xx_name).so -> lib$(ftd2xx_name).so.$($(os)-ftd2xx_vers)"
	@echo "$(span)"

darwin-install: $($(os)-outputs)
	@ln -sf lib$(ftd2xx_name).$($(os)-ftd2xx_vers).dylib $(install)/lib$(ftd2xx_name).dylib
	@echo "$(span)"
	@echo "libraries installed to \"$(install)/\":"
	@echo "	$(target).a"
	@echo "	$(target).dylib"
	@echo "	lib$(ftd2xx_name).a"
	@echo "	lib$(ftd2xx_name).$($(os)-ftd2xx_vers).dylib"
	@echo "	lib$(ftd2xx_name).dylib -> lib$(ftd2xx_name).$($(os)-ftd2xx_vers).dylib"
	@echo "$(span)"

clean: $(os)-clean

linux-clean:
	$(RM) -f *.o *.a *.so *.info

darwin-clean:
	$(RM) -f *.o *.a *.dylib *.info

$(target): $(os)-$(target)

linux-$(target): $(objects)
	$(CC) -o $(target).so $(objects) $(LDFLAGS)
	$(AR) -rcs $(target).a $(objects)

darwin-$(target): $(objects)
	$(CC) -o $(target).dylib $(objects) $(LDFLAGS)
	$(AR) -rcs $(target).a $(objects)

$($(os)-outputs): all
	@test -d $(install) || mkdir -p $(install)
	cp $@ $(install)

$(objects): %.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<
