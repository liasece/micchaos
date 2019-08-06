
export GOPATH=$(shell pwd)/../
export GOBIN=$(GOPATH)/bin

COMM_PATH=./comm/

EXCELJSON_PATH = ./exceljson 
CEHUA_SVNURL = https://192.168.150.238/svn/GiantCode/wxcat/cehua/$(USER)/

SUB_DIRS = ./chaos

all: debug 

debug:
	@cd github.com/liasece/micserver/tools && ./makeservermsg.sh
	@echo GOPATH:$(GOPATH)
	@go install -gcflags "-N -l" $(SUB_DIRS) || exit 1; 
	@echo Done

proto:
	protoc -I=$(COMM_PATH) --proto_path=$(COMM_PATH) --go_out=. $(COMM_PATH)/*.proto

clean:
	@rm -rf ../bin/src
	@find -name "*~" | xargs rm -f
	@find -name "*.swp" | xargs rm -f
	@rm -rf release/*
wc:
	@find . -iname \*.go -exec cat \{\} \; | wc -l

tags:
	@ctags -R

msg:
	@python3 github.com/liasece/micserver/tools/go2go.py -i ./command/command.go -o go 

.PHONY: all debug clean wc image



