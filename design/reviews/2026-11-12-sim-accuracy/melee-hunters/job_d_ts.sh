# Talent searches at 60 for the five specs, full reports to stdout (run from sim/ on the box).
set -u
export GOMAXPROCS=5
make -C .. simdb >/dev/null || exit 1
go build -o /tmp/ts.$$ ./cmd/talent-search || exit 1
keeps() { python3 -c "import json,sys;print(','.join(json.load(open('../design/reviews/talent-search/raid-keeps.json'))['keep'].get(sys.argv[1],[])))" "$1"; }
out=/tmp/tsout.$$; mkdir -p $out
for s in warrior-arms warrior-fury hunter-beast-mastery hunter-marksmanship hunter-survival; do
  for f in horde alliance; do
    nice -n 15 /tmp/ts.$$ -repo-root .. -spec $s -preset raid -faction $f -keep "$(keeps $s)" -out $out/raid-$f >/dev/null 2>&1
    nice -n 15 /tmp/ts.$$ -repo-root .. -spec $s -preset bare -faction $f -out $out/bare-$f >/dev/null 2>&1
  done
done
for f in $out/*/*.md; do echo "########## $f"; cat "$f"; done
rm -rf /tmp/ts.$$ $out
