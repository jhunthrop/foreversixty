#!/usr/bin/env bash
# db.sh "<sql>": run read-only SQL on prod through the local proxy; never prints the URL.
set -euo pipefail
url=$(gcloud secrets versions access latest --secret=DATABASE_URL --project=foreversixty)
url=$(python3 -c '
import sys,urllib.parse as u
p=u.urlsplit(sys.argv[1]); q=[(k,v) for k,v in u.parse_qsl(p.query) if not k.startswith("pool_") and k!="host"]
net=p.netloc.rsplit("@",1); auth=net[0]
print(u.urlunsplit((p.scheme, auth+"@127.0.0.1:6543", p.path, u.urlencode(q), "")))' "$url")
PGOPTIONS='-c default_transaction_read_only=on' psql "$url" -X -A -F $'\t' -c "$1"
