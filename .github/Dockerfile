# =============================================================================
#  Test Container for Various Go Versions
# =============================================================================
# Default version
ARG VARIANT="1.22-alpine"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM golang:${VARIANT}

ENV GOTOOLCHAIN=local

RUN apk add --no-cache \
    mesa-dev \
    pkgconfig \
    alpine-sdk \
    build-base

WORKDIR /workspaces

ENTRYPOINT ["sh", "-ec", "go test -mod=readonly -race ./... ./_example/3d && cd _example/2d && go test -mod=readonly -race ."]
