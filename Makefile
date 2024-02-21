TL_GOPRIVATE=github.com/Terralayr/*
GH_SSH_PK_PATH=/home/theo/.ssh/trlyr-github

tidy:
	@go env -w GOPRIVATE=${TL_GOPRIVATE}
	@git config --global url."git@github.com:Terralayr".insteadOf "https://github.com/Terralayr"
	@go mod tidy

build: tidy fmt
	@CGO_ENABLED=0 GOOS=linux go build -o target/main ./cmd

fmt:
	@golines --shorten-comments --base-formatter gofmt -m 100 -w .

docker-build: tidy
	docker build \
		--ssh default=${GH_SSH_PK_PATH} \
		.

up: down tidy docker-build
	SSH_AUTH_SOCK=${GH_SSH_PK_PATH} docker compose up --build

down:
	docker compose down --volumes

.PHONY: tidy build fmt docker-build up down
