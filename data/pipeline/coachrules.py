"""The 28 spec rule files the log coach reads, data/curated/coach/<spec>.json.

Each file lists, per spec, the cooldowns, auras, downranks and defensives the coach judges.
The choice of abilities is curated in SPECS below; every spell id and every `verified`
stamp is resolved against a build's trainables, spellconst and talent tables (and the
spec's own APL ids), so nothing here invents an id and a new client build restamps the
files with `python -m pipeline coach-rules --build <build>`.
`api/internal/coach/specrules_data_test.go` pins the result against the active build.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.spellranks import load_book_ids

MAX_LEVEL = 60  # a player's trainable ranks end here

# aura kinds
DOT, DEBUFF, SELF, DEF = "dot", "debuff", "self_buff", "defensive"

# spec -> lists. None = omitted (Couldn"t check); [] = honest empty.
SPECS = {
    # ---- DPS (20)
    "druid-balance": dict(
        core=[], auras=[("Insect Swarm", DOT)], downranks=[], defensives=["Barkskin"]
    ),
    "druid-feral": dict(
        core=["Shifting Power"],
        auras=[("Rip", DOT)],
        downranks=[],
        defensives=["Barkskin"],
    ),
    "hunter-beast-mastery": dict(
        core=["Bestial Wrath", "Rapid Fire", "Aimed Shot", "Multi-Shot"],
        auras=[("Serpent Sting", DOT), ("Aspect of the Hawk", SELF)],
        downranks=[],
        defensives=["Feign Death", "Deterrence"],
    ),
    "hunter-marksmanship": dict(
        core=["Rapid Fire", "Aimed Shot", "Multi-Shot"],
        auras=[("Aspect of the Hawk", SELF)],
        downranks=[],
        defensives=["Feign Death", "Deterrence"],
    ),
    "hunter-survival": dict(
        core=["Raptor Strike", "Strider Kick", "Immolation Trap"],
        auras=[("Lacerate", DOT)],
        downranks=[],
        defensives=["Feign Death", "Deterrence"],
    ),
    "mage-arcane": dict(
        core=["Arcane Power", "Presence of Mind"],
        auras=[],
        downranks=[],
        defensives=["Ice Block", "Ice Barrier"],
    ),
    "mage-fire": dict(
        core=["Combustion", "Fire Blast"],
        auras=None,
        downranks=[],
        defensives=["Ice Block", "Ice Barrier"],
    ),
    "mage-frost": dict(core=[], auras=[], downranks=[], defensives=["Ice Block", "Ice Barrier"]),
    "paladin-retribution": dict(
        core=["Judgement", "Holy Strike"],
        auras=[("Seal of Righteousness", SELF)],
        downranks=[],
        defensives=["Divine Shield", "Divine Protection", "Lay on Hands"],
    ),
    "priest-shadow": dict(
        core=["Shadow Word: Death", "Devouring Plague", "Inner Focus"],
        auras=[("Shadow Word: Pain", DOT), ("Shadowform", SELF)],
        downranks=[],
        defensives=["Desperate Prayer"],
    ),
    "rogue-assassination": dict(
        core=["Cold Blood"],
        auras=[("Slice and Dice", SELF)],
        downranks=[],
        defensives=["Evasion", "Vanish"],
    ),
    "rogue-combat": dict(
        core=["Adrenaline Rush", "Blade Flurry"],
        auras=[("Slice and Dice", SELF)],
        downranks=[],
        defensives=["Evasion", "Vanish"],
    ),
    "rogue-subtlety": dict(
        core=["Ghostly Strike"],
        auras=[("Rupture", DOT), ("Slice and Dice", SELF)],
        downranks=[],
        defensives=["Evasion", "Vanish"],
    ),
    "shaman-elemental": dict(core=["Lava Burst"], auras=[], downranks=[], defensives=[]),
    "shaman-enhancement": dict(core=["Stormstrike"], auras=[], downranks=[], defensives=[]),
    "warlock-affliction": dict(
        core=["Bane of Doom"],
        auras=[("Corruption", DOT), ("Bane of Doom", DOT), ("Immolate", DOT), ("Wrack", DOT)],
        downranks=[],
        defensives=["Death Coil"],
    ),
    "warlock-demonology": dict(
        core=["Soul Fire", "Bane of Doom"],
        auras=[("Corruption", DOT), ("Bane of Doom", DOT), ("Immolate", DOT)],
        downranks=[],
        defensives=["Death Coil"],
    ),
    "warlock-destruction": dict(
        core=["Conflagrate", "Shadowburn", "Bane of Doom"],
        auras=[("Corruption", DOT), ("Bane of Doom", DOT), ("Immolate", DOT)],
        downranks=[],
        defensives=["Death Coil"],
    ),
    "warrior-arms": dict(
        core=["Bloodrage", "Mortal Strike"],
        auras=[("Battle Shout", SELF), ("Rend", DOT)],
        downranks=[],
        defensives=["Last Stand"],
    ),
    # Fury's defensives are an honest empty list: Last Stand is a protection talent.
    "warrior-fury": dict(
        core=["Bloodrage", "Death Wish", "Bloodthirst", "Whirlwind"],
        auras=[("Battle Shout", SELF)],
        downranks=[],
        defensives=[],
    ),
    # ---- tanks (3)
    "druid-feral-bear": dict(
        core=["Primal Bite"],
        auras=[("Demoralizing Roar", DEF)],
        downranks=[],
        defensives=["Barkskin", "Frenzied Regeneration"],
    ),
    "paladin-protection": dict(
        core=[
            "Holy Shield",
            "Holy Strike",
            "Hammer of the Righteous",
            "Judgement",
            "Swift Judgement",
        ],
        auras=[("Holy Shield", DEF)],
        downranks=[],
        defensives=["Templar's Bulwark", "Divine Shield", "Divine Protection", "Lay on Hands"],
    ),
    "warrior-protection": dict(
        core=["Bloodrage", "Shield Slam"],
        auras=[("Shield Block", DEF), ("Demoralizing Shout", DEF), ("Thunder Clap", DEF)],
        downranks=[],
        defensives=["Shield Wall", "Last Stand"],
    ),
    # ---- healers (5)
    "druid-restoration": dict(
        core=[], efficiency=["Nature's Swiftness"], auras=[], downranks=[], defensives=["Barkskin"]
    ),
    "paladin-holy": dict(
        core=[],
        efficiency=[],
        auras=[],
        downranks=["Holy Shock", "Flash of Light"],
        defensives=["Divine Shield", "Divine Protection", "Lay on Hands"],
    ),
    "priest-discipline": dict(
        core=[],
        efficiency=["Inner Focus"],
        auras=[("Inner Fire", SELF)],
        downranks=["Power Word: Shield", "Prayer of Healing"],
        defensives=["Desperate Prayer"],
    ),
    "priest-holy": dict(
        core=[],
        efficiency=["Inner Focus"],
        auras=[("Inner Fire", SELF)],
        downranks=["Renew"],
        defensives=["Desperate Prayer"],
    ),
    "shaman-restoration": dict(
        core=[],
        efficiency=["Mana Tide Totem", "Nature's Swiftness"],
        auras=[("Water Shield", SELF)],
        downranks=[],
        defensives=[],
    ),
}


def _walk_apl(node: object, out: set[int]) -> None:
    """Collect every spell id an APL names under a `spellId` or `auraId` reference."""
    if isinstance(node, dict):
        for key, value in node.items():
            if key in ("spellId", "auraId") and isinstance(value, dict) and "spellId" in value:
                out.add(value["spellId"])
            else:
                _walk_apl(value, out)
    elif isinstance(node, list):
        for item in node:
            _walk_apl(item, out)


def _load_class(builds: Path, build: str, cls: str):
    root = builds / build
    spells = json.loads((root / "spellconst" / f"{cls}.json").read_text(encoding="utf-8"))
    ranks = json.loads((root / "trainables" / f"{cls}.json").read_text(encoding="utf-8"))
    trees = json.loads((root / "talents" / f"{cls}.json").read_text(encoding="utf-8"))
    sc = {int(k): v for k, v in spells["spells"].items()}
    tr = {t["name"]: t["ranks"] for t in ranks["trainables"]}
    talents = {t["name"] for tree in trees["trees"] for t in tree["talents"]}
    return sc, tr, talents


def _family_ids(name, sc, tr, apl_ids):
    ids = [r["id"] for r in tr.get(name, []) if r["level"] <= MAX_LEVEL and r["id"] in sc]
    ids += [i for i in sorted(apl_ids) if i in sc and sc[i]["name"] == name and i not in ids]
    if not ids:
        # a talent with one rank that neither the trainables nor this APL names: the
        # spellconst must hold exactly one spell of that name
        ids = [i for i, s in sc.items() if s["name"] == name]
        if len(ids) != 1:
            raise SystemExit(f"coach-rules: no single id for {name}: {ids}")
    for i in ids:
        if sc[i]["name"] != name:
            raise SystemExit(f"coach-rules: spell {i} is {sc[i]['name']!r}, not {name!r}")
    return ids


def _top_id(ids, sc, apl_ids, books=frozenset()):
    """The rank the coach judges against: the highest a trainer teaches. An Ahn'Qiraj
    book rank (curated/book-ranks.json) is not learnable in Phase 1, so it is the top only
    when the family has nothing else."""
    ids = [i for i in ids if i not in books] or ids
    best = max(sc[i]["rank"] for i in ids)
    tied = [i for i in ids if sc[i]["rank"] == best]
    preferred = [i for i in tied if i in apl_ids]
    return (preferred or sorted(tied))[0]


def _all_ranks(ids, top, sc):
    """ids plus every ranked (rank >= 1) spellconst spell of the same name, ordered by
    rank then id, with the stamped top rank last (F1's pickSpell falls back to the last id
    when nothing was cast). Unranked same-name spells (NPC copies, triggered effects) stay
    out: their cooldowns are not the player's."""
    name = sc[top]["name"]
    full = set(ids) | {i for i, s in sc.items() if s["name"] == name and s["rank"] >= 1}
    ordered = sorted(full - {top}, key=lambda i: (sc[i]["rank"], i))
    return [*ordered, top]


