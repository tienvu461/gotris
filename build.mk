.PHONY: build clean

NAME := gotris
ENV := local
GOOS := linux
GOARCH := amd64
HASH := $$(git rev-parse --short --verify HEAD)
DATE := $$(date -u '+%Y%m%dT%H%M%S')
GOVERSION = $$(go version)

build: $(NAME).$(ENV).$(GOOS).$(GOARCH)

$(NAME).$(ENV).$(GOOS).$(GOARCH):
	GOOS=$(GOOS) GOARCH=$(GOARCH) \
	    go build -tags=$(ENV) \
	        -o $(NAME) \
	        .
	mv $(NAME) $@

clean:
	-rm -rf $(NAME).$(ENV).$(GOOS).$(GOARCH)
