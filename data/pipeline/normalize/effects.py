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

`equip.item_effect_spells` takes an `item_ids` container specifically so a
caller can review the auras it names in advance for the item set it ships;
checking an item outside that set fails the whole run over an aura nobody
has reviewed (see its docstring). `.stats(item_id)` and `.text(item_id)`
below are therefore per-item and lazy: each call resolves the aura table
only for the single item asked about, so an item `build_class_items` never
keeps (a gamemaster or test row, or simply an item off this build's own
class lists) can never fail the run merely because *its* equip spell
happens to carry an aura nobody has classified yet.

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


class EffectIndex:
    """One item's equip stats and effect text, resolved lazily per item.

    Built once per build from the raw `ItemEffect`, `ItemXItemEffect` and
    `SpellEffect` tables and a `SpellText` for descriptions. `__init__` does
    no per-item work at all -- see the module docstring for why that is
    deferred to `.stats()`/`.text()`.
    """

    def __init__(
        self,
        item_effect_rows: Sequence[Mapping[str, str]],
        item_x_item_effect_rows: Sequence[Mapping[str, str]],
        spell_effect_rows: Sequence[Mapping[str, str]],
        spell_text: SpellText,
    ) -> None:
        self._item_effect_rows = list(item_effect_rows)
        self._item_x_item_effect_rows = list(item_x_item_effect_rows)
        self._effects_by_spell = index_spell_effects(list(spell_effect_rows))
        self._spell_text = spell_text

    def _spells_for(self, item_id: int, trigger_types: Container[str] | None) -> list[int]:
        return item_effect_spells(
            self._item_effect_rows,
            self._item_x_item_effect_rows,
            {item_id},
            trigger_types=trigger_types,
        ).get(item_id, [])

    def stats(self, item_id: int) -> dict[str, int]:
        """The item's on-equip spells' stats, summed and rounded to `int`.

        Raises `UnknownAuraError` if an on-equip spell uses an aura
        `pipeline.simdb.equip` has not classified -- see the module
        docstring for why that is not silently dropped.
        """
        equip_spell_ids = self._spells_for(item_id, _EQUIP_ONLY)
        if not equip_spell_ids:
            return {}
        bonus = spell_bonus(equip_spell_ids, self._effects_by_spell)
        return {key: int(round(amount)) for key, amount in bonus.stats.items()}

    def text(self, item_id: int) -> str:
        """The item's use/proc spells' descriptions, joined with a space.

        Equip auras are already stats (see `.stats()`); repeating them as
        prose here would double them on screen, so a spell counted there is
        left out here.
        """
        equip_spell_ids = set(self._spells_for(item_id, _EQUIP_ONLY))
        other_spell_ids = self._spells_for(item_id, None)
        parts = [
            self._spell_text.describe(spell_id)
            for spell_id in other_spell_ids
            if spell_id not in equip_spell_ids
        ]
        return " ".join(part for part in parts if part)
