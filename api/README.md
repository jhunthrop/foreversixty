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
| `DATABASE_URL` | yes | — | Runtime connection. Production: the Cloud SQL unix socket, with the pool size in the URL. See "Production database" below. |
| `MIGRATE_DATABASE_URL` | no | `DATABASE_URL` | Connection used only for running migrations. See "Migrations and connection poolers" below. |
| `RESEND_API_KEY` | yes | — | |
| `PUBLIC_BASE_URL` | yes | — | e.g. `https://foreversixty.gg`; confirm/unsubscribe redirects land here. |
| `API_BASE_URL` | yes | — | e.g. `https://api.foreversixty.gg`; used to build the confirmation link. |
| `MAIL_FROM` | no | `Forever Sixty <help@foreversixty.gg>` | `From:` address for outgoing mail. |
| `MAIL_REPLY_TO` | no | `help@foreversixty.gg` | `Reply-To:` on every mail; the help@ inbox is Google Workspace. |
| `TRUSTED_PROXY_HOPS` | no | `1` | How many reverse proxies in front of this service (e.g. Cloud Run) are trusted to append to `X-Forwarded-For`; used to resolve the real client IP for rate limiting. `0` ignores `X-Forwarded-For` entirely and rate-limits by the raw connection address. Must be an integer >= 0. |
| `R2_ACCOUNT_ID` | no | — | Cloudflare account id; the S3 endpoint is derived from it. |
| `R2_ACCESS_KEY_ID` | no | — | R2 API token id (Object Read & Write). |
| `R2_SECRET_ACCESS_KEY` | no | — | R2 API token secret. |
| `R2_BUCKET` | no | `foreversixty-logs` | The bucket every report file lives in. |
| `R2_ENDPOINT` | no | derived | S3 endpoint override; the tests point it at an in-process fake. |
| `SESSION_COOKIE_DOMAIN` | no | `.foreversixty.gg` | `Domain` on `fs_session` and `fs_csrf`. Set it to `none` for a local `http://localhost` run, which makes the cookies host-only. |
| `BNET_CLIENT_ID` | no | — | Battle.net OAuth client, the only sign-in the site offers. Without the id, the secret, and the redirect URL, the Battle.net routes answer 404. |
| `BNET_CLIENT_SECRET` | no | — | |
| `BNET_REDIRECT_URL` | no | — | Registered for production and for `http://localhost:8080/v1/auth/battlenet/callback`. |
| `BNET_PROFILE_GAME` | no | `forever` | The profile-game segment of a Blizzard namespace (`profile-<game>-<region>`, `dynamic-<game>-<region>`). Blizzard serves Forever as `profile-forever-<region>` since 2026-11-03. |
| `BNET_IMPORT_CHARACTERS` | no | on | The account character import on sign-in, the primary way characters reach the site. `false` is the incident switch: sign-in still works and characters come from the addon alone. |
| `EMAIL_SIGN_IN` | no | off | `true` mounts `POST /v1/auth/email`, the emailed sign-in link. Off (the default since launch), the route answers 404; the link landing `GET /v1/auth/email/callback` stays mounted so a link already sent still redeems. The Battle.net-outage fallback: on its own it brings back no button. See "Email sign-in: the Battle.net-outage fallback" below. |
| `BNET_REGIONS` | no | `us,eu,tw` | Comma-separated; which regional Blizzard hosts to try for an account during import and refresh. |
| `BNET_PROBE_GAMES` | no | `forever,classic1x` | Comma-separated; every game namespace segment the nightly refresh's namespace probe checks, logging one line per game (any game other than `BNET_PROFILE_GAME` that answers 200 is logged at WARN as `namespace_appeared`). |
| `PARSE_JOB_NAME` | no | `parse-report` | The Cloud Run job that parses a whole-file upload. |
| `PARSE_JOB_REGION` | no | `us-east1` | |
| `PARSE_JOB_PROJECT` | no | `foreversixty` | The Google Cloud project the job lives in. |
| `SIM_JOB_NAME` | no | `sim-run` | The Cloud Run job that runs one premium sim. |
| `SIM_JOB_REGION` | no | `us-east1` | |
| `SIM_JOB_PROJECT` | no | `foreversixty` | The Google Cloud project the sim jobs live in. |
| `STRIPE_SECRET_KEY` | no | — | Stripe secret or restricted key. Absent, every `/v1/billing/*` route answers 503 and the Premium page's buttons fail honestly; present, billing is on. Lives in Secret Manager (`STRIPE_SECRET_KEY`, live; `STRIPE_SECRET_KEY_TEST`, the sandbox). |
| `STRIPE_WEBHOOK_SECRET` | no | — | The signing secret of the `POST /v1/billing/webhook` endpoint in the same mode as the key. |
| `STRIPE_ENVIRONMENT` | no | — | `test` or `live`. Must be `live` exactly when `PUBLIC_BASE_URL` is the production origin, and the key's own prefix must agree; `config.ValidateStripeKeyEnvironment` refuses to start otherwise. |
| `STRIPE_CONNECT_WEBHOOK_SECRET` | no | — | The Connect endpoint's signing secret, on the same URL as `STRIPE_WEBHOOK_SECRET` (`POST /v1/billing/webhook`): `account.updated` for creators' payout accounts. The webhook accepts a signature from either secret. Without it, onboarding status is still re-read on the creator's return (`?done=1`) and by the owner's Refresh. |
| `STRIPE_PAYOUTS_KEY` | no | — | The creator program's own restricted key, platform-account permissions only (Balance read; Accounts, Login Links and Transfers write): every Connect and transfer call, never on a connected account's behalf. Subscriptions never use it. Missing, onboarding and payouts answer 503; with `CREATOR_PAYOUTS` on, the startup check says so on the admin pages. Secret Manager; never logged. |
| `CREATOR_PAYOUTS` | no | off | `true` lets the owner approve a creator's month, which makes a Stripe transfer. Off, statements read normally, Approve is disabled and `POST .../approve` answers 503 `payouts_off`. Leave it off until live Connect is set up and a first statement has been checked against the Dashboard. |
| `OUR_RENEWAL_EMAILS` | no | off | `true` has billing-daily send our own yearly renewal notice a week ahead. Off, it sends none: Stripe's upcoming-renewal email covers renewals (premium-trial spec 0C). |
| `PREMIUM_TRIAL` | no | off | `true` turns the free trial on (premium-trial spec): a first Premium checkout trials 7 days, billing-daily sends the 3-day trial reminder (Premium and guild) with its one-click cancel link, and the first charge after a trial gets a receipt, all to the Stripe customer's email read at send time and never stored or logged (the `billing-daily` job needs `STRIPE_SECRET_KEY` for that read; without it a reminder goes in the app alone). Off until the trial's pages ship; a creator's code gives its free month either way. |
| `NOTIFY_ENCRYPTION_KEY` | no | — | Base64 of 32 random bytes (`openssl rand -base64 32`); seals Discord webhook tokens at rest (AES-256-GCM). Absent, the Discord notifications feature is off; present but malformed, the service refuses to start. Lives in Secret Manager; never logged or printed. When set, the sim-run and parse-report jobs post one Discord embed per finished server sim or parsed upload; five consecutive counted failures (a 429 does not count) switch the webhook off. Migration 0047 adds `webhook_name` and backfills `upload_private` true for existing webhooks (new ones default off). |
| `DISCORD_CLIENT_ID` | no | — | The Discord application's Client ID (public). Together with `DISCORD_CLIENT_SECRET` it turns the one-click Discord connect on (needs `NOTIFY_ENCRYPTION_KEY` too); `/v1/me` then carries `notifications.connect: true`. Both unset leaves connect off (its routes answer 404); exactly one set refuses to start, naming the missing variable. The redirect URI is `API_BASE_URL` + `/v1/notifications/discord/callback`. |
| `DISCORD_CLIENT_SECRET` | no | — | The Discord application's Client Secret. Lives in Secret Manager; never logged, printed or returned. |
| `TREE_DATA_DIR` | no | `/data` | Directory holding one subdirectory per client build (`<build>/talents/*.json`, `<build>/items/*.json`, `<build>/{sets,classes,races,combos}.json`). The Docker image copies `data/builds` here, and `web/src/data/active-build.json` to `<dir>/active-build.json`: the API's active build (the Battle.net import's tables and FS1 stamp, spec inference, the default DPS-per-point and rotation build, the BiS build) is the build that file names, never merely the newest folder, so a build `data.yml` fetched but `tools/activate-build.sh` has not activated changes nothing. With no file beside the builds it reads the repository's `../../web/src/data/active-build.json` (so `TREE_DATA_DIR=../data/builds` works locally); with neither it falls back to the newest client build folder and logs it. A file naming a build that is not on disk fails the boot. A missing or pre-Phase-1 directory is logged at startup and simply has no data: the service still serves health, version, and subscribe, and every save fails validation on `tree_version`. |

