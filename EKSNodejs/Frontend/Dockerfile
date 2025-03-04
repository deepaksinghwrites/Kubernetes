FROM node:18-alpine

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json first (to leverage Docker caching)
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the entire application code
COPY . .

# Set environment variable for React app
ENV REACT_APP_BACKEND_URL=http://backend-service:8080/api/tasks

# Build the React app
RUN npm run build

# Install a lightweight web server for serving static files
RUN npm install -g serve

# Expose the port used by the React app
EXPOSE 3000

# Start the React app in production mode
CMD ["serve", "-s", "build", "-l", "3000"]
