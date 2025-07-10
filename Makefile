define \n


endef
VERSION := $(file < ./version.txt)
AIR_CONF := $(file < ./.air.toml)
PACKAGE := github.com/MrRainbow0704/ProgettoIride
LDFLAGS := -ldflags="-X '$(PACKAGE)/internal/version.version=$(VERSION)-dev'"
LDFLAGS_R := -ldflags="-X '$(PACKAGE)/internal/version.version=$(VERSION)' -H windowsgui"
LDFLAGS_L := $(subst /,\/,$(LDFLAGS_R))
AIR_CONF := $(subst @LDFLAGS@,$(subst ",\",$(LDFLAGS)),$(AIR_CONF))
AIR_CONF := $(subst ${\n}${\n},${\n},$(AIR_CONF))
.PHONY: build release clear go go_r

build: clear go

release: clear go_r

go:
	go mod download
	go mod tidy
	go tool rsrc -ico iride.ico -manifest iride.manifest
	bash -c "mv *.syso ./cmd/iride/"
	go build $(LDFLAGS) -o ./bin/ ./cmd/...
	bash -c "rm -f ./cmd/iride/*.syso"

go_r:
	go mod download
	go mod tidy
	go tool rsrc -ico iride.ico -manifest iride.manifest
	bash -c "mv *.syso ./cmd/iride/"
	go build $(LDFLAGS_R) -o ./bin/ ./cmd/...
	bash -c "rm -f ./cmd/iride/*.syso"

live:
	go mod download
	go mod tidy
	go tool rsrc -ico iride.ico -manifest iride.manifest
	bash -c "mv *.syso ./cmd/iride/"
	mkdir tmp
	echo. > ./tmp/.air.toml
	echo $(subst ${\n}, >> ./tmp/.air.toml ${\n}echo ,$(AIR_CONF)) >> ./tmp/.air.toml
	go tool air -c ./tmp/.air.toml

clear:
	bash -c "rm -rf ./bin"
	bash -c "rm -rf ./tmp"
	bash -c "rm -f ./cmd/iride/*.syso"
