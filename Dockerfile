FROM golang:1.21 as setup-base
RUN mkdir -p -m 0600 ~/.ssh && \
	ssh-keyscan github.com >> ~/.ssh/known_hosts

RUN git config --global url."git@github.com:".insteadOf "https://github.com/"
WORKDIR /app
COPY go.mod go.su[m] ./
RUN --mount=type=ssh go mod download

FROM setup-base as setup-builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/main ./cmd/main.go

FROM alpine:latest
COPY --from=setup-builder /app/build/main /usr/bin
EXPOSE 8080
ENTRYPOINT ["main"]
