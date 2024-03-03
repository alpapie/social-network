# Utilisation d'une image de base pour Go
FROM golang:latest AS backend

# Répertoire de travail pour le backend
WORKDIR /go/src/app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/*.go ./
RUN go build -o backend-app .

# Utilisation d'une image de base pour Node.js
FROM node:latest AS frontend

# Répertoire de travail pour le frontend
WORKDIR /usr/src/app
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ .
RUN npm run build

# Utilisation d'une image de base légère pour l'exécution
FROM alpine:latest

# Installation des dépendances nécessaires pour l'exécution des deux applications
RUN apk --no-cache add ca-certificates

# Copie des fichiers binaires des applications
COPY --from=backend /go/src/app/backend-app /usr/local/bin/
COPY --from=frontend /usr/src/app/build /app

# Commande par défaut pour démarrer les deux applications
CMD ["backend-app"]
