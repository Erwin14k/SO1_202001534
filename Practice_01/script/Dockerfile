# Build
# docker build -t script_practica1_202001534:1.00 .
# Tag
# docker tag script_practica1_202001534:1.00 erwin14k/script_practica1_202001534:1.00
# Push
# docker push erwin14k/script_practica1_202001534:1.00

FROM alpine:latest

# Add bash
RUN apk add --no-cache bash
# Install tzdata and set the timezone to America/Guatemala
RUN apk add --no-cache tzdata
ENV TZ=America/Guatemala
# Copy the script to the workdir
COPY reportsScript.sh /app/Backend/
# Set permissions to the script
RUN chmod +x /app/Backend/reportsScript.sh
# Execute the script
CMD ["/bin/bash", "-c", "/app/Backend/reportsScript.sh"]
