BIN_DIR := ${PWD}/bin
OS := linux
ARCH := amd64

# define deps
#   cd ${DEP_DIR}
#   git submodule init
#   git submodule update --remote
# endef

define gobuild
	cd cmd/; go mod tidy
	cd cmd/; GOOS=${OS} GOARCH=${ARCH} go build $(1).go
	cd cmd/; mv $(1) ${BIN_DIR}
endef

build:
	mkdir -p ${BIN_DIR}
	$(call gobuild,main)

clean:
	rm -rf ${BIN_DIR}
