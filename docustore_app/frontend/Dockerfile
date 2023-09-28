# Use the Node.js image
FROM node:alpine

# Set up the working directory
WORKDIR /usr/src/app

# Copy the project files
COPY ./public ./public
COPY server.js .

# Install Express
RUN npm install express

# Expose the server port
EXPOSE 3000


ENV APP_NAME "super_document_manager"
# Run the server
CMD ["node", "server.js"]
