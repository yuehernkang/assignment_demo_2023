# Base image with Node.js 14
FROM node:18-alpine AS base

# Set working directory
WORKDIR /app

# Install dependencies
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile --production

# Copy project files
COPY . .

# Build the Next.js application
RUN yarn build

# Production image
FROM node:18-alpine AS production

# Set working directory
WORKDIR /app

# Install only production dependencies
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile --production

# Copy build output from the base image
COPY --from=base /app/.next ./.next
COPY --from=base /app/public ./public

# Set environment variables
ENV NODE_ENV=production

# Expose the port used by the Next.js application (default: 3000)
EXPOSE 3000

# Start the Next.js application
CMD ["yarn", "start"]
