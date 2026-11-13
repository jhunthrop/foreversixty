#!/usr/bin/env bash
# r2env.sh <cmd...>: run cmd with R2 creds in env, never printed.
set -euo pipefail
eval "$(gcloud run services describe api --project=foreversixty --region=us-east1 --format=json | python3 -c '
import json,sys,shlex
d=json.load(sys.stdin)
for c in d["spec"]["template"]["spec"]["containers"]:
  for e in c.get("env",[]):
    if e["name"] in ("R2_ACCOUNT_ID","R2_BUCKET") and "value" in e: print("export %s=%s" % (e["name"], shlex.quote(e["value"])))
')"
export R2_ACCESS_KEY_ID=$(gcloud secrets versions access latest --secret=R2_ACCESS_KEY_ID --project=foreversixty)
export R2_SECRET_ACCESS_KEY=$(gcloud secrets versions access latest --secret=R2_SECRET_ACCESS_KEY --project=foreversixty)
export R2_BUCKET=${R2_BUCKET:-foreversixty-logs}
exec "$@"
