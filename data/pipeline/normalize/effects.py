"""ItemEffect -> Spell -> SpellEffect: the stats and text an item's own
equip, use and proc effects add on top of ItemSparse's own columns.

An adapter over `pipeline.simdb.equip`, not a second aura table. That
module's `STAT_AURAS`/`IGNORED_AURAS` tables were measured by hand against
build 1.60.1.69893's own item set (see its docstring); keeping a second,
independently maintained aura table here risks the two disagreeing, which is
exactly the failure `equip.py`'s own docstring warns about. `EffectIndex`
therefore composes `equip.index_spell_effects`, `equip.item_effect_spells`
and `equip.spell_bonus` for the stats, and `pipeline.spelltext.SpellText`
for the text -- the same `$`-token substitution talent text already gets.

`equip.spell_bonus` returns its stats as `dict[str, float]`, plus a
`weapon_skills` mapping and a flat `bonus_physical_damage`. `GearItem.stats`
is `dict[str, int]` and has no field for either of the other two, so
`.stats()` below converts the stat amounts to `int` and drops
`weapon_skills`/`bonus_physical_damage` entirely -- an equip spell that
grants only a weapon skill or only physical damage (Forever's re-itemised
world has both, per `equip.py`'s own docstring) contributes nothing to a
`GearItem` until one of those gets its own field, and none exists yet.

Two different things happen at two different times, and it matters which:

* **Grouping** an item's ItemEffect/ItemXItemEffect rows by item id and
  trigger type touches no aura data at all -- it is done once, eagerly, in
  `__init__`, over every item either table names (not only the build's
  shipped item set), because it cannot fail: nothing here is unclassified,
  there is only rows to sort into buckets.
* **Classifying** an aura -- turning a `SpellEffect` row into a stat, via
  `equip.spell_bonus` -- happens lazily, in `.stats()`/`.text()`, and only
  for the single item asked about. `equip.item_effect_spells`'s own
  docstring explains why: checking an item outside the caller's shipped set
  fails the whole run over an aura nobody has reviewed for it.
  `build_class_items` never keeps a gamemaster/test row or an item off this
  build's own class lists, so that item's equip spell -- however strange
  its aura -- can never fail the run merely because nobody asked about it.

An aura this table does not know raises rather than being dropped, by way
of `equip.EquipEffectError` (aliased here as `UnknownAuraError`, so a
caller catching either name catches the one exception `equip.py` actually
raises): a dropped stat is an item that scores lower than it should with
nothing on screen to say why, and the gear picker's whole claim is that its
numbers are honest.
"""

from __future__ import annotations

from collections.abc import Container, Mapping, Sequence

from pipeline.simdb.equip import (
    TRIGGER_ON_EQUIP,
    EquipEffectError,
    index_spell_effects,
    item_effect_spells,
    spell_bonus,
)
from pipeline.spelltext import SpellText

#: ItemEffect.TriggerType for "Equip:", as an int for callers that want to
#: compare it directly. `equip.py`'s own `TRIGGER_ON_EQUIP` is the string
#: form `item_effect_spells` compares CSV rows against.
EQUIP_TRIGGER = int(TRIGGER_ON_EQUIP)

#: An aura `pipeline.simdb.equip` will not guess at. An alias, not a
#: subclass: `equip.py` is the module that actually raises, so a caller
#: catching `UnknownAuraError` needs it to be the very class that comes out
#: of `.stats()`, not a sibling of it.
UnknownAuraError = EquipEffectError

_EQUIP_ONLY: Container[str] = frozenset({TRIGGER_ON_EQUIP})


def _named_item_ids(
    item_effect_rows: Sequence[Mapping[str, str]],
    item_x_item_effect_rows: Sequence[Mapping[str, str]],
) -> set[int]:
    """Every item id either table names, on either schema.

    Not a filter -- `item_effect_spells` takes an `item_ids` container to
    let a caller restrict *which* items' auras get reviewed (see the module
    docstring's "classifying" bullet); grouping has no auras to review, so
    this passes every item id the raw rows themselves already carry, which
    restricts nothing.
    """
    ids = {int(row["ParentItemID"]) for row in item_effect_rows if "ParentItemID" in row}
    ids.update(int(row["ItemID"]) for row in item_x_item_effect_rows)
    return ids


class EffectIndex:
    """Every item's equip stats and effect text, grouped once and
    classified lazily per item -- see the module docstring for the
    distinction and why each half is timed the way it is.
    """

    def __init__(
        self,
        item_effect_rows: Sequence[Mapping[str, str]],
        item_x_item_effect_rows: Sequence[Mapping[str, str]],
        spell_effect_rows: Sequence[Mapping[str, str]],
        spell_text: SpellText,
    ) -> None:
        item_effect_rows = list(item_effect_rows)
        item_x_item_effect_rows = list(item_x_item_effect_rows)
        item_ids = _named_item_ids(item_effect_rows, item_x_item_effect_rows)
        self._equip_spells_by_item = item_effect_spells(
            item_effect_rows, item_x_item_effect_rows, item_ids, trigger_types=_EQUIP_ONLY
        )
        self._all_spells_by_item = item_effect_spells(
            item_effect_rows, item_x_item_effect_rows, item_ids, trigger_types=None
        )
        self._effects_by_spell = index_spell_effects(list(spell_effect_rows))
        self._spell_text = spell_text

    def stats(self, item_id: int) -> dict[str, int]:
        """The item's on-equip spells' stats, summed and rounded to `int`.

        Raises `UnknownAuraError` if an on-equip spell uses an aura
        `pipeline.simdb.equip` has not classified -- see the module
        docstring for why that is not silently dropped.
        """
        equip_spell_ids = self._equip_spells_by_item.get(item_id, [])
        if not equip_spell_ids:
            return {}
        bonus = spell_bonus(equip_spell_ids, self._effects_by_spell)
        return {key: int(round(amount)) for key, amount in bonus.stats.items()}

    def text(self, item_id: int) -> str:
        """The item's spell descriptions that are not already counted as a stat.

        An on-equip spell that produced at least one stat in `.stats()` is
        left out here -- repeating it as prose would double it on screen.
        An on-equip spell that produced *no* stat (only a weapon skill, only
        physical damage, or an aura `equip.STAT_AURAS`/`IGNORED_AURAS`
        deliberately maps to nothing) is not excluded: nothing else on the
        item represents it, so its description is the only place it shows up
        at all. A use or proc spell (any other trigger type) is never
        excluded regardless.
        """
        equip_spell_ids = self._equip_spells_by_item.get(item_id, [])
        equip_ids_with_stats = {
            spell_id
            for spell_id in equip_spell_ids
            if spell_bonus([spell_id], self._effects_by_spell).stats
        }
        parts = [
            self._spell_text.describe(spell_id)
            for spell_id in self._all_spells_by_item.get(item_id, [])
            if spell_id not in equip_ids_with_stats
        ]
        return " ".join(part for part in parts if part)
