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
    # Shifting Power is mana-limited (55% of base mana) and held above 60 energy; F1 cannot read
    # either, so a ready one is not judged as a lost cast (round 26).
    "druid-feral": dict(
        core=[],
        auras=[("Rip", DOT)],
        downranks=[],
        defensives=["Barkskin"],
    ),
    # Aimed Shot takes the cooldown it shares with Multi-Shot every time it comes up; Multi-Shot is
    # for packs and moving (round 25), so a ready Multi-Shot is never a lost cast.
    # Aimed Shot shares its cooldown with Multi-Shot, which F1 does not restart on (round 26).
    "hunter-beast-mastery": dict(
        core=["Bestial Wrath", "Rapid Fire"],
        auras=[("Serpent Sting", DOT), ("Aspect of the Hawk", SELF)],
        downranks=[],
        defensives=["Feign Death", "Deterrence"],
    ),
    "hunter-marksmanship": dict(
        core=["Rapid Fire"],
        auras=[("Aspect of the Hawk", SELF)],
        downranks=[],
        defensives=["Feign Death", "Deterrence"],
    ),
    # Immolation Trap is 18 s from 38 (Survivalist's Discipline); F1 reads 30 s (round 26).
    "hunter-survival": dict(
        core=["Raptor Strike", "Strider Kick"],
        auras=[("Lacerate", DOT)],
        downranks=[],
        defensives=["Feign Death", "Deterrence"],
    ),
    # Only Frost carries Ice Block and Ice Barrier (mage index.md); neither Arcane nor Fire build
    # takes them, so their defensives are an honest empty list.
    "mage-arcane": dict(
        core=["Arcane Power", "Presence of Mind"],
        auras=[],
        downranks=[],
        defensives=[],
    ),
    "mage-fire": dict(
        core=["Combustion", "Fire Blast"],
        auras=None,
        downranks=[],
        defensives=[],
    ),
    "mage-frost": dict(core=[], auras=[], downranks=[], defensives=["Ice Block", "Ice Barrier"]),
    # With Unsealed Wrath the raid build twists Seal of Command and Righteousness, so no one Seal's
    # uptime is the spec's rule (round 25); a dropped Seal is left to rotation adherence.
    "paladin-retribution": dict(
        core=["Judgement", "Holy Strike"],
        auras=[],
        downranks=[],
        defensives=["Divine Shield", "Divine Protection", "Lay on Hands"],
    ),
    "priest-shadow": dict(
        core=["Shadow Word: Death", "Devouring Plague", "Inner Focus"],
        auras=[("Shadow Word: Pain", DOT), ("Shadowform", SELF)],
        downranks=[],
        # Desperate Prayer is the Dwarf priest spell (race mask 4): no defensive every priest has.
        defensives=[],
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
    # Lightning Mastery fires from the rotation's cooldowns line, never by name, so it is not
    # judged as a core ability; Flame Shock's uptime is.
    "shaman-elemental": dict(core=[], auras=[("Flame Shock", DOT)], downranks=[], defensives=[]),
    "shaman-enhancement": dict(core=["Stormstrike"], auras=[], downranks=[], defensives=[]),
    # The rotation keeps its DoTs and bane up (DoTs crit in Forever) and fills with Wrack; Amplify
    # Curse is saved for the Agony after the last Doom at sixty, so it is not judged as a cooldown.
    "warlock-affliction": dict(
        core=[],
        auras=[
            ("Corruption", DOT),
            ("Bane of Agony", DOT),
            ("Siphon Life", DOT),
            ("Immolate", DOT),
        ],
        downranks=[],
        defensives=["Death Coil"],
    ),
    # Soul Fire is skipped in the raid build ("never cast it in a raid", round 25).
    "warlock-demonology": dict(
        core=["Bane of Doom"],
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
        # Last Stand is a protection talent neither Arms build takes.
        defensives=[],
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
        # A warrior's Demoralizing Shout holds the same slot (feral-bear.md), so its time counts.
        auras=[("Demoralizing Roar", DEF, ("warrior", "Demoralizing Shout"))],
        downranks=[],
        defensives=["Barkskin", "Frenzied Regeneration"],
    ),
    # Hammer of the Righteous and Judgement are held for Holy Shield's mana reserve (protection.md),
    # which F1 cannot see, so a ready one is not a lost cast (round 26).
    "paladin-protection": dict(
        core=[
            "Holy Shield",
            "Holy Strike",
        ],
        # Swift Judgement is held for a tank swap (protection.md): a ready one is no lost cast.
        auras=[("Holy Shield", DEF)],
        downranks=[],
        defensives=["Templar's Bulwark", "Divine Shield", "Divine Protection", "Lay on Hands"],
    ),
    "warrior-protection": dict(
        core=["Bloodrage", "Shield Slam"],
        # Demoralizing Shout is left to a druid's Roar when one holds the slot (protection.md): the
        # Roar's time counts as the slot filled (covered_by), so it is judged again (round 26).
        auras=[
            ("Shield Block", DEF),
            ("Demoralizing Shout", DEF, ("druid", "Demoralizing Roar")),
            ("Thunder Clap", DEF),
        ],
        downranks=[],
        defensives=["Shield Wall", "Last Stand"],
    ),
    # ---- healers (5)
    "druid-restoration": dict(
        # Nature's Swiftness is saved for a tank about to die, never rolled on cooldown (round 25).
        core=[],
        efficiency=[],
        auras=[],
        downranks=[],
        defensives=["Barkskin"],
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
        # Desperate Prayer is the Dwarf priest spell (race mask 4): no defensive every priest has.
        defensives=[],
    ),
    "priest-holy": dict(
        core=[],
        efficiency=["Inner Focus"],
        auras=[("Inner Fire", SELF)],
        downranks=["Renew"],
        # Desperate Prayer is the Dwarf priest spell (race mask 4): no defensive every priest has.
        defensives=[],
    ),
    "shaman-restoration": dict(
        core=[],
        # Nature's Swiftness is kept for the tank under a quarter (restoration.md), not rolled.
        efficiency=["Mana Tide Totem"],
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


def _entries(names, ctx):
    if names is None:
        return None
    out = []
    for name in names:
        ids = _family_ids(name, ctx.sc, ctx.tr, ctx.apl_ids)
        top = _top_id(ids, ctx.sc, ctx.apl_ids, ctx.books)
        out.append(
            {
                "family": name,
                "spell_ids": _all_ranks(ids, top, ctx.sc),
                "talent": name in ctx.talents,
                "verified": ctx.stamp(top),
            }
        )
    return out


def _cover_ids(cover, ctx):
    """Every ranked id of another class's aura that holds the same slot (cover is
    (class, name)), so F2 counts its time as the slot filled."""
    cls, name = cover
    spells = json.loads(
        (ctx.builds / ctx.build / "spellconst" / f"{cls}.json").read_text(encoding="utf-8")
    )["spells"]
    ids = sorted(int(k) for k, v in spells.items() if v["name"] == name and v["rank"] >= 1)
    if not ids:
        raise SystemExit(f"coach-rules: no ranked {name} in {cls} spellconst")
    return ids


def _auras(entries, ctx):
    if entries is None:
        return None
    out = []
    for entry in entries:
        name, kind = entry[0], entry[1]
        ids = _family_ids(name, ctx.sc, ctx.tr, ctx.apl_ids)
        top = _top_id(ids, ctx.sc, ctx.apl_ids, ctx.books)
        aura = {
            "spell_id": top,
            "spell_ids": _all_ranks(ids, top, ctx.sc),
            "name": name,
            "kind": kind,
        }
        if name in ctx.talents:
            aura["talent"] = True
        if len(entry) > 2:
            aura["covered_by"] = _cover_ids(entry[2], ctx)
        aura["verified"] = ctx.stamp(top)
        out.append(aura)
    return out


class _Context:
    """One spec's resolved tables: what every list is built against."""

    def __init__(self, build, cls, sc, tr, talents, apl_ids, books=frozenset(), builds=Path()):
        self.build, self.cls = build, cls
        self.sc, self.tr, self.talents, self.apl_ids = sc, tr, talents, apl_ids
        self.books = books
        self.builds = builds

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
    ctx = _Context(
        build,
        cls,
        sc,
        tr,
        talents,
        apl_ids,
        load_book_ids(data_root / "curated"),
        data_root / "builds",
    )
    doc: dict = {"spec": spec, "build": build}
    for key, value in (
        ("core", _entries(lists.get("core"), ctx)),
        ("auras", _auras(lists.get("auras"), ctx)),
        # Every list names the top rank a trainer teaches: a book rank is not learnable in
        # Phase 1 (round 6 healers: Renew 25315, Prayer of Healing 25316; round 14: the
        # auras' Aspect of the Hawk 25296 and Serpent Sting 25295).
        ("downranks", _entries(lists.get("downranks"), ctx)),
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
