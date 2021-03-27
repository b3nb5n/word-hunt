FROM golang:alpine

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .

RUN go build -o main .

WORKDIR /build

RUN mv /src/main .

EXPOSE 8080

CMD ["/build/main"]
