FROM node:20-alpine

WORKDIR /app

# Copy package.json and package-lock.json for caching
COPY package*.json ./

RUN npm install

# Copy source code
COPY . .

# Build TypeScript
RUN npm run build

# Run compiled JS
CMD ["npm", "run", "start"]