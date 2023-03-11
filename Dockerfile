# Build stage
FROM golang:1.16-alpine AS build
WORKDIR /app
COPY . .
RUN apk add git
RUN go build -o temperature

# Run stage
FROM alpine:3.12
WORKDIR /app
COPY --from=build /app/temperature .
CMD ["./temperature"]
EXPOSE 8080
