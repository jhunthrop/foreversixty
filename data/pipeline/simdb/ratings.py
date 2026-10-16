"""Combat-rating conversion for simdb's item- and enchant-sourced stats.

Forever's 1.60 client itemises hit, crit, dodge, parry, block and defense
through the full retail `ItemModType` vocabulary -- `StatModifier_bonusStat_31`
is literally `ITEM_MOD_HIT_RATING`, 32 is `ITEM_MOD_CRIT_RATING`, and 12-15
are `ITEM_MOD_DEFENSE_SKILL_RATING`/`DODGE_RATING`/`PARRY_RATING`/
`BLOCK_RATING` -- rather than the flat percentages Classic states directly.
Lionheart Helm's 20 hit / 28 crit (item 12640) only means "2% hit, 2% crit"
once divided by `gametables/combatratings.txt`'s level-60 row: 20 rating
points buy 1% hit, so 20 / 10 = 2%, and 28 rating points buy 1% crit, so
28 / 14 = 2%.

The engine has no combat-rating system of its own (`sim/core/stats` reads
`SimItem.stats` and `SimEnchant.stats` hit/crit/dodge/parry/block/defense as
flat percentages; Forever's ruling is that the engine's own rating constants
are identity), so this conversion belongs entirely in the data pipeline, and
only on simdb's path: the planner's `items/<class-slug>.json` keeps the
rating number `normalize/gear.py`'s `resolve_item_values` already resolves,
matching what the client's own tooltip shows.

What is *not* converted, and why:

* An on-equip spell's flat stat (`pipeline.simdb.equip.STAT_AURAS`, e.g. aura
  54 "chance to hit") states its amount as a literal percentage already --
  Classic's older, pre-rating itemisation convention -- not a rating, so
  `equip.spell_bonus`'s output never passes through here.
* Haste (`ITEM_MOD_HASTE_RATING`, id 36) and expertise/armor penetration
  (37, 44) are mapped to `None` in `gear.STAT_BY_MODIFIER_ID`: the planner has
  no stat for them and simdb never sees them, rating or not.
* `spell_hit`/`spell_crit` have no `StatModifier_bonusStat` id of their own --
  the unified `hit`/`crit` ids (31/32) cover every school -- so they are only
  ever equip-spell sourced, and fall under the first bullet.
* Block *value* (`STAT_AURAS` 274/564, the engine's `StatBlockValue`) is a
  flat absorption amount, not a chance, and has no rating column here either.
"""

from __future__ import annotations

from collections.abc import Mapping
from pathlib import Path

from pipeline.gametables import parse_game_table

LEVEL_60 = "60"

#: Engine-bound stat key (as `pipeline.normalize.gear.STAT_BY_MODIFIER_ID`
#: names it) -> the `combatratings.txt` column(s) that convert it from rating
#: points to a percentage at level 60. `hit` and `crit` list all three
#: schools because Forever's engine unifies them into one `Stat` (see
#: `pipeline/simdb/statmap.py`); `load_rating_factors` requires every column
#: for a key to agree, so a build where a school's factor actually diverged
#: fails loudly instead of silently picking one.
RATING_STAT_COLUMNS: dict[str, tuple[str, ...]] = {
    "hit": ("Hit - Melee", "Hit - Ranged", "Hit - Spell"),
    "crit": ("Crit - Melee", "Crit - Ranged", "Crit - Spell"),
    "dodge": ("Dodge",),
    "parry": ("Parry",),
    "block": ("Block",),
    "defense": ("Defense Skill",),
}


class RatingFactorError(ValueError):
    """`combatratings.txt` does not have what simdb needs from it."""


def load_rating_factors(build_dir: Path) -> dict[str, float]:
    """Level-60 rating points per 1%, for every key in `RATING_STAT_COLUMNS`.

    Divide a `StatModifier_bonusStat`-sourced amount by `factors[key]` to turn
    a combat-rating amount into the flat percentage `SimItem`/`SimEnchant`
    stats are. Never hardcode these numbers in the pipeline -- read them from
    the build's own `gametables/combatratings.txt` so a build whose ratings
    differ (or, per the docstring above, actually diverge by school) is
    caught rather than silently mis-simmed.
    """
    path = build_dir / "gametables" / "combatratings.txt"
    if not path.exists():
        raise RatingFactorError(
            f"no {path}; run `python -m pipeline gametables` for this build first"
        )
    table = parse_game_table("combatratings.txt", path.read_text(encoding="utf-8"))
    if LEVEL_60 not in table.rows:
        raise RatingFactorError(f"{path} has no level 60 row")
    row = dict(zip(table.columns[1:], table.rows[LEVEL_60], strict=True))
    factors: dict[str, float] = {}
    for key, columns in RATING_STAT_COLUMNS.items():
        values = {float(row[column]) for column in columns}
        if len(values) != 1:
            raise RatingFactorError(
                f"combatratings.txt level 60 {columns} disagree ({sorted(values)}) "
                f"for {key!r}; the engine's unified stat needs one factor, not "
                f"per-column ones"
            )
        (factor,) = values
        if factor <= 0:
            raise RatingFactorError(
                f"combatratings.txt level 60 {key!r} factor is {factor}, not positive"
            )
        factors[key] = factor
    return factors


def convert_rating_stats(
    stats: Mapping[str, float], factors: Mapping[str, float]
) -> dict[str, float]:
    """`stats`, with every rating-denominated key divided by its factor.

    Returns a new mapping rather than mutating `stats`: the same dict this
    package builds from `resolve_item_values`/`STAT_BY_MODIFIER_ID` may still
    be read elsewhere (the planner's own output) exactly as the client states
    it.
    """
    return {
        key: (amount / factors[key] if key in factors else amount)
        for key, amount in stats.items()
    }
