# Use Red Hat Universal Base Image minimal as the base image
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

# Install necessary packages for building Go applications
RUN microdnf install -y golang && microdnf clean all && microdnf install iputils && microdnf install bind-utils && microdnf install iproute && microdnf install tcpdump && microdnf install wireshark-cli

# Set the working directory
WORKDIR /app

# Copy the Go application source code to the container
COPY shakeout.go .

# Build the Go application
RUN go build -o shakeout shakeout.go

# Expose port 9000
EXPOSE 9000

# Command to run the Go application
CMD ["./shakeout"]
