# Build
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /app/ipserver ./main.go

# Run (tiny image)
FROM scratch
COPY --from=build /app/ipserver /ipserver
EXPOSE 8080
ENTRYPOINT ["/ipserver"]
