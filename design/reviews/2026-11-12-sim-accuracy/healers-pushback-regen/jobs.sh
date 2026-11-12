#!/usr/bin/env bash
# jobs.sh <jobs file> [iterations]: builds spec-breakdown once and runs each job line in turn
# (one sim at a time, nice 15, GOMAXPROCS=5), printing one line per job: the label, effective
# HPS +/- error, how long the mana lasted and the casts and mana income rows of the healer's
# spells. It syncs the embedded simdb inputs first (make simdb). A job line is
# <label> <spec-breakdown flags...>; a hash starts a comment. Run from sim/ on the box through
# tools/remote-test.sh with --dir sim and the engine pinned (the pin, or an --engine-sha):
#   bash ../design/reviews/2026-11-12-sim-accuracy/healers-pushback-regen/jobs.sh ../design/reviews/2026-11-12-sim-accuracy/healers-pushback-regen/<jobs>.txt
set -uo pipefail
jobs=$1; iters=${2:-10000}
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
rows='Flash Heal|Greater Heal|Prayer of Healing|Prayer of Mending|Renew|Penance|Power Word: Shield|Lesser Heal|Heal \(|Lightwell|Dark Sacrifice|Litany of Light|Inner Focus|Healing Touch|Regrowth|Rejuvenation|Swiftmend|Tranquility|Innervate|Nature.s Swiftness|Clearcasting|Healing Wave|Lesser Healing Wave|Riptide|Chain Heal|Mana Tide|Water Shield|Mana Spring|Mana Regen'
while read -r label rest; do
	[[ -z $label || $label == \#* ]] && continue
	# shellcheck disable=SC2086
	out=$(nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -heal -iterations "$iters" $rest 2>&1) || { echo "$label FAILED: $(echo "$out" | tail -3 | tr '\n' ' ')"; continue; }
	hps=$(echo "$out" | sed -n 's/^| effective HPS | \(.*\) |$/\1/p')
	mana=$(echo "$out" | sed -n 's/^| mana lasts (s) | \(.*\) |$/\1/p')
	casts=$(echo "$out" | grep -E "^\| ($rows)" | sed -n 's/^| \([^|]*\) | \([^|]*\) |.*/\1=\2/p' | tr -d ' ' | paste -sd, -)
	echo "$label hps=$hps mana=$mana $casts"
done < "$jobs"
