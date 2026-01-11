.PHONY: build install link

build:
	go build -o walpha .

install:
	go install .

link: build
	sudo ln -sf $(PWD)/walpha /usr/local/bin/walpha
