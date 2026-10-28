# Log Coach runbook

How the log coach is deployed, switched on and switched off. The design is
`docs/superpowers/specs/2026-10-27-log-coach-design.md`; the experience spec is
`design/specs/2026-10-28-log-coach.md`.

## What runs where

- **The API service** (`api`) serves the coach routes when it has a bucket and a readable
  `COACH_DATA_DIR`. Without them it logs one `coach` warn line and the routes are not mounted.
  - `GET /v1/reports/{id}/fights/{n}/coach?player=` is the Coach panel. With `line=1` it is the row strip, which only reads.
  - `GET /v1/reports/{id}/coach/night?player=` is the night line.
  - `GET /v1/me/coach/recent` is the account's last-raid line and the newest reviews.
  - `GET /v1/characters/{region}/{ruleset}/{name}/coach/history?cursor=` is a character's stored reviews.
  - `GET /v1/admin/coach/samples?day=` is for moderators.
- **Every rated fight** writes `coach_profiles` rows (the top-logs pool). The writer is the rating store's hook:
  - on the service, it is set in `api/cmd/api/coach.go` `mountCoach`;
  - on the re-rating jobs (`rating-backfill`, `rerate-flagged`, `rating-rebuild-preparation`), it comes from `ratingStoreWithCoach`.
- **Cloud Run jobs.** All three run from the API image, one job per command. Each API runner is bound to one job by its name, in `SIM_JOB_PROJECT` / `SIM_JOB_REGION`.
  - `coach-review <report> <fight> <player_key> <account_id> <holder>`: queued by the fight route. It writes one review.
  - `coach-simref <engine> <build> <spec> <band> <duration>`: queued on a sim-reference miss. It sims the spec's BiS-ladder character.
  - `coach-samples [YYYY-MM-DD]`: nightly. It samples 20 of yesterday's (UTC) reviews for the moderators.

  An unreachable runner is not fatal. The service logs one `jobs` warn line and carries on:
  - with no `coach-review` job, nothing is queued and every review is the rule-only rendering;
  - with no `coach-simref` job, no spec waits on a sim, and the top logs are the only basis.
- **The image.** It carries `data/curated/coach/` at `/coach`: the spec rule files and the personas. `data/builds` is at `/data`.

### Creating the jobs (once, by hand)

The deploy workflow only moves jobs that already exist to the new image (`.github/workflows/api.yml`,
"Point the jobs at the new image"; `TestDeployMovesEveryJobInTheJobsTable` pins the list). The four
jobs this lane adds are created once, the way `api/README.md` creates `parse-report`: the API image,
the job's command as its first argument, and the service's job environment (`.env.job`: the
database, the bucket, the `LLM_*` variables and caps from the table below, with `LLM_API_KEY` as a
mounted secret, never a value). `coach-review` and `coach-simref` are executed by the API through
the Cloud Run Admin API, so the API's runtime service account needs `run.developer` on them, as it
has on `parse-report`.

```bash
IMAGE=us-east1-docker.pkg.dev/foreversixty/api/api:latest
ENV="$(tr '\n' ',' < .env.job)"

gcloud run jobs create coach-review --image "$IMAGE" --region us-east1 \
  --args coach-review --cpu 1 --memory 1Gi --task-timeout 10m \
  --set-env-vars "$ENV" --set-secrets LLM_API_KEY=LLM_API_KEY:latest

gcloud run jobs create coach-simref --image "$IMAGE" --region us-east1 \
  --args coach-simref --cpu 4 --memory 4Gi --task-timeout 15m \
  --set-env-vars "$ENV"

gcloud run jobs create coach-samples --image "$IMAGE" --region us-east1 \
  --args coach-samples --cpu 1 --memory 1Gi --task-timeout 10m \
  --set-env-vars "$ENV"
gcloud scheduler jobs create http coach-samples-nightly --location us-east1 \
  --schedule "20 3 * * *" --time-zone UTC \
  --uri "https://us-east1-run.googleapis.com/apis/run.googleapis.com/v1/namespaces/foreversixty/jobs/coach-samples:run" \
  --http-method POST --oauth-service-account-email "$SCHEDULER_SA"

gcloud run jobs create rating-rebuild-preparation --image "$IMAGE" --region us-east1 \
  --args rating-rebuild-preparation --cpu 2 --memory 2Gi --task-timeout 60m \
  --set-env-vars "$ENV"
```

