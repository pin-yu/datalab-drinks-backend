FROM golang:1.16.3

WORKDIR /app
COPY . .

RUN go build src/main.go

ENV GIN_MODE=release

# set timezone
ENV TZ=Asia/Taipei

EXPOSE 5002

ENTRYPOINT [ "./main" ]