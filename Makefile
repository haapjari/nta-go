MAIN_MODULE             := cmd/main.go
IMAGE_VERSION           := latest
SERVICE_NAME 			?= network-traffic-analyzer
OUTPUT_PATH 			?= bin/nta-go
OUTPUT_FILE 			?= report.txt

##  
## Commands
##

.PHONY: run dev test compile tools

tools: 
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

lint: tools
	golangci-lint run > $(OUTPUT_FILE) || { echo "Linting errors found"; cat $(OUTPUT_FILE); exit 1; }

fmt: 
	goimports -w -l .
	gofmt -w -l .

run: compile
	./${OUTPUT_PATH}

dev:
	go run -race cmd/main.go

test:
	go test -v -count=1 ./...

compile:
	go build -o ${OUTPUT_PATH} ${MAIN_MODULE}
