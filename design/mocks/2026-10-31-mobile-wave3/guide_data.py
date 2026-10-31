"""The Elemental Shaman guide as the site ships it: frontmatter and sections from web/src/content/guides/shaman/elemental.md,
the rotation from data/curated/apl/shaman-elemental.json, talent and spell words from the build's own files. Nothing here is typed by hand
except the order the screens show the rows in."""
import json
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[3]
GUIDE = ROOT / "web/src/content/guides/shaman/elemental.md"
APL = ROOT / "data/curated/apl/shaman-elemental.json"
BUILD = ROOT / "data/builds/1.60.1.70658"


def _sections() -> tuple[dict, dict[str, str]]:
    text = GUIDE.read_text()
    _, head, body = text.split("---", 2)
    fm = {}
    for key in ("statPriority", "recommendedRaces", "updated"):
        m = re.search(rf"^{key}: (.+)$", head, re.M)
        fm[key] = m.group(1).strip() if m else ""
    fm["statPriority"] = [s.strip() for s in fm["statPriority"].strip("[]").split(",")]
    parts = re.split(r"^(#{2,3}) (.+)$", body, flags=re.M)
    secs = {parts[i + 2]: parts[i + 3].strip() for i in range(0, len(parts) - 3, 3)}
    return fm, secs


FM, SECS = _sections()


def bullets(name: str) -> list[str]:
    return [ln[2:].strip() for ln in SECS[name].splitlines() if ln.startswith("- ")]


def overview() -> str:
    return SECS["Overview"].splitlines()[0].strip()


def key_talents() -> list[str]:
    """The bold names of the guide's talent paragraph, in its order, plus every talent of the raid build that the paragraph skips and the spec casts through."""
    return re.findall(r"\*\*(.+?)\*\*", SECS["Talents and builds"]) + ["Elemental Mastery"]


def talent_words() -> dict[str, dict]:
    """Client words per talent: the top rank's description."""
    tal = json.loads((BUILD / "talents/shaman.json").read_text())
    out = {}
    def walk(o):
        if isinstance(o, dict):
            if "ranks" in o and "name" in o and isinstance(o["ranks"], list) and o["ranks"] and isinstance(o["ranks"][-1], dict):
                out[o["name"]] = {"icon": o["icon"], "max_rank": o["max_rank"], "text": o["ranks"][-1]["description"]}
            for v in o.values():
                walk(v)
        elif isinstance(o, list):
            for v in o:
                walk(v)
    walk(tal)
    return out


def rotation() -> list[tuple[str, str]]:
    """(spell name, the file's own note) in priority order, from the APL, for every castSpell line with a note; Flame Shock is the guide's explicit non-line."""
    spells = {s["id"]: s["name"] for s in json.loads((BUILD / "spells.json").read_text())}
    out = []
    for line in json.loads(APL.read_text())["rotation"]["priorityList"]:
        cast = line.get("action", {}).get("castSpell")
        if cast and line.get("notes"):
            out.append((spells.get(cast["spellId"]["spellId"], "?"), line["notes"]))
    return out
