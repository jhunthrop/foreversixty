"""Copy every written curated rotation into the two places that run it.

`make engine-pin` writes the engine version and NOTHING else - it does
not carry rotations across, which was believed for a while and is why
two copies drifted. This is the target that carries them: for each
data/curated/apl/<spec>.json with state "written", its `rotation` block
is written to the engine fork's ui/<ui dir>/apls/forever_<spec>.apl.json
and to sim/request/apl/<spec>.apl.json, which both artifacts embed.

The copies are the curated bytes, dedented one level (see
apl_paths.extract_rotation), so running this against an unchanged
curated file leaves both trees untouched.

A fork copy that already runs the same rotation (it differs from the
curated one in display fields only, apl_paths.DISPLAY_KEYS) is left as
it is: notes are the curated file's alone, and rewording one must never
dirty the fork or need an engine pin. The site's own copy is always
refreshed, since it lives in this repository.

Run through `make apl-sync`; `make apl-check` then proves it.
"""

import json
import os
import pathlib
import sys

import apl_paths


def write_if_changed(path: pathlib.Path, text: str) -> bool:
    """Write only a real change, so an unchanged sync dirties no tree."""
    if path.exists() and path.read_text() == text:
        return False
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(text)
    return True


def same_rotation(path: pathlib.Path, text: str) -> bool:
    """Whether the copy at `path` already runs the rotation `text` holds."""
    if not path.exists():
        return False
    try:
        copy = json.loads(path.read_text())
    except json.JSONDecodeError:
        return False
    return apl_paths.semantics(copy) == apl_paths.semantics(json.loads(text))


def main() -> int:
    engine_dir = pathlib.Path(os.environ["ENGINE_DIR"])
    curated_dir = pathlib.Path(os.environ["CURATED_APL_DIR"])
    specs_json = pathlib.Path(os.environ["CURATED_SPECS_JSON"])
    repo_root = pathlib.Path.cwd()

    written = list(apl_paths.written_specs(curated_dir, specs_json))
    if not written:
        print(f"apl-sync: no curated rotation in {curated_dir} is marked written", file=sys.stderr)
        return 1

    for spec, spec_slug, curated_path in written:
        text = apl_paths.extract_rotation(curated_path)
        for path in (
            apl_paths.fork_path(engine_dir, spec, spec_slug),
            apl_paths.site_path(repo_root, spec),
        ):
            # The fork's ui/<class> tree has to exist already: creating
            # one would mean inventing a class directory in a repository
            # this target is otherwise only allowed to copy into.
            if path.is_relative_to(engine_dir) and not path.parent.is_dir():
                print(f"apl-sync: {engine_dir} has no {path.parent}", file=sys.stderr)
                return 1
            if path.is_relative_to(engine_dir) and same_rotation(path, text):
                print(f"apl-sync: unchanged {path} (notes stay curated-only)")
                continue
            verb = "wrote" if write_if_changed(path, text) else "unchanged"
            print(f"apl-sync: {verb} {path}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
