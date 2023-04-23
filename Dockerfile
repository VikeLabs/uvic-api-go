# syntax=docker/dockerfile:1

FROM golang:1.20.3-alpine3.17
WORKDIR /api
COPY . .
RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest
EXPOSE 8000
CMD ["air"]
