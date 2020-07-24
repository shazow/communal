VERSION := $(shell git describe --tags --dirty --always 2> /dev/null || echo "dev")
LDFLAGS = "-X main.Version=$(VERSION) -w -s"
SOURCES = $(shell find . -type f -name '*.go')

BINARY = $(notdir $(PWD))
RUN = ./$(BINARY)

all: $(BINARY)

$(BINARY): $(SOURCES)
	go build -ldflags $(LDFLAGS) -o "$@"

build: $(BINARY)

clean:
	rm $(BINARY)

run: $(BINARY)
	$(RUN) --help

test:
	go test -vet "all" -timeout 5s -race ./...

deploy: build
	rsync -bavz --progress $(BINARY) ip.shazow.net:projects/news/$(BINARY)
	ssh ip.shazow.net fuser -s -k -HUP projects/news/$(BINARY)~ || echo "binary did not change"
