"""The level a talent-taught spell is learned at: the first level its talent can be taken.

A talent tree opens one tier every five points from level 10 (tier t at 10 + 5t), and a
talent's own spells (its rank spells and the spell it teaches) are learned when the talent
is. The client's SpellLevels row says otherwise for some of them: since 1.60.1.70753 it
gives Stormstrike 17364 level 40, a tier-3 talent taken at 25 that the engine casts from
there. `spellranks` and `trainables` key learn levels off this, so the cards, the ladder
and the addon agree with the tree.
"""

import json
from collections.abc import Mapping
from pathlib import Path

#: The level the first talent point is spent at, and the points a tier asks for.
FIRST_TALENT_LEVEL = 10
POINTS_PER_TIER = 5


def tier_level(tier: int) -> int:
    """The first level a talent on `tier` (0-based) can be taken."""
    return FIRST_TALENT_LEVEL + POINTS_PER_TIER * tier


def talent_spell_levels(trees: list[dict]) -> dict[int, int]:
    """Spell id -> its talent's first level, over every talent of a class's trees."""
    levels: dict[int, int] = {}
    for tree in trees:
        for talent in tree["talents"]:
            if not isinstance(talent.get("tier"), int):
                continue
            ids = {talent.get("spell_id")} | {
                rank.get("spell_id") for rank in talent.get("ranks", [])
            }
            for spell_id in ids:
                if isinstance(spell_id, int) and spell_id > 0:
                    levels[spell_id] = tier_level(talent["tier"])
    return levels


def load_talent_spell_levels(talents_dir: Path | None, class_slug: str) -> dict[int, int]:
    """`talent_spell_levels` for `talents_dir/<class_slug>.json`, or none without one."""
    path = None if talents_dir is None else talents_dir / f"{class_slug}.json"
    if path is None or not path.exists():
        return {}
    return talent_spell_levels(json.loads(path.read_text(encoding="utf-8"))["trees"])


#: The published talent orders (`sim/cmd/addon-builds`): per spec, the leveling and raid
#: builds, each point as [tree, tier, column], 1-based, the i-th spent at level 10 + i.
ADDON_BUILDS_FILE = "addon-builds.json"


def floor_pairs(pairs: list[tuple[int, int]], levels: Mapping[int, int]) -> list[tuple[int, int]]:
    """(level, id) pairs with each id `levels` names no earlier than that level and every
    other pair no earlier than the lowest of them: a trainer rank above a talent rank waits for
    the talent. Never earlier than the pair's own learn level (spellranks' floor)."""
    floors = [levels[spell_id] for _, spell_id in pairs if spell_id in levels]
    if not floors:
        return pairs
    floor = min(floors)
    return sorted(
        ((max(level, levels.get(spell_id, floor)), spell_id) for level, spell_id in pairs),
        key=lambda pair: pair[0],
    )


def build_point_levels(order: list[list[int]], trees: list[dict]) -> dict[int, int]:
    """Spell id -> the level `order` spends its first point in the talent that teaches it."""
    by_position = {
        (tree_index + 1, talent["tier"] + 1, talent["column"] + 1): talent
        for tree_index, tree in enumerate(trees)
        for talent in tree["talents"]
        if isinstance(talent.get("tier"), int) and isinstance(talent.get("column"), int)
    }
    levels: dict[int, int] = {}
    for index, point in enumerate(order):
        talent = by_position.get(tuple(point))
        if talent is None:
            continue
        ids = {talent.get("spell_id")} | {rank.get("spell_id") for rank in talent.get("ranks", [])}
        for spell_id in ids:
            if isinstance(spell_id, int) and spell_id > 0:
                levels.setdefault(spell_id, FIRST_TALENT_LEVEL + index)
    return levels


def load_spec_point_levels(build_dir: Path, spec: str, class_slug: str) -> dict[int, int]:
    """Spell id -> the level the spec's published leveling build spends its talent's first
    point, or the raid build's for a talent the leveling build never takes. A card row for a
    talent's spell starts there, not at its tier's first level (round 22: Survival spends
    Counterattack at 32, seven levels after its tier opens at 25). Empty with
    no published build."""
    builds_path = build_dir / ADDON_BUILDS_FILE
    talents_path = build_dir / "talents" / f"{class_slug}.json"
    if not builds_path.exists() or not talents_path.exists():
        return {}
    builds = json.loads(builds_path.read_text(encoding="utf-8")).get(spec, {})
    trees = json.loads(talents_path.read_text(encoding="utf-8"))["trees"]
    levels: dict[int, int] = {}
    for kind in ("leveling", "raid"):
        if kind in builds:
            for spell_id, level in build_point_levels(builds[kind]["order"], trees).items():
                levels.setdefault(spell_id, level)
    return levels
