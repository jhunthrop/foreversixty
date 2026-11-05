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
those apart from the learnable ones. Only AcquireMethod 0 (trainer) and 2
(learned with the skill line: every class's starting rank 1s, Battle Stance,
Auto Shot, Sniper Shot rank 3 at 58) rows with a learn level above 0 count as
learnable: method 3 rows are Season of Discovery runes. A method-2 row that
names no class (mask 0) counts only as a rank of an ability with a trainer row:
the unmasked ones of no such ability are rune effects (Metamorphosis : Demon
Charge).

**Ranks.** An ability is every spell sharing a name within the class, merged
with every `SupercedesSpell` link: the client reuses one id chain per rank for
most abilities and a fresh unlinked id per rank for others (Frostfire Bolt).
Some ranks carry a second learn row for the spell the castable one fires
(Lightning Shield 324 and its damage 26364, Explosive Trap 13813 and 13812,
Shadowguard, Tame Beast, Reincarnation): an inactive spell on the same rank
and level as an active one of its ability is that effect, not a rank.

**Active.** A trainable is active when any rank has a power cost, a cast time
or a cooldown -- the abilities a rotation can press. Passives (weapon skills,
proficiencies, auras) are listed with `active: false` and the engine report
ignores them.

**Talent or trainer.** The client's SkillLineAbility rows cannot tell a talent
from a trainer spell: a talent's spell carries the same learn row (Penance,
Holy Shock), and so does an Era talent Forever made a trainer spell (Aimed
Shot, Blessing of Kings, Lightwell). The build's live tree (`talents/`, the
hotfix-aware trees) decides: an ability a live talent teaches, by name or by
any rank's spell, is `source: "talent"`; everything else keeps its learn row's
source. A live talent with an active spell that neither source finds (Cold
Snap, Adrenaline Rush: no learn row, no rank) is listed as a talent too, so
every ability a player can press is in the file exactly once. A spell whose
talent node a hotfix repointed (Voice of Truth, Binding Heal in 1.60.1.70707)
keeps its learn row and is a trainer spell: Wowhead's live spellbook lists it
at that level.

