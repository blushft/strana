FROM golang:1.14-alpine AS builder
RUN apk add --no-cache gcc musl-dev

WORKDIR /build
COPY . .
RUN go mod download
RUN go generate ./...
RUN go build -o strana ./cmd/cli/main.go 

FROM alpine
LABEL maintainer="Blushft"

RUN apk add ca-certificates

WORKDIR /app
COPY --from=builder /build/strana /app/strana

EXPOSE 8863
EXPOSE 4442
EXPOSE 4443

CMD ["/app/strana"]
