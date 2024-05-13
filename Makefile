PROJECT:=go-app

swag:
	@swag i -g init_router.go -dir app/admin/router --instanceName admin --parseDependency -o docs/admin

test:
	@go test ./... -v
	@go vet

pre-push:
    @commit_message=`git log -1 --pretty=format:"%s"`; \
    if [ -z "$$commit_message" ]; then \
        echo "Error: No commit message provided. Please use 'git commit -m' to add a commit message."; \
        exit 1; \
    fi

git-commit:
	@git add .
	@git commit -m "$(COMMIT)"

git-push: pre-push
	@git push origin $(BRANCH)

# make build-linux
compile:
	@# 32-Bit Systems
	@# FreeBDS
	@GOOS=freebsd GOARCH=386 go build -o target/$(PROJECT)-freebsd-386 main.go
	@# MacOS
	@GOOS=darwin GOARCH=386 go build -o target/$(PROJECT)-darwin-386 main.go
	@# Linux
	@GOOS=linux GOARCH=386 go build -o target/$(PROJECT)-linux-386 main.go
	@# Windows
	@GOOS=windows GOARCH=386 go build -o target/$(PROJECT)-windows-386 main.go
	@# 64-Bit
	@# FreeBDS
	@GOOS=freebsd GOARCH=amd64 go build -o target/$(PROJECT)-freebsd-amd64 main.go
	@# MacOS
	@GOOS=darwin GOARCH=amd64 go build -o target/$(PROJECT)-darwin-amd64 main.go
	@# Linux
	@GOOS=linux GOARCH=amd64 go build -o target/$(PROJECT)-linux-amd64 main.go
	@# Windows
	@GOOS=windows GOARCH=amd64 go build -o target/$(PROJECT)-windows-amd64 main.go

build:
	@docker build -t $(IMAGE):$(VERSION) .
	@echo "build successful"

docker-push:
	docker push $(IMAGE):$(VERSION)

## pre: swag test
pre: swag test

## git commit: git-commit git-push
git: pre git-commit git-push

## all: build、docker-push
all: build docker-push

.PHONY: build docker-push
