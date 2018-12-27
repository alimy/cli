.PHONY: default clean

LDFLAGS += -X "github.com/alimy/cli/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/cli/version.BuildGitHash=$(shell git rev-parse HEAD)"

default:
	go build -ldflags '$(LDFLAGS)' .

clean:
	rm cli