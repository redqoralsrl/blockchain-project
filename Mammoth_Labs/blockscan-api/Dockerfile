FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy
RUN go mod download

COPY . .

RUN go get github.com/air-verse/air
RUN go install github.com/air-verse/air

WORKDIR /app/cmd

CMD ["air", "-c", "air.toml"]