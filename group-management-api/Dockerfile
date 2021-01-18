FROM golang:1.15.6-alpine3.12
ENV PORT=3000
WORKDIR /group-management-api
COPY . .
RUN go build main.go
EXPOSE $PORT
ENTRYPOINT ["/group-management-api/main"]