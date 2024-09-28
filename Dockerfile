FROM golang:1.23

WORKDIR /cmd

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server

EXPOSE 3001

CMD ["./server"]