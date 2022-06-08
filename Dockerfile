FROM golang:1.18 AS builder

RUN mkdir -p /app /out
WORKDIR /app
ADD go.mod go.sum /app/
RUN go mod download
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -a -o /out/exporter ./

FROM alpine:3.16.0
COPY --from=builder /out/exporter ./
EXPOSE 3093
CMD ["/exporter"]
