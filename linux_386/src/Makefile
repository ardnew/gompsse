target  = libMPSSE
install = ../lib

ftd2xx_name = ftd2xx
ftd2xx_root = lib$(ftd2xx_name)-linux-386-1.4.8

includes = $(addprefix -I,$(ftd2xx_root) .)
libpaths = $(addprefix -L,$(ftd2xx_root))
libnames = $(addprefix -l,$(ftd2xx_name))

i2c_src  = ftdi_i2c.c
spi_src  = ftdi_spi.c
sources  = ftdi_infra.c ftdi_mid.c $(i2c_src) $(spi_src)
objects  = $(sources:.c=.o)

#defines = -DINFRA_debug_ENABLE
defines  =
optimize = -Og
debug    = -g
warnings = -Wall
cflags   = -fPIC
ldflags  = -shared

cross   =
objdump = $(cross)objdump
CC      = $(cross)gcc
AR      = $(cross)ar
LD      = $(cross)ld
CFLAGS  = $(debug) $(optimize) $(warnings) $(defines) $(includes) $(cflags)
LDFLAGS = $(libpaths) $(libnames) $(ldflags)

span := $(shell perl -e 'print "-"x80')
buildtime := $(shell perl -e 'printf "\tBuild date:\t%s", scalar localtime()')

.PHONY: all install clean

all: $(target)
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

install: all
	cp $(target).a $(target).a.info $(target).so $(target).so.info $(install)

clean:
	$(RM) -f *.o *.a *.so *.info

$(target): $(objects)
	$(CC) $(LDFLAGS) -o $(target).so $(objects)
	$(AR) -rcs $(target).a $(objects)

$(objects): %.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<

