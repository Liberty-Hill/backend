FROM golang:1.23

WORKDIR /cmd

COPY go.mod go.sum ./

RUN go mod download

COPY *.go .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server

EXPOSE ${PORT}

CMD ["./server"]