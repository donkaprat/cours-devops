# Étape de construction
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o cours-devops

EXPOSE 8080

CMD ["./cours-devops"]
