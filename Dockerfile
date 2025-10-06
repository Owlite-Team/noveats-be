FROM golang:1.23-alpine AS buildstage

WORKDIR /app

RUN apk add --no-cache ca-certificates git

# Set GOTOOLCHAIN to auto to allow downloading the required Go version
ENV GOTOOLCHAIN=auto

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=buildstage /app/bin/server .
COPY --from=buildstage /app/.env ./

EXPOSE 8080

CMD ["./server"]
