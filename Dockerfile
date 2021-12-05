FROM golang:alpine as builder

RUN apk --no-cache add git
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build

FROM alpine
WORKDIR /
COPY --from=builder /build/tracker .
EXPOSE 8000
CMD ["./tracker"]
