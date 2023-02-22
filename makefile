CMD=go
GO_PATH=/usr/local/go/bin/
BIN_PATH=bin
SRC_PATH=src

all:clean build install

build:

		$(GO_PATH)/$(CMD) build -o $(BIN_PATH)/mydocker $(SRC_PATH)/*
install: 
		cp bin/mydocker /usr/bin/mydocker
		cp bin/mydocker /usr/local/bin/mydocker	
		mkdir -p /var/lib/mydocker/images
		mkdir -p /var/lib/mydocker/volumes
		mkdir -p /var/lib/mydocker/containers
		
uninstall:
		rm -rf /usr/bin/mydocker /usr/local/bin/docker
		rm -rf /bin/docker
clean: uninstall