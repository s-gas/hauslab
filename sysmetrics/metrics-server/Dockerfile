FROM golang:1.26

WORKDIR /usr/src/app

EXPOSE 1024

COPY . .
RUN go build -o /usr/local/bin/metrics-server ./...

CMD ["metrics-server"]
