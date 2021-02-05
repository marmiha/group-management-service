FROM golang:1.15.6-alpine3.12
WORKDIR /group-management-api
ENV PORT=3000
COPY . .
RUN go build cmd/start/main.go
EXPOSE $PORT
ENTRYPOINT ["/group-management-api/main"]