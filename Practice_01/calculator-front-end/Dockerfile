# Build
# docker build -t frontend_practica1_202001534:1.00 .
# Tag
# docker tag frontend_practica1_202001534:1.00 erwin14k/frontend_practica1_202001534:1.00
# Push
# docker push erwin14k/frontend_practica1_202001534:1.00

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
