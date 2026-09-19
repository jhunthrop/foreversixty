"""Hold the engine fork's rotations to the curated ones.

A rotation is edited in exactly one place, data/curated/apl/<spec>.json's
`rotation` block. Two copies are derived from it: the engine fork's
ui/<class>/apls/forever_<spec>.apl.json, which the fork's own spec tests
run, and sim/request/apl/<spec>.apl.json, which both artifacts embed. A
copy that drifts is a sim measuring a rotation nobody wrote down.

Only the fork's copies are checked here, because the site's are still
the engine's Era presets by ruling: bringing them into line moves both
adapter fixtures, so it happens in the same pass as the regeneration.
That exception is named in EXPECT_DRIFT so it is a listed debt rather
than a silence.

Run through `make apl-check`.
"""

import json
import os
import pathlib
import sys

# Site-side copies knowingly not yet regenerated from the curated
# source. Remove a spec from here in the commit that regenerates it.
EXPECT_DRIFT = {"sim/request/apl"}


def fork_copies(engine_dir: pathlib.Path):
    """Yield (spec, path) for every forever_*.apl.json the fork carries."""
    for path in sorted(engine_dir.glob("ui/*/apls/forever_*.apl.json")):
        klass = path.parent.parent.name
        spec = path.name[len("forever_") : -len(".apl.json")].replace("_", "-")
        yield f"{klass}-{spec}", path


def main() -> int:
    engine_dir = pathlib.Path(os.environ["ENGINE_DIR"])
    curated_dir = pathlib.Path(os.environ["CURATED_APL_DIR"])

    copies = list(fork_copies(engine_dir))
    if not copies:
        print(f"apl-check: {engine_dir} carries no ui/*/apls/forever_*.apl.json", file=sys.stderr)
        return 1

    failures = []
    for spec, path in copies:
        curated_path = curated_dir / f"{spec}.json"
        if not curated_path.exists():
            failures.append(f"{path}: no curated source at {curated_path}")
            continue
        curated = json.loads(curated_path.read_text())["rotation"]
        copy = json.loads(path.read_text())
        if curated != copy:
            failures.append(
                f"{path} differs from {curated_path}'s rotation; "
                f"edit the curated file and copy it over, never the other way"
            )
        else:
            print(f"apl-check: {spec} matches {curated_path}")

    for note in sorted(EXPECT_DRIFT):
        print(f"apl-check: {note} is a known-stale copy, regenerated with the adapter fixtures")

    for failure in failures:
        print(f"apl-check: {failure}", file=sys.stderr)
    return 1 if failures else 0


if __name__ == "__main__":
    sys.exit(main())