Every Phase 3 variable is optional and degrades honestly: with no R2 credentials the service
serves everything that does not touch object storage and logs that the ingest, upload, and
report-file routes are not mounted; with no Battle.net client nobody can sign in (the email link
is behind `EMAIL_SIGN_IN`, below); with no Cloud Run credentials it serves the live ingest but not whole-file uploads.

### Migrations and connection poolers

`db.Migrate` (via golang-migrate) takes a session-level Postgres advisory lock while migrations
run, so it needs a session-pooled or direct (non-pooled) connection. A transaction-pooled
connection (e.g. Neon's default pooled/PgBouncer endpoint) cannot hold that lock reliably.

When `MIGRATE_DATABASE_URL` is unset, `DATABASE_URL` is used for migrations too. Production
connects directly (no pooler), so both URLs name the same socket. `MIGRATE_DATABASE_URL` stays
separate because golang-migrate's lib/pq driver rejects pgx's `pool_*` parameters.

### Production database

Since 2026-11-04 production runs on the Cloud SQL instance `commish-pg` (project `trycommish`,
Postgres 18, db-g1-small, us-east1). That instance is shared with the Commish app, so Forever
Sixty is fenced into its own space:

- Database `foreversixty`, owned by the login role `foreversixty_app`. That role isn't a member of
  `cloudsqlsuperuser`, has `CONNECTION LIMIT 30`, and holds no grants in the `commish` database.
  `PUBLIC` connect on `foreversixty` is revoked, so Commish's roles can't open it. Inside
  `foreversixty`, `PUBLIC` has no `CONNECT`, `TEMP` or schema `public` rights, and there are no
  default privileges.
- The fence runs both ways (since 2026-11-04). `PUBLIC` `CONNECT` is revoked on `commish`,
  `postgres` and `template1`, so `foreversixty_app` gets `permission denied for database`. The
  Commish roles (`commish_app`, `commish_research`), `postgres`, `cloudsqlsuperuser` and the Cloud
  SQL service roles (`cloudsqlagent`, `cloudsqlimportexport`, `cloudsqllogical`,
  `cloudsqlobservability`, `cloudsqlreplica`, `cloudsqlconnpooladmin`, `cloudsqliam*`) hold
  explicit `CONNECT` on those three. `PUBLIC` keeps `TEMP`. A database created later with
  `gcloud sql databases create` starts with `PUBLIC` connect again, so revoke it there too. To
  undo the fence, `grant connect on database commish, postgres, template1 to public`.
- The fence has a known limit. `commish_app` and `commish_research` belong to
  `cloudsqlsuperuser` and have `CREATEROLE`, so they can still alter `foreversixty_app`, including
  its password. Removing that needs changes on the Commish side, so it's left to the owner.
- Instance state on 2026-11-04: public IP is on with no authorized networks, so connections come
  only through the connector or proxy. SSL isn't required (`ALLOW_UNENCRYPTED_AND_ENCRYPTED`),
  and both apps use the connector's socket with `sslmode=disable`. Daily backups (7 kept) and
  point-in-time recovery (7 days of logs) are on.
- The service and every job mount the instance through the Cloud Run Cloud SQL connector
  (`--add-cloudsql-instances` / `--set-cloudsql-instances trycommish:us-east1:commish-pg`). The
  runtime account `api-runtime@foreversixty` holds `roles/cloudsql.client` on `trycommish`.
- `DATABASE_URL` is `postgres://foreversixty_app:…@/foreversixty?host=/cloudsql/trycommish:us-east1:commish-pg&sslmode=disable&pool_max_conns=6&pool_max_conn_idle_time=10m`.
  `MIGRATE_DATABASE_URL` is the same without the `pool_*` parameters. Both live only in Secret
  Manager (version 2; version 1, the old Neon URL, is destroyed).
- Connection budget: 6 per process × 3 service instances (`maxScale` 3) = 18. That leaves 12 of
  the role's 30 for jobs and deploy overlap. Raising `maxScale` or the pool means raising the
  role limit with it: `alter role foreversixty_app connection limit N`. `foreversixty_app` can't
  change its own limit, so run that as a temporary built-in admin. Create one with `gcloud sql
  users create <name> --instance commish-pg --project trycommish --password=<generated>` (it
  joins `cloudsqlsuperuser`), connect through the proxy to the `postgres` database, and delete it
  afterwards with `gcloud sql users delete`. Never change the passwords of the existing roles.
  Never resize or restart `commish-pg` without the owner; it interrupts Commish.

A tool that reads the production `DATABASE_URL` from Secret Manager can't use the
`/cloudsql/...` socket from a laptop. Run `cloud-sql-proxy --gcloud-auth --port 6543
trycommish:us-east1:commish-pg` and swap the host for `127.0.0.1:6543`. The rule that only the
deployed API migrates still holds.

**Recovery** (Neon is deleted; Cloud SQL backups are the only copy). `commish-pg` keeps 7 daily
automated backups (08:00 UTC) and 7 days of transaction logs for point-in-time recovery. Both are
instance-wide. Never restore a backup or PITR onto `commish-pg` itself: it would roll Commish back
too and restart the instance. The first daily backup holding `foreversixty` is 2026-11-04 08:00
UTC; before that only point-in-time recovery covers it. Recover into a throwaway instance and
copy only `foreversixty`:

```bash
T=--project=trycommish
gcloud sql backups list --instance commish-pg $T                     # pick a backup, or use a time
gcloud sql instances clone commish-pg fs-recovery $T --point-in-time '2026-11-05T12:00:00Z'
# Or restore a daily backup into a new instance instead:
#   gcloud sql instances create fs-recovery $T --database-version POSTGRES_18 --tier db-g1-small --region us-east1
#   gcloud sql backups restore <BACKUP_ID> --restore-instance fs-recovery --backup-instance commish-pg $T
cloud-sql-proxy --gcloud-auth --port 6545 trycommish:us-east1:fs-recovery &
pg_dump -Fc -d 'host=127.0.0.1 port=6545 user=foreversixty_app dbname=foreversixty' -f fs.dump
```

Stop writes before loading: close the API (`gcloud run services update api --project
foreversixty --region us-east1 --ingress internal`) and pause the Forever schedulers. Then restore
`fs.dump` into `commish-pg`'s `foreversixty` as `foreversixty_app` through the usual proxy with
`pg_restore --clean --if-exists --no-owner`. Reopen with `--ingress all`, resume the schedulers,
and delete the clone with
`gcloud sql instances delete fs-recovery --project trycommish`. The clone carries Commish's data
too, so don't keep it.

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

## Discord application

The one-click connect needs a Discord application (https://discord.com/developers/applications). One application serves every environment.

- Name it `Forever Sixty` and use our seal as the icon (the same image the posts carry as their avatar).
- OAuth2 > Redirects: add both URIs, exactly: `https://api.foreversixty.gg/v1/notifications/discord/callback` (production) and `http://localhost:8080/v1/notifications/discord/callback` (local runs — Discord accepts plain `http` only for `localhost`, not `127.0.0.1`; set `API_BASE_URL=http://localhost:8080` for a local connect). A staging deployment adds its own `API_BASE_URL` + `/v1/notifications/discord/callback`. Paste each URI without a trailing space: the portal rejects the whole save with "You provided an invalid redirect URI" when any entry has one. The Cloud Logging `_Default` sink carries an exclusion for the callback request line, because Cloud Run's request log would otherwise store Discord's one-time `code`; the app logs the callback outcome itself.
- Client ID is public: set `DISCORD_CLIENT_ID`. The Client Secret goes in Secret Manager as `DISCORD_CLIENT_SECRET`.
- No bot, no interactions endpoint, no privileged intents: the only scope is `webhook.incoming`.

## Cache headers

Every `GET` route sets its `Cache-Control` through one of two `httpx` helpers
(`api/internal/httpx/cache.go`), never by hand:

- `httpx.CachePublic(w, maxAge, staleWhileRevalidate)` — a shared cache and the browser may
  both keep the answer. Four values are in use:
  - **Live public boards** (rankings, ratings, the reports feed, a guild or character's public
    page): `public, max-age=30, stale-while-revalidate=300`.
  - **Public per-object reads** (a public report's meta/fights, a saved sim result, a public
    character's sim input): `public, max-age=60, stale-while-revalidate=600`.
  - **Reference** (spec lists, builds, phase boundaries): today's per-route value, no
    `stale-while-revalidate`.
- `httpx.CachePrivate(w)` — `Cache-Control: private, no-cache` plus `Vary: Cookie,
  Authorization`, so a shared cache never keys one caller's answer for another. Every
  signed-in-only read (`/v1/me`, devices, your reports, your sims, guild settings/home,
  private/guild reports) uses this.

A route whose answer depends on who is asking — a report that is public for one viewer and
private for another — branches between the two at request time rather than picking one
statically; see `reports.Service.get` or `sims.Service.simInput` for the pattern. The guild
centre's Progression tab (`GET /v1/guilds/{id}/progression`) is public, and its counts depend
on the viewer: a verified member of the guild reads the guild's own numbers, which count its
private nights (never a flagged report), in a private answer; anyone else reads the public
numbers the public guild page prints, in a public answer. Both carry `Vary: Cookie,
Authorization`.

Every successful (2xx) `GET` also carries `ETag: W/"<16 hex chars>"`, a weak tag computed over
the envelope's `data` field alone. A matching `If-None-Match` (weak or strong, comma-separated)
answers `304 Not Modified` with an empty body and the same `ETag`/`Cache-Control` the `200`
would have carried — the cheap revalidation path a client can use once it already holds a
cached answer, regardless of which class that route is in.

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

The router-wide limit keeps two buckets by credential kind: a request with a device bearer
(`Authorization: Bearer fsd_...`) is counted against that token's own 120-per-minute bucket (keyed by
the token's hash), never against the per-IP bucket, so a companion cannot starve a browser on the
same address; a token the API has answered 401 is cut to 10 per minute. Requests without a device
bearer stay in the 120-per-minute per-IP bucket.

Builds are validated against the tree data for their `tree_version` before they are stored; the rules
are in `internal/builds/validate.go` and come from the Phase 1 interface contract, which has the web
planner mirror them in `web/src/lib/planner/rules.ts`.

## Simulator and phase endpoints

The simulator's full surface (`POST /v1/sims`, `POST /v1/sims/run`, `GET /v1/sims/{id}`,
`GET /v1/sims/{id}/progress`, `POST /v1/sims/{id}/seen`, `POST /v1/sims/{id}/rerun`, `GET /v1/specs`, and the sim-input route) is documented in
`openapi.yaml`, not restated here. The rows below are the ones that do not belong to any other table in this file:

| Route | Notes |
|---|---|
| `GET /v1/sims?mine=1&kind=&lane=` | The caller's own sims, newest first, optionally narrowed to one tool (`run`, `gear`, `talents`, `drops`, `weights`) and to one lane (`server`, `browser`). Each row carries its kind, a composed one-line headline, its lane, `finished_at`, `failure_reason` and `seen`; queued and running rows carry `estimate_seconds`. An unknown kind or lane is 400. |
| `POST /v1/sims/{id}/seen` | Marks the caller's finished server run as looked at (204); 404 when it is not theirs or not done. |
| `POST /v1/sims/{id}/rerun` | Premium. Starts a failed run's stored request again as a new run (202 `{sim_id}`) under the checks `POST /v1/sims/run` applies; 404 when it is not the caller's failed run. |
| `PUT /v1/notifications/discord` | Premium. Saves the Discord webhook (`webhook_url?`, `events`: all of `sim_finished`, `upload_parsed`, `upload_private`; a private report posts only when `upload_parsed` and `upload_private` are both on) and answers the block `/v1/me` carries under `notifications.discord`; a save with a URL also stores Discord's `name` for the webhook (returned as `name`, null when none); an events-only change omits the URL, keeps the name (404 when none saved). A bad URL is 400, a URL Discord does not accept 422, both `invalid`; 10 per minute per member. Mounted only when `NOTIFY_ENCRYPTION_KEY` is set; otherwise `/v1/me` omits `notifications`. |
| `DELETE /v1/notifications/discord` | Removes the webhook (204, idempotent, allowed after a lapse). A connected webhook is also deleted at Discord first, best effort, and only while the member is entitled; a pasted one never is. |
| `POST /v1/notifications/discord/test` | Premium. Posts the test message (204); 502 `upstream` with Discord's own message when it answers with an error, 504 `upstream` when it does not answer; 5 per minute per member. |
| `GET /v1/notifications/discord/connect` | Premium, session cookie. Sets the `fs_discord_state` cookie (httpOnly, `SameSite=Lax`, path `/v1/notifications/discord/callback`, 10 minutes) and answers 302 to Discord's `oauth2/authorize` with `scope=webhook.incoming`. 302 to `/account?discord=failed#discord` without the entitlement; 404 when connect is off. |
| `GET /v1/notifications/discord/callback` | Session cookie. Clears the state cookie, then 302 to `/account?discord=` `connected`, `denied` (no `code`, or `error=access_denied`) or `failed` (bad state, failed exchange, unusable webhook, store error) `#discord`. A connect stores the sealed webhook with `source = connected` (no name), keeps an existing webhook's events (defaults sim on, log on, private off), and removes the replaced connected webhook from Discord. The code, token and URL never reach a header or log. |
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
| `GET /v1/characters/{region}/{ruleset}/{name}/history` | The owner's score and parse history, every rated fight in every report they can see. Owner with the history entitlement only: 401 anonymous, 402 `premium_required`, 404 for anyone else and for an anonymised character. `private, no-store`. |
| `GET /reports/{id}/card.png` | The unfurl card, cached five minutes — `private` and `Vary: Cookie` when the report is not public or unlisted, so a shared cache never serves one reader's card to another. |

Sign-in is Battle.net only (the email magic link is dark behind `EMAIL_SIGN_IN`, below); the
companion pairs with a code from `POST /v1/devices/pair` and uploads with `Authorization: Bearer fsd_…`. Browser
sessions are opaque cookies with double-submit CSRF (`fs_csrf` plus `X-CSRF-Token`); device
tokens are exempt, because a companion is not a browser.

### Email sign-in: the Battle.net-outage fallback

Battle.net has been the only sign-in since launch (4 Nov 2026). The emailed magic link's code is
kept, dark: `POST /v1/auth/email` is mounted only when `EMAIL_SIGN_IN=true`, and answers the
router's 404 otherwise. Its landing, `GET /v1/auth/email/callback`, is always mounted, so turning
the switch off never breaks a link already in someone's inbox (links last twenty minutes).

Turning the switch on alone brings nothing back for players: the web build has no email entry
point any more, by design, so re-enabling is a deliberate change. During a Battle.net outage:

1. Restore a web entry point. The "Email me a sign-in link" form on `/login` (and, if wanted, the
   header selector's and `/connect`'s) was removed in the commit `feat: Battle.net is the only
   sign-in`; reverting its `web/` part restores `requestEmailLink` in `web/src/lib/account/api.ts`,
   the forms in `Account.svelte` and `Connect.svelte`, and their copy. Push it so the site deploys.
2. Turn the route on: `gcloud run services update api --region us-east1 --update-env-vars
   EMAIL_SIGN_IN=true`. The deploy workflow passes no env vars, so the setting survives later
   deploys until it is removed with `--remove-env-vars EMAIL_SIGN_IN`.
3. Check `RESEND_API_KEY` and `MAIL_FROM` still send (the subscribe confirmation uses the same
   mailer), and restore the privacy policy's line that an email address is collected for sign-in
   (`web/src/content/pages/privacy.md`).

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

## Stripe: going live

Billing is off until the three `STRIPE_*` variables exist on the Cloud Run service; the
deploy workflow touches no Stripe secret or switch (the creator program's `STRIPE_PAYOUTS_KEY` and
`CREATOR_PAYOUTS` included), so each is set once by hand and survives every deploy.
The account runs Stripe's Managed Payments (Stripe is the merchant of record and handles
sales tax), which is why `stripe-setup` gives every product a tax code. In order:

1. Activate the Stripe account (identity, bank account, statement descriptor `FOREVERSIXTY`).
2. Put the live secret key in Secret Manager, never in a shell history or a chat:
   `gcloud secrets create STRIPE_SECRET_KEY --project foreversixty --data-file=<file>` (then
   delete the file). The sandbox key is `STRIPE_SECRET_KEY_TEST`.
3. Create the products, prices and the Customer Portal configuration in live mode with the
   same command the sandbox got, from a shell whose `STRIPE_SECRET_KEY` is read from Secret
   Manager and `STRIPE_ENVIRONMENT=live`; `PUBLIC_BASE_URL=https://foreversixty.gg` so the
   portal links the real legal pages: `go run ./cmd/api stripe-setup`. It is idempotent.
4. Create the webhook endpoint for `https://api.foreversixty.gg/v1/billing/webhook` through the
   API, not the Dashboard, with `api_version` set to the version stripe-go was built for
   (`stripe.APIVersion`, `2025-08-27.basil` for v82): the library rejects an event of any other
   version as unverifiable, and a Dashboard-made endpoint gets the account's newer default.
   Events: `checkout.session.completed`, `customer.subscription.created`,
   `customer.subscription.updated`, `customer.subscription.deleted`, `invoice.paid` and
   `invoice.payment_failed`. Put its signing secret in Secret Manager as
   `STRIPE_WEBHOOK_SECRET` without it ever touching a terminal.
5. Wire the service, which restarts it with billing on:
   `gcloud run services update api --region us-east1 --update-secrets STRIPE_SECRET_KEY=STRIPE_SECRET_KEY:latest,STRIPE_WEBHOOK_SECRET=STRIPE_WEBHOOK_SECRET:latest --update-env-vars STRIPE_ENVIRONMENT=live`
6. Create the `stripe-reconcile` job and its nightly schedule (below), then buy and cancel one
   real Premium month yourself and refund it in the Dashboard: the webhook lands, the account
   page shows Premium, the portal cancels it at period end.

The entitlements/billing lane adds two more, both cheap and short-lived — `stripe-setup` is
run by hand once per mode (see "First-time setup" above); `stripe-reconcile` is scheduled
nightly the same way `sim-validate` already is:

    gcloud run jobs create stripe-reconcile \
      --image <the API image> --region us-east1 --args stripe-reconcile \
      --cpu 1 --memory 512Mi --task-timeout 5m

    gcloud scheduler jobs create http stripe-reconcile-nightly \
      --schedule "0 4 * * *" --uri "https://us-east1-run.googleapis.com/apis/run.googleapis.com/v1/namespaces/foreversixty/jobs/stripe-reconcile:run" \
      --http-method POST --oauth-service-account-email api-runtime@foreversixty.iam.gserviceaccount.com

## Creator program

The creator program (`internal/creators`, `design/specs/2026-11-08-creator-program.md`) runs on the
deployment's own Stripe key: the sandbox until the owner sets up live Connect. Creators are paid $2
per returning player and half of their players' net subscription revenue for 12 months, monthly,
through Stripe Connect Express. Nothing is paid until the owner presses Approve on
`/admin/creators/payouts`, and only while `CREATOR_PAYOUTS=true`. The owner's pages are
`/admin/creators` and `/admin/creators/payouts`; a creator's are `/creators/onboarding` (Stripe's
return and refresh URLs) and `/creators/stats`; the public page is `/creators` with
`/creators/apply`.

