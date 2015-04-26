
GO := go
GOPATH := $(PWD):$(GOPATH)

TARGETS := bomber-win64.exe bomber-osx64 bomber-linux64
SOURCES := src/bomber.go src/bom/bom.go

all: $(TARGETS)

test:
	$(GO) test -v src/bom/*.go
	$(GO) test -i src/*.go
	$(GO) test -v src/*.go

clean:
	rm bomber-*

bomber-win32.exe: $(SOURCES)
	GOOS=windows GOARCH=386 $(GO) build -o $@ $<

bomber-win64.exe: $(SOURCES)
	GOOS=windows GOARCH=amd64 $(GO) build -o $@ $<

bomber-osx32: $(SOURCES)
	GOOS=darwin GOARCH=386 $(GO) build -o $@ $<

bomber-osx64: $(SOURCES)
	GOOS=darwin GOARCH=amd64 $(GO) build -o $@ $<

bomber-linux32: $(SOURCES)
	GOOS=linux GOARCH=386 $(GO) build -o $@ $<

bomber-linux64: $(SOURCES)
	GOOS=linux GOARCH=amd64 $(GO) build -o $@ $<


