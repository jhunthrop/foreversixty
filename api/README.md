# api

Go service for api.foreversixty.gg, deployed as a container on Google Cloud Run.

## Commands

```
make test   # go test -cover ./...
make build  # go build -o bin/api ./cmd/api
make run    # build then run (needs DATABASE_URL, RESEND_API_KEY, PUBLIC_BASE_URL, API_BASE_URL)
make lint   # go vet ./...
```

Run against the checked-in tree data (rather than the image's `/data`):

```bash
TREE_DATA_DIR=../data/builds \
DATABASE_URL='postgres://forever:forever@localhost:5434/forever_test?sslmode=disable' \
RESEND_API_KEY=x PUBLIC_BASE_URL=https://foreversixty.gg API_BASE_URL=https://api.foreversixty.gg \
make run
```

## Environment variables

| Variable | Required | Default | Notes |
|---|---|---|---|
| `PORT` | no | `8080` | |
| `DATABASE_URL` | yes | — | Runtime connection; may point at a pooler (e.g. Neon's PgBouncer endpoint). |
| `MIGRATE_DATABASE_URL` | no | `DATABASE_URL` | Connection used only for running migrations. See "Migrations and connection poolers" below. |
| `RESEND_API_KEY` | yes | — | |
| `PUBLIC_BASE_URL` | yes | — | e.g. `https://foreversixty.gg`; confirm/unsubscribe redirects land here. |
| `API_BASE_URL` | yes | — | e.g. `https://api.foreversixty.gg`; used to build the confirmation link. |
| `MAIL_FROM` | no | `Forever Sixty <hello@foreversixty.gg>` | `From:` address for outgoing mail. |
| `TRUSTED_PROXY_HOPS` | no | `1` | How many reverse proxies in front of this service (e.g. Cloud Run) are trusted to append to `X-Forwarded-For`; used to resolve the real client IP for rate limiting. `0` ignores `X-Forwarded-For` entirely and rate-limits by the raw connection address. Must be an integer >= 0. |
| `R2_ACCOUNT_ID` | no | — | Cloudflare account id; the S3 endpoint is derived from it. |
| `R2_ACCESS_KEY_ID` | no | — | R2 API token id (Object Read & Write). |
| `R2_SECRET_ACCESS_KEY` | no | — | R2 API token secret. |
| `R2_BUCKET` | no | `foreversixty-logs` | The bucket every report file lives in. |
| `R2_ENDPOINT` | no | derived | S3 endpoint override; the tests point it at an in-process fake. |
| `SESSION_COOKIE_DOMAIN` | no | `.foreversixty.gg` | `Domain` on `fs_session` and `fs_csrf`. Set it to `none` for a local `http://localhost` run, which makes the cookies host-only. |
| `BNET_CLIENT_ID` | no | — | Battle.net OAuth client. Without the id, the secret, and the redirect URL, only the email magic link is offered. |
| `BNET_CLIENT_SECRET` | no | — | |
| `BNET_REDIRECT_URL` | no | — | Registered for production and for `http://localhost:8080/v1/auth/battlenet/callback`. |
| `BNET_PROFILE_GAME` | no | `classic1x` | The profile-game segment of a Blizzard namespace (`profile-<game>-<region>`, `dynamic-<game>-<region>`). Forever's real namespace is one env change on launch day — see "The Battle.net character import" below. |
| `BNET_REGIONS` | no | `us,eu` | Comma-separated; which regional Blizzard hosts to try for an account during import and refresh. |
| `BNET_PROBE_GAMES` | no | `classic1x,classic,classic-forever,classicforever,forever,classic60,anniversary` | Comma-separated; every game namespace segment the nightly refresh's namespace probe checks, logging one line per game (any game other than `BNET_PROFILE_GAME` that answers 200 is logged at WARN as `namespace_appeared`). |
| `PARSE_JOB_NAME` | no | `parse-report` | The Cloud Run job that parses a whole-file upload. |
| `PARSE_JOB_REGION` | no | `us-east1` | |
| `PARSE_JOB_PROJECT` | no | `foreversixty` | The Google Cloud project the job lives in. |
| `SIM_JOB_NAME` | no | `sim-run` | The Cloud Run job that runs one premium sim. |
| `SIM_JOB_REGION` | no | `us-east1` | |
| `SIM_JOB_PROJECT` | no | `foreversixty` | The Google Cloud project the sim jobs live in. |
| `TREE_DATA_DIR` | no | `/data` | Directory holding one subdirectory per client build (`<build>/talents/*.json`, `<build>/items/*.json`, `<build>/{sets,classes,races,combos}.json`). The Docker image copies `data/builds` here. A missing or pre-Phase-1 directory is logged at startup and simply has no data: the service still serves health, version, and subscribe, and every save fails validation on `tree_version`. |

Every Phase 3 variable is optional and degrades honestly: with no R2 credentials the service
serves everything that does not touch object storage and logs that the ingest, upload, and
report-file routes are not mounted; with no Battle.net client it offers the email magic link
alone; with no Cloud Run credentials it serves the live ingest but not whole-file uploads.

### Migrations and connection poolers

`db.Migrate` (via golang-migrate) takes a session-level Postgres advisory lock while migrations
run, so it needs a session-pooled or direct (non-pooled) connection. A transaction-pooled
connection (e.g. Neon's default pooled/PgBouncer endpoint) cannot hold that lock reliably.

Point `DATABASE_URL` at Neon's pooled endpoint for normal runtime traffic, and point
`MIGRATE_DATABASE_URL` at Neon's direct (unpooled) endpoint so migrations run against a
connection that can actually hold the advisory lock. When `MIGRATE_DATABASE_URL` is unset,
`DATABASE_URL` is used for migrations too.

Tests need a Postgres instance:

```bash
docker compose -f docker-compose.test.yml up -d   # Postgres for tests on host port 5434
export TEST_DATABASE_URL='postgres://forever:forever@localhost:5434/forever_test?sslmode=disable'
make test
```

The database-backed packages share that one Postgres and each truncates the tables it uses, so
the whole suite runs with `-p 1`:

```bash
go test -race -cover -p 1 ./...
```

## Build endpoints

| Route | Notes |
|---|---|
| `POST /v1/builds` | Saves a build. 201 with `{id, url}`, or 200 with the same body when that id already exists (ids are content hashes). Body max 8 KB, 20 saves per IP per hour. 400 carries `error.fields` keyed `class_id`, `race_id`, `tree_version`, `title`, `point_order`, `point_order[i]`, `gear.<slot>`. |
| `GET /v1/builds/{id}` | The record, cached a day. 404 in the envelope. |
| `GET /v1/builds?mine=1` | The signed-in player's own saved builds, newest first. A build saved anonymously has no owner; a later signed-in save of the same build (ids are content hashes) claims it if nobody owns it yet. |
| `GET /b/{id}` | The server-rendered build page, with the site's chrome, Open Graph tags, and the record inlined for the planner island. A missing or unloadable build gets an HTML message page (404 or 500) linking to the planner instead. |
| `GET /b/{id}/card.png` | 1200×630 preview PNG, cached a week. A build whose own card cannot be drawn falls back to the static card, cached five minutes. |

Both `GET /b/...` routes are exempt from the router-wide per-IP rate limit: they reach this service
through the site's Cloudflare Worker, whose egress IP would otherwise collapse every visitor into one
bucket, and they are read-only and cached.

Builds are validated against the tree data for their `tree_version` before they are stored; the rules
are in `internal/builds/validate.go` and come from the Phase 1 interface contract, which has the web
planner mirror them in `web/src/lib/planner/rules.ts`.

## Simulator and phase endpoints

The simulator's full surface (`POST /v1/sims`, `POST /v1/sims/run`, `GET /v1/sims/{id}`,
`GET /v1/sims/{id}/progress`, `GET /v1/specs`, and the sim-input route) is documented in
`openapi.yaml`, not restated here. The two rows below are the ones the router mounts
unconditionally that do not belong to any other table in this file:

| Route | Notes |
|---|---|
| `GET /v1/sims?mine=1&kind=` | The caller's own sims, newest first, optionally narrowed to one tool (`run`, `gear`, `talents`, `drops`, `weights`). Each row carries its kind and a composed one-line headline. An unknown kind is 400. |
| `GET /v1/phases` | The content phase boundaries, cached an hour. The table is compiled in; `api/internal/phase/boundaries_test.go` holds it to `data/curated/phases.json`. |

## Keeping the page chrome in step with the site

`GET /b/{id}` renders the site's header and footer from `internal/site/chrome.html`, a copy of the
markup the Astro site emits. `internal/site/testdata/chrome.html` is the recorded site output.
`TestChromeTemplateMatchesTheSnapshot` always runs and compares the two;
`TestChromeSnapshotMatchesTheBuiltSite` compares the recording to `web/dist/about.html` and skips when
the site has not been built, so the api suite never depends on the web build. After changing
`web/src/components/{Header,Footer}.astro`, refresh both files:

```bash
(cd ../web && npm run build) && go test ./internal/site -run TestChromeSnapshot -update
```

## The logs product

The API is the ingest, the index, and the rankings for combat logs. The engine itself is the
`logs/` module, which this module requires through a `replace ../logs`; the repository root's
`go.work` lists both. `internal/engine` is the one place the engine is configured, so the
companion, the whole-file job, and the ingest's verification all parse the same way.

| Route | Notes |
|---|---|
| `POST /v1/reports` | Starts a live report. Device or session. |
| `PUT /v1/reports/{id}/fights/{n}` | One closed fight as multipart: `summary`, `events` (Parquet), `metrics`, `raw_range`. The API rebuilds the metrics from the events and answers 409 with `error.fields.metrics` when they disagree. 200 when the same fight with the same raw hash is already stored. |
| `PUT /v1/reports/{id}/fights/{n}/live` | The running summary while a fight is open. |
| `PUT /v1/reports/{id}/raw?offset=N` | One zstd chunk of the original log, at most 8 MiB, hashed by `X-Raw-SHA256` over the decoded bytes. Offsets are report-relative. |
| `POST /v1/reports/{id}/complete` | Ends the report and schedules the raw-sample check. |
| `POST /v1/uploads`, `POST /v1/uploads/{id}/complete` | Signed R2 multipart URLs for a whole-file upload, then the parse job. |
| `GET /v1/reports/{id}` | The report, its fights, and `data_base_url`. |
| `GET /v1/reports/{id}/visibility` | What the site Worker asks before serving files off the bucket. Cached 60 s. |
| `GET /v1/reports/{id}/access`, `GET /v1/reports/{id}/files/{path}` | A private or guild report's files, as signed redirects valid ten minutes. |
| `GET /v1/rankings`, `/v1/rankings/percentile`, `/v1/rankings/guilds` | Leaderboards and percentiles, cached 30 s. |
| `GET /v1/characters/{region}/{ruleset}/{name}`, `GET /v1/guilds/...` | Character and guild pages. |
| `GET /reports/{id}/card.png` | The unfurl card, cached five minutes — `private` and `Vary: Cookie` when the report is not public or unlisted, so a shared cache never serves one reader's card to another. |

Sign-in is Battle.net first and an email magic link as the fallback; the companion pairs with a
code from `POST /v1/devices/pair` and uploads with `Authorization: Bearer fsd_…`. Browser
sessions are opaque cookies with double-submit CSRF (`fs_csrf` plus `X-CSRF-Token`); device
tokens are exempt, because a companion is not a browser.

### What the `anonymize` flag covers

`PATCH /v1/me {"anonymize": true}` replaces your name on report pages: `owner.battletag` in
`GET /v1/reports/{id}` reads as `user-<id>` instead of your battletag, for every reader of
every report you own. That is its whole scope today.

It does **not** touch ranking rows. A row in `/v1/rankings`, on a character page or on a guild
page still carries the character name you logged under — as `player.name` and inside
`player.key`, which is `<region>/<ruleset>/<name-slug>`. The key is the character's identity:
the region and ruleset filters, the character page's own URL, and the links between a
leaderboard and a character page are all keyed on it, so substituting a pseudonym for the name
while publishing the key beside it would look like privacy without being any.

Separately, and regardless of the flag: no route shows an account's email address to anyone
but that account, on `GET /v1/me`. An
account that signed in by magic link has no battletag, and a report owner with no battletag
reads as `user-<id>` whether or not they set `anonymize`.

### The parse job

A whole-file upload is parsed by the same image, run as a Cloud Run job with different
arguments. Create it once:

```bash
gcloud run jobs create parse-report \
  --image us-east1-docker.pkg.dev/foreversixty/api/api:latest \
  --region us-east1 --cpu 2 --memory 2Gi --task-timeout 30m \
  --args parse-report \
  --set-env-vars "$(tr '\n' ',' < .env.job)"
```

The API executes it per upload through the Cloud Run Admin API, which needs `run.developer` on
the job for the API's runtime service account:

```bash
gcloud run jobs add-iam-policy-binding parse-report --region us-east1 \
  --member "serviceAccount:<api runtime service account>" --role roles/run.developer
```

Deploys update the job's image automatically (see `.github/workflows/api.yml`); the job needs
the same secrets and variables as the service, minus `PORT`.

### The simulator's Cloud Run jobs

Both run the same image as the API, dispatched on their first argument.

    gcloud run jobs create sim-run \
      --image <the API image> --region us-east1 --args sim-run \
      --cpu 4 --memory 4Gi --task-timeout 15m

    gcloud run jobs create sim-validate \
      --image <the API image> --region us-east1 --args sim-validate \
      --cpu 4 --memory 4Gi --task-timeout 30m

The entitlements/billing lane adds two more, both cheap and short-lived — `stripe-setup` is
run by hand once per mode (see "First-time setup" above); `stripe-reconcile` is scheduled
nightly the same way `sim-validate` already is:

    gcloud run jobs create stripe-reconcile \
      --image <the API image> --region us-east1 --args stripe-reconcile \
      --cpu 1 --memory 512Mi --task-timeout 5m

    gcloud scheduler jobs create http stripe-reconcile-nightly \
      --schedule "0 4 * * *" --uri "https://us-east1-run.googleapis.com/apis/run.googleapis.com/v1/namespaces/foreversixty/jobs/stripe-reconcile:run" \
      --http-method POST --oauth-service-account-email api-runtime@foreversixty.iam.gserviceaccount.com

`stripe-reconcile` heals any webhook Stripe's own three-day retry window never successfully
delivered (design spec §2.8): it is a backstop, not the primary path — the webhook itself
(`POST /v1/billing/webhook`) does the real-time work.

The performance ratings lane adds `rating-backfill`: it rates every fight in every
complete report that has no rating yet, in bounded batches, and is safe to re-run and to
run beside live ingest (new uploads are rated as they land). Run it once after the
ratings deploy, and again after any engine or curated-data change that bumps the model
version:

    gcloud run jobs create rating-backfill \
      --image <the API image> --region us-east1 --args rating-backfill \
      --cpu 2 --memory 2Gi --task-timeout 60m \
      --set-env-vars "$(tr '\n' ',' < .env.job)"

    gcloud run jobs execute rating-backfill --region us-east1 --wait

### The data-addon job

`data-addon` (`api/internal/dataaddon`) aggregates every public rated character's last 90
days and every verified guild's roster into the Forever Sixty Data addon's `Data.lua`, and
publishes it to a Cloud Storage bucket for `.github/workflows/addon-data-release.yml` to
package nightly. Unlike every other job in this file, its bucket address is read directly
via `os.Getenv("DATA_ADDON_BUCKET")` in `main.go` rather than through `config.Config` — see
`docs/superpowers/plans/2026-09-21-data-addon.md`'s Task 9 for why.

One-time setup (in addition to "1. Google Cloud project setup" above):

    gcloud services enable storage.googleapis.com
    gcloud storage buckets create gs://foreversixty-addon-data --location us-east1 --uniform-bucket-level-access

    # api-runtime (the service's own service account) needs to write the nightly file:
    gcloud storage buckets add-iam-policy-binding gs://foreversixty-addon-data \
      --member serviceAccount:api-runtime@foreversixty.iam.gserviceaccount.com \
      --role roles/storage.objectAdmin

    # the deployer service account (see "Then set up Workload Identity Federation..." above)
    # needs read access, since addon-data-release.yml downloads through the same WIF identity
    # api.yml's deploy job already uses:
    gcloud storage buckets add-iam-policy-binding gs://foreversixty-addon-data \
      --member serviceAccount:<deployer service account email> \
      --role roles/storage.objectViewer

    gcloud run jobs create data-addon \
      --image us-east1-docker.pkg.dev/foreversixty/api/api:latest \
      --region us-east1 --args data-addon \
      --cpu 1 --memory 512Mi --task-timeout 10m \
      --set-env-vars "$(tr '\n' ',' < .env.job),DATA_ADDON_BUCKET=foreversixty-addon-data" \
      --service-account api-runtime@foreversixty.iam.gserviceaccount.com

    gcloud scheduler jobs create http data-addon-nightly \
      --schedule "0 4 * * *" \
      --uri "https://us-east1-run.googleapis.com/apis/run.googleapis.com/v1/namespaces/foreversixty/jobs/data-addon:run" \
      --http-method POST --oauth-service-account-email api-runtime@foreversixty.iam.gserviceaccount.com

Then, as a GitHub repository variable (Settings → Secrets and variables → Actions →
Variables — the same place `GCP_WIF_PROVIDER` and `GCP_DEPLOYER_SA` already live):

    DATA_ADDON_BUCKET = foreversixty-addon-data

`addon-data-release.yml` is gated on this variable existing (`if: vars.DATA_ADDON_BUCKET != ''`),
the same way `api.yml`'s own `deploy` job is gated on `GCP_WIF_PROVIDER`.

Like `parse-report`, `sim-run` and `sim-validate`, `data-addon` is in `.github/workflows/api.yml`'s
"Point the jobs at the new image" list, so every deploy repoints it at the new image.

### The bnet-refresh job

`bnet-refresh` (`api/internal/bnetimport`) re-syncs every `characters` row with
`source = 'bnet'` and `refreshed_at` older than 20 hours against Blizzard's public profile and
guild-roster data (the app's own client-credentials token — no user's OAuth token is needed,
or kept, for this job), at most 4 Blizzard calls per second, stopping the run on the first
`429`. It then runs the namespace probe (`BNET_PROBE_GAMES`), logging one line per game and a
WARN for any game other than `BNET_PROFILE_GAME` that answers 200 — the owner's signal to flip
`BNET_PROFILE_GAME` on launch day. It needs the same Battle.net credentials the service itself
does (`BNET_CLIENT_ID`, `BNET_CLIENT_SECRET`) and is a no-op with none configured.

    gcloud run jobs create bnet-refresh \
      --image us-east1-docker.pkg.dev/foreversixty/api/api:latest \
      --region us-east1 --args bnet-refresh \
      --cpu 1 --memory 512Mi --task-timeout 15m \
      --set-env-vars "$(tr '\n' ',' < .env.job)" \
      --service-account api-runtime@foreversixty.iam.gserviceaccount.com

    gcloud scheduler jobs create http bnet-refresh-nightly \
      --schedule "30 3 * * *" \
      --uri "https://us-east1-run.googleapis.com/apis/run.googleapis.com/v1/namespaces/foreversixty/jobs/bnet-refresh:run" \
      --http-method POST --oauth-service-account-email api-runtime@foreversixty.iam.gserviceaccount.com

`bnet-refresh` is in `.github/workflows/api.yml`'s "Point the jobs at the new image" list too,
so every deploy repoints it at the new image.

`sim-run` is executed by the API for one premium run and takes the sim
id as a second argument; `sim-validate` is scheduled nightly by Cloud
Scheduler and takes none. Both need `/engine/forever-sim` in the image
to do real work — see "The engine binary" below — and fall back to the
checked-in fixture result when it is absent.

The `--cpu 4 --task-timeout 15m` on `sim-run` is load-bearing, not just a
ceiling. A Top Gear, Droptimizer or talent-compare submit — any request
that carries a bulk block — is sized before it is queued: the API asks
the binary to expand the request without running it (`forever-sim
-plan`), divides the precision ladder's total iterations by the
engine's measured rate times those four CPUs, and refuses anything past
`sims.BulkBudget` (840 seconds, one minute inside the task timeout)
with `400 too_large` and the estimate. A stat-weights submit is not
sized this way: the module's planner is bulk-only today (it refuses a
request with no bulk block), and a weights request carries
`req.Weights` instead, so it is accepted and queued unchecked. Changing
the job's CPU count means changing `simJobCPUs` in
`api/internal/sims/simdep.go` in the same commit, or every estimate is
wrong. The rate itself is `measure.NativeIterationsPerCPUSecond`, which
`sim/measure` publishes from its own benchmark — never restate it here.
Jobs are created by hand, so nothing enforces this but this paragraph.

### The engine binary

The simulator jobs run `/engine/forever-sim`. The image builds it from
the `sim/` module against the sha in `sim/enginever/version.go`, in its
own stage, rewriting that module's development `replace` to the
published fork — so the engine the jobs run is the engine the pin
names, and re-pinning is a one-line change plus a rebuild.

A deployment whose image somehow lacks the binary still serves: the
jobs log that they are answering from the checked-in fixture result and
every sim comes back with the fixture's numbers. The deploy workflow
does not rely on that fallback going unnoticed, though: right after
`docker build` it runs the image's `/engine/forever-sim -version` and
fails the deploy unless the output equals the pin in
`sim/enginever/version.go` — a stale build-cache layer serving an old
`sim/` tree at the right paths, with the wrong sha baked into its
ldflags, would still make the binary run; it would not make it match.

`forever-sim` resolves `item:<id>` consumables through the build's
`data/builds/<build>/simconsumes.json`, which the image already carries
at `/data`.

The engine stage's own build-time cost is kept independent of how many
client builds `data/builds/` holds: `sim/internal/simdb` needs only the
active build's `simdb.bin`, so a plain local build (no extra arguments)
copies every build's data and lets a `RUN` step pick the active one out
— the same as `make simdb` does on a developer's machine — while the
deploy workflow passes `--build-arg ENGINE_SIMDB_SOURCE=engine-source-active
--build-arg ACTIVE_BUILD=<the active build>` (read from
`web/src/data/active-build.json`, never typed) so the image copies
exactly that one file instead. Both paths land at the same place before
the build runs, so re-pinning or switching the active build never
needs a Dockerfile change.

### Granting premium

Premium is a flag on the account, set by hand until payments are
designed. There is no code path that turns it on.

    update users set premium = true where email = '<the tester>';

### The bucket's CORS rule

A whole-file upload goes straight from the browser to R2 through the signed part URLs, and the
browser has to read each part's `ETag` back to complete the upload. Both need a CORS rule on the
bucket, applied once with the AWS CLI pointed at R2:

```bash
cat > /tmp/logs-cors.json <<'JSON'
{"CORSRules":[{
  "AllowedOrigins":["https://foreversixty.gg"],
  "AllowedMethods":["PUT","GET","HEAD"],
  "AllowedHeaders":["content-type","content-length"],
  "ExposeHeaders":["ETag"],
  "MaxAgeSeconds":3600
}]}
JSON
AWS_ACCESS_KEY_ID=$R2_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY=$R2_SECRET_ACCESS_KEY \
aws s3api put-bucket-cors --bucket foreversixty-logs \
  --endpoint-url "https://$R2_ACCOUNT_ID.r2.cloudflarestorage.com" --region auto \
  --cors-configuration file:///tmp/logs-cors.json
```

Without `ExposeHeaders: ["ETag"]` every part uploads and the completion call then fails, because
the browser cannot see the ETag it has to send back.

### Fight numbering

Fight indexes are the engine's own, and the engine numbers fights **from 1**. Every route and
column keyed on a fight index — `PUT /v1/reports/{id}/fights/{n}`, `fights.fight_index`,
`fight_metrics.fight_index` — carries that number unchanged. (The interface contract's
Identifiers section says 0-based; the engine's segmenter starts at one, and the site and the
companion both use the engine's numbers, so 1-based is what ships.)

### Running the job locally

```bash
TREE_DATA_DIR=../data/builds DATABASE_URL=… R2_ACCESS_KEY_ID=… R2_SECRET_ACCESS_KEY=… \
  R2_ACCOUNT_ID=… go run ./cmd/api parse-report <report_id>
```

## Docker

Build the image locally. The build context is the repository root, not `api/`, because the image
copies `data/builds` to `/data`:

```bash
cd "$(git rev-parse --show-toplevel)"   # the build context is the repository root
docker build --build-arg VERSION=local -f api/Dockerfile -t foreversixty-api:local .
```

Run it and hit the health check:

```bash
docker run -d --name foreversixty-api-smoke \
  -e DATABASE_URL='postgres://forever:forever@host.docker.internal:5434/forever_test?sslmode=disable' \
  -e RESEND_API_KEY=x \
  -e PUBLIC_BASE_URL=https://foreversixty.gg \
  -e API_BASE_URL=https://api.foreversixty.gg \
  -p 18080:8080 \
  foreversixty-api:local
curl -s http://localhost:18080/health
docker stop foreversixty-api-smoke && docker rm foreversixty-api-smoke
```

## Deploy

Pushes to `main` run tests, build the image, push to Artifact Registry, and deploy to Cloud Run
(see `.github/workflows/api.yml`). Secrets live in Secret Manager: `DATABASE_URL`,
`MIGRATE_DATABASE_URL`, `RESEND_API_KEY`, `R2_ACCESS_KEY_ID`, `R2_SECRET_ACCESS_KEY`,
`BNET_CLIENT_SECRET`. Other env vars (`PORT`, `PUBLIC_BASE_URL`,
`API_BASE_URL`, `MAIL_FROM`, `TRUSTED_PROXY_HOPS`) are set on the Cloud Run service.

The image is built from the repository root (`-f api/Dockerfile .`) because it copies `data/builds`
to `/data`. A change under `data/builds/**` therefore triggers the workflow too: new client data ships
as a new image, not as a separate deploy step.

Migrations run at container startup (`db.Migrate` in `main.go`), so a deploy applies them before
serving traffic; Cloud Run only routes to the new revision once `/health` passes. The path is `/health`, not `/healthz`: Google's front end answers `/healthz` on run.app hosts itself, before the request reaches the container.

### Shutdown and Cloud Run's termination grace period

Cloud Run sends `SIGTERM` and then forcibly kills the container after a fixed 10-second
termination grace period (not configurable, and unrelated to the `--timeout 30` request timeout
above, which bounds an individual request instead). `main.go`'s shutdown budget is sized to fit
inside that: `srv.Shutdown` gets a 10-second deadline and an in-flight confirmation send is
capped at `mailSendTimeout` (8 seconds, see `internal/subscribe/service.go`), so a send that was
already running when `SIGTERM` arrives has a realistic chance to finish before the process is
killed. Running `gcloud run services update api --no-cpu-throttling` is not needed for this — it
only affects CPU allocation between requests, not the termination grace period.

## First-time setup (manual)

These steps are run once, by hand, before the CI/CD workflow can deploy anything. They are not
automated because they provision cloud resources and secrets.

### 1. Google Cloud project setup

```bash
gcloud projects create foreversixty --name "Forever Sixty"
gcloud config set project foreversixty
gcloud services enable run.googleapis.com artifactregistry.googleapis.com secretmanager.googleapis.com iamcredentials.googleapis.com
gcloud artifacts repositories create api --repository-format docker --location us-east1
printf '%s' '<neon pooled url>'   | gcloud secrets create DATABASE_URL --data-file=-
printf '%s' '<neon direct url>'   | gcloud secrets create MIGRATE_DATABASE_URL --data-file=-
printf '%s' '<resend key>'        | gcloud secrets create RESEND_API_KEY --data-file=-
gcloud iam service-accounts create api-runtime
gcloud secrets add-iam-policy-binding DATABASE_URL --member serviceAccount:api-runtime@foreversixty.iam.gserviceaccount.com --role roles/secretmanager.secretAccessor
gcloud secrets add-iam-policy-binding MIGRATE_DATABASE_URL --member serviceAccount:api-runtime@foreversixty.iam.gserviceaccount.com --role roles/secretmanager.secretAccessor
gcloud secrets add-iam-policy-binding RESEND_API_KEY --member serviceAccount:api-runtime@foreversixty.iam.gserviceaccount.com --role roles/secretmanager.secretAccessor
```

Then set up Workload Identity Federation for GitHub Actions following the
`google-github-actions/auth` README: a `deployer` service account with `roles/run.admin`,
`roles/artifactregistry.writer`, and `roles/iam.serviceAccountUser` on `api-runtime`. Record the
provider resource name and deployer email as GitHub repository variables `GCP_WIF_PROVIDER` and
`GCP_DEPLOYER_SA`.

### 2. First deploy

Like the local build, this runs from the repository root so the context includes `data/builds`:

```bash
cd "$(git rev-parse --show-toplevel)"   # the build context is the repository root
docker build --platform linux/amd64 --build-arg VERSION=$(git rev-parse --short HEAD) \
  -f api/Dockerfile -t us-east1-docker.pkg.dev/foreversixty/api/api:$(git rev-parse --short HEAD) .
docker push us-east1-docker.pkg.dev/foreversixty/api/api:$(git rev-parse --short HEAD)
# Public access: the derbee.ai organization policy forbids allUsers IAM bindings, so the service
# runs with the invoker IAM check disabled instead of --allow-unauthenticated.
gcloud run deploy api \
  --image us-east1-docker.pkg.dev/foreversixty/api/api:$(git rev-parse --short HEAD) \
  --region us-east1 --platform managed --no-invoker-iam-check \
  --service-account api-runtime@foreversixty.iam.gserviceaccount.com \
  --set-env-vars PUBLIC_BASE_URL=https://foreversixty.gg,API_BASE_URL=https://api.foreversixty.gg \
  --set-secrets DATABASE_URL=DATABASE_URL:latest,MIGRATE_DATABASE_URL=MIGRATE_DATABASE_URL:latest,RESEND_API_KEY=RESEND_API_KEY:latest \
  --min-instances 0 --max-instances 3 --cpu 1 --memory 512Mi --concurrency 80 --timeout 30
gcloud run domain-mappings create --service api --domain api.foreversixty.gg --region us-east1
```

`--memory 512Mi`, not the 256Mi the service ran on before the simulator arrived: the same image
carries `/engine/forever-sim` (see `api/Dockerfile`), and `POST /v1/sims/run` shells it with
`-plan` to size a bulk request before queueing it. `sim/runner`'s own note puts that subprocess's
JSON at roughly 22 MB at the server lane's cap, and it is built in this container, on one shared
CPU, at `--concurrency 80`. 512Mi is a conservative bound chosen from that 22 MB figure — it is
not a measurement, and nobody has watched the service's RSS under a full-cap plan. Replace it
with a measured full-cap `-plan` footprint when there is one; a real measurement may well argue
for more, or for taking the plan off the request path entirely.

`MAIL_FROM` and `TRUSTED_PROXY_HOPS` are deliberately left out of `--set-env-vars` above: both
take their documented defaults (`Forever Sixty <hello@foreversixty.gg>` and `1`, matching a
single Cloud Run edge in front of the service) unless explicitly set on the service.

Add the DNS records the last command prints to Cloudflare as DNS-only (grey cloud) so Google's
certificate validates. Verify:

```bash
curl -s https://api.foreversixty.gg/health
```

Expected: `{"ok":true,"data":{"status":"ok"},...}`.

After this, every push to `main` that touches `api/**` or `data/builds/**` runs tests, builds, and
deploys automatically.

### 3. Stripe (added by the entitlements/billing lane, done once test mode is ready)

Not required for the API to start or serve anything else — every billing route answers
`503 billing_unavailable` until these are set (`config.Config.StripeConfigured`). When ready:

```bash
printf '%s' '<stripe restricted key, sk_test_... or sk_live_...>' | gcloud secrets create STRIPE_SECRET_KEY --data-file=-
printf '%s' '<stripe webhook signing secret, whsec_...>'          | gcloud secrets create STRIPE_WEBHOOK_SECRET --data-file=-
gcloud secrets add-iam-policy-binding STRIPE_SECRET_KEY --member serviceAccount:api-runtime@foreversixty.iam.gserviceaccount.com --role roles/secretmanager.secretAccessor
gcloud secrets add-iam-policy-binding STRIPE_WEBHOOK_SECRET --member serviceAccount:api-runtime@foreversixty.iam.gserviceaccount.com --role roles/secretmanager.secretAccessor
gcloud run services update api --update-secrets STRIPE_SECRET_KEY=STRIPE_SECRET_KEY:latest,STRIPE_WEBHOOK_SECRET=STRIPE_WEBHOOK_SECRET:latest \
  --update-env-vars STRIPE_ENVIRONMENT=live
```

`STRIPE_ENVIRONMENT` must be `live` exactly when `PUBLIC_BASE_URL` is `https://foreversixty.gg`,
and the key's own prefix (`sk_live_`/`rk_live_` vs `sk_test_`/`rk_test_`) must agree with it —
the service refuses to start otherwise (`config.ValidateStripeKeyEnvironment`). Neither secret is
ever logged. See the entitlements/payments design spec §3 for the full threat model, and create
the key as a Restricted API key (not the unrestricted default secret key) with the scopes listed
there.

Run `stripe-setup` once per mode after the secrets are set, to create the four Products/Prices
(idempotent, safe to re-run):

```bash
gcloud run jobs execute api --region us-east1 --args stripe-setup --wait
```

## Logs

```bash
gcloud run services logs read api --region us-east1 --limit 100
```
