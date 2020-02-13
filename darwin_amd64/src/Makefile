TARGET  = libMPSSE
INSTALL = ../lib

FTD2XX_DIR = libftd2xx-darwin-amd64-1.4.16
INCLUDES   = -I$(FTD2XX_DIR) -I.

#DEFINES = -DINFRA_DEBUG_ENABLE
DEFINES  = 
OPTIMIZE = -O3
DEBUG    = -g
WARNINGS = -Wall
#FLAGS    = -fprofile-arcs -ftest-coverage
FLAGS    =

CROSS = 
CC    = $(CROSS)gcc
AR    = $(CROSS)ar
LD    = $(CROSS)ld

CFLAGS  = $(DEBUG) $(OPTIMIZE) $(WARNINGS) $(FLAGS) $(DEFINES) $(INCLUDES)
LDFLAGS = -L$(FTD2XX_DIR) -lftd2xx

I2C_SRC = ftdi_i2c.c
SPI_SRC = ftdi_spi.c
SOURCES = ftdi_infra.c ftdi_mid.c $(I2C_SRC) $(SPI_SRC)
OBJECTS = $(SOURCES:.c=.o)

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) $(LDFLAGS) -dynamiclib -o $(TARGET).dylib $(OBJECTS)
	$(AR) -rcs $(TARGET).a $(OBJECTS)

$(OBJECTS): %.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<

install: $(TARGET)
	cp $(TARGET).a $(TARGET).dylib $(INSTALL)

clean:
	$(RM) -f *.o *.a *.dylib
