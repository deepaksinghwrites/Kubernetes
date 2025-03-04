# Use official Node.js LTS image
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy only package.json (remove package-lock.json)
COPY package.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application
COPY . .

# Set environment variable for API URL
ENV REACT_APP_BACKEND_URL=http://backend-service:8080/api/tasks

# Build the React app
RUN npm run build

# Install a lightweight static server
RUN npm install -g serve

# Expose port 3000
EXPOSE 3000

# Start the app
CMD ["serve", "-s", "build", "-l", "3000"]
