# Baselines at the pin and the talent searches (sim-accuracy 2026-11-12, melee-hunters lane).
# Run from sim/ on the box: tools/remote-test.sh --dir sim HEAD -- bash ../design/reviews/2026-11-12-sim-accuracy/melee-hunters/job_a_base.sh
set -u
R=../design/reviews/2026-11-12-sim-accuracy/melee-hunters
export GOMAXPROCS=5
make -C .. simdb >/dev/null || exit 1
go build -o /tmp/sb.$$ ./cmd/spec-breakdown && go build -o /tmp/ts.$$ ./cmd/talent-search || exit 1
sb() { echo "### sb $*"; nice -n 15 /tmp/sb.$$ -repo-root .. -iterations 20000 -seed 1 "$@" 2>&1 | grep -E "DPS|out of mana" | head -3; }
for s in warrior-arms warrior-fury hunter-beast-mastery hunter-marksmanship hunter-survival; do
  for f in horde alliance; do sb -spec $s -preset raid -band 60 -faction $f; sb -spec $s -preset bare -band 60 -faction $f; done
done
keeps() { python3 -c "import json,sys;print(','.join(json.load(open('$R/../../talent-search/raid-keeps.json'))['keep'].get(sys.argv[1],[])))" "$1"; }
mkdir -p /tmp/tsout.$$
for s in warrior-arms warrior-fury hunter-beast-mastery hunter-marksmanship hunter-survival; do
  echo "### ts $s raid"; nice -n 15 /tmp/ts.$$ -repo-root .. -spec $s -preset raid -faction horde -keep "$(keeps $s)" -out /tmp/tsout.$$/raid 2>&1 | tail -2
  echo "### ts $s bare"; nice -n 15 /tmp/ts.$$ -repo-root .. -spec $s -preset bare -faction horde -out /tmp/tsout.$$/bare 2>&1 | tail -2
done
for f in /tmp/tsout.$$/*/*.md; do echo "#### $f"; sed -n 1,40p "$f"; done
rm -rf /tmp/sb.$$ /tmp/ts.$$ /tmp/tsout.$$
