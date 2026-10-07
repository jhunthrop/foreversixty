"""Per-class trainable abilities: what a Forever player learns, from the client.

The simulator only registers the spells its class packages declare, so a
trainable ability the engine lacks (Frostfire Bolt, Icy Veins, the new Curse of
the Elements) is invisible to a report that starts from the engine's spellbook.
This module is the other side of that comparison: the client's own list of what
each class can learn, written to `builds/<build>/trainables/<class>.json` for
the fork's conformance report to diff against the engine.

**Sources.** `SkillLineAbility` rows on the class skill lines
(`SkillLine.CategoryID == 7`: Arms, Fire, Affliction, ...) are the primary
source. A row names its class in `ClassMask`, or leaves it 0 (the Affliction
Curse of the Elements ranks), in which case the spell's own
`SpellClassOptions.SpellClassSet` names the class; a row for which neither
names exactly one class (a pet's Lacerate under Survival) is not a player
ability and is dropped. The client does not list every learnable spell there:
Unstable Affliction and Hydra Shot have no SkillLineAbility row in build
1.60.1.70009, yet are gated by a level in `SpellLevels` like any trained spell.
So a class-family spell (SpellClassSet) with a rank of 1 or more, a level and
an active shape is also listed, tagged `source: "class_spell"`; the report keeps
those apart from the learnable ones. Only AcquireMethod 0 rows with a learn
level above 0 count as learnable: method 3 rows are Season of Discovery runes.

**Ranks.** An ability is every spell sharing a name within the class, merged
with every `SupercedesSpell` link: the client reuses one id chain per rank for
most abilities and a fresh unlinked id per rank for others (Frostfire Bolt).

**Active.** A trainable is active when any rank has a power cost, a cast time
or a cooldown -- the abilities a rotation can press. Passives (weapon skills,
proficiencies, auras) are listed with `active: false` and the engine report
ignores them.
"""

from __future__ import annotations

import logging
import re
import shutil
from collections import defaultdict
from collections.abc import Callable
from pathlib import Path

from pydantic import BaseModel

from pipeline.csvio import read_csv
from pipeline.normalize import write_model
from pipeline.proficiency import CLASS_MASK_BIT, CLASS_SLUG_BY_ID
from pipeline.simconst import SPELL_FAMILY_BY_CLASS_SLUG, _by_spell_id, _int, _rank

logger = logging.getLogger(__name__)

TRAINABLES = "trainables"

#: SkillLine.CategoryID of the class skill lines.
CLASS_SKILL_CATEGORY = 7

#: The class skill lines a player learns spells on, by client display name.
#: Engraving, Pet - *, Beast Training, Runes and the other category-7 lines
#: carry engravable runes and pet or companion spells, which are not the
#: spellbook this audit covers.
CLASS_SKILL_LINES = frozenset(
    {
        "Arms", "Fury", "Protection", "Fire", "Frost", "Arcane", "Affliction",
        "Demonology", "Destruction", "Holy", "Shadow Magic", "Discipline",
        "Balance", "Feral Combat", "Restoration", "Assassination", "Combat",
        "Subtlety", "Beast Mastery", "Marksmanship", "Survival", "Retribution",
        "Elemental Combat", "Enhancement",
    }
)  # fmt: skip

#: SkillLineAbility.AcquireMethod of a spell a player learns. Method 3 rows are
#: Season of Discovery runes (learn level 0 or 1), which Forever does not teach.
LEARN_ACQUIRE_METHOD = 0

#: Skill line reported for a spell found only through its class family.
FAMILY_SKILL_LINE = ""

SOURCE_SKILL_LINE = "skill_line_ability"
SOURCE_CLASS_SPELL = "class_spell"

#: Name markers for test, duplicate and placeholder spells that are never learned.
_EXCLUDED_NAME_MARKERS = ("DNT", "Copy of ", "zzOLD", "UNUSED")

#: A development spell: "Test Maul", "Runecarving Test - Crippling Poison", "Tyler Test Imbue".
_TEST_NAME = re.compile(r"\bTest\b|\bTesting\b")

_SLUG_BY_CLASS_MASK = {CLASS_MASK_BIT[cid]: slug for cid, slug in CLASS_SLUG_BY_ID.items()}
_SLUG_BY_FAMILY = {family: slug for slug, family in SPELL_FAMILY_BY_CLASS_SLUG.items()}


class TrainableRank(BaseModel):
    id: int
    rank: int
    level: int


class Trainable(BaseModel):
    name: str
    skill_line: str
    source: str
    active: bool
    #: Of the highest-level rank: the figures that say why the ability matters.
    cost: int
    cost_type: int
    cast_time_ms: int
    cooldown_ms: int
    ranks: list[TrainableRank]


class ClassTrainables(BaseModel):
    build: str
    class_slug: str
    trainables: list[Trainable]


class _Candidate(BaseModel):
    spell_id: int
    class_slug: str
    skill_line: str
    source: str
    supersedes: int


