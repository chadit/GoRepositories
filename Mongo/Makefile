# GoRepositories.Mongo Makefile

# Variables
GO=go
COMMIT=$(shell git rev-parse HEAD)

# Variables to be overridden on the command line, e.g.
# make test PKG=./v1
PKG:= ./...
BUILDMODE:=build
TESTFLAGS:=
COVEROUT:=cover.out
COVERTXT:=cover.txt

define GIT_CONFIG
[url "git@github.com:"]
	    insteadOf = https://github.com/
endef

export GIT_CONFIG 

test: TESTFLAGS += $(shell ./dev_variables.sh "test")
test:
	$(GO) test -v $(PKG) $(TESTFLAGS)
.PHONY: test

race: TESTFLAGS += $(shell ./dev_variables.sh "test")
race:
	$(GO) test -race $(PKG) $(TESTFLAGS)
.PHONY: race

bench: TESTFLAGS += $(shell ./dev_variables.sh "test")
bench:
	$(GO) test -bench $(PKG) $(TESTFLAGS)
.PHONY: bench

# Creates a temporary cover profile file per package and merges them into one file.
getcover:
	$(GO) get github.com/wadey/gocovmerge
	C=0;$(GO) list $(PKG) | while read n; do C=$$(($$C+1)); go test $$n $(TESTFLAGS) -coverprofile "$$C.coverprofile"; done
	gocovmerge `ls *.coverprofile` > $(COVEROUT)

cover: TESTFLAGS += $(shell ./dev_variables.sh "test")
cover: getcover
	echo "Tested On $(shell date -u +"%D %r %Z")" > $(COVERTXT)
	$(GO) tool cover -func=$(COVEROUT) >> $(COVERTXT)
	rm *.coverprofile $(COVEROUT)
.PHONY: cover

html: TESTFLAGS += $(shell ./dev_variables.sh "test")
html: getcover
	$(GO) tool cover -html=$(COVEROUT)
	rm *.coverprofile $(COVEROUT)
.PHONY: html