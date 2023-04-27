# syntax=docker/dockerfile:1

FROM golang:1.20
WORKDIR /api
COPY . .
#RUN go get .
#RUN go install github.com/cosmtrek/air@latest
RUN CGO_ENABLED=1 GOOS=linux go build -o /app
EXPOSE 8080
CMD ["/app"]
