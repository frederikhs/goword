all: download vet lint test

download:
	go mod download

vet:
	go vet ./...

lint:
	golangci-lint run

test:
	go test ./... -v

cover:
	go test -cover -coverprofile=coverage.out -v ./...
	go tool cover -html=coverage.out

clean:
	git clean -fxd -e .idea

act: act-lint-122 act-test-122

act-lint-122:
	act push --workflows ./.github/workflows/lint.yml

act-test-122:
	act push --workflows ./.github/workflows/test.yml --matrix go-version:1.22.x
