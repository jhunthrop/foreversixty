"""Hold both derived copies of a rotation to the curated one.

A rotation is edited in exactly one place, data/curated/apl/<spec>.json's
`rotation` block. Two copies are derived from it by `make apl-sync`: the
engine fork's ui/<class>/apls/forever_<spec>.apl.json, which the fork's
own spec tests run, and sim/request/apl/<spec>.apl.json, which both
artifacts embed. A copy that drifts is a sim measuring a rotation nobody
wrote down.

Every written curated spec is checked here - a missing copy fails, so
forgetting `make apl-sync` cannot pass - and so is every forever_*.apl.json
the fork carries, so a copy whose curated source was renamed away is
caught too.

The comparison is of PARSED json, not bytes: what the copies must agree
on is the rotation, and a difference in content is what a drifting copy
means.

Run through `make apl-check`.
"""

import json
import os
import pathlib
import sys

import apl_paths


def compare(label: str, path: pathlib.Path, curated_path: pathlib.Path, curated) -> str | None:
    """Return a failure line, or None when the copy matches."""
    if not path.exists():
        return f"{path}: no {label} copy of {curated_path}; run `make apl-sync`"
    try:
        copy = json.loads(path.read_text())
    except json.JSONDecodeError as err:
        return f"{path}: not valid json ({err})"
    if copy != curated:
        return (
            f"{path} differs from {curated_path}'s rotation; "
            f"edit the curated file and run `make apl-sync`, never the other way"
        )
    return None


def main() -> int:
    engine_dir = pathlib.Path(os.environ["ENGINE_DIR"])
    curated_dir = pathlib.Path(os.environ["CURATED_APL_DIR"])
    specs_json = pathlib.Path(os.environ["CURATED_SPECS_JSON"])
    repo_root = pathlib.Path.cwd()

    written = list(apl_paths.written_specs(curated_dir, specs_json))
    if not written:
        print(f"apl-check: no curated rotation in {curated_dir} is marked written", file=sys.stderr)
        return 1

    failures = []
    checked_fork = set()
    for spec, class_slug, spec_slug, curated_path in written:
        curated = json.loads(curated_path.read_text())["rotation"]
        fork = apl_paths.fork_path(engine_dir, class_slug, spec_slug)
        checked_fork.add(fork)
        for label, path in (("fork", fork), ("site", apl_paths.site_path(repo_root, spec))):
            failure = compare(label, path, curated_path, curated)
            if failure:
                failures.append(failure)
            else:
                print(f"apl-check: {spec} {label} copy matches {curated_path}")

    # A fork copy with no written curated source is the other direction of
    # the same drift: the rotation the fork runs is one nobody maintains.
    for path in sorted(engine_dir.glob("ui/*/apls/forever_*.apl.json")):
        if path not in checked_fork:
            failures.append(
                f"{path}: no curated source marked written at "
                f"{curated_dir / (apl_paths.fork_spec_key(path) + '.json')}"
            )

    for failure in failures:
        print(f"apl-check: {failure}", file=sys.stderr)
    return 1 if failures else 0


if __name__ == "__main__":
    sys.exit(main())