A creator's code is our own, not a Stripe promotion code: `POST /v1/billing/checkout` takes
`creator_code`, and a live code gives the subscription a 30-day trial on Premium or a guild plan
(once per account, `premium_trials`; once per guild, `guild_trials`). The buyer becomes the creator's
player when `checkout.session.completed` lands.

The free trial (`design/specs/2026-11-08-premium-trial.md`) sits behind `PREMIUM_TRIAL=true`, off by
default: on, a first Premium checkout trials `catalog.PremiumTrialDays` (7), billing-daily sends the
3-day trial reminder (Premium and guild) with the one-click cancel link, and the first charge after a
trial gets a receipt. The reminder and receipt go to the Stripe customer's email, read at send time
and never stored. Its server side is landed (`POST /v1/billing/checkout/terms`, `/trial/cancel`,
`/trial/keep`, `/cancel-link/preview`, `/cancel-link`, migration `0101_trial_reminders`); the
chooser, the account page's trial row and the `/premium/cancel` page are not built yet, so the
switch stays off until they ship. A creator's code gives its free month either way.

Going live, in order (the owner's Stripe account, never a key in a shell history):

1. Turn on Connect for the live account (Express, US recipients) and finish the platform profile.
   Live Connect is on (an accounts list answers, with 0 accounts); check that transfers work under
   Managed Payments, which the sandbox could not show.
2. Make the payouts key; the subscription key is never widened. The subscription key stays the
   restricted `rk_live_…` key in Secret Manager project `foreversixty`, secret `STRIPE_SECRET_KEY`
   (it also reads a creator's player's invoices, invoice payments, payment intents, charges,
   refunds and balance transactions for the revenue share, which billing already reads). Creator
   payouts use a separate restricted key with exactly these four platform-account permissions,
   and none "in connected accounts":
   - **Balance: Read**;
   - **Accounts: Write** (makes the Express account, reads its status, lists accounts for the
     startup check; it covers Account Links, the onboarding link);
   - **Login Links: Write** (the creator's Express dashboard link);
   - **Transfers: Write** (the approved payout, and reading a transfer back by its group).

   No call is ever made on a connected account's behalf (no `Stripe-Account` header; a test holds
   it): what a creator sees comes from our own records, and their Express dashboard covers the rest.

   Put it in Secret Manager project `foreversixty` as `STRIPE_PAYOUTS_KEY` without it touching a
   terminal, and wire it: `gcloud run services update api --region us-east1 --update-secrets
   STRIPE_PAYOUTS_KEY=STRIPE_PAYOUTS_KEY:latest`. Every Connect and transfer call goes through it.
   In tests and the sandbox both keys are the Connect sandbox's `STRIPE_SANDBOX_SECRET_KEY`
   (project `trycommish`).

   With `CREATOR_PAYOUTS=true` the API lists one connected account with the payouts key at
   startup. A missing key, or one without Connect access, logs `creators op=connect_check
   state="the stripe payouts key is missing or lacks connect permissions"` at ERROR, both admin
   pages say `Stripe payouts key missing or lacks Connect permissions`, and Approve answers 503
   `connect_denied` until the key is fixed and the service restarts.
3. Add `charge.refunded` to the live account endpoint's events (refunds of a creator's player's
   payment become refund lines), keeping the six it has.
