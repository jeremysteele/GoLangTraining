FROM golang:1.12 as build-env

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o api ./cmd/worker/main.go

FROM alpine

WORKDIR /app

COPY --from=build-env /src/api /app/

CMD "/app/api" 