# Stage 1: Build the Go binary
FROM golang:1.20 AS build

WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download the Go module dependencies
RUN go mod download

# Copy the application source code
COPY . .

ENV GOARCH=amd64

# Build the Go binary
RUN go build -o myapp

# Stage 2: Create the final Docker image
FROM gcr.io/distroless/base-debian11

WORKDIR /app
# Copy the Go binary from the build stage
COPY --from=build /app/myapp .

EXPOSE 8081

USER nonroot:nonroot

# Set the entrypoint for the Docker image
ENTRYPOINT ["./myapp"]
