BINFILE = wpc
TESTED = .tested.out
COVERED = coverage.out
HTMLCOVERED = coverage.html.out
GOFILES = $(shell find pkg/ -type f -name '*.go')

$(TESTED): $(GOFILES)
	make test

$(BINFILE): $(TESTED) $(COVERED) $(GOFILES)
	go build

$(COVERED): $(TESTED) $(GOFILES)
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out | grep -v '100.0%'

$(HTMLCOVERED): $(TESTED) $(GOFILES)
	go tool cover -html=coverage.html.out 
	

test:
	go test ./...
	touch $(TESTED)

build:
	make $(BINFILE)

cover: 
	make $(COVERED)

cover-html: cover
	make $(HTMLCOVERED)

clean:
	-rm $(TESTED) $(COVERED) $(HTMLCOVERED) $(BINFILE)