4. Create a second endpoint for the same URL, `https://api.foreversixty.gg/v1/billing/webhook`,
   with `connect: true`, the same `api_version` as the first (`2025-08-27.basil`) and the one event
   `account.updated`. Put its signing secret in Secret Manager as `STRIPE_CONNECT_WEBHOOK_SECRET`
   and wire it with `--update-secrets` as the first one was.
5. Turn payouts on by hand on the service, never from the deploy (the deploy step merges its own
   variables and leaves these alone): `gcloud run services update api --region us-east1
   --update-env-vars CREATOR_PAYOUTS=true --update-secrets STRIPE_PAYOUTS_KEY=STRIPE_PAYOUTS_KEY:latest`
   (the runtime account `api-runtime@foreversixty` holds secretAccessor on that secret, granted
   2026-11-08). Until then the program runs with payouts off: applications, approvals, invites,
   links, codes and statements all work, Approve is disabled, both admin pages say `Payouts are
   switched off on the server. Statements can be read but not paid.`, and the API answers 503
   `payouts_off`. Payouts are still one statement at a time, only on the owner's Approve. After the
   switch, `/admin/creators` shows neither `Payouts are switched off` nor the payouts key notice.
6. Approve checks the platform's available USD balance first (Balance read on the payouts key) and
   the dialog shows `Available in Stripe: $X. This payout: $Y.`; a short balance turns Pay off with
   `Not enough in your Stripe balance. Add funds in Stripe (Balances → Add to balance) or wait for
   subscription payments to clear.` The owner's balance pays out to the bank automatically, so it is
   often low. A transfer Stripe still refuses with `balance_insufficient` leaves the statement unpaid
   with nothing recorded (the claim is undone); the same Approve later is safe (same transfer group
   and idempotency key).

