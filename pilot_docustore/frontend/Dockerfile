# Step 1: Build the Nuxt application
FROM node:18.18.0-alpine3.18 AS build

# Set the working directory
WORKDIR /usr/src/app

# Copy package.json and pnpm-lock.yaml
COPY package.json pnpm-lock.yaml ./

# Install pnpm
RUN npm install -g pnpm

# Install dependencies
RUN pnpm i

# Copy the local files to the container
COPY . .

# Build the app
RUN pnpm build

FROM node:18

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/.output/ .output/

# Optional, only needed if you rely on unbundled dependencies
# COPY --from=build /src/node_modules /src/node_modules

CMD [ "node", ".output/server/index.mjs" ]