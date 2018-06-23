 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME=swivel-server
OUTFOLDER=./out
BINARY_UNIX=$(OUTFOLDER)/$(BINARY_NAME)_unix
BINARY_WINDOWS=$(OUTFOLDER)/$(BINARY_NAME)_windows
BINARY_DARWIN=$(OUTFOLDER)/$(BINARY_NAME)_darwin
BINARY_LINUX=$(OUTFOLDER)/$(BINARY_NAME)_linux
apps_folder=apps/main

all: test build
build:
	@echo "Building windows executable $@"
	@GOOS=windows GOARCH=amd64 go build -o $(BINARY_WINDOWS) ./$(apps_folder)/$* 
	chmod 755 $(BINARY_WINDOWS)
	@echo "Building darwin executable $@"
	@GOOS=darwin GOARCH=amd64 go build -o $(BINARY_DARWIN) ./$(apps_folder)/$*
	chmod 755 $(BINARY_DARWIN)
	@echo "Building linux executable $@"
	@GOOS=linux GOARCH=amd64 go build -o $(BINARY_LINUX) ./$(apps_folder)/$*
	chmod 755 $(BINARY_LINUX)
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -rf $(OUTFOLDER)
run:
	$(GORUN) $(apps_folder)/main.go
deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop

