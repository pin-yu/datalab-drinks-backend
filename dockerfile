FROM golang:1.16.3

WORKDIR /app
COPY . .

RUN go build src/main.go

ENV GIN_MODE=release
EXPOSE 5002

ENTRYPOINT [ "./main" ]