
export GOPATH=$(shell pwd)/../
export GOBIN=$(GOPATH)/bin

COMM_PATH=./comm/

EXCELJSON_PATH = ./exceljson 
CEHUA_SVNURL = https://192.168.150.238/svn/GiantCode/wxcat/cehua/$(USER)/

SUB_DIRS = chaos testclient

all: debug 

debug:
	@echo GOPATH:$(GOPATH)
	@for dir in $(SUB_DIRS); do \
		go install -gcflags "-N -l" ./$$dir || exit 1; \
	done
	@echo Done

proto:
	protoc -I=$(COMM_PATH) --proto_path=$(COMM_PATH) --go_out=. $(COMM_PATH)/*.proto

clean:
	@for bdir in $(SUB_DIRS); do \
		rm -rf ../bin/$$bdir
	done
	@find -name "*~" | xargs rm -f
	@find -name "*.swp" | xargs rm -f
	@rm -rf release/*
wc:
	@find . -iname \*.go -exec cat \{\} \; | wc -l

tags:
	@ctags -R

msg:
	@cd github.com/liasece/micserver/tools && ./makeservermsg.sh
	@python3 github.com/liasece/micserver/tools/go2go.py -i ./command/command.go -o go --onlynames

.PHONY: all debug clean wc image



