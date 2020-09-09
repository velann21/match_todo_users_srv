FROM golang:latest as builder
WORKDIR /app/backend
ADD . /app/backend
RUN go mod download
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /users_srv /app/backend/main.go


# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY ./pkg/migration_scripts/* ./
COPY --from=builder /users_srv ./
RUN chmod +x ./users_srv
ENTRYPOINT ["./users_srv"]
EXPOSE 8081