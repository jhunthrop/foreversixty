"""Per-class spell constants, keyed by spell id.

Design section 2.3: a regenerated constants file per class carries base damage,
costs, cooldowns and durations, and the engine's ability files reference it
rather than a literal, so a beta patch that changes a number is a pipeline run
rather than a code edit. Forever has already moved several: Bloodthirst rank 4
states 48 and level 60 on build 1.60.1.69893 where Classic Era stated 45 and
level 40.

`SpellClassOptions.SpellClassSet` partitions the spell universe by class, and
the mapping is stable across the Classic lineage. On build 1.60.1.69893 that is
5,658 spells over nine files.

What is deliberately absent: there is no `SpellScaling` table on this lineage
(research/07-simulator.md 5.3), so nothing here is scaled. `sp_coefficient` and
`ap_coefficient` are `EffectBonusCoefficient` and `BonusCoefficientFromAP`
exactly as the table states them, zeros included. Vanilla coefficients are a
convention, not data; the engine lane owns the convention and its per-spell
overrides.

Two client quirks worth knowing: rage costs are stored times ten (Bloodthirst
is `cost: 300` with `cost_type: 1`), and the real cooldown of most abilities is
`CategoryRecoveryTime`, not `RecoveryTime`.
"""

from __future__ import annotations

import logging
import re
import shutil
from collections import defaultdict
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.manifest import refresh_manifest
from pipeline.models import ClassSpellConstants, SpellConstant, SpellEffectConstant
from pipeline.normalize import write_model
from pipeline.spelltext import effect_amount

logger = logging.getLogger(__name__)

SPELLCONST = "spellconst"

#: Class slug -> SpellClassOptions.SpellClassSet. Sets 0, 1 and 13 exist and
#: are not player classes; their spells are dropped.
SPELL_FAMILY_BY_CLASS_SLUG: dict[str, int] = {
    "mage": 3,
    "warrior": 4,
    "warlock": 5,
    "priest": 6,
    "druid": 7,
    "rogue": 8,
    "hunter": 9,
    "paladin": 10,
    "shaman": 11,
}
_SLUG_BY_FAMILY = {family: slug for slug, family in SPELL_FAMILY_BY_CLASS_SLUG.items()}
_FAMILY_MASK_COLUMNS = range(4)

_RANK = re.compile(r"Rank (\d+)")


def _rank(subtext: str) -> int:
    match = _RANK.search(subtext or "")
    return int(match.group(1)) if match else 0


def _by_spell_id(rows: list[dict[str, str]]) -> dict[int, dict[str, str]]:
    """First row per spell id, ignoring the non-default difficulties."""
    indexed: dict[int, dict[str, str]] = {}
    for row in rows:
        if row.get("DifficultyID", "0") not in ("0", ""):
            continue
        indexed.setdefault(int(row["SpellID"]), row)
    return indexed


def _int(row: dict[str, str] | None, column: str) -> int:
    """One column of an optional row, or 0.

    A spell that appears in SpellClassOptions and in no cooldown, power or
    level table is normal -- most passives are exactly that -- so a missing row
    or a missing column is zero rather than an error. This is not the
    fail-fast path: nothing here is guessed at, only absent.
    """
    if not row:
        return 0
    value = row.get(column, "")
    return int(value) if value not in ("", None) else 0


def build_spell_constants(build: str, raw: Path) -> list[ClassSpellConstants]:
    names = {int(row["ID"]): row["Name_lang"] for row in read_csv(raw / "SpellName.csv")}
    subtexts = {int(row["ID"]): row["NameSubtext_lang"] for row in read_csv(raw / "Spell.csv")}
    misc = _by_spell_id(read_csv(raw / "SpellMisc.csv"))
    cast_times = {int(row["ID"]): row for row in read_csv(raw / "SpellCastTimes.csv")}
    durations = {int(row["ID"]): row for row in read_csv(raw / "SpellDuration.csv")}
    cooldowns = _by_spell_id(read_csv(raw / "SpellCooldowns.csv"))
    levels = _by_spell_id(read_csv(raw / "SpellLevels.csv"))
    powers = _by_spell_id(read_csv(raw / "SpellPower.csv"))

    effects: dict[int, list[dict[str, str]]] = defaultdict(list)
    for row in read_csv(raw / "SpellEffect.csv"):
        if row.get("DifficultyID", "0") not in ("0", ""):
            continue
        effects[int(row["SpellID"])].append(row)

    spells_by_slug: dict[str, dict[str, SpellConstant]] = {
        slug: {} for slug in SPELL_FAMILY_BY_CLASS_SLUG
    }
    for row in read_csv(raw / "SpellClassOptions.csv"):
        slug = _SLUG_BY_FAMILY.get(int(row["SpellClassSet"]))
        if slug is None:
            continue
        spell_id = int(row["SpellID"])
        spell_misc = misc.get(spell_id)
        cast = cast_times.get(_int(spell_misc, "CastingTimeIndex"))
        duration = durations.get(_int(spell_misc, "DurationIndex"))
        power = powers.get(spell_id)
        spells_by_slug[slug][str(spell_id)] = SpellConstant(
            name=names.get(spell_id, ""),
            rank=_rank(subtexts.get(spell_id, "")),
            school_mask=_int(spell_misc, "SchoolMask"),
            cast_time_ms=_int(cast, "Base"),
            gcd_ms=_int(cooldowns.get(spell_id), "StartRecoveryTime"),
            cooldown_ms=_int(cooldowns.get(spell_id), "RecoveryTime"),
            category_cooldown_ms=_int(cooldowns.get(spell_id), "CategoryRecoveryTime"),
            duration_ms=_int(duration, "Duration"),
            cost=_int(power, "ManaCost"),
            cost_type=_int(power, "PowerType"),
            spell_level=_int(levels.get(spell_id), "SpellLevel"),
            family_mask=[int(row[f"SpellClassMask_{index}"]) for index in _FAMILY_MASK_COLUMNS],
            effects=[
                SpellEffectConstant(
                    index=int(effect["EffectIndex"]),
                    effect=int(effect["Effect"]),
                    aura=int(effect["EffectAura"]),
                    amount=effect_amount(effect),
                    sp_coefficient=float(effect.get("EffectBonusCoefficient") or 0),
                    ap_coefficient=float(effect.get("BonusCoefficientFromAP") or 0),
                    period_ms=_int(effect, "EffectAuraPeriod"),
                    misc_value=_int(effect, "EffectMiscValue_0"),
                    trigger_spell=_int(effect, "EffectTriggerSpell"),
                )
                for effect in sorted(effects.get(spell_id, []), key=lambda e: int(e["EffectIndex"]))
            ],
        )
    return [
        ClassSpellConstants(
            build=build,
            class_slug=slug,
            family=SPELL_FAMILY_BY_CLASS_SLUG[slug],
            spells=dict(sorted(spells.items(), key=lambda pair: int(pair[0]))),
        )
        for slug, spells in sorted(spells_by_slug.items())
    ]


def write_spell_constants(build: str, root: Path = Path("builds")) -> Path:
    build_dir = root / build
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    out = build_dir / SPELLCONST
    # Rebuilt from scratch so a class that disappears between builds leaves no
    # stale file behind, exactly as normalize does for talents/ and items/.
    shutil.rmtree(out, ignore_errors=True)
    total = 0
    for record in build_spell_constants(build, raw):
        write_model(record, out / f"{record.class_slug}.json")
        total += len(record.spells)
    refresh_manifest(build_dir)
    logger.info("wrote %s: %d spells over %d classes", out, total, len(SPELL_FAMILY_BY_CLASS_SLUG))
    return out
