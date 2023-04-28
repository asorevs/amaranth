# Base image
FROM mongo:latest

# Set environment variables
ENV MONGODB_USERNAME=""
ENV MONGODB_PASSWORD=""
ENV MONGODB_HOST="localhost"
ENV MONGODB_PORT="27017"
ENV MONGODB_DATABASE=""

# Copy the initialization script to the container
COPY init-mongodb.sh /docker-entrypoint-initdb.d/

# Set execute permissions for the initialization script
RUN chmod +x /docker-entrypoint-initdb.d/init-mongodb.sh

