# Build the manager binary
FROM golang:1.19 as builder

WORKDIR /app

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the source code to the container
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -o docustoreserver .

FROM scratch
WORKDIR /app
COPY --from=builder /app/docustoreserver /app/

ENTRYPOINT ["./docustoreserver"]
