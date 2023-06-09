# syntax=docker/dockerfile:1

FROM golang:1.20
WORKDIR /api
COPY . .
RUN go build -o /app -buildvcs=false
EXPOSE 8080
CMD ["/app"]
