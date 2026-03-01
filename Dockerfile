# Backend
FROM golang:1.21 AS backend
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 9090
CMD ["./main"]

# Frontend
FROM node:18 AS frontend
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
EXPOSE 9091
CMD ["npm", "start"]
