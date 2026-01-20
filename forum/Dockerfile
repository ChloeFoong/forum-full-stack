FROM golang:1.25.6-bookworm

# Install build tools and libc/sqlite headers
RUN apt-get update && apt-get install -y \
    build-essential \
    libsqlite3-dev \
    pkg-config \
    git \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy Go modules and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy remaining source
COPY . .

# Build binary with CGO enabled (needed by sqlite3)
RUN CGO_ENABLED=1 go build -o main .

# Expose port your app uses
EXPOSE 8080

# Run your app
CMD ["./main"]
