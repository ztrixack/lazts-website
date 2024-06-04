############################
# STEP 1 prepare the source
############################
FROM golang:1.22-alpine AS builder

# Set the environment variables for the go command
# Let it default to the host architecture
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Copy go.mod, go.sum and dependencies if exists
COPY go.mod go.sum? vendor? ./

# Install dependencies if the vendor folder is not present
RUN if [ ! -d "vendor" ]; then go mod tidy; fi

# Import the code from the context.
COPY . .

# Build the Go application
RUN go build -o /bin/app ./cmd/app

############################
# STEP 2 the running container
############################
FROM alpine:3.16 AS runner
LABEL maintainer="Tanawat Hongthai <ztrixack.th@gmail.com>"

# Create a non-root user and group with a specific UID and GID
RUN addgroup -g 1000 appgroup && adduser -u 1000 -G appgroup -D appuser

# Copy the compiled Go application and set ownership to appuser
COPY --from=builder --chown=appuser:appgroup /bin/app /app

# Ensure all directories exist in the source context or they are properly ignored if not necessary
COPY web/ /web

# Switch to the non-root user
USER appuser

# Start the application
CMD ["/app"]
