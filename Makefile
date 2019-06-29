IMPORT_PATH=${PWD}

BINARY_LOCATION=${PWD}
BINARY_NAME=arthemis

MAIN_LOCATION=${PWD}
MAIN_LDFLAGS="-X $(IMPORT_PATH).Version=0.0.1"

all: build

build:
	go build -v -o $(BINARY_LOCATION)/$(BINARY_NAME) -ldflags $(MAIN_LDFLAGS) $(MAIN_LOCATION)/main.go

run:
	$(BINARY_LOCATION)/$(BINARY_NAME)

clean:
	rm -f $(BINARY_LOCATION)/$(BINARY_NAME)

test:
	go test -v -race -coverprofile=coverage.out ./...
