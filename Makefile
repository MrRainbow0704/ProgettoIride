VERSION = $(shell cat ./version.txt)
PACKAGE = github.com/MrRainbow0704/ProgettoIride
LDFLAGS = -ldflags="-H windowsgui -X '$(PACKAGE)/internal/version.version=$(VERSION)-dev'"
LDFLAGS_R = -ldflags="-H windowsgui -X '$(PACKAGE)/internal/version.version=$(VERSION)'"


.PHONY: build release clear go go_r

build: clear go

release: clear go_r

go:
	go mod download
	go mod tidy
	go tool rsrc -ico iride.ico
	bash -c "mv *.syso ./cmd/iride/"
	go build $(LDFLAGS) -o ./bin/ ./cmd/...

go_r:
	go mod download
	go mod tidy
	go tool rsrc -ico iride.ico
	bash -c "mv *.syso ./cmd/iride/"
	go build $(LDFLAGS_R) -o ./bin/ ./cmd/...

clear:
	bash -c "rm -f ./bin/*"
	bash -c "rm -f ./cmd/iride/*.syso"
