FROM golang:1.18.4-alpine3.16 as builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o chaos-monkey .

FROM alpine:3.16

ENV USER=chaos
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /app
COPY --from=builder /app/chaos-monkey .
USER chaos

CMD ["sh", "-c", "/app/chaos-monkey"]
