run: 
	go build -ldflags "-s -w" -o ./bin/getset cmd/*.go;
	./bin/getset