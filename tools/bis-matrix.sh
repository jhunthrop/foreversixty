#!/usr/bin/env bash
# bis-matrix.sh: bis.yml's rank matrix, one leg per class.
#
#   tools/bis-matrix.sh <repo-root> [<spec>,<spec>,...]
#
# Prints a GitHub Actions matrix (`{"include":[{"class":"druid","specs":"druid-balance,..."}]}`)
# of every spec with a written rotation (data/curated/specs.json's slugs whose
# data/curated/apl/<spec>.json has state "written", the rule sim/cmd/leveling-bis's
# writtenSpecs reads; TestBISMatrixScriptAgreesWithWrittenSpecs keeps the two equal), grouped
# by specs.json's class_slug in specs.json order. Each leg runs `leveling-bis -all -specs
# <its list>`.
#
# One leg per class because the ranking's wall clock is the slowest class: the priests (Holy
# 44 to 67 minutes, Discipline 34 to 52) are about a third of a night, so a finer split
# would not shorten it. A second argument narrows the matrix to those specs (bis.yml's
# `specs` dispatch input, for a proof run of a couple of quick specs); a name that is not
# written fails, as the ranker's own -specs does, so a typo never ranks nothing quietly.
#
# Needs jq. Exit 2 on a usage or input error, with the reason on stderr.
set -euo pipefail

main() {
	[[ $# -ge 1 && $# -le 2 ]] || { echo "usage: $0 <repo-root> [<spec>,<spec>,...]" >&2; exit 2; }
	local root=$1 only=${2:-} specs_json apl_dir
	specs_json=$root/data/curated/specs.json
	apl_dir=$root/data/curated/apl
	[[ -f $specs_json ]] || { echo "$specs_json: no such file" >&2; exit 2; }

	# One line per written spec: "<class_slug> <spec>", in specs.json order.
	local written
	written=$(written_specs "$specs_json" "$apl_dir")
	[[ -n $written ]] || { echo "no spec under $apl_dir has a written rotation" >&2; exit 2; }

	if [[ -n $only ]]; then
		written=$(narrow "$written" "$only")
	fi

	# Group by class, keeping the written order inside each leg and the order the
	# classes first appear in specs.json across legs.
	printf '%s\n' "$written" | jq -R -s -c '
		[split("\n")[] | select(length > 0) | split(" ") | {class: .[0], spec: .[1]}]
		| group_by(.class)
		| {include: map({class: .[0].class, specs: (map(.spec) | join(","))})}' |
		order_like "$written"
}

# written_specs prints "<class_slug> <spec>" for every spec whose rotation is written.
written_specs() {
	local specs_json=$1 apl_dir=$2 class spec state
	while read -r class spec; do
		[[ -f $apl_dir/$spec.json ]] || continue
		state=$(jq -r '.state // ""' "$apl_dir/$spec.json")
		[[ $state == written ]] && printf '%s %s\n' "$class" "$spec"
	done < <(jq -r '.[] | "\(.class_slug) \(.spec)"' "$specs_json")
	return 0
}

# narrow keeps the written lines whose spec the comma list names; an unknown name fails.
narrow() {
	local written=$1 only=$2 name kept=""
	IFS=',' read -r -a names <<<"$only"
	for name in "${names[@]}"; do
		name=${name// /}
		[[ -n $name ]] || continue
		grep -q " $name\$" <<<"$written" ||
			{ echo "$name is not a spec with a written rotation" >&2; exit 2; }
	done
	while read -r line; do
		for name in "${names[@]}"; do
			[[ ${line#* } == "${name// /}" ]] && { kept+="$line"$'\n'; break; }
		done
	done <<<"$written"
	printf '%s' "$kept"
}

# order_like re-sorts the legs to the order their classes first appear in the written list
# (jq's group_by sorts by class name; the specs.json order is the one people read).
order_like() {
	local written=$1 classes
	classes=$(printf '%s\n' "$written" | awk '!seen[$1]++ {print $1}' | jq -R -s -c 'split("\n") | map(select(length > 0))')
	jq -c --argjson order "$classes" '.include |= sort_by(. as $leg | $order | index($leg.class))'
}

main "$@"