def is_excluded_name(name: str) -> bool:
    return (
        not name
        or any(marker in name for marker in _EXCLUDED_NAME_MARKERS)
        or _TEST_NAME.search(name) is not None
    )


def _class_from_mask(mask: int) -> str | None:
    return _SLUG_BY_CLASS_MASK.get(mask)


def _skill_line_names(raw: Path) -> dict[int, str]:
    return {
        int(row["ID"]): row["DisplayName_lang"]
        for row in read_csv(raw / "SkillLine.csv")
        if row["CategoryID"] == str(CLASS_SKILL_CATEGORY)
        and row["DisplayName_lang"] in CLASS_SKILL_LINES
    }


def _ability_candidates(
    raw: Path, family_by_spell: dict[int, int], level_of: Callable[[int], int]
) -> list[_Candidate]:
    lines = _skill_line_names(raw)
    found: list[_Candidate] = []
    for row in read_csv(raw / "SkillLineAbility.csv"):
        line = lines.get(int(row["SkillLine"]))
        if line is None:
            continue
        spell_id = int(row["Spell"])
        if int(row["AcquireMethod"] or 0) != LEARN_ACQUIRE_METHOD or level_of(spell_id) < 1:
            continue
        slug = _class_from_mask(int(row["ClassMask"])) or _SLUG_BY_FAMILY.get(
            family_by_spell.get(spell_id, -1)
        )
        if slug is None:
            continue
        found.append(
            _Candidate(
                spell_id=spell_id,
                class_slug=slug,
                skill_line=line,
                source=SOURCE_SKILL_LINE,
                supersedes=int(row["SupercedesSpell"] or 0),
            )
        )
    return found


def _family_by_spell(raw: Path) -> dict[int, int]:
    return {
        int(row["SpellID"]): int(row["SpellClassSet"])
        for row in read_csv(raw / "SpellClassOptions.csv")
    }


class _SpellTables:
    """The client rows that describe one spell, indexed once."""

    def __init__(self, raw: Path) -> None:
        self.names = {int(r["ID"]): r["Name_lang"] for r in read_csv(raw / "SpellName.csv")}
        self.subtexts = {int(r["ID"]): r["NameSubtext_lang"] for r in read_csv(raw / "Spell.csv")}
        self.misc = _by_spell_id(read_csv(raw / "SpellMisc.csv"))
        self.cast_times = {int(r["ID"]): r for r in read_csv(raw / "SpellCastTimes.csv")}
        self.cooldowns = _by_spell_id(read_csv(raw / "SpellCooldowns.csv"))
        self.levels = _by_spell_id(read_csv(raw / "SpellLevels.csv"))
        self.powers = _by_spell_id(read_csv(raw / "SpellPower.csv"))

    def rank(self, spell_id: int) -> int:
        return _rank(self.subtexts.get(spell_id, ""))

    def level(self, spell_id: int) -> int:
        return _int(self.levels.get(spell_id), "SpellLevel")

    def cast_time_ms(self, spell_id: int) -> int:
        cast = self.cast_times.get(_int(self.misc.get(spell_id), "CastingTimeIndex"))
        return _int(cast, "Base")

    def cooldown_ms(self, spell_id: int) -> int:
        row = self.cooldowns.get(spell_id)
        return max(_int(row, "RecoveryTime"), _int(row, "CategoryRecoveryTime"))

    def cost(self, spell_id: int) -> tuple[int, int]:
        power = self.powers.get(spell_id)
        return _int(power, "ManaCost"), _int(power, "PowerType")

    def is_active(self, spell_id: int) -> bool:
        cost, _ = self.cost(spell_id)
        return (
            cost > 0
            or _float_positive(self.powers.get(spell_id), "PowerCostPct")
            or self.cast_time_ms(spell_id) > 0
            or self.cooldown_ms(spell_id) > 0
        )


def _float_positive(row: dict[str, str] | None, column: str) -> bool:
    return bool(row) and float(row.get(column) or 0) > 0


def _pet_spell_names(raw: Path, names: dict[int, str]) -> set[str]:
    """Names listed on a non-class category-7 line (Pet - *, Beast Training, Engraving, ...).

    A pet ability carries its owner's SpellClassSet, so the family alone cannot
    tell Hunter Growl (a pet's) from a player spell.
    """
    class_lines = _skill_line_names(raw)
    non_class = {
        int(row["ID"])
        for row in read_csv(raw / "SkillLine.csv")
        if row["CategoryID"] == str(CLASS_SKILL_CATEGORY) and int(row["ID"]) not in class_lines
    }
    return {
        names[int(row["Spell"])]
        for row in read_csv(raw / "SkillLineAbility.csv")
        if int(row["SkillLine"]) in non_class and int(row["Spell"]) in names
    }


