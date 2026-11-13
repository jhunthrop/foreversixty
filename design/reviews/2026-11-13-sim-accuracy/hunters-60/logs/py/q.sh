#!/usr/bin/env bash
# q.sh: read-only psql against prod through the local proxy (127.0.0.1:6543). Never prints the URL.
U=$(gcloud secrets versions access latest --secret=DATABASE_URL --project foreversixty 2>/dev/null) || exit 1
U=$(python3 -c 'import sys,urllib.parse as up
u=up.urlsplit(sys.argv[1]); cred=u.netloc.rsplit("@",1)[0]
q="&".join(p for p in u.query.split("&") if p and not p.startswith(("pool_","host=")))
print(up.urlunsplit(("postgresql",cred+"@127.0.0.1:6543",u.path,q,"")))' "$U")
PGOPTIONS='-c default_transaction_read_only=on' exec psql "$U" -v ON_ERROR_STOP=1 "$@"
