VERSION := 1.0.0
APP_NAME := xxl-job-tester
BUILD_DIR := build

.PHONY: all clean build-windows build-linux build-macos package

all: clean build-windows build-linux build-macos package

clean:
	rm -rf $(BUILD_DIR)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME).exe

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-macos

package:
	cd $(BUILD_DIR) && zip $(APP_NAME)-$(VERSION)-windows.zip $(APP_NAME).exe
	cd $(BUILD_DIR) && zip $(APP_NAME)-$(VERSION)-linux.zip $(APP_NAME)-linux
	cd $(BUILD_DIR) && zip $(APP_NAME)-$(VERSION)-macos.zip $(APP_NAME)-macos