FROM debian:bookworm

RUN apt-get update -y && apt-get install -y ca-certificates

WORKDIR /app

COPY train-api-arm64 /app/train-api

RUN chmod +x /app/train-api

CMD ["/app/train-api"]