`rating-rebuild-preparation` is executed by hand, once, right after the API deploy that carries a
new rating model version, in a quiet window (see "Deploy order"). Pass `-dry-run` first.

## Environment

| Variable | Where | Value |
| --- | --- | --- |
| `COACH_LLM_ENABLED` | service and jobs | `false` (the default) until the switch is flipped |
| `COACH_DATA_DIR` | service and jobs | `/coach` (the default) |
| `LLM_BASE_URL` | service and jobs | defaults to the Vercel AI Gateway |
| `LLM_MODEL` | service and jobs | the confirmed gateway slug |
| `LLM_PROVIDERS` | service and jobs | comma list pinning the gateway's provider routing |
| `LLM_PRICE_IN_PER_MTOK`, `LLM_PRICE_OUT_PER_MTOK` | service and jobs | USD per million tokens; the caps count spend with them |
| `LLM_REASONING` | service and jobs | `off` (the default), `low`, `medium` or `high`; `off` sends the gateway `reasoning: {enabled: false}`, the others `{effort: level}`. Raise it only on a coacheval result |
| `LLM_API_KEY` | service and jobs | mounted from Secret Manager, never in a file or a log |
| `COACH_DAILY_SPEND_CAP_USD` | service and jobs | `25` |
| `COACH_DAILY_REVIEWS_PER_ACCOUNT` | service and jobs | `30` |
| `COACH_MONTHLY_SPEND_CAP_PREMIUM_USD` | service and jobs | `1.00` (the default): a Premium account's model spend in a UTC calendar month |
| `COACH_MONTHLY_SPEND_CAP_GUILD_USD` | service and jobs | `2.50` (the default): an officer of a Guild-plan guild |
| `COACH_MONTHLY_SPEND_CAP_RAID_GUILD_USD` | service and jobs | `4.00` (the default): an officer of a Raid-guild-plan guild |
| `COACH_DAILY_SPEND_CAP_PER_ACCOUNT_USD` | service and jobs | `0.15` (the default): any account's model spend in a UTC day |
| `COACH_TEMPERATURE` | service and jobs | `0.3` (the default) |

While `COACH_LLM_ENABLED` is on, the config refuses to start if any of these are missing or
not positive: `LLM_MODEL`, `LLM_API_KEY`, `LLM_PROVIDERS`, either price, or any of the six caps.
While it is off, none of them is needed, and the routes answer with the rule-only review.

## Deploy order

1. Make sure migration 0054 is live before 0055 ships. golang-migrate never applies a version below the live one.
2. Migrate through the deployed API only. Deploy the API image: the service migrates at start, so `schema_migrations` reads 56 (0056 adds the per-account spend). Never run a job from a newer image before the API that carries its migration is deployed: a job migrates too (incident 2026-10-04).
3. Set the environment above on the service and on each coach job, with `COACH_LLM_ENABLED=false`.
4. Create the Cloud Run jobs `coach-review`, `coach-simref` and `coach-samples` from the deployed image:
   - shape them like `sim-run`;
   - pass the command as the first container argument;
   - add a Cloud Scheduler trigger for `coach-samples` daily at 06:00 UTC.

   Every later deploy moves these three and `rating-rebuild-preparation` to the new image (the "Point the jobs at the new image" step in `.github/workflows/api.yml`), so create them under exactly these names. A test in `api/cmd/api` fails when a job command has no line in that step.
