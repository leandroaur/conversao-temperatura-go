# Build stage
FROM golang:1.16-alpine AS build
WORKDIR /app
COPY . .
RUN apk add git
RUN go mod download
RUN go build -o convert

# Run stage
FROM alpine:3.12
WORKDIR /app
COPY --from=build /app/convert .
CMD ["./convert"]
EXPOSE 8080

