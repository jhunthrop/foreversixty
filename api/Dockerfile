# api/Dockerfile
# The build context is the REPOSITORY ROOT, not api/:
#   docker build --build-arg VERSION=$(git rev-parse --short HEAD) -f api/Dockerfile -t foreversixty-api:local .
# The API serves the planner's talent data out of the image, and it
# builds against the engine in logs/ through the replace in api/go.mod,
# so the context has to include data/builds, logs/ and sim/.
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
COPY sim/go.mod ./sim/
WORKDIR /src/api
RUN go mod download
WORKDIR /src
COPY logs/ ./logs/
COPY sim/ ./sim/
COPY api/ ./api/
WORKDIR /src/api
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.version=${VERSION}" -o /out/api ./cmd/api

# The native simulator binary, built from the sim module against the
# engine sha in sim/enginever/version.go. It cannot be built from the
# development replace in sim/go.mod - that points at a local clone,
# which is not in this context - so the replace is rewritten to the
# published fork at the pinned sha, exactly as the sim workflow does.
FROM golang:1.25-alpine AS engine
RUN apk add --no-cache git
WORKDIR /src
ENV GOTOOLCHAIN=auto
ENV GOWORK=off
COPY logs/ ./logs/
COPY sim/ ./sim/
# sim/internal/simdb embeds simdb.bin at compile time. The copy is
# git-ignored, so the build reproduces `make simdb`: the active build
# named by web/src/data/active-build.json names the source under
# data/builds/, which is copied here rather than typed in.
COPY data/builds/ ./data/builds/
COPY web/src/data/ ./web/src/data/
WORKDIR /src/sim
ARG ENGINE_MODULE=github.com/jhunthrop/wowsims-forever
RUN active=$(sed -n 's/.*"build"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' ../web/src/data/active-build.json) && \
    test -n "$active" && \
    cp "../data/builds/${active}/simdb.bin" internal/simdb/simdb.bin && \
    sha=$(grep 'const Version' enginever/version.go | cut -d'"' -f2) && \
    test -n "$sha" && \
    go mod edit -replace github.com/wowsims/classic=${ENGINE_MODULE}@${sha} && \
    go mod tidy && \
    CGO_ENABLED=0 go build -ldflags "-s -w -X main.Version=${sha}" \
      -o /out/forever-sim ./cmd/forever-sim

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/api /api
# TREE_DATA_DIR defaults to /data; one directory per client build.
COPY data/builds /data
# The native simulator binary. The jobs run /engine/forever-sim; the
# runner falls back to the checked-in fixture if it is ever absent.
COPY --from=engine /out/forever-sim /engine/forever-sim
EXPOSE 8080
# The same image is the Cloud Run job: `parse-report <report_id>` as
# the container arguments parses one whole-file upload and exits.
ENTRYPOINT ["/api"]
