# Dockerfile References: https://docs.docker.com/engine/reference/builder/
#
# Build and bundle the Vue.js frontend SPA 
#
# FROM node:14-alpine AS vue-build
# WORKDIR /build

# COPY client/front-end/package*.json ./
# RUN npm install

# COPY client/front-end/ .
# RUN npm run build

#
# Build and run the Go backend 
#
FROM golang:alpine
ADD . /go/src/app
WORKDIR /go/src/app

RUN go build

CMD ["go", "run", "main.go"]

#
# Assemble the server binary and Vue bundle into a single app
#
# FROM scratch
# WORKDIR /app 
# COPY --from=vue-build /build/dist .

# EXPOSE 5050