Adding a creator (from the form or by approving an application) gives the creator's own account
Premium, as an entitlements row with source `creator`; it never replaces a paid subscription, and
removing the creator ends only that row. The Raid guild plan for a creator's guild stays a hand
grant: when the owner asks, the coordinator runs `api grant --guild <id> --plan guild --note
"creator program"` from the deployed sha (the grant CLI migrates at start).

`stripe-reconcile` heals any webhook Stripe's own three-day retry window never successfully
delivered (design spec §2.8): it is a backstop, not the primary path — the webhook itself
(`POST /v1/billing/webhook`) does the real-time work. It also deletes `stripe_events` rows
received more than 30 days ago (raw payloads carry the buyer's name, email and address, and
Stripe resends nothing older). A row that fails is logged with its subscription id and the job
carries on; it exits non-zero naming every failure.

A live grant holds its row against Stripe: a subscription billing beside a whole-plan grant is
recorded as a `grant_conflict` anomaly to cancel or refund by hand (one already set to end at
its period end is only logged), reconcile re-reads it nightly so it takes the row the night the
grant lapses, and `grant` refuses a
subject whose Stripe subscription still bills (cancel it first; one set to end at its period end
is fine).

`raid-posts` is the guild raids' clock: deferred posts (a raid saved more than a week ahead
posts a week before) and the day-before reminders go out every ten minutes, even when the
service is cold. It runs the raid schedule, then up to five bot drain passes (only when the bot
is configured), logs `posted`, `reminded` and `failed` counts, and exits non-zero on a
schedule error. It migrates at start like every job, so it must run the deployed image; CI only
moves an existing job to the new image, so create it once by hand. It takes the service's
environment and secrets by name (never paste a value):

    gcloud run jobs create raid-posts \
      --image <the API image> --region us-east1 --args raid-posts \
      --cpu 1 --memory 512Mi --task-timeout 5m \
      --service-account api-runtime@foreversixty.iam.gserviceaccount.com \
      --set-env-vars "$(tr '\n' ',' < .env.job)" \
      --set-secrets DATABASE_URL=DATABASE_URL:latest,MIGRATE_DATABASE_URL=MIGRATE_DATABASE_URL:latest,NOTIFY_ENCRYPTION_KEY=NOTIFY_ENCRYPTION_KEY:latest,DISCORD_CLIENT_SECRET=DISCORD_CLIENT_SECRET:latest,DISCORD_BOT_TOKEN=DISCORD_BOT_TOKEN:latest

    gcloud scheduler jobs create http raid-posts-every-10m \
      --schedule "*/10 * * * *" --uri "https://us-east1-run.googleapis.com/apis/run.googleapis.com/v1/namespaces/foreversixty/jobs/raid-posts:run" \
      --http-method POST --oauth-service-account-email api-runtime@foreversixty.iam.gserviceaccount.com

