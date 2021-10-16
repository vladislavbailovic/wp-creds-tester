BINFILE = wpc
TESTED = .tested.out
COVERED = coverage.out
HTMLCOVERED = .coverage.html.out
GOFILES = $(shell find . -type f -name '*.go')

$(TESTED): $(GOFILES)
	make test

$(BINFILE): $(TESTED) $(COVERED) $(GOFILES)
	go build

$(COVERED): $(TESTED) $(GOFILES)
	go test ./... -coverprofile=$(COVERED)
	go tool cover -func=$(COVERED) | grep -v '100.0%'

$(HTMLCOVERED): $(TESTED) $(COVERED) $(GOFILES)
	go tool cover -html=$(COVERED) 
	touch $(HTMLCOVERED)
	

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
