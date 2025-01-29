# Tkgo

## _Token Pool for tracking, Simulation_

Tkgo is a quick, runtime memory, offline-storage for Token pool

- Written in Golang
- Uses golang's self sufficient standard library
- Quick, read-write safe, easy to use

## Features

- Create a new user with a predefined number of tokens and a simulation time limit.
- Retrieve the least-used token for a given user.
- Automatically stop token allocation once the simulation limit is reached.
- Uses an in-memory storage system for fast operations.

## Tech

Tkgo uses a number of projects to work properly:

- [Golang] - Golang language
- [net/http] - mux, servering
- [Zap (Logging)] - Logging info and errors for better behaviour understanding

## Installation

Tkgo requires [Golang](https://go.dev/) to run.

Install the dependencies and devDependencies and start the server.

```sh
# Clone the repository
git clone https://github.com/briheet/Tkgo.git
cd tkgo

# Normal build and run
make
```

## Docker

Tkgo is very easy to use and deploy in a Docker container.

By default, the Docker will expose port 8080, so change this within the
Dockerfile if necessary. When ready, simply use the Dockerfile to
build the image.

```sh
# Enter the project directory
cd Tkgo

# Directly build the image
docker build -t tkgo:multistage -f Dockerfile.multistage .

# Or use Makefile
make docker-build
```

This will create the Tkgo image and pull in the necessary dependencies.

Once done, run the Docker image and map the port to whatever you wish on
your host. For now, we simply map port 8080 of the host to
port 8080 of the Docker (or whatever port was exposed in the Dockerfile):

```sh
# Directly run the image
docker run -p 8080:8080 tkgo:multistage

# Or use the Makefile
make docker-run
```

Verify the deployment by navigating to your server address in
your preferred browser.

```sh
http://localhost:8080/health
```

## Development

Going on. Want to contribute? Make a pr :)

[//]: # "These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax"
[net/http]: https://pkg.go.dev/net/http
[Zap (Logging)]: https://github.com/uber-go/zap
[Golang]: http://go.dev
