FROM golang:go1.20.4

WORKDIR /cmd

COPY go.mod go.sum ./

RUN go mod download

# Copy whatever is left
COPY . .

RUN go build -o ./cmd

EXPOSE 8080

CMD ["./cmd"]