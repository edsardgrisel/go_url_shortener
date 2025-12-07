MAIN = cmd/app/main.go
EXECUTABLE_PATH = bin/main

all:
	go build -o $(EXECUTABLE_PATH) $(MAIN)

run:
	./$(EXECUTABLE_PATH)

dev:
	go run $(MAIN)

.PHONY: all run dev clean re