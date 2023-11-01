run:
	rm -rf build
	mkdir build
	cd http && go build -mod vendor -o ../build/fileserver
	cd build && ./fileserver

run-ipfs:
	ipfs daemon

generate-swagger:
	swag init -g http/server.go