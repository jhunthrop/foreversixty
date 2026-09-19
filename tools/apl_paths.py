"""Where a curated rotation's copies live, and how one is cut out.

A rotation is edited in exactly one place, data/curated/apl/<spec>.json's
`rotation` block. Two copies are derived from it: the engine fork's
ui/<class>/apls/forever_<spec_slug>.apl.json, which the fork's own spec
tests run, and sim/request/apl/<spec>.apl.json, which both artifacts
embed.

`make apl-sync` (tools/apl_sync.py) writes the copies and `make
apl-check` (tools/apl_check.py) proves they still say what the curated
file says. Both need the same answers about paths and about which specs
have a rotation at all, so the answers live here once.
"""

import json
import pathlib

# Curated specs whose `rotation` block is a real rotation. Anything else
# is a placeholder the data lane has not written yet, and copying it
# would hand the engine an empty priority list.
WRITTEN = "written"

SITE_APL_DIR = pathlib.Path("sim/request/apl")


def written_specs(curated_dir: pathlib.Path, specs_json: pathlib.Path):
    """Yield (spec, class_slug, spec_slug, curated_path), canonical order.

    The class a spec belongs to is read from data/curated/specs.json, the
    same file sim/specs is generated from, rather than split out of the
    slug: "hunter-beast-mastery" happens to split correctly today, and
    nothing guarantees the next spec key will.
    """
    for row in json.loads(specs_json.read_text()):
        path = curated_dir / f"{row['spec']}.json"
        if not path.exists():
            continue
        if json.loads(path.read_text()).get("state") != WRITTEN:
            continue
        yield row["spec"], row["class_slug"], row["spec_slug"], path


def fork_path(engine_dir: pathlib.Path, class_slug: str, spec_slug: str) -> pathlib.Path:
    """The fork's copy. Its file names use underscores, the site's hyphens."""
    return engine_dir / "ui" / class_slug / "apls" / f"forever_{spec_slug.replace('-', '_')}.apl.json"


def fork_spec_key(path: pathlib.Path) -> str:
    """The spec key a fork copy's path names - the inverse of fork_path."""
    class_slug = path.parent.parent.name
    spec_slug = path.name[len("forever_") : -len(".apl.json")].replace("_", "-")
    return f"{class_slug}-{spec_slug}"


def site_path(repo_root: pathlib.Path, spec: str) -> pathlib.Path:
    """The copy sim/request embeds, named by the spec key it is served for."""
    return repo_root / SITE_APL_DIR / f"{spec}.apl.json"


def extract_rotation(curated_path: pathlib.Path) -> str:
    """Cut the `rotation` block out of a curated file as TEXT, dedented.

    The copies are the curated bytes, not a re-print of the parsed value.
    A JSON printer would have to reproduce the curated file's own layout -
    which keeps a short object on one line and expands a long one - and no
    two printers agree on where that line falls, so a re-print would churn
    the fork's tree on formatting alone. Taking the text keeps the copy
    byte-stable as long as the source is, and the caller checks that what
    came out parses to the same value the source holds.
    """
    lines = curated_path.read_text().split("\n")
    opens = [i for i, ln in enumerate(lines) if ln.strip().startswith('"rotation"')]
    if len(opens) != 1:
        raise ValueError(f"{curated_path}: expected one `rotation` key, found {len(opens)}")
    start = opens[0]
    indent = len(lines[start]) - len(lines[start].lstrip(" "))
    closer = " " * indent + "}"
    end = next((j for j in range(start + 1, len(lines)) if lines[j].startswith(closer)), None)
    if end is None:
        raise ValueError(f"{curated_path}: the `rotation` block is not closed at indent {indent}")

    body = [" " * indent + "{"] + lines[start + 1 : end + 1]
    out = [ln[indent:] if ln.startswith(" " * indent) else ln for ln in body]
    out[-1] = out[-1].rstrip(",")
    text = "\n".join(out) + "\n"

    want = json.loads(curated_path.read_text())["rotation"]
    if json.loads(text) != want:
        raise ValueError(f"{curated_path}: the extracted rotation text does not parse back to the rotation")
    return text
