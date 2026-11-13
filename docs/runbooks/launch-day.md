# Launch day

WoW Forever launches 4 Nov 2026 at 15:00 PST (23:00 UTC). This runbook covers what to watch, how
to scale, how to shed load, how to roll back, and what to do when the shared database runs hot.
The site keeps serving if the API goes down: it's static Cloudflare Worker assets. Only signed-in
and live features stop.

## The shape of the system

- **Site**: Cloudflare Worker `foreversixty` (web/wrangler.jsonc), static assets only.
- **API**: Cloud Run service `api`, project `foreversixty`, region `us-east1`. 1 vCPU and 1 GiB
  per instance, concurrency 80, always-on CPU, startup CPU boost on. Cold start is about 4 s.
  `api.foreversixty.gg` is proxied by Cloudflare, which caches the public reads that send
  `Cache-Control: public` (reports feed, rankings, specs, go links).
- **Database**: database `foreversixty` on Cloud SQL `commish-pg` (project `trycommish`). It's a
  db-g1-small (shared core, about 1.7 GiB), and it's **shared with Commish production**. Our
  role `foreversixty_app` has `CONNECTION LIMIT 30`. See api/README.md, "Production database".
- **Jobs**: 13 Cloud Run jobs use the same `DATABASE_URL:latest` secret, so each execution can
  open up to the pool size (6). On a schedule: `raid-posts` every 10 min, and nightlies between
  03:20 and 04:00 UTC. On demand: `sim-run` (premium server runs) and `parse-report` (uploads).

### The connection budget

| Consumer | Connections |
| --- | --- |
| api: pool 6 × maxScale 3 | 18 |
| Jobs and deploy overlap | 12 |
| Role limit | 30 |

When the limit is reached, new connections fail with `too many connections for role`. Pools open
connections lazily and close them after 10 idle minutes, so normal usage is far lower (2 to 6 on
4 Nov before launch).

### Baseline (24 h to 4 Nov 04:00 UTC, before launch)

- API: about 1.7k requests a day, no 5xx. `GET /v1/me` is about 65% of requests (every page view
  calls it; anonymous visitors get a 401 with no DB read). p50 latency is 2 to 30 ms. The p95/p99
  hourly worst case is about 5 s, from cold starts and Battle.net callbacks.
- API instances: 1 to 3. CPU p99 at most 20%, memory p99 at most 32%, at most 5 concurrent
  requests per instance.
- commish-pg: CPU averaged 17 to 66% over an hour, memory 48 to 86%, mostly from Commish. The
  `commish` database has 4 to 11 backends; `foreversixty` has 2 to 6.
- The `foreversixty` database is 16 MB, so it fits in memory. Postgres CPU on the shared core is
  the ceiling, not disk.

## What to watch

The alerts email jhunthrop@gmail.com (channel "Forever Sixty owner") and help@foreversixty.gg
(channel "Forever Sixty help inbox"). Each one links back here.

| Alert (project foreversixty) | Fires when |
| --- | --- |
| API 5xx over 2% | 5xx share > 2% **and** > 10 errors in 5 min |
| API p95 latency over 3 s | p95 > 3 s for 10 min |
| API at max instances | 3 active instances for 15 min |
| API memory over 85% | an instance over 85% of 1 GiB for 10 min |
| api.foreversixty.gg unreachable | uptime check on `/health` failing 3 min |
| foreversixty.gg unreachable | uptime check on `/` failing 3 min |
| commish-pg CPU over 80% | mean > 80% for 10 min |
| commish-pg memory over 92% | mean > 92% for 10 min |
| foreversixty connections over 24 of 30 | backends on `foreversixty` > 24 for 5 min |

The Cloud SQL alerts can live in `foreversixty` because `trycommish` is linked into its metrics
scope. That link is read-only and changes nothing in Commish.

Dashboards:

- Cloud Run: https://console.cloud.google.com/run/detail/us-east1/api/metrics?project=foreversixty
- Cloud SQL: https://console.cloud.google.com/sql/instances/commish-pg/overview?project=trycommish
- Alerts: https://console.cloud.google.com/monitoring/alerting?project=foreversixty

From a terminal:

```bash
P='--project foreversixty --region us-east1'
gcloud run services describe api $P --format='value(status.latestReadyRevisionName,status.traffic)'
# Errors in the last 15 minutes
gcloud logging read 'resource.type="cloud_run_revision" AND resource.labels.service_name="api" AND severity>=ERROR' \
  --project foreversixty --freshness=15m --limit 50 --format='value(timestamp,jsonPayload.op,jsonPayload.err)'
# 429s (rate limiter) and 5xx by path
gcloud logging read 'resource.type="cloud_run_revision" AND resource.labels.service_name="api" AND httpRequest.status>=429' \
  --project foreversixty --freshness=15m --limit 200 --format='value(httpRequest.status,httpRequest.requestUrl)' | sort | uniq -c | sort -rn | head
# Battle.net imports that were cut short or rate limited
gcloud logging read 'jsonPayload.op=~"battlenet_import" AND severity>=WARNING' --project foreversixty --freshness=1h --limit 20
```

