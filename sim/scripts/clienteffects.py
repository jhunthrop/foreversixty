"""Regenerate sim/cmd/leveling-bis/clienteffects_generated.go from the
client's own item-effect tables.

The ranker flags a pick whose proc or equip effect the engine does not
simulate (effect_unmodelled). It used to read that off the item's
`effect_text` alone, but `effect_text` is empty whenever the client spell
has no readable description, so a chance-on-hit weapon whose spell text is
blank (Lord General's Sword, Kindling Stave, Iceblade Hacker) was neither
ranked as a proc item nor flagged. This table is built from the effect rows
themselves instead:

    ItemXItemEffect.csv  ItemID -> ItemEffectID
    ItemEffect.csv       TriggerType (0 on use, 1 on equip, 2 chance on hit), SpellID
    SpellEffect.csv      Effect, EffectAura

An item is listed when, among the items this build's class files carry, it
has any of:

* a chance-on-hit spell (TriggerType 2);
* an on-use spell (TriggerType 0) with an effect that is not an aura, a
  taught spell or an engraved rune (a heal, damage or energize), or that
  carries a behaviour aura;
* an on-equip spell (TriggerType 1) whose effects include a behaviour aura
  no stat can hold: periodic damage (3), periodic trigger (23), proc trigger
  (42), spell-family modifiers (107, 108), class-script (112), periodic
  dummy (226).

A plain stat aura (attack power, crit, spell power) is not listed: the
pipeline already folds it into the item's stats. The dummy aura (4) is not a
behaviour aura here either, and a use spell that is only auras is not listed:
the client carries passive stat items (Blackhand's Breadth's +2% crit, Eye
of the Beast) as exactly those, so listing them would flag a stat stick as an
unsimulated proc.

Usage: python3 -I sim/scripts/clienteffects.py <build dir> [--check]
  <build dir> is data/builds/<build>; its raw/ holds the CSVs (gitignored,
  so run it from the checkout that has them) and its items/ the class files.
"""
from __future__ import annotations

import csv
import json
import sys
from collections import defaultdict
from pathlib import Path

HERE = Path(__file__).resolve().parent
OUT = HERE.parent / "cmd" / "leveling-bis" / "clienteffects_generated.go"

TRIGGER_ON_USE = "0"
TRIGGER_ON_EQUIP = "1"
TRIGGER_CHANCE_ON_HIT = "2"
EFFECT_APPLY_AURA = "6"
EFFECT_LEARN_SPELL = "36"
EFFECT_ENGRAVE = "54"
NOT_A_USE_EFFECT = frozenset({EFFECT_APPLY_AURA, EFFECT_LEARN_SPELL, EFFECT_ENGRAVE})
BEHAVIOUR_AURAS = frozenset({"3", "23", "42", "107", "108", "112", "226"})

HEADER = """package main

// clientEffectItemIDs is every equippable item this build's client gives a
// proc, use or equip-behaviour effect that no item stat can express: the
// ids rank.go's carriesEffect reads, so a pick whose effect the engine does
// not simulate is flagged effect_unmodelled even when its effect_text is
// empty. Generated from the client's ItemXItemEffect, ItemEffect and
// SpellEffect tables by sim/scripts/clienteffects.py (rules in its
// docstring); regenerate it with each build.
"""


def read_csv(raw: Path, name: str) -> list[dict[str, str]]:
    with open(raw / name, encoding="utf-8", newline="") as handle:
        return list(csv.DictReader(handle))


def equippable_item_ids(items_dir: Path) -> set[int]:
    ids: set[int] = set()
    for path in sorted(items_dir.glob("*.json")):
        ids.update(item["id"] for item in json.loads(path.read_text(encoding="utf-8"))["items"])
    return ids


def effect_carrying_ids(build: Path) -> set[int]:
    raw = build / "raw"
    item_effects = {row["ID"]: row for row in read_csv(raw, "ItemEffect.csv")}
    effects_by_spell: dict[str, list[dict[str, str]]] = defaultdict(list)
    for row in read_csv(raw, "SpellEffect.csv"):
        if row.get("DifficultyID", "0") in ("0", ""):
            effects_by_spell[row["SpellID"]].append(row)
    wanted = equippable_item_ids(build / "items")
    ids: set[int] = set()
    for link in read_csv(raw, "ItemXItemEffect.csv"):
        item_id = int(link["ItemID"])
        effect = item_effects.get(link["ItemEffectID"])
        if item_id not in wanted or effect is None:
            continue
        rows = effects_by_spell.get(effect["SpellID"], [])
        trigger = effect["TriggerType"]
        if trigger == TRIGGER_CHANCE_ON_HIT:
            ids.add(item_id)
        elif trigger == TRIGGER_ON_USE and any(
            r["Effect"] not in NOT_A_USE_EFFECT or r["EffectAura"] in BEHAVIOUR_AURAS for r in rows
        ):
            ids.add(item_id)
        elif trigger == TRIGGER_ON_EQUIP and any(r["EffectAura"] in BEHAVIOUR_AURAS for r in rows):
            ids.add(item_id)
    return ids


def render(ids: set[int], build_name: str) -> str:
    rows = sorted(ids)
    lines = ["\t" + " ".join(f"{n}: true," for n in rows[i : i + 10]) for i in range(0, len(rows), 10)]
    return (
        f"{HEADER}// Generated from build {build_name}.\n"
        f"var clientEffectItemIDs = map[int]bool{{\n" + "\n".join(lines) + "\n}\n"
    )


def main() -> int:
    args = [a for a in sys.argv[1:] if not a.startswith("--")]
    if len(args) != 1:
        print(__doc__, file=sys.stderr)
        return 2
    build = Path(args[0]).resolve()
    rendered = render(effect_carrying_ids(build), build.name)
    if "--check" in sys.argv:
        if not OUT.exists() or OUT.read_text(encoding="utf-8") != rendered:
            print("clienteffects_generated.go is stale; run sim/scripts/clienteffects.py", file=sys.stderr)
            return 1
        return 0
    OUT.write_text(rendered, encoding="utf-8")
    print(f"wrote {OUT} ({len(effect_carrying_ids(build))} ids) from {build.name}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
