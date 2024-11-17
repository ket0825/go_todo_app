# Build stage
# FROM golang:1.23.1 AS deploy-builder

# WORKDIR /app

# COPY go.mod. go.sum ./
# RUN go mod download

# COPY . .
# RUN go build -trimpath -ldflags "-w -s" -o app

# Build stage
FROM golang:1.23.1 AS deploy-builder

WORKDIR /app

COPY . .
RUN go mod download && \
    go build -trimpath -ldflags "-w -s" -o app

# Deployment stage

FROM debian:bullseye-slim as deploy

RUN apt-get update

# 빌드된 바이너리를 복사
COPY --from=deploy-builder /app/app .

# 실행 명령
CMD ["./app"]

# Development stage with Air hot reload
FROM golang:1.23.1 AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]