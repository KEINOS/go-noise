# =============================================================================
#  Test Container for Vaious Go Versions
# =============================================================================
# Default version
ARG VARIANT="1.17-alpine"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM golang:${VARIANT}

RUN apk add --no-cache \
    mesa-dev \
    pkgconfig \
    alpine-sdk \
    build-base

WORKDIR /workspaces

ENTRYPOINT go mod download && go test -race ./...
