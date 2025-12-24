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

mysql-start:
	@echo "$(BLUE)Starting MySQL database$(NC)"
	brew services start mysql

mysql-stop:
	@echo "$(RED)Stopping MySQL database$(NC)"
	brew services stop mysql

redis-start:
	@echo "$(BLUE)Starting Redis server$(NC)"
	brew services start redis

redis-stop:
	@echo "$(RED)Stopping Redis server$(NC)"
	brew services stop redis

services-start: mysql-start redis-start
	@echo "$(GREEN)All services started$(NC)"

services-stop: mysql-stop redis-stop
	@echo "$(GREEN)All services stopped$(NC)"

run: all services-start
	@echo "$(BLUE)Starting application$(NC)"
	./$(EXECUTABLE_PATH)

dev: services-start
	go run $(MAIN)

clean:
	@echo "$(YELLOW)Cleaning build files$(NC)"
	rm -rf bin

re: clean all

.PHONY: all run dev clean re