def _entries(names, ctx, books=frozenset()):
    if names is None:
        return None
    out = []
    for name in names:
        ids = _family_ids(name, ctx.sc, ctx.tr, ctx.apl_ids)
        top = _top_id(ids, ctx.sc, ctx.apl_ids, books)
        out.append(
            {
                "family": name,
                "spell_ids": _all_ranks(ids, top, ctx.sc),
                "talent": name in ctx.talents,
                "verified": ctx.stamp(top),
            }
        )
    return out


def _auras(pairs, ctx):
    if pairs is None:
        return None
    out = []
    for name, kind in pairs:
        ids = _family_ids(name, ctx.sc, ctx.tr, ctx.apl_ids)
        top = _top_id(ids, ctx.sc, ctx.apl_ids)
        aura = {
            "spell_id": top,
            "spell_ids": _all_ranks(ids, top, ctx.sc),
            "name": name,
            "kind": kind,
        }
        if name in ctx.talents:
            aura["talent"] = True
        aura["verified"] = ctx.stamp(top)
        out.append(aura)
    return out


class _Context:
    """One spec's resolved tables: what every list is built against."""

    def __init__(self, build, cls, sc, tr, talents, apl_ids, books=frozenset()):
        self.build, self.cls = build, cls
        self.sc, self.tr, self.talents, self.apl_ids = sc, tr, talents, apl_ids
        self.books = books

    def stamp(self, spell_id: int) -> str:
        return f"{self.build}/spellconst/{self.cls}.json#{spell_id}"


