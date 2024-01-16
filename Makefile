BUILD := build
CMD := ./cmd
APP := shortener
MODULE := github.com/aleks0ps/url-service

.PHONY: all
all: build

go.mod:
	go mod init $(MODULE)

.PHONY: build
build: go.mod
	mkdir -vp $(BUILD)/$(APP)
	mkdir -vp $(BUILD)/client
	# use local path, otherwise 'go build' will lookup global dir /usr/local/go/src/cmd/ 
	go build -o $(BUILD)/$(APP)/$(APP) $(CMD)/$(APP)
	go build -o $(BUILD)/client/client $(CMD)/client

.PHONY: clean
clean:
	rm -rvf $(BUILD)
