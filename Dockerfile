FROM golang:1.19

WORKDIR /usr/src/app

#ENV GIN_MODE=release

COPY . .

RUN go mod init github.com/DoanRii/go-arkcon
RUN go mod tidy && go mod verify

RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 9222

CMD ["app"]