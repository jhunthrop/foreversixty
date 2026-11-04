"""The game's own spell tooltips, for the site's hover cards: one file per class and a shared one.

A guide names spells (its rotation, cooldowns, prose), racials, enchants and consumables.
The site shows each one as the client would: name, rank, cost, range, cast time, cooldown
and the description with its numbers filled in. Every figure here is a client row:

- the spells are every rank of every ability in the build's `trainables/<class>.json`
  (what a Forever player learns, talent-taught abilities included);
- `SpellName`/`Spell.NameSubtext_lang` give the name and rank, `SpellPower` the cost,
  `SpellMisc` the cast time and range, `SpellCooldowns` the cooldown, and the description
  is `Spell.Description_lang` resolved by `pipeline.spelltext` (the same text the talents
  carry);
- racials are the spells on the client's racial skill lines (`pipeline.racialicons`'
  rule), by name, the lowest spell id taking the name;
- an enchant is its `enchants.json` row's own spell;
- a consumable is the item `simbuffs.json` names, matched to `items.json` by exact name,
  and the spell its `ItemEffect` row triggers on use.

Written to `builds/<build>/spelltips/`, after `trainables` and `loot` (it reads both); the
data job's later `simdb` re-hashes the manifest with it.
"""

from __future__ import annotations

import json
import logging
import shutil
from pathlib import Path

from pydantic import BaseModel

from pipeline.csvio import read_csv
from pipeline.hotfix_push import require_layered
from pipeline.icons import icon_names, resolve_icon
from pipeline.normalize import write_model
from pipeline.racialicons import RACIAL_LINE_WORD, RACIAL_SKILL_CATEGORY
from pipeline.simconst import _by_spell_id, _int, _rank
from pipeline.spelltext import ExtraRows, SpellText, load_spell_text

logger = logging.getLogger(__name__)

SPELLTIPS = "spelltips"
SHARED = "shared"

#: `ItemEffect.TriggerType` of an on-use effect.
ON_USE_TRIGGER = 0

#: `SpellMisc.Attributes_0`: the spell is a passive (SPELL_ATTR0_PASSIVE).
PASSIVE_BIT = 0x40
#: `SpellMisc.Attributes_1`: the spell is channeled (SPELL_ATTR1_CHANNELED_1 and _2).
CHANNELED_BITS = 0x04 | 0x40


class SpellTip(BaseModel):
    name: str
    #: 0 for a spell the client gives no rank.
    rank: int
    #: In the power's own client units (rage is tenths); 0 with `cost_pct` for a percent cost.
    cost: int
    #: `SpellPower.PowerType`: 0 mana, 1 rage, 2 focus, 3 energy.
    cost_type: int
    #: A percent of base mana, 0 when the cost is flat.
    cost_pct: float
    cast_time_ms: int
    channeled: bool
    passive: bool
    cooldown_ms: int
    #: Maximum range in yards, 0 for self.
    range_yd: float
    description: str
    #: The spell's own icon (SpellMisc.SpellIconFileDataID through ManifestInterfaceData), the
    #: client's placeholder when it has none. Appended last: key order is the file's contract.
    icon: str = ""


class ClassSpellTips(BaseModel):
    build: str
    class_slug: str
    spells: dict[str, SpellTip]


class ConsumableTip(BaseModel):
    item_id: int
    spell_id: int
    #: The on-use spell's description, the "Use:" line of the item's tooltip.
    use: str


#: Every raw table a tooltip reads: each must be the hotfix-layered copy.
INPUT_TABLES: tuple[str, ...] = (
    "ItemEffect",
    "ItemXItemEffect",
    "ManifestInterfaceData",
    "SkillLine",
    "SkillLineAbility",
    "Spell",
    "SpellAuraOptions",
    "SpellCastTimes",
    "SpellCooldowns",
    "SpellDuration",
    "SpellEffect",
    "SpellMisc",
    "SpellName",
    "SpellPower",
    "SpellRadius",
    "SpellRange",
    "SpellTargetRestrictions",
)


