FROM golang:1.20.0-alpine3.16 AS builder
RUN apk add --no-cache git gcc musl-dev
WORKDIR /go/src/portfolio_service
COPY . .
RUN go mod download
RUN go build -o /go/bin/portfolio_service

FROM alpine:3.16.0
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /go/bin/portfolio_service /portfolio_service
RUN mkdir /operation_data
ENTRYPOINT /portfolio_service
LABEL Name=PortfolioService Version=2.0
EXPOSE 2023
CMD ["portfolio_service"]