MAIN = cmd/app/main.go
EXECUTABLE_PATH = bin/main

RED = \033[0;31m
GREEN = \033[0;32m
BLUE = \033[0;34m
YELLOW = \033[1;33m
NC = \033[0m

all:
	@echo "$(BLUE)Compiling$(NC)"
	mkdir -p bin
	go build -o $(EXECUTABLE_PATH) $(MAIN)

run: all
	@echo "$(BLUE)Starting MySQL database$(NC)"
	brew services start mysql
	@echo "$(BLUE)Starting application$(NC)"
	./$(EXECUTABLE_PATH)

dev:
	go run $(MAIN)

stop:
	@echo "$(RED)Stopping MySQL database$(NC)"
	brew services stop mysql

clean:
	@echo "$(YELLOW)Cleaning build files$(NC)"
	rm -rf bin

re: clean all

.PHONY: all run dev clean re