class SharedSpellTips(BaseModel):
    build: str
    #: Racial name, as the client spells it -> its spell.
    racials: dict[str, SpellTip]
    #: `enchants.json` id -> the enchant's spell.
    enchants: dict[str, SpellTip]
    #: `simbuffs.json` id -> the consumable's item and its on-use line.
    consumables: dict[str, ConsumableTip]
    #: Each pushed input table -> the sha256 of its hotfix-layered copy these tips read
    #: (`pipeline.hotfix_push.require_layered`); tools/activate-build.sh refuses without it.
    hotfix_layer: dict[str, str]


class _Tables:
    """The client rows a tooltip reads, indexed once."""

    def __init__(self, raw: Path) -> None:
        spell_rows = read_csv(raw / "Spell.csv")
        name_rows = read_csv(raw / "SpellName.csv")
        misc_rows = read_csv(raw / "SpellMisc.csv")
        self.names = {int(r["ID"]): r["Name_lang"] for r in name_rows}
        self.subtexts = {int(r["ID"]): r["NameSubtext_lang"] for r in spell_rows}
        self.misc = _by_spell_id(misc_rows)
        self.icons = icon_names(read_csv(raw / "ManifestInterfaceData.csv"))
        self.cast_times = {int(r["ID"]): r for r in read_csv(raw / "SpellCastTimes.csv")}
        self.cooldowns = _by_spell_id(read_csv(raw / "SpellCooldowns.csv"))
        self.powers = _by_spell_id(read_csv(raw / "SpellPower.csv"))
        range_rows = read_csv(raw / "SpellRange.csv")
        self.ranges = {int(r["ID"]): float(r["RangeMax_0"] or 0) for r in range_rows}
        self.text: SpellText = load_spell_text(
            spell_rows,
            misc_rows,
            read_csv(raw / "SpellEffect.csv"),
            read_csv(raw / "SpellDuration.csv"),
            ExtraRows(
                aura_options=_optional(raw / "SpellAuraOptions.csv"),
                radius=_optional(raw / "SpellRadius.csv"),
                range=range_rows,
                names=name_rows,
                target_restrictions=_optional(raw / "SpellTargetRestrictions.csv"),
            ),
        )

    def tip(self, spell_id: int) -> SpellTip | None:
        name = self.names.get(spell_id)
        if name is None:
            return None
        misc = self.misc.get(spell_id)
        power = self.powers.get(spell_id)
        cooldown = self.cooldowns.get(spell_id)
        cast = self.cast_times.get(_int(misc, "CastingTimeIndex"))
        return SpellTip(
            name=name,
            rank=_rank(self.subtexts.get(spell_id, "")),
            cost=_int(power, "ManaCost"),
            cost_type=_int(power, "PowerType"),
            cost_pct=float((power or {}).get("PowerCostPct") or 0),
            cast_time_ms=max(_int(cast, "Base"), 0),
            channeled=bool(_int(misc, "Attributes_1") & CHANNELED_BITS),
            passive=bool(_int(misc, "Attributes_0") & PASSIVE_BIT),
            cooldown_ms=max(
                _int(cooldown, "RecoveryTime"), _int(cooldown, "CategoryRecoveryTime"), 0
            ),
            range_yd=self.ranges.get(_int(misc, "RangeIndex"), 0.0),
            description=self.text.describe(spell_id).strip(),
            icon=resolve_icon(self.text.icon_file_id(spell_id), self.icons, f"spell {spell_id}"),
        )


def _optional(path: Path) -> list[dict[str, str]]:
    return read_csv(path) if path.exists() else []


def _read(path: Path) -> object:
    return json.loads(path.read_text(encoding="utf-8"))


def class_spell_tips(build: str, trainables: dict, tables: _Tables) -> ClassSpellTips:
    """Every rank of every trainable ability the class file lists."""
    ids = sorted({rank["id"] for t in trainables["trainables"] for rank in t["ranks"]})
    spells = {str(spell_id): tip for spell_id in ids if (tip := tables.tip(spell_id)) is not None}
    return ClassSpellTips(build=build, class_slug=trainables["class_slug"], spells=spells)


