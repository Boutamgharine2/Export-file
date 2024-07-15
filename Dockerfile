# Utilisation de l'image officielle de Go comme image de base
FROM golang:1.23rc1-alpine3.19

# Définition du répertoire de travail à l'intérieur de l'image Docker
WORKDIR /app

# Copie du code source dans le répertoire de travail de l'image Docker
COPY . .

# Commande pour compiler l'application Go
RUN apk --no-cache add bash

RUN go build -o main .

LABEL ProjectName="ASCII-ART-DOCKERIZE"

LABEL Version="1.0"

# Commande pour exécuter l'application une fois que le conteneur Docker est lancé
CMD ["./main"]
