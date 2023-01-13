# from example template: https://hub.docker.com/_/golang
FROM golang:1.19

WORKDIR /usr/src/hello

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/hello ./...

CMD ["hello"]