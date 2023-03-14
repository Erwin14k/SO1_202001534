# Build
# docker build -t so1_practice02_api:1.00 .
# Tag
# docker tag so1_practice02_api:1.00 erwin14k/so1_practice02_api:1.00
# Push
# docker push erwin14k/so1_practice02_api:1.00

# Nonde alpine image
FROM node:alpine

# Container directory
WORKDIR /app

# Copy important files
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the code to the container
COPY . .

# Expone el puerto en el que se ejecutará nuestra aplicación
EXPOSE 8080

# Inicia nuestra aplicación
CMD ["npm", "start"]
