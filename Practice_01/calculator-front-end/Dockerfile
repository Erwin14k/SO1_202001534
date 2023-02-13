## Build
# docker build -t calculator-front-end:1.00 .

## Run
# docker run -p 3000:3000 -d calculator-front-end:1.00

# Tag
# docker tag calculator-front-end:1.00 erwin14k/calculator-front-end:1.00
# Push
# docker push erwin14k/calculator-front-end:1.00

## Enter to the component
# docker run -it full-front:0.1.0 /bin/bash

# Nonde alpine image
FROM node:alpine

# Container directory
WORKDIR /app

# Copy the code to the container
COPY . .

# Install all packages
RUN npm install

# Build application
RUN npm run build

# Expose the port
EXPOSE 3000

# Init the application
CMD ["npm", "start"]
