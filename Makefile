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

.PHONY: test
test:
	@mkdir -vp test
	@if ! test -f test/shortenertest; then \
	  wget -P test/ https://github.com/Yandex-Practicum/go-autotests/releases/download/v0.10.3/shortenertest; \
	  chmod +x test/shortenertest; \
	fi
	@cd test && ./shortenertest -test.v -test.run=^TestIteration1$$ -binary-path=../build/shortener/shortener

.PHONY: clean
clean:
	rm -rvf $(BUILD)