The raid post's roster marks each sign-up with its spec's icon, uploaded once as the
application's own emoji (`fs_{class}_{spec}`, one per row of `data/curated/specs.json`; 28 today).
`discord-emoji` builds 128 px PNGs from the active build's icons (the build
`web/src/data/active-build.json` names) under `web/public/data/<build>/icons` (generated web data, so run it from a checkout whose
`web/public/data` holds the active build, not from the image; `--icons <dir>` names another folder) and uploads only the names the application lacks; a second run uploads
nothing, and `--replace` deletes and uploads them all again after an icon changes. It prints
names and counts, never the token. Run it once per application (production, then the test
application), with the service's two values in the shell's environment, taken by name and never
pasted into a command line or a file:

    cd api && DISCORD_APPLICATION_ID=… DISCORD_BOT_TOKEN=… go run ./cmd/api discord-emoji

(set both from Secret Manager and the service's variables, for example with
`export DISCORD_BOT_TOKEN="$(gcloud secrets versions access latest --secret DISCORD_BOT_TOKEN)"`
in a shell whose history is off). The bot reads the list once at its first post and again at
most hourly while any is missing (`emoji_missing` in the logs); until then those lines read
`{Name} · {Spec} {Class}`. A build that adds a spec adds its row to `specs.json`, and the next
run uploads only that one.

