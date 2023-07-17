# --------------- BUILD STAGE --------------- #
FROM golang:1.19-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# ------------------- DEV ------------------- #

FROM builder AS DEV
WORKDIR /app

RUN apk add --no-cache make
RUN go install github.com/cespare/reflex@latest

COPY . .

ENTRYPOINT ["/app/scripts/docker-entry.sh"]
CMD ["make watch"]

# ----------------- DOCKER ------------------ #
FROM builder as prod_builder

COPY . .

RUN go build -ldflags "-w -s" -o arcadia_server

FROM alpine:3.17 AS DOCKER
WORKDIR /app

COPY --from=prod_builder /app/arcadia_server .
COPY --from=prod_builder /app/config.json .
COPY --from=prod_builder /app/.env .
COPY --from=prod_builder /app/scripts/docker-entry.sh .

ENTRYPOINT ["/app/docker-entry.sh"]
CMD ["/app/arcadia_server"]
