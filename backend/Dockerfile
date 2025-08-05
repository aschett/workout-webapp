FROM golang:1.24-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o workout-backend ./cmd/web

EXPOSE 4000

CMD ["./workout-backend"]