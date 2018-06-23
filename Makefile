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
top_packages := $(glide nv)
top_package_patterns = $(filter-out .,$(filter-out ./gen/...,$(top_packages)))
top_package_names = $(top_package_patterns:./%/...=%)
all_package_paths := $(go list ${top_package_patterns})
all_package_names = $(all_package_paths:$(package_basename)/%=%)
top_package_tests = $(top_package_names:%=test.%)
top_package_test_reports = $(top_package_names:%=$(out_test_path)/%.func.txt)
all_test_report = $(out_test_path)/all.func.txt
all_test_report_html = $(out_test_path)/all.func.html
minimum_coverage_percent := $(ENV_MIN_COVERAGE)

build: test fmt build

generate:
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
	@echo "Cleaing the output folders"
	$(GOCLEAN)
	rm -rf $(OUTFOLDER)

run:
	$(GORUN) $(apps_folder)/main.go

fmt:
	@echo "Checking formatting of go sources"
	@result=$$(gofmt -d -l -e $(top_package_names) 2>&1); \
		if [[ "$$result" ]]; then \
			echo "$$result"; \
			echo 'gofmt failed!'; \
			exit 1; \
		fi

.PHONY: fixfmt
fixfmt:
	@echo "Fixing format of go sources"
	@gofmt -l -e $(top_package_names) 2>&1; \
		if [[ "$$?" != 0 ]]; then \
		    echo "gofmt failed! (exit-code: '$$?')"; \
		    exit 1; \
		fi

coverage:
	$(GOTEST) -cover -v ./...