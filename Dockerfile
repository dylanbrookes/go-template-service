FROM --platform=linux/amd64 golang:alpine AS builder

ARG GH_TOKEN

# Build Stage
RUN apk add git ca-certificates
RUN git config --global url."https://${GH_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN go env -w GOPRIVATE=github.com/external-repo/*
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app

# Deploy Stage
FROM alpine:latest
WORKDIR /template-service
COPY --from=builder /build/app .
EXPOSE 5000
CMD [ "/template-service/app" ]
