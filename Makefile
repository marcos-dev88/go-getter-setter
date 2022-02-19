run-file: 
	go build -ldflags "-s -w" -o ./bin/getset cmd/*.go;
	./bin/getset gbf

run-cli:
	go build -ldflags "-s -w" -o ./bin/getset cmd/*.go
	./bin/getset gbc -path="$(path)" -fn="$(fn)"

test-all:
	go test -v ./...