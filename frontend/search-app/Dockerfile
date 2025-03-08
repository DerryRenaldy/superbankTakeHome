# Use the official Node.js image as a build stage
FROM node:18 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json files
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the source code into the container
COPY . .

# Build the Next.js app
RUN npm run build

# Start a new stage from scratch
FROM node:18 AS production

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/package.json /app/package-lock.json ./
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public

# Install only production dependencies
RUN npm install --only=production

EXPOSE 3000

# Command to run the Next.js app
CMD ["npm", "start"]