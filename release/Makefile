APP=providers
ENV=release
VERSION=v2.0.0
DATE=`date +%FT%T%z`
EXECUTE=/usr/local/go/bin/go
GOARCH=`go env GOARCH`

ifeq ($(shell uname), Darwin)
	PRINT=@echo
else
	PRINT=@echo -e
endif

define begin
	$(PRINT) "\t->当前系统: \t\033[1;36m[ $(1) ]\033[0m\t";
	$(PRINT) "\t->系统平台: \t\033[1;36m[ $(2) ]\033[0m\t";
	$(PRINT) "\t->正在编译: \t\033[1;36m[ ... ]\033[0m\t";
endef

define finish
	$(PRINT) "\t->编译程序:\t\033[1;32m[ OK ]\033[0m";
endef

define compile
	@GO111MODULE=on GOOS=$(1) GOARCH=$(2) $(EXECUTE) build -o build/$(1)/$(2)/$(3)/bin/$(3) -ldflags "-X main.ENV=$(ENV) -X main.VERSION=$(VERSION) -X main.APP=$(APP)" ./
endef

default:
	$(PRINT) "你可以使用: \n\t make all \t 编译所有平台 \n\t make linux \t 来编译 linux 平台数据 \n\t make window \t 编译 window 平台执行文件 \n\t make darwin \t 编译 MacOS 平台执行文件"

darwin:
	$(call begin,darwin,$(GOARCH))
	$(call compile,darwin,$(GOARCH),$(APP))
	$(call finish)

linux:
	$(call begin,linux,386)
	$(call compile,linux,386,$(APP))
	$(call finish)
	$(call begin,linux,amd64,$(APP))
	$(call compile,linux,amd64,$(APP))
	$(call finish)

window:
	$(call begin,windows,386)
	$(call compile,windows,386,$(APP).exe)
	$(call finish)
	$(call begin,windows,amd64)
	$(call compile,windows,amd64,$(APP).exe)
	$(call finish)

.PHONY:
	$(PRINT) "编译所有平台可执行程序\t\033[1;32m[ 请等待 ]\033[0m";

all: .PHONY darwin linux window
	$(PRINT) "所有平台可执行程序编译\t\033[1;32m[ 完成 ]\033[0m";

clear:
	@rm -rf build

install:
	@cp build/darwin/$(APP) `go env GOPATH`/bin/