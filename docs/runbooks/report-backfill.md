# Backfilling a live report's fights

`api backfill-report` parses a live report again from the raw chunks its companion uploaded
and writes the fights through the live path's own fight writer (`parse.Backfill`, the same
`writeFight`/`writeReport` the whole-file parse uses). It exists because from 10 Nov 2026 the
API answered 400 "the bundle has no summary part" to every companion fight bundle
(`reports.partBytes` read form fields with `FormFile`; fixed at `0ed1abe86`), so that week's
live reports completed with no fights while their raw chunks landed intact.

What it does, per report:

- streams the stored raw chunks from R2 in offset order, decoded, as one log; stops at the
  first chunk that does not continue the one before it and reports `gap_at`;
- runs the engine over the stream with the layout the report's health record names, from
  the report's own creation time as the clock seed;
- **dry run (default)**: counts fights, encounters, kills, lines and parse errors; writes
  nothing, not even the status;
- **`--apply`**: writes every fight (files, `fights` row, rankings), sets the status
  `complete` with the engine's health, writes the final `report.json`, and runs the
  completion hook (log zone, guild attribution, raid match, site feed line);
- never announces: no notifier is wired and the guilds hook has no event router, so no
  Discord post, push or in-app notification follows a backfill;
- refuses a report that already holds fights, a report with no raw chunks, and a run without
  `--report`. One report per run; there is no sweep.

## Safety rules

1. **Only the deployed API migrates.** The command connects without running migrations, and
   it is still run from a worktree checked out at the deployed image's sha, so the binary
   reads the schema production has (incident 2026-10-04, migration 0029).
2. **Through the Cloud SQL proxy**, as `foreversixty_app`, with every `pool_*` query
   parameter dropped from `DATABASE_URL` (one connection needs none of them).
3. **Never print a secret.** The URL and the R2 keys go from Secret Manager straight into the
   environment of the one command.
4. **Dry run first, always**, and read the counts against the report's line count before
   `--apply`. Stop if the dry run shows zero fights on a large log, parse errors above a few
   percent, or a `gap_at`.
5. The owner approves each `--apply` (production write), the same as a grant.

## The commands

```bash
# 1. The deployed sha is the api service's image tag; the worktree sits at it.
gcloud run services describe api --region us-east1 --project foreversixty \
  --format 'value(spec.template.spec.containers[0].image)'          # ...:<sha>
git -C ~/code/forever worktree add /tmp/forever-deployed <sha>

# 2. The proxy, in its own terminal, if it is not already up on 6543.
cloud-sql-proxy --gcloud-auth --port 6543 trycommish:us-east1:commish-pg

# 3. The data directories the image carries, laid out the same way (api/Dockerfile).
D=/tmp/forever-deployed-data && mkdir -p $D/trees $D/curated/raids
ln -s /tmp/forever-deployed/data/builds/* $D/trees/
cp /tmp/forever-deployed/web/src/data/active-build.json $D/trees/
cp /tmp/forever-deployed/data/curated/instance-levels.json $D/curated/
cp /tmp/forever-deployed/data/curated/raids/first-tier.json $D/curated/raids/

# 4. One report, dry run. The URL's host becomes the proxy and its pool_* params go.
cd /tmp/forever-deployed/api
secret() { gcloud secrets versions access latest --secret "$1" --project foreversixty; }
# R2_ACCOUNT_ID and R2_BUCKET are plain values on the service, not secrets.
svcenv() { gcloud run services describe api --region us-east1 --project foreversixty --format json |
  python3 -c "import json,sys; print(next(e['value'] for e in json.load(sys.stdin)['spec']['template']['spec']['containers'][0]['env'] if e['name']=='$1'))"; }
env -i PATH="$PATH" HOME="$HOME" \
  DATABASE_URL="$(secret DATABASE_URL | sed -E 's#@[^/@]+/#@127.0.0.1:6543/#; s/[?&]pool_[a-z_]+=[^&]*//g')" \
  RESEND_API_KEY=unused PUBLIC_BASE_URL=https://foreversixty.gg API_BASE_URL=https://api.foreversixty.gg \
  R2_ACCOUNT_ID="$(svcenv R2_ACCOUNT_ID)" R2_BUCKET="$(svcenv R2_BUCKET)" \
  R2_ACCESS_KEY_ID="$(secret R2_ACCESS_KEY_ID)" R2_SECRET_ACCESS_KEY="$(secret R2_SECRET_ACCESS_KEY)" \
  TREE_DATA_DIR=$D/trees CURATED_DIR=$D/curated \
  go run ./cmd/api backfill-report --report <id>

# 5. The same command with --apply, once the dry run's counts read right and the owner said so.
  ... go run ./cmd/api backfill-report --report <id> --apply
```

The `sed` keeps the user, password and database and swaps only the host and port. Check the
result's shape by piping it through `sed -E 's#//[^@]+@#//…@#'`, never with a bare `echo`. If
`DATABASE_URL` in Secret Manager carries the host in another form, fix the expression, not the
secret. The secret and env names above are the api service's as of 2026-11-12; confirm them
with `gcloud run services describe api --region us-east1 --project foreversixty` before a run.

## Reading the dry run

The one `backfill-report` line carries `chunks`, `bytes`, `fights`, `encounters`, `kills`,
`lines`, `parse_errors`, `layout`, `missing_header` and, when the stream broke, `gap_at`.

- `fights` of zero on a log of thousands of lines means the engine saw no combat worth a
  fight, or read it with the wrong layout: look at `layout` and `parse_errors` before anything
  else.
- `parse_errors` of more than a few percent of `lines` is a layout problem (a report whose
  companion started mid-file, so no `COMBAT_LOG_VERSION` header reached the engine, infers its
  layout; companion 0.2.45 reads the header back out of the file for that case). Backfilling
  it writes what the inferred layout can read, no better than the live parse would have.
- `gap_at` means a raw chunk never arrived. The fights before the gap are whole and are
  written; the night after it is not recoverable from the bucket.

## Verifying afterwards, read-only

```sql
select r.id, r.status, r.engine_version, r.completed_at,
       (select count(*) from fights f where f.report_id = r.id) as fights
from reports r where r.id in ('<id>', ...);
```

Then open the report page: the fights list fills from `report.json`, which `--apply` rewrote.
No notification row, Discord post or push should exist for the report's completion.

## The reports of 10 to 12 Nov 2026

The five zero-fight live reports the bug left are the ids beginning `3sdrcw`, `edcgw6`,
`eygg26`, `ctxicu` and `4q5el2`. Dry-run each, then apply one at a time.
