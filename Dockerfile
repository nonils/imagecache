FROM golang:latest
LABEL maintainer="Sebastian Bogado <seebogado@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 9090
CMD ["./main"]