Signals to watch for:

- **`rate_limited` 429s on ordinary GETs**: the router allows 300 requests per minute per IP
  (per instance). A few IPs at the cap is normal abuse protection. Many distinct IPs at the cap
  means the budget is too low.
- **`battlenet_import ... rate_limited=true`**: Blizzard answered 429 (about 36k calls an hour
  and 100 a second per client). Sign-in still works and characters are listed. Their gear and
  talents fill in on the nightly `bnet-refresh` (03:30 UTC) or the player's next sign-in.
- **`bnetimport op=needs_ruleset`**: no source named a Battle.net character's ruleset (realm
  type, profile `ruleset`, the account's own addon row, the character's own earlier row, or the
  realm's agreed player answers), so the character was imported under the ruleset `unknown`
  (`fallback=ask`) and /setup asks its player "Which realm type is ... on?". The answer moves the
  character and is stored in `realm_ruleset_answers`; two or more players agreeing on a realm with
  none disagreeing resolve it for everyone (`op=realm_answer ... resolved=true`), and a
  disagreement logs `op=realm_answers_disagree` and resolves nothing. The line lists
  `fields_checked`. A trickle is new characters whose profile Blizzard has not published yet. A
  burst means Blizzard moved the profile's `ruleset` field: read a stored
  `characters.bnet_profile`. It is counted by the
  log-based metric `bnet_needs_ruleset`, created once:

  ```bash
  gcloud logging metrics create bnet_needs_ruleset --project foreversixty \
    --description='Battle.net characters no source named a ruleset for (imported as unknown, player asked)' \
    --log-filter='resource.type="cloud_run_revision" AND jsonPayload.msg="bnetimport" AND jsonPayload.op="needs_ruleset"'
  ```
- **Rising latency with low CPU**: requests are waiting on the database (pool or commish-pg). See
  "commish-pg is hot".

## Scaling

Change one setting at a time. After each change, check that the new revision is Ready and that
`curl -s -o /dev/null -w '%{http_code}\n' https://api.foreversixty.gg/health` returns 200. Every
`gcloud run services update` makes a new revision with the same image. `api.yml` deploys swap only
the image, so these settings survive deploys.

```bash
P='--project foreversixty --region us-east1'
gcloud run services update api $P --min-instances 2     # launch setting (was 1 before 4 Nov)
gcloud run services update api $P --max-instances 3     # do not raise without the pool math below
```

- **minScale** is 2 from 4 Nov: two warm instances, so a spike or one wedged instance doesn't
  wait on a 4 s cold start. Set it back to 1 a few days after launch to save cost.
- **maxScale** stays at 3. The pool and the database limit capacity, not the instance count
  (CPU p99 is 20% at baseline, with 80 concurrent requests per instance available). To go to 4,
  first lower the pool to 5 so that 4 × 5 = 20 still fits the budget. The pool size lives in the
  `DATABASE_URL` secret and every job reads it too:

  ```bash
  # Replace pool_max_conns=6 with 5 without printing the secret, then add it as a new version.
  gcloud secrets versions access latest --secret DATABASE_URL --project foreversixty \
    | sed 's/pool_max_conns=6/pool_max_conns=5/' \
    | gcloud secrets versions add DATABASE_URL --project foreversixty --data-file=-
  gcloud run services update api $P --max-instances 4    # new revision picks up the new version
  ```

  Never raise the role's connection limit and never resize commish-pg without the owner: both
  affect Commish. The `MIGRATE_DATABASE_URL` secret must not carry `pool_*` parameters.
- **Don't deploy during the peak.** During a rollout, old and new instances both hold pools,
  which can briefly double the API's connections.

## Shedding load

Each switch is an env update, which makes a new revision in about 30 s. Use
`--update-env-vars`, never `--set-env-vars`, which would wipe the rest. Undo the change with the
opposite value.

| Switch | Effect when flipped |
| --- | --- |
| `BNET_IMPORT_CHARACTERS=false` | Sign-in still works; no characters are imported from Battle.net (characters come from the addon). Removes the heaviest per-sign-in DB and Blizzard work. |
| `COACH_LLM_ENABLED=false` | Turns off log-coach LLM reviews (cost and latency). `COACH_DAILY_SPEND_CAP_USD` already caps spend at $25 a day. |
| `APP_FEATURES` | Mobile app switches (comma list; every key defaults to off in the apps). Drop a key to turn that feature off in both apps at once. The list holds commas, so use a custom delimiter: `--update-env-vars '^@^APP_FEATURES=sign_in,character_home,...'`. |
| `APP_SIGN_IN=false` | Native-app sign-in routes answer 404. |
| `EMAIL_SIGN_IN=true` | Not a shedding switch: it's the Battle.net-outage fallback. It's off at launch, and the web build has no email button (see api/README.md). |

