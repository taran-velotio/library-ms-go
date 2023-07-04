FROM golang:1.20.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy whatever is left
COPY . .

RUN go build -o main cmd/main.go

CMD ["/app/main"]