"""Every stat key a committed build emits must be one the site can render.

`items/<class>.json` rows carry a `stats` map whose keys are the planner's
vocabulary. The web reads them through `web/src/lib/planner/types.ts`'s
`STAT_KEYS`/`STAT_LABELS`: a key with no label renders as `undefined` and the
gear panel's sort crashes on `undefined.localeCompare`.

`test_normalize_effects.py` already asserts that everything
`pipeline.simdb.equip` can emit is a real `PROTO_STAT_ALIASES` key. That is a
weaker gate than it looks, because `PROTO_STAT_ALIASES` is much wider than the
site's table: `equip` can also produce `arcane_power`, `block_value`,
`fire_power`, `frost_power`, `holy_power`, `melee_haste`, `nature_power`,
`shadow_power`, `spell_crit`, `spell_haste` and `spell_hit`, none of which the
site has a label for. No item on any committed build carries one today, so the
first Forever re-itemisation that lands one would reach the browser unnoticed.

This is the gate that catches it: it reads what the committed builds actually
emit rather than what the tables could emit, so it fails in the data suite, at
the point the new key first ships, with a message naming the web table to
update in the same change.
"""

import json
from pathlib import Path

BUILDS_DIR = Path("builds")

#: Transcribed from `web/src/lib/planner/types.ts`'s `STAT_KEYS`, which is also
#: the key set of `STAT_LABELS` beside it (TypeScript enforces that pairing).
#: Transcribed rather than derived on purpose: the point of this test is to
#: catch the pipeline emitting a key the *web* has not been taught, so reading
#: the web's list through some shared file would defeat it. If this test fails,
#: the fix is to add the key to both `STAT_KEYS` and `STAT_LABELS` there and to
#: this list, in the same change -- not to delete the assertion.
SITE_RENDERABLE_STAT_KEYS = frozenset(
    {
        "strength",
        "agility",
        "stamina",
        "intellect",
        "spirit",
        "armor",
        "crit",
        "hit",
        "spell_power",
        "healing",
        "attack_power",
        "ranged_attack_power",
        "defense",
        "dodge",
        "parry",
        "block",
        "mp5",
        "spell_penetration",
        "fire_res",
        "frost_res",
        "nature_res",
        "shadow_res",
        "arcane_res",
    }
)


def _emitted_stat_keys(items_dir: Path) -> set[str]:
    keys: set[str] = set()
    for path in sorted(items_dir.glob("*.json")):
        payload = json.loads(path.read_text(encoding="utf-8"))
        for item in payload["items"]:
            keys.update(item["stats"])
    return keys


def _builds_with_items() -> list[Path]:
    return sorted(
        build_dir for build_dir in BUILDS_DIR.iterdir() if (build_dir / "items").is_dir()
    )


def test_every_emitted_gear_stat_key_has_a_web_label():
    """Pinned against every committed build, the active 1.60 one included."""
    build_dirs = _builds_with_items()
    assert build_dirs, "no committed build ships items/; this test would pass vacuously"
    for build_dir in build_dirs:
        unlabelled = sorted(_emitted_stat_keys(build_dir / "items") - SITE_RENDERABLE_STAT_KEYS)
        assert not unlabelled, (
            f"{build_dir}/items/ emits stat key(s) {unlabelled} that the site cannot "
            "render. Add each one to STAT_KEYS and STAT_LABELS in "
            "web/src/lib/planner/types.ts, and to SITE_RENDERABLE_STAT_KEYS here, in "
            "the same change -- otherwise the gear panel sorts on an undefined label."
        )
