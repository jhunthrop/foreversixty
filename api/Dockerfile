# api/Dockerfile
# The build context is the REPOSITORY ROOT, not api/:
#   docker build --build-arg VERSION=$(git rev-parse --short HEAD) -f api/Dockerfile -t foreversixty-api:local .
# The API serves the planner's talent data out of the image, and it
# builds against the engine in logs/ through the replace in api/go.mod,
# so the context has to include data/builds and logs/.
FROM golang:1.25-alpine AS build
WORKDIR /src
ENV GOTOOLCHAIN=auto
# GOWORK=off on purpose: the repository's go.work lists companion/,
# which is not in this image's context. The api module's own
# `replace ../logs` resolves the engine, so the build needs no
# workspace.
ENV GOWORK=off
COPY api/go.mod api/go.sum ./api/
COPY logs/go.mod logs/go.sum ./logs/
WORKDIR /src/api
RUN go mod download
WORKDIR /src
COPY logs/ ./logs/
COPY api/ ./api/
WORKDIR /src/api
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.version=${VERSION}" -o /out/api ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/api /api
# TREE_DATA_DIR defaults to /data; one directory per client build.
COPY data/builds /data
EXPOSE 8080
# The same image is the Cloud Run job: `parse-report <report_id>` as
# the container arguments parses one whole-file upload and exits.
ENTRYPOINT ["/api"]