5. Run `rating-rebuild-preparation` once from the deployed sha, in a quiet window right after the deploy. A fight ingested live while the digests are wiped can be double-counted. `-dry-run` counts first; a crashed run resumes with the `-after` cursor its last progress line logged.
6. Verify with the switch off:
   - on a real Forever report, the owner's row shows the rule-only review;
   - a stranger gets 404;
   - `GET /v1/me/coach/recent` answers.

## The kill switch

The switch is `COACH_LLM_ENABLED` on the API service and on the `coach-review` job. Flipping it is Task 42's checklist.

**On:**
1. Set it to `true` on both, and deploy a new revision.
2. Open one review as the owner.
3. Confirm one `coach_reviews` row with `source = 'model'`, its tokens, its cost and the gateway request id.
4. Confirm one `llm_spend` row for the day.

**Off:**
1. Set it to `false` on both, and deploy a new revision.
2. From the next request on, every review is the rule-only rendering, and no stored model review is served.
3. Nothing is deleted: turning it back on serves the stored model reviews again.

The caps act on their own without the switch:
- A day's spend at `COACH_DAILY_SPEND_CAP_USD`, or an account at `COACH_DAILY_REVIEWS_PER_ACCOUNT`, answers the rule-only review marked busy and queues nothing.
- Spend is counted per UTC day in `llm_spend`.
- An account past its own spend cap gets the same rule-only review, with a line naming the day or the month instead of the busy line (see "Spend caps").

## Spend caps

No subscriber may spend more of the model budget than the margin on their plan (owner,
2026-10-28). Each model call's cost (the client's reported cost, or its estimate times 1.25) is
written in one transaction to the day's `llm_spend` row and to the `coach_account_day` row
(`usd`) of the account that opened the review: an officer opening a member's review spends
their own budget. An account's month is the sum of its day rows in the UTC calendar month.

The monthly cap follows the account's plan tier, from the same entitlements the access rules
read: a verified officer or leader of a guild paying (Stripe) for the Raid guild or Guild plan
uses that plan's cap (the higher wins); a guild grant (comp, or the grandfathered
`officer_views` scope) raises nothing; a Premium holder (personal plan or a held seat) uses
Premium's; an account with no plan (a moderator, a grant-only officer) uses Premium's. Each monthly cap is about a sixth of the
plan's monthly price ($6, $15, $25); a review costs about $0.0005 at the pinned prices, so the
caps allow thousands of reviews and only stop abuse. The daily cap ($0.15) stops a burst.

A capped account's review is the rule-only rendering, stored with the reason `account_cap`
and retried once the cap lifts; the panel says `capAccountDay` or `capAccountMonth` (spec
§8.2) where the busy line goes. The teaser and the row strip are unaffected. To see who is near
a cap:

```sql
select user_id, sum(usd) as month_usd from coach_account_day
where day >= date_trunc('month', now() at time zone 'utc')::date
group by user_id order by month_usd desc limit 20;
```

## Checks

- Spend today: `select * from llm_spend order by day desc limit 7;`
- Failed validations: `select reason, count(*) from coach_reviews where generated_at > now() - interval '1 day' group by reason;`
  - `validation`: the model failed validation twice. `model_error`: the gateway failed; retried on the next view.
  - `timeout`: the generation ran past its 4-minute deadline (inside the 5-minute lease); not retried.
  - `privacy`: the pack carried a roster name or realm, so no call was made. The log line names the finding codes only. Any count above zero is a pack bug to fix.
  - After the switch flips, watch two rates: `privacy` rows, and retries caused by lowercase spell names (`name` problems in `validation` on rows with `source = 'model'`, or rows with reason `validation`). A lowercase name is a common English word that is also a class spell missing from `ordinaryProseWords`.
- Spend counts every attempt the gateway may have billed (timeouts, unreadable answers) at the worst case, so `llm_spend` can read above the gateway's invoice during an incident.
- Leases stuck past their time: `select count(*) from coach_review_leases where leased_until < now();` (taken over on the next request).
- Moderator samples: `GET /v1/admin/coach/samples?day=YYYY-MM-DD`. An anonymised character's sample has no name and no review.
