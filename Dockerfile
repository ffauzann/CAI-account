
FROM golang:1.21.1 as builder 

LABEL maintainer="ffauzann"

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Fetch dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix -cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk update && apk add bash && apk --no-cache add tzdata
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy mandatory files from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/internal/app/user.config.yaml internal/app/user.config.yaml
COPY --from=builder /app/internal/app/auth.config.yaml internal/app/auth.config.yaml
COPY --from=builder /app/internal/migration internal/migration

# Command to run the executable
CMD [ "./main" ]