"""Where a curated rotation's copies live, and how one is cut out.

A rotation is edited in exactly one place, data/curated/apl/<spec>.json's
`rotation` block. Two copies are derived from it: the engine fork's
ui/<ui dir>/apls/forever_<spec_slug>.apl.json, which the fork's own spec
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
    """Yield (spec, spec_slug, curated_path), canonical order.

    The spec slug is read from data/curated/specs.json, the same file
    sim/specs is generated from, rather than split out of the spec key:
    "hunter-beast-mastery" happens to split correctly today, and nothing
    guarantees the next spec key will.
    """
    for row in json.loads(specs_json.read_text()):
        path = curated_dir / f"{row['spec']}.json"
        if not path.exists():
            continue
        if json.loads(path.read_text()).get("state") != WRITTEN:
            continue
        yield row["spec"], row["spec_slug"], path


#: The fork's UI package each spec's rotation belongs to.
#:
#: The fork does NOT have one directory per class. Where one Go package
#: serves every spec of a class - hunter, mage, rogue, warlock - the UI is
#: one directory named for the class and three rotations land side by side
#: in it. Where the fork models a spec on its own, the directory is named
#: <spec>_<class> (balance_druid, shadow_priest, tank_warrior), and
#: "warrior" is the DPS warrior's directory with "tank_warrior" beside it.
#: So the directory cannot be derived from the spec key, and guessing it
#: from the class slug wrote rotations into ui/druid, a directory the fork
#: does not have.
#:
#: Every spec on data/curated/specs.json is listed, written or not, so the
#: day a tank's or a healer's rotation is written this table already knows
#: where it goes. A spec missing from it is an error, not a guess.
FORK_UI_DIRS = {
    "druid-balance": "balance_druid",
    "druid-feral": "feral_druid",
    "druid-restoration": "restoration_druid",
    "hunter-beast-mastery": "hunter",
    "hunter-marksmanship": "hunter",
    "hunter-survival": "hunter",
    "mage-arcane": "mage",
    "mage-fire": "mage",
    "mage-frost": "mage",
    "paladin-holy": "holy_paladin",
    "paladin-protection": "protection_paladin",
    "paladin-retribution": "retribution_paladin",
    "priest-discipline": "healing_priest",
    "priest-holy": "healing_priest",
    "priest-shadow": "shadow_priest",
    "rogue-assassination": "rogue",
    "rogue-combat": "rogue",
    "rogue-subtlety": "rogue",
    "shaman-elemental": "elemental_shaman",
    "shaman-enhancement": "enhancement_shaman",
    "shaman-restoration": "restoration_shaman",
    "warlock-affliction": "warlock",
    "warlock-demonology": "warlock",
    "warlock-destruction": "warlock",
    "warrior-arms": "warrior",
    "warrior-fury": "warrior",
    "warrior-protection": "tank_warrior",
}


def fork_path(engine_dir: pathlib.Path, spec: str, spec_slug: str) -> pathlib.Path:
    """The fork's copy. Its file names use underscores, the site's hyphens."""
    try:
        ui_dir = FORK_UI_DIRS[spec]
    except KeyError:
        raise KeyError(
            f"{spec} has no fork UI directory in apl_paths.FORK_UI_DIRS; "
            f"add the one the fork actually carries rather than guessing one"
        ) from None
    return engine_dir / "ui" / ui_dir / "apls" / f"forever_{spec_slug.replace('-', '_')}.apl.json"


def fork_spec_keys(engine_dir: pathlib.Path, specs_json: pathlib.Path) -> dict[pathlib.Path, str]:
    """Every spec's fork path, indexed by that path - the inverse of fork_path.

    Built by walking data/curated/specs.json rather than by splitting a
    file name back apart: two specs can share a UI directory, and the spec
    slug is read from the same file sim/specs is generated from for the
    same reason written_specs reads it there.
    """
    rows = json.loads(specs_json.read_text())
    return {fork_path(engine_dir, row["spec"], row["spec_slug"]): row["spec"] for row in rows}


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
