run: build
	@./bin

build:
	@go build -o bin

docker-build:
	docker build -t tkgo:multistage -f Dockerfile.multistage .

docker-run:
	docker run -p 8080:8080 tkgo:multistage
