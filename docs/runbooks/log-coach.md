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

## Environment

| Variable | Where | Value |
| --- | --- | --- |
| `COACH_LLM_ENABLED` | service and jobs | `false` (the default) until the switch is flipped |
| `COACH_DATA_DIR` | service and jobs | `/coach` (the default) |
| `LLM_BASE_URL` | service and jobs | defaults to the Vercel AI Gateway |
| `LLM_MODEL` | service and jobs | the confirmed gateway slug |
| `LLM_PROVIDERS` | service and jobs | comma list pinning the gateway's provider routing |
| `LLM_PRICE_IN_PER_MTOK`, `LLM_PRICE_OUT_PER_MTOK` | service and jobs | USD per million tokens; the caps count spend with them |
| `LLM_API_KEY` | service and jobs | mounted from Secret Manager, never in a file or a log |
| `COACH_DAILY_SPEND_CAP_USD` | service and jobs | `25` |
| `COACH_DAILY_REVIEWS_PER_ACCOUNT` | service and jobs | `30` |
| `COACH_TEMPERATURE` | service and jobs | `0.3` (the default) |

While `COACH_LLM_ENABLED` is on, the config refuses to start if any of these are missing or
not positive: `LLM_MODEL`, `LLM_API_KEY`, `LLM_PROVIDERS`, either price, or either cap.
While it is off, none of them is needed, and the routes answer with the rule-only review.

## Deploy order

1. Make sure migration 0054 is live before 0055 ships. golang-migrate never applies a version below the live one.
2. Migrate through the deployed API only. Deploy the API image: the service migrates at start, so `schema_migrations` reads 55. Never run a job from a newer image before the API that carries its migration is deployed: a job migrates too (incident 2026-10-04).
3. Set the environment above on the service and on each coach job, with `COACH_LLM_ENABLED=false`.
4. Create the Cloud Run jobs `coach-review`, `coach-simref` and `coach-samples` from the deployed image:
   - shape them like `sim-run`;
   - pass the command as the first container argument;
   - add a Cloud Scheduler trigger for `coach-samples` daily at 06:00 UTC.
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

## Checks

- Spend today: `select * from llm_spend order by day desc limit 7;`
- Failed validations: `select reason, count(*) from coach_reviews where generated_at > now() - interval '1 day' group by reason;`
- Leases stuck past their time: `select count(*) from coach_review_leases where leased_until < now();` (taken over on the next request).
- Moderator samples: `GET /v1/admin/coach/samples?day=YYYY-MM-DD`.
