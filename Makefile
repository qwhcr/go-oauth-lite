all: build run

build:
	go build -o auth-server *.go


run:
	./auth-server

clean:
	rm ./auth-server