def racial_tips(raw: Path, tables: _Tables) -> dict[str, SpellTip]:
    """Racial name -> the tooltip of its lowest spell id on a racial skill line."""
    lines = {
        int(row["ID"])
        for row in read_csv(raw / "SkillLine.csv")
        if int(row["CategoryID"] or 0) == RACIAL_SKILL_CATEGORY
        and RACIAL_LINE_WORD in row["DisplayName_lang"]
    }
    spells = sorted(
        int(row["Spell"])
        for row in read_csv(raw / "SkillLineAbility.csv")
        if int(row["SkillLine"] or 0) in lines
    )
    out: dict[str, SpellTip] = {}
    for spell_id in spells:
        tip = tables.tip(spell_id)
        if tip is not None and tip.name not in out:
            out[tip.name] = tip
    return dict(sorted(out.items()))


def enchant_tips(enchants: list[dict], tables: _Tables) -> dict[str, SpellTip]:
    out: dict[str, SpellTip] = {}
    for record in enchants:
        spell_id = int(record.get("spell_id") or 0)
        tip = tables.tip(spell_id) if spell_id else None
        if tip is not None:
            out[str(record["id"])] = tip
    return out


def consumable_tips(
    raw: Path, simbuffs: dict[str, dict], items: list[dict], tables: _Tables
) -> dict[str, ConsumableTip]:
    """Each simbuffs entry whose name is exactly one client item with an on-use spell."""
    by_name: dict[str, list[int]] = {}
    for item in items:
        by_name.setdefault(item["name"], []).append(int(item["id"]))
    effects = {
        int(row["ID"]): int(row["SpellID"])
        for row in read_csv(raw / "ItemEffect.csv")
        if int(row["TriggerType"] or 0) == ON_USE_TRIGGER and int(row["SpellID"] or 0) > 0
    }
    use_spell: dict[int, int] = {}
    for row in read_csv(raw / "ItemXItemEffect.csv"):
        spell_id = effects.get(int(row["ItemEffectID"]))
        if spell_id is not None:
            use_spell.setdefault(int(row["ItemID"]), spell_id)
    out: dict[str, ConsumableTip] = {}
    for buff_id, entry in sorted(simbuffs.items()):
        candidates = [i for i in by_name.get(entry.get("name", ""), []) if i in use_spell]
        if len(candidates) != 1:
            continue
        item_id = candidates[0]
        use = tables.text.describe(use_spell[item_id]).strip()
        if use:
            out[buff_id] = ConsumableTip(item_id=item_id, spell_id=use_spell[item_id], use=use)
    return out


def write_spell_tips(build: str, root: Path = Path("builds")) -> Path:
    """Write `builds/<build>/spelltips/<class>.json` and `spelltips/shared.json`."""
    build_dir = root / build
    raw = build_dir / "raw"
    if not (raw / "Spell.csv").exists():
        raise SystemExit(f"no {raw / 'Spell.csv'}; run `python -m pipeline fetch` first")
    trainables_dir = build_dir / "trainables"
    if not trainables_dir.is_dir():
        raise SystemExit(f"no {trainables_dir}; run `python -m pipeline trainables` first")
    hotfix_layer = require_layered(raw, INPUT_TABLES)
    tables = _Tables(raw)
    out = build_dir / SPELLTIPS
    shutil.rmtree(out, ignore_errors=True)
    for path in sorted(trainables_dir.glob("*.json")):
        record = class_spell_tips(build, _read(path), tables)  # type: ignore[arg-type]
        write_model(record, out / f"{record.class_slug}.json")
    simbuffs_path = build_dir / "simbuffs.json"
    simbuffs = _read(simbuffs_path)["entries"] if simbuffs_path.exists() else {}  # type: ignore[index]
    shared = SharedSpellTips(
        build=build,
        racials=racial_tips(raw, tables),
        enchants=enchant_tips(_read(build_dir / "enchants.json"), tables),  # type: ignore[arg-type]
        consumables=consumable_tips(raw, simbuffs, _read(build_dir / "items.json"), tables),  # type: ignore[arg-type]
        hotfix_layer=hotfix_layer,
    )
    write_model(shared, out / f"{SHARED}.json")
    logger.info("wrote %s", out)
    return out
