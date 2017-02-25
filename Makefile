.PHONY: test uninstall
all: build

install: build
	@go install

uninstall:

clean:
	@go clean

build: clean
	@goimports -w .
	@gofmt -w .
	@go build .

test: build
	@go test -v ./...

run: build
	-@rm hoge.zip
	-@rm -rf bbbb
	@./spm install tzmfreedom/spm/sample/repositories/no_dependencies -u ${SF_USERNAME} -p ${SF_PASSWORD} -d bbbb
