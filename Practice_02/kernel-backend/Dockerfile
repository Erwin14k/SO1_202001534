# Build
# docker build -t so1_practice02_backend:1.00 .
# Tag
# docker tag so1_practice02_backend:1.00 erwin14k/so1_practice02_backend:1.00
# Push
# docker push erwin14k/so1_practice02_backend:1.00


# The Alpine base image for a smaller size of the resulting image
FROM golang:alpine

# Define a directory for our code inside the container
WORKDIR /app

# Copy the code of our application in the container
COPY . .

# Install the necessary dependencies for our application
RUN go mod download

# Compile the application
RUN go build -o ./app .

# Specify the command to start our application
CMD ["./app"]