FROM --platform=$BUILDPLATFORM golang:1.22.2 AS build

WORKDIR /workspace

# clone repo via git
RUN git clone https://github.com/go-mc/server.git .
RUN go mod download

# Automatically provided by the buildkit
ARG TARGETOS
ARG TARGETARCH

# Build
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -ldflags="-s -w" -a -o app main.go

# Move binary into final image
FROM --platform=$BUILDPLATFORM gcr.io/distroless/static-debian11 AS app
COPY --from=build /workspace/app /
COPY ./config.toml /config.toml
COPY ./world /world
CMD ["/app"]