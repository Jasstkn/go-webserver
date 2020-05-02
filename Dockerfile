FROM golang:alpine3.11

RUN mkdir /app

COPY main.go /app/

WORKDIR /app

RUN addgroup -g 1001 -S app && \
    adduser -u 1002 -G app -H -h /app -s /sbin/nologin app -D && \
    chown -R app:app /app

USER app

RUN go build -o main .

CMD ["/app/main"]