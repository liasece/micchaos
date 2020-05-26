
all: update debug 

debug:
	@mkdir -p bin
	@go build -o bin/ github.com/liasece/micchaos/...

update: 
	@go get -u github.com/liasece/micserver

# go get -u github.com/mailru/easyjson...
msg:
	@python3 tools/go2go.py -i ./ccmd/ccmd.go -o go --onlynames
	@easyjson -all ./ccmd/ccmd.go

.PHONY: all debug msg update



