FROM golang:1.21

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main .
EXPOSE 8080

CMD ["/app/main"]
