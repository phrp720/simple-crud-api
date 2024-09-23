FROM golang:1.23.0-alpine3.13

WORKDIR /app

# Copy the source code
COPY . .

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o api .

#EXPOSE the port
EXPOSE 8080
# Run the executable
CMD ["./api"]