`.env.job` must carry the non-secret variables the service has (`PUBLIC_BASE_URL`,
`API_BASE_URL`, `DISCORD_CLIENT_ID`, `DISCORD_APPLICATION_ID`, `DISCORD_PUBLIC_KEY`). Without
the bot's variables the job still posts through webhooks.

Accepted, not swept: if an admin deletes a bot post in Discord after its day-before reminder
reply went out, that reply's buttons stay pressable and nothing withdraws it.

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

`rating-rerate-flagged` brings the stored Forever Score back into line once reports are
flagged (tampered or unverified). The score never reads a flagged report, but the snapshot
the rater rolled onto a character's fights may have counted it. The job finds every character
whose score a flagged report moved, in three groups, each character in the first it fits:

- the characters with a rated row in a flagged report (its fights, discipline and parses);
- the guild-mates of the flagged night, whose attendance counted it among the guild's nights,
  there or not: the guild's verified members (the membership the score's attendance counts
  nights for), and every character with a rated fight in the last 90 days (the attendance
  window) in a report of that guild;
- the characters linked to the same account as a character in the flagged report, whose
  experience counted its kills.

For each it re-rates the character's newest clean fight through the backfill's own path and
rolls the score up onto it again. It logs, per group, how many characters and fights it found,
re-rated and failed, and exits non-zero when any fight failed (every fight is still tried), so
the job's failure alert fires. It is idempotent, so run it after the deploy that carries it and
again whenever reports are flagged (or schedule it):

    gcloud run jobs create rating-rerate-flagged \
      --image <the API image> --region us-east1 --args rating-rerate-flagged \
      --cpu 2 --memory 2Gi --task-timeout 60m \
      --set-env-vars "$(tr '\n' ',' < .env.job)"

    gcloud run jobs execute rating-rerate-flagged --region us-east1 --wait

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