def render_spec(spec: str, build: str, data_root: Path = Path()) -> str:
    """The JSON text of one spec's rule file for `build`; `data_root` is the data/ dir."""
    lists = SPECS[spec]
    cls = spec.split("-")[0]
    sc, tr, talents = _load_class(data_root / "builds", build, cls)
    apl_ids: set[int] = set()
    apl = json.loads((data_root / "curated" / "apl" / f"{spec}.json").read_text(encoding="utf-8"))
    _walk_apl(apl["rotation"], apl_ids)
    ctx = _Context(build, cls, sc, tr, talents, apl_ids, load_book_ids(data_root / "curated"))
    doc: dict = {"spec": spec, "build": build}
    for key, value in (
        ("core", _entries(lists.get("core"), ctx)),
        ("auras", _auras(lists.get("auras"), ctx)),
        # A downrank is judged against the top rank a trainer teaches: a book rank is
        # not learnable in Phase 1 (round 6 healers: Renew 25315, Prayer of Healing 25316).
        ("downranks", _entries(lists.get("downranks"), ctx, ctx.books)),
        ("defensives", _entries(lists.get("defensives"), ctx)),
        ("efficiency", _entries(lists.get("efficiency", []), ctx)),
    ):
        if value is not None:
            doc[key] = value
    return json.dumps(doc, indent=2) + "\n"


def rule_file_path(spec: str, data_root: Path = Path()) -> Path:
    return data_root / "curated" / "coach" / f"{spec}.json"


def write_rule_files(build: str, data_root: Path = Path()) -> list[Path]:
    """Regenerate every spec's rule file for `build`; returns the paths written."""
    paths = []
    for spec in SPECS:
        path = rule_file_path(spec, data_root)
        path.write_text(render_spec(spec, build, data_root), encoding="utf-8")
        paths.append(path)
    return paths


def drifted_rule_files(build: str, data_root: Path = Path()) -> list[Path]:
    """The rule files whose committed text differs from a regeneration for `build`."""
    return [
        rule_file_path(spec, data_root)
        for spec in SPECS
        if not rule_file_path(spec, data_root).is_file()
        or rule_file_path(spec, data_root).read_text(encoding="utf-8")
        != render_spec(spec, build, data_root)
    ]