**Removed learn rows.** A hotfix push that removes a spell's learn rows (status
2) leaves them in `raw/shipped/`, so the no-learn-row fallback never lists that
spell for the build the push targets. The next client bakes the removal into
its shipped files and the trace is gone: 1.60.1.70790 shipped without the Lava
Burst and Binding Heal rows 70755's push had removed, and both came back as
`class_spell`. The previous committed build is the memory the shipped tables
lost: a class-family spell with no learn row is listed only when that build
already listed it without a learn row (Unstable Affliction, Hydra Shot, and
Turn Evil, a class spell the live tree tags a talent), or did not know the
spell at all (a new ability). A spell the previous build knew and did not list
that way had learn rows that are now gone, and stays gone.
"""

from __future__ import annotations

import json
import logging
import re
import shutil
from collections import defaultdict
from collections.abc import Callable, Mapping
from dataclasses import dataclass
from pathlib import Path
from types import MappingProxyType

from pydantic import BaseModel

from pipeline.csvio import read_csv
from pipeline.hotfix_push import SHIPPED_DIR
from pipeline.normalize import write_model
from pipeline.proficiency import CLASS_MASK_BIT, CLASS_SLUG_BY_ID
from pipeline.simconst import SPELL_FAMILY_BY_CLASS_SLUG, _by_spell_id, _int, _rank
from pipeline.talent_levels import load_talent_spell_levels

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

#: SkillLineAbility.AcquireMethod of a spell a player buys from a trainer. Method 3
#: rows are Season of Discovery runes (learn level 0 or 1), which Forever does not teach.
LEARN_ACQUIRE_METHOD = 0
#: AcquireMethod of a spell learned with the skill line itself, at its level (a starting
#: rank 1, Battle Stance, Auto Shot, Sniper Shot rank 3).
AUTO_LEARN_ACQUIRE_METHOD = 2
LEARNABLE_ACQUIRE_METHODS = frozenset({LEARN_ACQUIRE_METHOD, AUTO_LEARN_ACQUIRE_METHOD})

#: Skill line reported for a spell found only through its class family.
FAMILY_SKILL_LINE = ""

SOURCE_SKILL_LINE = "skill_line_ability"
SOURCE_CLASS_SPELL = "class_spell"
#: An ability a talent of the build's live tree teaches.
SOURCE_TALENT = "talent"

#: Name markers for test, duplicate and placeholder spells that are never learned
#: ("(DND)": do not display, the client's hidden class passives).
_EXCLUDED_NAME_MARKERS = ("DNT", "(DND)", "Copy of ", "zzOLD", "UNUSED")

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
    #: An auto-learned row naming no class: a rank only of an ability with a trainer row.
    needs_trainer_row: bool = False


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
    raw: Path,
    family_by_spell: dict[int, int],
    level_of: Callable[[int], int],
    names_of: Callable[[int], str],
) -> list[_Candidate]:
    lines = _skill_line_names(raw)
    found: list[_Candidate] = []
    for row in read_csv(raw / "SkillLineAbility.csv"):
        line = lines.get(int(row["SkillLine"]))
        if line is None:
            continue
        spell_id = int(row["Spell"])
        method = int(row["AcquireMethod"] or 0)
        if method not in LEARNABLE_ACQUIRE_METHODS or level_of(spell_id) < 1:
            continue
        masked = _class_from_mask(int(row["ClassMask"]))
        slug = masked or _SLUG_BY_FAMILY.get(family_by_spell.get(spell_id, -1))
        if slug is None:
            continue
        found.append(
            _Candidate(
                spell_id=spell_id,
                class_slug=slug,
                skill_line=line,
                source=SOURCE_SKILL_LINE,
                supersedes=int(row["SupercedesSpell"] or 0),
                needs_trainer_row=method == AUTO_LEARN_ACQUIRE_METHOD and masked is None,
            )
        )
    return _with_trainer_rows(found, names_of)


def _with_trainer_rows(found: list[_Candidate], names_of: Callable[[int], str]) -> list[_Candidate]:
    """Drop the unmasked auto-learned rows of abilities no trainer row teaches."""
    trained = {
        (c.class_slug, names_of(c.spell_id))
        for c in found
        if c.source == SOURCE_SKILL_LINE and not c.needs_trainer_row
    }
    return [
        c
        for c in found
        if not c.needs_trainer_row or (c.class_slug, names_of(c.spell_id)) in trained
    ]


def _family_by_spell(raw: Path) -> dict[int, int]:
    return {
        int(row["SpellID"]): int(row["SpellClassSet"])
        for row in read_csv(raw / "SpellClassOptions.csv")
    }


def _spells_with_learn_rows(raw: Path) -> set[int]:
    """Every spell with a SkillLineAbility row (a rune's included), in the layered table or
    in the shipped one under `raw/shipped/`: a row a hotfix push removed (status 2) still
    says the spell was a learn-row spell, so the no-learn-row fallback never lists it."""
    tables = [raw / "SkillLineAbility.csv", raw / SHIPPED_DIR / "SkillLineAbility.csv"]
    return {int(row["Spell"]) for path in tables if path.exists() for row in read_csv(path)}


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


def _without_effect_rows(group: list[_Candidate], tables: _SpellTables) -> list[_Candidate]:
    """Drop an inactive learn-row spell holding the same rank and level as an active one
    of its ability: it is the effect that spell fires (Lightning Shield 26364)."""
    active = {
        (tables.rank(c.spell_id), tables.level(c.spell_id))
        for c in group
        if tables.is_active(c.spell_id)
    }
    return [
        c
        for c in group
        if tables.is_active(c.spell_id)
        or (tables.rank(c.spell_id), tables.level(c.spell_id)) not in active
    ]


def _build_trainable(group: list[_Candidate], tables: _SpellTables) -> Trainable:
    group = _without_effect_rows(group, tables)
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


class _LiveTree(BaseModel):
    """One class's live talents (`talents/<class>.json`): names, and each talent's spells."""

    tree_by_talent: dict[str, str] = {}
    spells_by_talent: dict[str, list[int]] = {}

    @property
    def spell_ids(self) -> set[int]:
        return {spell for spells in self.spells_by_talent.values() for spell in spells}

    def teaches(self, trainable: Trainable) -> bool:
        return trainable.name in self.tree_by_talent or any(
            rank.id in self.spell_ids for rank in trainable.ranks
        )


def _live_tree(talents_dir: Path | None, class_slug: str) -> _LiveTree:
    """The class's live tree, or an empty one when the build has no `talents/` for it."""
    path = None if talents_dir is None else talents_dir / f"{class_slug}.json"
    if path is None or not path.exists():
        return _LiveTree()
    trees = json.loads(path.read_text(encoding="utf-8"))["trees"]
    talents = [(tree["name"], talent) for tree in trees for talent in tree["talents"]]
    return _LiveTree(
        tree_by_talent={talent["name"]: tree for tree, talent in talents},
        spells_by_talent={
            talent["name"]: sorted(
                ({talent.get("spell_id") or 0} | {r.get("spell_id") or 0 for r in talent["ranks"]})
                - {0}
            )
            for _, talent in talents
        },
    )


def _unlisted_talents(
    live: _LiveTree, listed: list[Trainable], tables: _SpellTables, class_slug: str
) -> list[_Candidate]:
    """Active spells of live talents that no listed trainable covers (by name or by spell)."""
    names = {t.name for t in listed}
    ids = {rank.id for t in listed for rank in t.ranks}
    found: list[_Candidate] = []
    for talent, spells in live.spells_by_talent.items():
        if talent in names or ids.intersection(spells):
            continue
        if not any(spell in tables.names and tables.is_active(spell) for spell in spells):
            continue
        found.extend(
            _Candidate(
                spell_id=spell,
                class_slug=class_slug,
                skill_line=live.tree_by_talent[talent],
                source=SOURCE_TALENT,
                supersedes=0,
            )
            for spell in spells
            if spell in tables.names
        )
    return found


def _classified(trainable: Trainable, live: _LiveTree) -> Trainable:
    if not live.teaches(trainable):
        return trainable
    return trainable.model_copy(update={"source": SOURCE_TALENT})


def _unique_by_spell(candidates: list[_Candidate]) -> list[_Candidate]:
    """One candidate per (class, spell): a spell on two learn rows of a line is one rank."""
    seen: set[tuple[str, int]] = set()
    unique: list[_Candidate] = []
    for c in candidates:
        if (c.class_slug, c.spell_id) not in seen:
            seen.add((c.class_slug, c.spell_id))
            unique.append(c)
    return unique


def _with_talent_levels(trainable: Trainable, talent_levels: Mapping[int, int]) -> Trainable:
    """A rank a talent teaches is learned when its talent can first be taken
    (`pipeline.talent_levels`), and a trainer rank above it no earlier: the client's spell
    level says 40 for Stormstrike, a tier-3 talent taken at 25."""
    floors = [talent_levels[r.id] for r in trainable.ranks if r.id in talent_levels]
    if not floors:
        return trainable
    floor = min(floors)
    ranks = [
        r.model_copy(update={"level": talent_levels.get(r.id, max(r.level, floor))})
        for r in trainable.ranks
    ]
    return trainable.model_copy(update={"ranks": ranks})


def _class_trainables(
    candidates: list[_Candidate],
    tables: _SpellTables,
    live: _LiveTree,
    class_slug: str,
    talent_levels: Mapping[int, int] = MappingProxyType({}),
) -> list[Trainable]:
    listed = [
        _build_trainable(group, tables)
        for group in _group_by_ability(candidates, tables.names).values()
    ]
    talents = _unlisted_talents(live, listed, tables, class_slug)
    listed += [
        _build_trainable(group, tables)
        for group in _group_by_ability(talents, tables.names).values()
    ]
    return sorted(
        (_with_talent_levels(_classified(t, live), talent_levels) for t in listed),
        key=lambda t: (t.name, t.ranks[0].id),
    )


#: A client build directory name, and the number that orders them.
_BUILD_DIR = re.compile(r"^\d+\.\d+\.\d+\.(\d+)$")
#: The committed spell list every build carries (`normalize`): one {id, name} per spell.
SPELLS_FILE = "spells.json"


@dataclass(frozen=True)
class PreviousBuild:
    """What the previous committed build knew (the module's "Removed learn rows" note):
    every spell id its `spells.json` lists, and the ids it listed as trainable without a
    learn row (the empty family skill line: `source: "class_spell"`, or a talent the live
    tree re-tagged from one)."""

    build: str
    known: frozenset[int]
    no_learn_row: frozenset[int]

    def keeps_without_learn_row(self, spell_id: int) -> bool:
        """Whether a no-learn-row class spell of this build may be listed: it was already
        one in the previous build, or that build never knew the spell."""
        return spell_id in self.no_learn_row or spell_id not in self.known


def _build_number(name: str) -> int | None:
    match = _BUILD_DIR.match(name)
    return int(match.group(1)) if match else None


def load_previous_build(root: Path, build: str) -> PreviousBuild | None:
    """The newest build under `root` older than `build` whose spell list and trainables are
    committed, or None when there is none (the first build, or a bare checkout)."""
    number = _build_number(build)
    if number is None:
        return None
    candidates = [
        (n, path)
        for path in root.iterdir()
        if (n := _build_number(path.name)) is not None
        and n < number
        and (path / SPELLS_FILE).is_file()
        and (path / TRAINABLES).is_dir()
    ]
    if not candidates:
        return None
    _, path = max(candidates)
    known = {int(s["id"]) for s in json.loads((path / SPELLS_FILE).read_text(encoding="utf-8"))}
    # A class spell keeps the family's empty skill line when the live tree re-tags it a
    # talent (Turn Evil), so the empty line, not the source, says "no learn row".
    no_learn_row = {
        rank["id"]
        for record_path in sorted((path / TRAINABLES).glob("*.json"))
        for trainable in json.loads(record_path.read_text(encoding="utf-8"))["trainables"]
        if trainable["skill_line"] == FAMILY_SKILL_LINE
        for rank in trainable["ranks"]
    }
    return PreviousBuild(path.name, frozenset(known), frozenset(no_learn_row))


def build_trainables(
    build: str,
    raw: Path,
    talents_dir: Path | None = None,
    previous: PreviousBuild | None = None,
) -> list[ClassTrainables]:
    """Every class's trainables. With `talents_dir` (the build's committed, hotfix-aware
    `talents/`), each ability a live talent teaches is tagged `source: "talent"` and the
    live talents no learn row lists are added. With `previous` (`load_previous_build`), a
    no-learn-row class spell the previous build knew and did not list that way is a removed
    learn row and is dropped (see the module notes)."""
    tables = _SpellTables(raw)
    family_by_spell = _family_by_spell(raw)
    abilities = [
        c
        for c in _ability_candidates(
            raw, family_by_spell, tables.level, lambda spell: tables.names.get(spell, "")
        )
        if not is_excluded_name(tables.names.get(c.spell_id, ""))
    ]
    learned_names = {(c.class_slug, tables.names.get(c.spell_id, "")) for c in abilities}
    extras = [
        c
        for c in _family_candidates(
            tables,
            family_by_spell,
            _spells_with_learn_rows(raw),
            _pet_spell_names(raw, tables.names),
        )
        if not is_excluded_name(tables.names.get(c.spell_id, ""))
        # A no-learn-row copy of a learn-row ability (Aimed Shot 27632, Primal Bite 1238073)
        # is not one of its ranks: the learn rows are.
        and (c.class_slug, tables.names.get(c.spell_id, "")) not in learned_names
        # A learn row the previous build still had, or a push had removed, is gone for good.
        and (previous is None or previous.keeps_without_learn_row(c.spell_id))
    ]
    by_class: dict[str, list[_Candidate]] = defaultdict(list)
    for c in _unique_by_spell([*abilities, *extras]):
        by_class[c.class_slug].append(c)
    return [
        ClassTrainables(
            build=build,
            class_slug=slug,
            trainables=_class_trainables(
                by_class[slug],
                tables,
                _live_tree(talents_dir, slug),
                slug,
                load_talent_spell_levels(talents_dir, slug),
            ),
        )
        for slug in sorted(SPELL_FAMILY_BY_CLASS_SLUG)
    ]


def write_trainables(build: str, root: Path = Path("builds")) -> Path:
    """Write `builds/<build>/trainables/<class>.json`.

    The manifest is not refreshed here: the next `simconst` or `levels` run
    re-hashes the whole build directory, trainables included.
    """
    build_dir = root / build
    raw = build_dir / "raw"
    if not (raw / "SkillLineAbility.csv").exists():
        raise SystemExit(f"no {raw / 'SkillLineAbility.csv'}; run `python -m pipeline fetch` first")
    if not (build_dir / "talents").is_dir():
        raise SystemExit(f"no {build_dir / 'talents'}; run `python -m pipeline normalize` first")
    previous = load_previous_build(root, build)
    if previous is None:
        logger.info("no previous build under %s: every no-learn-row class spell is listed", root)
    else:
        logger.info("removed learn rows carried from build %s", previous.build)
    out = build_dir / TRAINABLES
    shutil.rmtree(out, ignore_errors=True)
    total = 0
    for record in build_trainables(build, raw, build_dir / "talents", previous):
        write_model(record, out / f"{record.class_slug}.json")
        total += len(record.trainables)
    logger.info("wrote %s: %d trainable abilities", out, total)
    return out
