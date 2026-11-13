#!/usr/bin/env bash
# fetch_own.sh <report_id>: public report files from foreversixty.gg/logs-data into data/own/<id>.
set -eu
UA='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36'
id=$1; base=https://foreversixty.gg/logs-data/reports/$id; out=data/own/$id
mkdir -p "$out"; [ -s "$out/report.json" ] || curl -sf -A "$UA" -o "$out/report.json" "$base/report.json"
for n in $(python3 -c "import json,sys;print(' '.join(str(f['index']) for f in json.load(open(sys.argv[1]))['fights']))" "$out/report.json"); do
  mkdir -p "$out/fights/$n"
  for f in summary.json events.parquet; do [ -s "$out/fights/$n/$f" ] || curl -sf -A "$UA" -o "$out/fights/$n/$f" "$base/fights/$n/$f" || echo "miss $n $f"; done
done