def _family_candidates(
    tables: _SpellTables, family_by_spell: dict[int, int], known: set[int], pet_names: set[str]
) -> list[_Candidate]:
    """Ranked, levelled, active class-family spells no SkillLineAbility row lists."""
    found: list[_Candidate] = []
    for spell_id, family in family_by_spell.items():
        slug = _SLUG_BY_FAMILY.get(family)
        if slug is None or spell_id in known or tables.names.get(spell_id) in pet_names:
            continue
        if tables.rank(spell_id) < 1 or tables.level(spell_id) < 1:
            continue
        if not tables.is_active(spell_id):
            continue
        found.append(
            _Candidate(
                spell_id=spell_id,
                class_slug=slug,
                skill_line=FAMILY_SKILL_LINE,
                source=SOURCE_CLASS_SPELL,
                supersedes=0,
            )
        )
    return found


def _group_by_ability(
    candidates: list[_Candidate], names: dict[int, str]
) -> dict[int, list[_Candidate]]:
    """Union candidates sharing a name or a SupercedesSpell link; key = lowest id."""
    parent: dict[int, int] = {c.spell_id: c.spell_id for c in candidates}

    def find(x: int) -> int:
        while parent[x] != x:
            parent[x] = parent[parent[x]]
            x = parent[x]
        return x

    def union(a: int, b: int) -> None:
        if a in parent and b in parent:
            parent[find(a)] = find(b)

    first_by_name: dict[str, int] = {}
    for c in candidates:
        union(c.spell_id, first_by_name.setdefault(names[c.spell_id], c.spell_id))
        if c.supersedes:
            union(c.spell_id, c.supersedes)
    groups: dict[int, list[_Candidate]] = defaultdict(list)
    for c in candidates:
        groups[find(c.spell_id)].append(c)
    return groups


def _build_trainable(group: list[_Candidate], tables: _SpellTables) -> Trainable:
    ranks = sorted(
        (
            TrainableRank(
                id=c.spell_id, rank=tables.rank(c.spell_id), level=tables.level(c.spell_id)
            )
            for c in group
        ),
        key=lambda r: (r.rank, r.level, r.id),
    )
    top = max(group, key=lambda c: (tables.level(c.spell_id), tables.rank(c.spell_id), c.spell_id))
    cost, cost_type = tables.cost(top.spell_id)
    lead = min(group, key=lambda c: (c.source != SOURCE_SKILL_LINE, c.spell_id))
    return Trainable(
        name=tables.names[top.spell_id],
        skill_line=lead.skill_line,
        source=lead.source,
        active=any(tables.is_active(c.spell_id) for c in group),
        cost=cost,
        cost_type=cost_type,
        cast_time_ms=tables.cast_time_ms(top.spell_id),
        cooldown_ms=tables.cooldown_ms(top.spell_id),
        ranks=ranks,
    )


def build_trainables(build: str, raw: Path) -> list[ClassTrainables]:
    tables = _SpellTables(raw)
    family_by_spell = _family_by_spell(raw)
    abilities = [
        c
        for c in _ability_candidates(raw, family_by_spell, tables.level)
        if not is_excluded_name(tables.names.get(c.spell_id, ""))
    ]
    # A spell with any SkillLineAbility row (a rune included) is never a no-learn-row extra.
    known = {int(row["Spell"]) for row in read_csv(raw / "SkillLineAbility.csv")}
    extras = [
        c
        for c in _family_candidates(
            tables, family_by_spell, known, _pet_spell_names(raw, tables.names)
        )
        if not is_excluded_name(tables.names.get(c.spell_id, ""))
    ]
    by_class: dict[str, list[_Candidate]] = defaultdict(list)
    for c in [*abilities, *extras]:
        by_class[c.class_slug].append(c)
    records = []
    for slug in sorted(SPELL_FAMILY_BY_CLASS_SLUG):
        groups = _group_by_ability(by_class[slug], tables.names)
        trainables = sorted(
            (_build_trainable(group, tables) for group in groups.values()),
            key=lambda t: (t.name, t.ranks[0].id),
        )
        records.append(ClassTrainables(build=build, class_slug=slug, trainables=trainables))
    return records


def write_trainables(build: str, root: Path = Path("builds")) -> Path:
    """Write `builds/<build>/trainables/<class>.json`.

    The manifest is not refreshed here: the next `simconst` or `levels` run
    re-hashes the whole build directory, trainables included.
    """
    build_dir = root / build
    raw = build_dir / "raw"
    if not (raw / "SkillLineAbility.csv").exists():
        raise SystemExit(f"no {raw / 'SkillLineAbility.csv'}; run `python -m pipeline fetch` first")
    out = build_dir / TRAINABLES
    shutil.rmtree(out, ignore_errors=True)
    total = 0
    for record in build_trainables(build, raw):
        write_model(record, out / f"{record.class_slug}.json")
        total += len(record.trainables)
    logger.info("wrote %s: %d trainable abilities", out, total)
    return out
