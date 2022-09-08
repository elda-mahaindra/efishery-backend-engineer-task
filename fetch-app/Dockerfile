# TypeScript Compiler stage
FROM node:16-alpine3.15 as ts-compiler
WORKDIR /app
COPY package*.json ./
COPY tsconfig*.json ./
RUN npm install
COPY . ./
RUN npm run build

# TypeScript Remover stage
FROM node:14-alpine3.10 as ts-remover
WORKDIR /app
COPY --from=ts-compiler /app/package*.json ./
COPY --from=ts-compiler /app/dist ./
RUN npm install --only=production

# Distroless Production stage
FROM gcr.io/distroless/nodejs:16
WORKDIR /app
COPY --from=ts-remover /app ./
COPY .env ./

EXPOSE 4000
CMD ["index.js"]