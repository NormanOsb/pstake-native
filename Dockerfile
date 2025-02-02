FROM golang:1.19-alpine AS build-env

# Set up dependencies
RUN apk add --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/persistenceOne/pstake-native

# Add source files
COPY . .

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN make install

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates jq bash curl perl

# Set working directory
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/pstaked /usr/bin/pstaked

EXPOSE 26657
