# api/Dockerfile
# The build context is the REPOSITORY ROOT, not api/:
#   docker build --build-arg VERSION=$(git rev-parse --short HEAD) -f api/Dockerfile -t foreversixty-api:local .
# The API serves the planner's talent data out of the image, so the context
# has to include data/builds.
FROM golang:1.25-alpine AS build
WORKDIR /src
ENV GOTOOLCHAIN=auto
COPY api/go.mod api/go.sum ./
RUN go mod download
COPY api/ ./
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.version=${VERSION}" -o /out/api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/api /api
# TREE_DATA_DIR defaults to /data; one directory per client build.
COPY data/builds /data
EXPOSE 8080
ENTRYPOINT ["/api"]
