FROM golang:1.25.5-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o generator

CMD ["./generator", "-file", "-path", "Pirates.csv"]