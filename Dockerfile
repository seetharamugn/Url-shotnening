# Use an official Golang runtime as a parent image
FROM golang:1.17

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8082

# Define environment variables
ENV DB_HOST=db
ENV DB_PORT=3306
ENV DB_USER=root
ENV DB_PASSWORD=root
ENV DB_NAME=urlshort

# Run the command when the container starts
CMD ["./main"]