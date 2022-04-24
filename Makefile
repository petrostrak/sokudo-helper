BINARY_NAME=sokudoApp

build:
	@go mod vendor
	@echo "Building Sokudo..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Sokudo built!"

run: build
	@echo "Starting Sokudo..."
	@./tmp/${BINARY_NAME} &
	@echo "Sokudo started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start_compose:
	sudo docker-compose up -d

stop_compose:
	sudo docker-compose down

start: run

stop:
	@echo "Stopping Sokudo..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Sokudo!"

restart: stop start