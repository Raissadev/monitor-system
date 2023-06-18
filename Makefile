GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOINSTALL := $(GOCMD) install

BUILD_TARGET := kenbunshoku-haki
INSTALL_PATH := /usr/bin/$(BUILD_TARGET)

all: build

build:
	$(GOBUILD) -o ./src/bin/$(BUILD_TARGET)

install:
	$(GOBUILD) -o ./src/bin/$(BUILD_TARGET)
	sudo ln -sf "$(CURDIR)/src/bin/$(BUILD_TARGET)" $(INSTALL_PATH)
	img2txt ./etc/mug.png

clean:
	$(GOCLEAN)
	rm -f ./bin/$(BUILD_TARGET)
	rm -f $(INSTALL_PATH)