.PHONY: build clean version run windows

BINARY=openmv-netcam

SRC_DIR=.
DIST_DIR=./dist
ASSETS_DIR=./assets

BUILD_ARCH=arm arm64 386 amd64 ppc64le riscv64 \
	mips mips64le mips64 mipsle loong64 s390x
BUILD_FLAGS=-s -w

build: clean $(BUILD_ARCH)
$(BUILD_ARCH):
	@mkdir -p $(DIST_DIR)/$@
	@rm -rf $(DIST_DIR)/$@/*
	@GOOS=linux GOARCH=$@ go build -ldflags="$(BUILD_FLAGS)" \
		-o $(DIST_DIR)/$@/$(BINARY) $(SRC_DIR)/*.go
	@cp -r $(ASSETS_DIR) $(DIST_DIR)/$@

windows:
	@mkdir -p $(DIST_DIR)/windows
	@rm -rf $(DIST_DIR)/windows/*
	@GOOS=windows GOARCH=amd64 go build -ldflags="$(BUILD_FLAGS)" \
		-o $(DIST_DIR)/windows/$(BINARY).exe $(SRC_DIR)/*.go
	@cp -r $(ASSETS_DIR) $(DIST_DIR)/windows

run:
	@go run $(SRC_DIR)/*.go --config $(ASSETS_DIR)/config.json

clean:
	@rm -rf $(DIST_DIR)
