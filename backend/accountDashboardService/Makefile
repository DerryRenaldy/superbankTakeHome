# Variables
DOCKER_CMD = docker run -it --name dashboard_db -e MYSQL_ROOT_PASSWORD=root -e MYSQL_PASSWORD=root -d -v ~/Desktop/DatabaseDocker/AuthDB:/var/lib/mysql -p 33061:3306 mysql:latest

# Default target
all: build

# Build target for your Go project
build:
	go build -o myapp .

# Target to run the Docker command
run-mysql:
	$(DOCKER_CMD)

run-test:
	GO_TEST=1 go test -v ./...

# Clean up Docker container and Go build files
clean:
	docker stop sql_user || true
	docker rm sql_user|| true
	rm -f myapp

.PHONY: all build run-mysql clean
