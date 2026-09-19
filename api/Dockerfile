# api/Dockerfile
# The build context is the REPOSITORY ROOT, not api/:
#   docker build --build-arg VERSION=$(git rev-parse --short HEAD) -f api/Dockerfile -t foreversixty-api:local .
# The API serves the planner's talent data out of the image, and it
# builds against the engine in logs/ through the replace in api/go.mod,
# so the context has to include data/builds, logs/ and sim/.
#
# ENGINE_SIMDB_SOURCE picks which of the engine-source-* stages below
# supplies the item database the engine stage embeds. A FROM line can
# only read an ARG declared before the file's first FROM, so it has to
# live here rather than near the stages it selects between. Its
# default keeps a plain local build (no extra arguments) copying the
# whole data/builds tree, which is always correct; the deploy workflow
# overrides it (with ACTIVE_BUILD, below) to copy exactly one build's
# simdb.bin instead of every build's ~83MB. See the two engine-source-*
# stages for why the choice has to be made this way rather than with a
# single COPY.
ARG ENGINE_SIMDB_SOURCE=engine-source-all

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
#
# git is here because rewriting the replace resolves the fork's
# pseudo-version from GitHub. The three stages below split the item
# database it embeds from the rest of the engine build: sim/internal/
# simdb embeds simdb.bin at compile time (the copy is git-ignored, so
# this reproduces `make simdb`), and Docker's COPY can't take a path
# computed at RUN time, so there is no single COPY that picks out only
# the active build's file. engine-source-all is the argument-free
# fallback (copy every build, let the RUN step below pick one out, the
# same as `make simdb` does on a developer's machine);
# engine-source-active is what the deploy workflow selects once it
# already knows the active build, copying only that file. Either way
# the RUN step's own active-build.json lookup is what actually decides
# which build gets embedded - never a typed id, and never trusting the
# two sources to agree without checking.
FROM golang:1.25-alpine AS engine-base
RUN apk add --no-cache git
WORKDIR /src
ENV GOTOOLCHAIN=auto
ENV GOWORK=off
COPY logs/ ./logs/
COPY sim/ ./sim/
COPY web/src/data/ ./web/src/data/

FROM engine-base AS engine-source-all
COPY data/builds/ ./data/builds/

# ACTIVE_BUILD defaults to a glob, not to a typed build id: the classic
# (non-BuildKit) builder evaluates every stage regardless of which one
# ENGINE_SIMDB_SOURCE ultimately selects, so this stage has to build
# successfully even when it is the one being discarded. The glob still
# resolves to exactly today's one build with a simdb.bin; the deploy
# workflow always overrides ACTIVE_BUILD with the exact id anyway, so
# production never depends on the glob matching only one file.
FROM engine-base AS engine-source-active
ARG ACTIVE_BUILD=*
COPY data/builds/${ACTIVE_BUILD}/simdb.bin ./data/builds/${ACTIVE_BUILD}/simdb.bin

FROM ${ENGINE_SIMDB_SOURCE} AS engine
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