`bnet-refresh` (`api/internal/bnetimport`) re-syncs every due `characters` row with
`source = 'bnet'` against Blizzard's public profile and guild-roster data (the app's own
client-credentials token; no user token is needed for this job). A row is due once its
`refreshed_at` is older than 3 hours for an account active in the last 14 days
(`users.bnet_synced_at`), or 20 hours for everyone else. It runs at most 4 characters per
second and stops on the first `429`. It then runs the namespace probe (`BNET_PROBE_GAMES`),
logging one line per game and a WARN for any game other than `BNET_PROFILE_GAME` that answers
200. It needs the same Battle.net credentials the service itself does (`BNET_CLIENT_ID`,
`BNET_CLIENT_SECRET`) and is a no-op with none configured.

The API runs the same refresh itself (2026-11-07), so the job and its nightly schedule are the
safety net, runnable by hand at any time:

- **The in-process loop** (`bnetimport.RefreshLoop`): at startup and then hourly, the one
  instance holding advisory lock `79_000_002` refreshes every due character. Each pass logs
  `msg="bnetimport" op="refresh_loop"` with `considered`, `refreshed`, `skipped` and
  `rate_limited`.
- **Refresh on visit** (`bnetimport.Service.Visited`): any signed-in browser or app request
  whose account last synced more than an hour ago starts a background refresh; the request
  never waits. The visit claims `users.bnet_synced_at` with a compare-and-set, which is the
  per-account cooldown and the dedupe across instances. While the sign-in's access token is
  still valid (Battle.net tokens last 24 hours and come with no refresh token) the refresh
  re-runs the sign-in import with it, sealed at rest in `users.bnet_token_box` with
  `NOTIFY_ENCRYPTION_KEY`; once it has expired, or Battle.net refuses it, the token is
  forgotten and the account's characters are refreshed one by one from public profiles.
  Signing out forgets the token too. The token is never logged and no endpoint returns it.
  Each visit logs `op="visit_refresh"` with `path` (`account` or `public`).
- **429 backoff**: a rate limit pauses this process's automatic refreshes for 30 s, doubling
  per further 429 up to 15 minutes, and resets after a batch that finishes clean.

Every refresh stamps `users.bnet_imported_at`, which the site shows as "Updated 2 h ago from
Battle.net". A refresh that brings new gear writes a new Blizzard build, and the auto-sim on
import (`ImportSimmer`, deduped per gear) queues a run for it.

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

### Granting officer tools

    api grant --guild <id> --features officer_views --until 2026-12-12 --note "reason"

Gives one guild a time-boxed grant that unlocks only the officer views (`--plan` may be omitted
and must be `guild` if given; `--until` is required; a later grant without `--features` replaces the scope with the whole plan).

    api grandfather-officer-tools --until 2026-12-12 --note "launch grandfathering" [--apply]

Gives that same scoped grant to every claimed guild without a live plan, printing one line per guild
and a totals line; it is a dry run unless `--apply` is passed. Order: deploy the API carrying
migration 0049 first, then run this command from the deployed sha (it migrates at startup).

Rolling back after `--apply`: a scoped row with its `features` column gone reads as the whole guild
plan, which would hand every grandfathered guild's members Premium. The 0049 down migration cancels
every scoped row before dropping the column; an image rollback to a pre-0049 build needs the same
`update entitlements set status = 'canceled' where features is not null` run first.

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
printf '%s' '<runtime url>'       | gcloud secrets create DATABASE_URL --data-file=-   # see "Production database"
printf '%s' '<migrate url>'       | gcloud secrets create MIGRATE_DATABASE_URL --data-file=-
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

`MAIL_FROM`, `MAIL_REPLY_TO` and `TRUSTED_PROXY_HOPS` are deliberately left out of
`--set-env-vars` above: they take their documented defaults (`Forever Sixty
<help@foreversixty.gg>`, `help@foreversixty.gg` and `1`, matching a single Cloud Run edge in
front of the service) unless explicitly set on the service.

Every email is one branded HTML shell with a plain-text alternative (`internal/mail`: a
`Letter` drawn by `Compose`), its words from `apps/shared/copy/server.json` (`mail.*`) through
`internal/copygen`. `go test ./internal/billingmail ./internal/mail` checks both parts.

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
