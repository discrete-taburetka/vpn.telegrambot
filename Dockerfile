# syntax=docker/dockerfile:1

# BUILD FRONTEND
FROM node:21.7.0-alpine3.18 AS build-frontend
WORKDIR /frontend
COPY frontend/package.json ./
COPY frontend/tsconfig*.json ./
COPY frontend/src src
COPY frontend/public public
RUN npm install
RUN npm run build

#BUILD SERVER AND COPY STATIC
FROM golang:1.22
WORKDIR /backend
COPY backend/go.mod ./
COPY backend/go.sum ./
RUN go mod download
COPY backend/*.go ./
COPY --from=build-frontend /frontend/build/ static/
RUN CGO_ENABLED=0 GOOS=linux go build -o /bot-vpn-server
EXPOSE 8080

# Run
CMD ["/bot-vpn-server"]