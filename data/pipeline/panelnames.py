"""Which class spells a rotation panel line may name at a level.

The addon's rotation panel prints one condition line per cast (`panel_when` in the curated
APLs). A line must never name a spell the character cannot have at that band (coordinator
ruling, 2026-11-08): Beast Mastery's Arcane Shot read "when Multi-Shot is 1.5+ s from ready"
at level 10, eight levels before a hunter trains Multi-Shot. This module says what a class
can name (every trainer spell, talent and spellranks family of the class) and what a
character has at a level (trainer spells learned by then, talents the spec's published build
has taken by then), and finds the names a line uses that the character lacks.
"""

from __future__ import annotations

import json
import re
from collections.abc import Iterable, Mapping
from dataclasses import dataclass
from pathlib import Path

from pipeline.talent_levels import first_point_levels

#: The trainables `source` of a spell a talent teaches: known only once the build takes it.
TALENT_SOURCE = "talent"


@dataclass(frozen=True)
class ClassSpellNames:
    """One class's spell names: trainer spells by first learn level, talents by the spell
    ids that teach them, and names the class can never learn (a Season of Discovery rune
    spellranks carries)."""

    trained: Mapping[str, int]
    talents: Mapping[str, frozenset[int]]
    unlearnable: frozenset[str]

    @property
    def all_names(self) -> frozenset[str]:
        return frozenset(self.trained) | frozenset(self.talents) | self.unlearnable


def load_class_spell_names(
    build_dir: Path,
    class_slug: str,
    spellranks: dict,
    owned_auras: Mapping[int, int] | None = None,
) -> ClassSpellNames:
    """`ClassSpellNames` for `class_slug` off `build_dir`'s trainables and talents files and
    the build's `spellranks`. A family neither file teaches is a talent's when one of its ids
    is an aura a talent gives (`owned_auras`, curated/talent-owned-auras.json: Clearcasting
    from Elemental Focus), else unlearnable."""
    talents: dict[str, set[int]] = {}
    talents_path = build_dir / "talents" / f"{class_slug}.json"
    if talents_path.exists():
        for tree in json.loads(talents_path.read_text(encoding="utf-8"))["trees"]:
            for talent in tree["talents"]:
                ids = {talent.get("spell_id")} | {
                    r.get("spell_id") for r in talent.get("ranks", [])
                }
                talents.setdefault(talent["name"], set()).update(
                    i for i in ids if isinstance(i, int) and i > 0
                )
    trained: dict[str, int] = {}
    trainables_path = build_dir / "trainables" / f"{class_slug}.json"
    if trainables_path.exists():
        for entry in json.loads(trainables_path.read_text(encoding="utf-8"))["trainables"]:
            ids = {rank["id"] for rank in entry["ranks"]}
            if entry.get("source") == TALENT_SOURCE or entry["name"] in talents:
                talents.setdefault(entry["name"], set()).update(ids)
                continue
            levels = [rank["level"] for rank in entry["ranks"] if rank["level"] > 0]
            trained[entry["name"]] = min(levels, default=1)
    families = spellranks.get("classes", {}).get(class_slug, {})
    for name, ranks in families.items():
        auras = {rank["id"] for rank in ranks} & set(owned_auras or {})
        if auras and name not in trained:
            talents.setdefault(name, set()).update(auras)
    return ClassSpellNames(
        trained=trained,
        talents={name: frozenset(ids) for name, ids in talents.items()},
        unlearnable=frozenset(set(families) - set(trained) - set(talents)),
    )


def known_names(
    names: ClassSpellNames, level: int, rank_levels: Mapping[tuple[int, int], int]
) -> frozenset[str]:
    """The class spells a character of the spec has at `level`: every trainer spell learned
    by then and every talent the published build (`rank_levels`, load_spec_rank_levels) has
    put its first point in by then."""
    taken = first_point_levels(rank_levels)
    trained = {name for name, first in names.trained.items() if first <= level}
    talents = {
        name
        for name, ids in names.talents.items()
        if any(spell_id in taken and taken[spell_id] <= level for spell_id in ids)
    }
    return frozenset(trained | talents)


def named_spells(line: str, names: Iterable[str]) -> list[str]:
    """The names `line` uses, as whole words with their case, in order. A longer name wins
    over one inside it ("Improved Rend" is not also "Rend")."""
    found: list[tuple[int, str]] = []
    taken = [False] * len(line)
    for name in sorted(set(names), key=len, reverse=True):
        for match in re.finditer(rf"(?<!\w){re.escape(name)}(?!\w)", line):
            if not any(taken[match.start() : match.end()]):
                taken[match.start() : match.end()] = [True] * (match.end() - match.start())
                found.append((match.start(), name))
    return [name for _, name in sorted(found)]


def unavailable_names(
    line: str, own_name: str, known: frozenset[str], all_names: frozenset[str]
) -> list[str]:
    """The class spells `line` names that the character lacks: every name in `all_names` it
    uses except `own_name` (the spell the row casts) and the `known` ones."""
    return [
        name for name in named_spells(line, all_names) if name != own_name and name not in known
    ]


@dataclass(frozen=True)
class PanelNameViolation:
    """A panel line naming class spells the character lacks at the band it prints in."""

    spec: str
    level: int
    spell: str
    condition: str
    missing: tuple[str, ...]

    def __str__(self) -> str:
        return (
            f"{self.spec} {self.level} {self.spell}: {self.condition!r} names {list(self.missing)}"
        )


def panel_name_violations(
    root: Path, build: str, curated_dir: Path = Path("curated")
) -> list[PanelNameViolation]:
    """Every published band line, over every curated spec, whose condition names a class
    spell the character lacks at the band's level. A band's own level is the lowest that
    reads it (Rotation.lua bandFor), and what a character has only grows with level."""
    from pipeline.addonrotation import _class_slug, build_rotations
    from pipeline.talent_levels import load_spec_rank_levels, load_talent_owned_auras

    spellranks = json.loads((root / build / "spellranks.json").read_text(encoding="utf-8"))
    violations: list[PanelNameViolation] = []
    for spec, bands in build_rotations(root, build, curated_dir).items():
        class_slug = _class_slug(spec)
        owned = load_talent_owned_auras(curated_dir, class_slug)
        names = load_class_spell_names(root / build, class_slug, spellranks, owned)
        rank_levels = load_spec_rank_levels(root / build, spec, class_slug, owned)
        for band in bands:
            known = known_names(names, band.level, rank_levels)
            for line in band.lines:
                missing = unavailable_names(line.condition, line.name, known, names.all_names)
                if missing:
                    violations.append(
                        PanelNameViolation(
                            spec, band.level, line.name, line.condition, tuple(missing)
                        )
                    )
    return violations
