FROM mcr.microsoft.com/devcontainers/base:ubuntu-24.04

# Install golang
# https://go.dev/dl/
ARG GOLANG_VERSION="1.23.1"
ARG GOLANG_CHECKSUM="49bbb517cfa9eee677e1e7897f7cf9cfdbcf49e05f61984a2789136de359f9bd"
RUN wget --https-only "https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz" -O /tmp/go.tar.gz && \
    sha256sum /tmp/go.tar.gz | grep "${GOLANG_CHECKSUM}" && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz
ENV PATH="${PATH}:/usr/local/go/bin"
ENV GOBIN="/usr/local/bin"

# golang IDE tooling, install latest versions
RUN go install -v golang.org/x/tools/gopls@latest && \
    go install -v honnef.co/go/tools/cmd/staticcheck@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest

# Project tooling, use specific versions
RUN go install -v github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0 && \
    go install -v github.com/pressly/goose/v3/cmd/goose@v3.22.0

# Install node
# https://nodejs.org/en/download/prebuilt-binaries
ARG NODE_VERSION="20.17.0"
ARG NODE_CHECKSUM="a24db3dcd151a52e75965dba04cf1b3cd579ff30d6e0af9da1aede4d0f17486b"
RUN wget --https-only "https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-x64.tar.xz" -O /tmp/node.tar.xz && \
    sha256sum /tmp/node.tar.xz | grep "${NODE_CHECKSUM}" && \
    tar -C /usr/local -xf /tmp/node.tar.xz --strip-components=1 --exclude="CHANGELOG.md" --exclude="LICENSE" --exclude="README.md" && \
    rm /tmp/node.tar.xz && \
    corepack enable && \
    corepack install --global pnpm@9.8.0
ENV COREPACK_ENABLE_DOWNLOAD_PROMPT=0

# Install just
# https://github.com/casey/just/releases
ARG JUST_VERSION="1.34.0"
ARG JUST_CHECKSUM="42df46f7ce727bdb2c3b5357017d885bd2033e8dbb154ca8e8f229e75093aeca"
RUN wget --https-only "https://github.com/casey/just/releases/download/${JUST_VERSION}/just-${JUST_VERSION}-x86_64-unknown-linux-musl.tar.gz" -O /tmp/just.tar.gz && \
    sha256sum /tmp/just.tar.gz | grep "${JUST_CHECKSUM}" && \
    tar -C /usr/local/bin -xzf /tmp/just.tar.gz && \
    rm /tmp/just.tar.gz