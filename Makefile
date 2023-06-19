PACKAGE_NAME=homebot
PLATFORM ?=aarch64
VERSION ?=0.0.1
BUILD_DIR=build/
BUILD_FILE=$(BUILD_DIR)$(PACKAGE_NAME)-${PLATFORM}-${VERSION}

make: clean test build

test:
	go test ./...

build:
	go build -o $(BUILD_FILE)

clean:
	rm -rf $(BUILD_DIR)

run:
	$(BUILD_FILE)