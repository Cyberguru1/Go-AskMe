# Use the official Golang base image for building the bot
FROM golang:1.18 as build

# Set the working directory inside the container
WORKDIR /app

# Copy your Go source code into the container
COPY . .

# Build the Go application
RUN go build -o askme .

# Create a minimal runtime image
FROM gcr.io/distroless/base-debian10

# Set the working directory inside the runtime container
WORKDIR /app

# Copy the bot binary from the build container
COPY --from=build /app/askme .

# Copy your .env file into the container
COPY .env .env

# Run the bot
CMD ["./askme"]
