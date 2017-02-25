.PHONY: test install uninstall

test: 
	@go test -v ./...

install: build
	@go install

uninstall:

clean:
	@go clean

build: clean
	@goimports -w .
	@gofmt -w .
	@go build .

run: build
	@./spm install tzmfreedom/spm/sample/repositories/no_dependencies -u ${SF_USERNAME} -p ${SF_PASSWORD}