```bash
gcloud run services update api --project foreversixty --region us-east1 --update-env-vars BNET_IMPORT_CHARACTERS=false
```

Jobs can be paused without a deploy:

```bash
gcloud scheduler jobs pause raid-posts-every-10m --location us-east1 --project foreversixty
gcloud scheduler jobs resume raid-posts-every-10m --location us-east1 --project foreversixty
```

## Rolling back a bad deploy

**API (Cloud Run)**: move traffic back to the previous Ready revision. This is instant and makes
no new build.

```bash
P='--project foreversixty --region us-east1'
gcloud run revisions list --service api $P --limit 5
gcloud run services update-traffic api $P --to-revisions <previous-revision>=100
# When main is fixed and deployed, the deploy sends 100% to the new revision again. To return early:
gcloud run services update-traffic api $P --to-latest
```

A revision whose build added a database migration can't simply be rolled back. The older build
refuses to start against a newer `schema_migrations` version ("no migration found for version
N"). Fix forward instead, or follow api/README.md and the 2026-10-04 incident note: set
`schema_migrations` back to the deployed max, but only when the newer migration is idempotent.

**Site (Cloudflare Worker)**:

```bash
cd web
npx wrangler deployments list          # recent versions, newest first
npx wrangler rollback <version-id>     # or run with no id to roll back to the previous version
```

Then revert the bad commit on main so the next deploy doesn't bring it back.

## commish-pg is hot

commish-pg runs Commish production. **Protect Commish first.** If the database is under strain,
Forever Sixty gives up capacity before Commish does. Never restart, resize, fail over or restore
`commish-pg`, and never touch Commish's roles or databases. Those calls belong to the owner.

1. **Confirm who is using it.** In the Cloud SQL console, look at CPU and "Active connections" by
   database. From the CLI (num_backends per database):

   ```bash
   T=$(gcloud auth print-access-token); NOW=$(date -u +%FT%TZ); AGO=$(date -u -v-30M +%FT%TZ)
   curl -s -G -H "Authorization: Bearer $T" https://monitoring.googleapis.com/v3/projects/trycommish/timeSeries \
     --data-urlencode 'filter=metric.type="cloudsql.googleapis.com/database/postgresql/num_backends"' \
     --data-urlencode "interval.startTime=$AGO" --data-urlencode "interval.endTime=$NOW"
   ```

   Query Insights is off on commish-pg. Turning it on is an instance change, so leave that to the
   owner.
2. **Look at our own queries.** Through the proxy, as `foreversixty_app`
   (`cloud-sql-proxy --gcloud-auth --port 6543 trycommish:us-east1:commish-pg`, then swap the
   host in `MIGRATE_DATABASE_URL`; never print it):

   ```sql
   select pid, now() - query_start as age, state, left(query, 120)
   from pg_stat_activity where usename = 'foreversixty_app' and state <> 'idle' order by age desc;
   -- Cancel our own long runners (our role can cancel only its own backends):
   select pg_cancel_backend(pid) from pg_stat_activity
   where usename = 'foreversixty_app' and state = 'active' and now() - query_start > interval '30 seconds';
   ```

3. **Shed Forever Sixty load, in this order**, checking the commish-pg CPU after each step:
   1. Pause the scheduled jobs (`raid-posts-every-10m`, and the nightlies if they're near).
   2. `BNET_IMPORT_CHARACTERS=false`.
   3. Cap the API: `--max-instances 2`. That's 12 connections and less parallel load. Then lower
      the pool (see Scaling).
   4. As a last resort, `--max-instances 1`. The site stays up: pages are static and public
      reads are edge-cached for up to 5 min stale.
4. **Restore in reverse**, one step at a time, once CPU stays below 60% for 15 minutes.

## The launch-capacity changes (4 Nov)

- `minScale` raised from 1 to 2 (revision `api-00401-pf4`). The previous revision was
  `api-00399-996`.
- Alerts, uptime checks and the owner's email channel were created in `foreversixty`, and
  `trycommish` was linked into its metrics scope.
- Branch `land/launch-capacity` changes three things:
  - Sign-in Battle.net enrichment is capped at half the pool (3 of 6) per instance. Each
    enrichment holds a connection inside a transaction across about 5 Blizzard calls, and a burst
    would otherwise pin the whole pool.
  - An import stops at Blizzard's first 429 and isn't continued in the background.
  - The router per-IP budget goes from 120 to 300 requests a minute, for campus and carrier NATs.
