FROM public.ecr.aws/docker/library/golang:1.24-alpine AS builder

RUN apk add --update --no-cache ca-certificates git

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image
# and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o apiserver ./cmd/api

FROM public.ecr.aws/docker/library/alpine:3.21

# COPY the binary from the builder stage.
COPY --from=builder ["/build/apiserver", "/"]

# Export necessary port.
EXPOSE 8080

# Command to run when starting the container.
CMD [ "/apiserver" ]
