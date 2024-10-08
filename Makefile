up:
	@echo "Building images..."
	docker-compose up -d
	@echo "Images built."

down:
	@echo "Stopping compose..."
	docker-compose down
	@echo "Done!"

build:
	@echo "Stopping existing docker images"
	docker-compose down
	@echo "Building images..."
	docker-compose up --build -d
	@echo "Images were built successfully!"