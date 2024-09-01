build:
	go build -o goshell ./cmd/goshell
	chmod +x goshell

test-unit: build
	go test ./tests/unit

test-integration: build
	go test ./tests/integration

test: test-unit test-integration

clean:
	rm -f goshell
	rm -f testfile.txt
