FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY *.go ./

RUN go build -o /go-arkcon

#EXPOSE 8080

#CMD [ "/go-arkcon" ]