# Use a base image
FROM golang:1.20-alpine

# RUN mkdir /app/backend

# Set the working directory inside the container
WORKDIR /app/backend

# Définissez la variable CGO_ENABLED
ENV CGO_ENABLED=1

# Copy the necessary files from the host to the container
COPY . /app/backend

RUN apk add --no-cache gcc musl-dev

#install packages
RUN go mod download 

# Build the Go application
RUN go build -o server


# Expose a port for the Go application
EXPOSE 8080

# Set metadata labels
LABEL version="1.0" \
      authors="Pdione; papegueye; alpapie; mirak; papediop" \
      description="social-network-backend"

# Set the command to run when the container starts
CMD ["./server"]
