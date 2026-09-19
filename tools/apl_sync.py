"""Copy every written curated rotation into the two places that run it.

`make engine-pin` writes the engine version and NOTHING else - it does
not carry rotations across, which was believed for a while and is why
two copies drifted. This is the target that carries them: for each
data/curated/apl/<spec>.json with state "written", its `rotation` block
is written to the engine fork's ui/<class>/apls/forever_<spec>.apl.json
and to sim/request/apl/<spec>.apl.json, which both artifacts embed.

The copies are the curated bytes, dedented one level (see
apl_paths.extract_rotation), so running this against an unchanged
curated file leaves both trees untouched.

Run through `make apl-sync`; `make apl-check` then proves it.
"""

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


def main() -> int:
    engine_dir = pathlib.Path(os.environ["ENGINE_DIR"])
    curated_dir = pathlib.Path(os.environ["CURATED_APL_DIR"])
    specs_json = pathlib.Path(os.environ["CURATED_SPECS_JSON"])
    repo_root = pathlib.Path.cwd()

    written = list(apl_paths.written_specs(curated_dir, specs_json))
    if not written:
        print(f"apl-sync: no curated rotation in {curated_dir} is marked written", file=sys.stderr)
        return 1

    for spec, class_slug, spec_slug, curated_path in written:
        text = apl_paths.extract_rotation(curated_path)
        for path in (
            apl_paths.fork_path(engine_dir, class_slug, spec_slug),
            apl_paths.site_path(repo_root, spec),
        ):
            # The fork's ui/<class> tree has to exist already: creating
            # one would mean inventing a class directory in a repository
            # this target is otherwise only allowed to copy into.
            if path.is_relative_to(engine_dir) and not path.parent.is_dir():
                print(f"apl-sync: {engine_dir} has no {path.parent}", file=sys.stderr)
                return 1
            verb = "wrote" if write_if_changed(path, text) else "unchanged"
            print(f"apl-sync: {verb} {path}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
