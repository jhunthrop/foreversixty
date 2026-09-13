# api

Go service for api.foreversixty.gg, deployed as a container on Google Cloud Run.

## Commands

```
make test   # go test -cover ./...
make build  # go build -o bin/api ./cmd/api
make run    # build then run (needs DATABASE_URL, RESEND_API_KEY, PUBLIC_BASE_URL, API_BASE_URL)
make lint   # go vet ./...
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

## Docker

Build the image locally:

```bash
docker build --build-arg VERSION=local -t foreversixty-api:local .
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
`MIGRATE_DATABASE_URL`, `RESEND_API_KEY`. Other env vars (`PORT`, `PUBLIC_BASE_URL`,
`API_BASE_URL`, `MAIL_FROM`, `TRUSTED_PROXY_HOPS`) are set on the Cloud Run service.

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

```bash
cd api
docker build --platform linux/amd64 --build-arg VERSION=$(git rev-parse --short HEAD) -t us-east1-docker.pkg.dev/foreversixty/api/api:$(git rev-parse --short HEAD) .
docker push us-east1-docker.pkg.dev/foreversixty/api/api:$(git rev-parse --short HEAD)
# Public access: the derbee.ai organization policy forbids allUsers IAM bindings, so the service
# runs with the invoker IAM check disabled instead of --allow-unauthenticated.
gcloud run deploy api \
  --image us-east1-docker.pkg.dev/foreversixty/api/api:$(git rev-parse --short HEAD) \
  --region us-east1 --platform managed --no-invoker-iam-check \
  --service-account api-runtime@foreversixty.iam.gserviceaccount.com \
  --set-env-vars PUBLIC_BASE_URL=https://foreversixty.gg,API_BASE_URL=https://api.foreversixty.gg \
  --set-secrets DATABASE_URL=DATABASE_URL:latest,MIGRATE_DATABASE_URL=MIGRATE_DATABASE_URL:latest,RESEND_API_KEY=RESEND_API_KEY:latest \
  --min-instances 0 --max-instances 3 --cpu 1 --memory 256Mi --concurrency 80 --timeout 30
gcloud run domain-mappings create --service api --domain api.foreversixty.gg --region us-east1
```

`MAIL_FROM` and `TRUSTED_PROXY_HOPS` are deliberately left out of `--set-env-vars` above: both
take their documented defaults (`Forever Sixty <hello@foreversixty.gg>` and `1`, matching a
single Cloud Run edge in front of the service) unless explicitly set on the service.

Add the DNS records the last command prints to Cloudflare as DNS-only (grey cloud) so Google's
certificate validates. Verify:

```bash
curl -s https://api.foreversixty.gg/health
```

Expected: `{"ok":true,"data":{"status":"ok"},...}`.

After this, every push to `main` that touches `api/**` runs tests, builds, and deploys automatically.

## Logs

```bash
gcloud run services logs read api --region us-east1 --limit 100
```
