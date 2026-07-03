MODULE_PATH := $(shell go list -m)
VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
ifndef RELEASE
	VERSION := $(VERSION)-dev
endif
LDFLAGS := -X '$(MODULE_PATH)/internal/version.version=$(VERSION)' -H windowsgui

.PHONY: run before pre build post

run: before pre build post

build:
	go build -ldflags "$(LDFLAGS)" -o ./bin/ ./cmd/...

before:
	go mod tidy
	go vet ./...

pre:
	cp iride.manifest ./cmd/iride/
	cd ./cmd/iride; sed -i 's/$$(VERSION)/$(shell echo $(VERSION) | cut -c2-)/g' iride.manifest
	cd ./cmd/iride; go tool rsrc -ico ./../../internal/gui/assets/iride.ico -manifest ./iride.manifest
	cd ./internal/gui/assets; go tool fyne bundle -prefix Resource -pkg assets -o ico.go iride.ico

post:
	cd ./cmd/iride; rm -f *.syso iride.manifest