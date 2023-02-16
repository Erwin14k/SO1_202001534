# Build
# docker build -t backend_practica1_202001534:1.00 .
# Tag
# docker tag backend_practica1_202001534:1.00 erwin14k/backend_practica1_202001534:1.00
# Push
# docker push erwin14k/backend_practica1_202001534:1.00


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

# Expose the port
EXPOSE 8080

# Specify the command to start our application
CMD ["./app"]


