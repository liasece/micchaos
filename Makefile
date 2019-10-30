
export GOPATH=$(shell pwd)/../
export GOBIN=$(GOPATH)/bin

COMM_PATH=./comm/

SUB_DIRS = chaos testclient

all: debug 

debug:
	@for dir in $(SUB_DIRS); do \
		go install -gcflags "-N -l" ./$$dir || exit 1; \
	done

wc:
	@find . -iname \*.go -exec cat \{\} \; | wc -l

tags:
	@ctags -R

msg:
	@cd github.com/liasece/micserver/tools && ./makeservermsg.sh
	@python3 github.com/liasece/micserver/tools/go2go.py -i ./ccmd/ccmd.go -o go --onlynames
	@../bin/easyjson -all ./ccmd/ccmd.go

.PHONY: all debug clean wc image



