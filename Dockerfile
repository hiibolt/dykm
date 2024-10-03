# syntax=docker/dockerfile:1

FROM golang:1.23.2-bookworm

# Install Dependencies
RUN apt-get update
RUN apt-get install -y pkg-config curl
RUN apt-get install -y libssl-dev openssl

# Change Directory To Our App
WORKDIR /app

# Copy All Files
COPY . .

# Run The Go App
WORKDIR /app/backend/src/dykm
CMD ["go", "run", "."]

# Expose The Application
ARG PORT
EXPOSE $PORT