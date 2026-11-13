#!/bin/sh
# db.sh "<sql>": read-only query through the Cloud SQL proxy on 127.0.0.1:6543. The URL stays in a variable.
U=$(gcloud secrets versions access latest --secret=DATABASE_URL --project=foreversixty 2>/dev/null) || { echo "no secret" >&2; exit 1; }
U=$(printf '%s' "$U" | python3 -c 'import sys,urllib.parse as p;u=p.urlsplit(sys.stdin.read().strip());q=[(k,v) for k,v in p.parse_qsl(u.query) if not k.startswith("pool_") and k!="host"];n=u.netloc.rsplit("@",1);print(p.urlunsplit((u.scheme,n[0]+"@127.0.0.1:6543",u.path,p.urlencode(q),"")))')
PGOPTIONS='-c default_transaction_read_only=on' psql "$U" -X -A -F '|' -c "$1"
