"""The build's own enchant ids over the engine fork's enchant table.

`enchants.json` starts from the fork's UI table (`gear.build_enchants`), which
is Classic Era's. Forever re-tunes enchants by pointing an old enchanting
spell at a new SpellItemEnchantment row: on 1.60.1.70790 "Enchant Gloves -
Greater Agility" (20012) teaches 8206, +10 Agility, where the fork still says
1887, +7. A picker or ranker that trusts the fork's id applies the wrong
enchant. So every fork row whose spell the client teaches through
SPELL_EFFECT_ENCHANT_ITEM (SpellEffect.Effect 53) takes the client's id, and
its stats are the client row's, read through the same `build_sim_enchants`
simdb.bin carries.

Forever also adds enchants the fork has never heard of (Enchant 2H Weapon -
Spellblasting, the necklace enchants). Those come from
`data/curated/client-enchants.json`, which names each one by its teaching
spell with the evidence for where a player gets it; everything else about
the row (name, icon, slots, shape, stats) is read from the client here.
"""

from __future__ import annotations

import json
from collections.abc import Callable, Iterable, Mapping
from dataclasses import dataclass
from pathlib import Path

from pipeline.csvio import read_csv
from pipeline.forkdb import ForkDatabase
from pipeline.loot.gear import (
    EquipRestriction,
    build_enchants,
    enchant_display_names,
    equip_restrictions,
)
from pipeline.models import EnchantRecord
from pipeline.simdb.enchants import (
    build_sim_enchants,
    melee_weapon_only_spells,
    ranged_weapon_only_spells,
)
from pipeline.simdb.equip import index_spell_effects
from pipeline.simdb.ratings import load_rating_factors
from pipeline.simdb.statmap import stat_keys

#: SpellEffect.Effect SPELL_EFFECT_ENCHANT_ITEM: EffectMiscValue_0 is the
#: SpellItemEnchantment id the spell applies.
EFFECT_ENCHANT_ITEM = "53"

CURATED_FILE = "client-enchants.json"

#: `buffs.client_tables`' icon resolver: (item id, spell id) -> icon name.
IconFor = Callable[[int, int], str]

#: SpellEquippedItems' EquippedItemClass values.
ITEM_CLASS_WEAPON = 2
ITEM_CLASS_ARMOR = 4

#: The weapon subclasses that are two-handed: axe 1, mace 5, polearm 6,
#: sword 8, staff 10. A weapon enchant whose mask stays inside these is a
#: two-hand enchant.
TWO_HAND_SUBCLASS_MASK = (1 << 1) | (1 << 5) | (1 << 6) | (1 << 8) | (1 << 10)
#: Armor subclass 6 is the shield.
SHIELD_SUBCLASS_BIT = 1 << 6

#: InventoryType bit -> the planner's slot names, for armor enchants.
SLOTS_BY_INVENTORY_BIT: dict[int, tuple[str, ...]] = {
    1: ("head",),
    2: ("neck",),
    3: ("shoulder",),
    5: ("chest",),
    7: ("legs",),
    8: ("feet",),
    9: ("wrist",),
    10: ("hands",),
    14: ("main_hand", "off_hand"),
    16: ("back",),
    20: ("chest",),
    23: ("off_hand",),
}
SHIELD_INVENTORY_BIT = 14
WEAPON_SLOTS = ("main_hand", "off_hand")


class ClientEnchantError(SystemExit):
    """The client does not say what the curated list or the fork claims."""


@dataclass(frozen=True)
class CuratedEnchant:
    """One `client-enchants.json` row: an enchant the fork lacks."""

    spell_id: int
    #: The kit or arcanum item that casts the spell; 0 for an enchanter's spell.
    item_id: int


def taught_enchants(spell_effect_rows: Iterable[Mapping[str, str]]) -> dict[int, int]:
    """Spell id -> the SpellItemEnchantment id it applies (Effect 53)."""
    return {
        int(row["SpellID"]): int(row["EffectMiscValue_0"])
        for row in spell_effect_rows
        if row["Effect"] == EFFECT_ENCHANT_ITEM and int(row["EffectMiscValue_0"])
    }


def load_curated(curated_dir: Path) -> list[CuratedEnchant]:
    path = curated_dir / CURATED_FILE
    document = json.loads(path.read_text(encoding="utf-8"))
    return [
        CuratedEnchant(spell_id=int(row["spell_id"]), item_id=int(row.get("item_id", 0)))
        for row in document["enchants"]
    ]


def client_stats(
    enchant_rows: list[dict[str, str]],
    effects_by_spell: dict[int, list[dict[str, str]]],
    rating_factors: dict[str, float],
    ranged_weapon_spells: frozenset[int] = frozenset(),
    melee_weapon_spells: frozenset[int] = frozenset(),
) -> dict[int, dict[str, float]]:
    """Enchant id -> its planner stats, exactly what simdb.bin's row carries."""
    return {
        row.effect_id: stat_keys(row.stats)
        for row in build_sim_enchants(
            enchant_rows,
            effects_by_spell,
            rating_factors,
            ranged_weapon_spells,
            melee_weapon_spells,
        )
    }


def rebind_to_client(
    records: Iterable[EnchantRecord],
    taught: Mapping[int, int],
    stats: Mapping[int, dict[str, float]],
) -> list[EnchantRecord]:
    """Each fork row with the id its spell applies in this client, and that
    row's stats. A row whose spell the client does not teach keeps the
    fork's id. A taught id the client has no row for is refused."""
    out: list[EnchantRecord] = []
    for record in records:
        client_id = taught.get(record.spell_id)
        if client_id is None or client_id == record.id:
            out.append(record)
            continue
        if client_id not in stats:
            raise ClientEnchantError(
                f"spell {record.spell_id} ({record.name}) applies enchant {client_id}, "
                f"which SpellItemEnchantment does not carry"
            )
        out.append(record.model_copy(update={"id": client_id, "stats": stats[client_id]}))
    return out


def _slots_and_type(rule: EquipRestriction, item_id: int) -> tuple[list[str], str]:
    """Where an enchant goes and its EnchantType word, from its spell's
    SpellEquippedItems row."""
    if rule.item_class == ITEM_CLASS_WEAPON:
        two_hand = rule.subclass_mask and not rule.subclass_mask & ~TWO_HAND_SUBCLASS_MASK
        return list(WEAPON_SLOTS), "two_hand" if two_hand else "normal"
    if rule.item_class != ITEM_CLASS_ARMOR:
        raise ClientEnchantError(f"equip class {rule.item_class} is not a weapon or armor")
    if (
        rule.subclass_mask == SHIELD_SUBCLASS_BIT
        or rule.inventory_type_mask == 1 << SHIELD_INVENTORY_BIT
    ):
        return list(WEAPON_SLOTS), "shield"
    slots = sorted(
        {
            slot
            for bit, names in SLOTS_BY_INVENTORY_BIT.items()
            if rule.inventory_type_mask & (1 << bit)
            for slot in names
        }
    )
    if not slots:
        raise ClientEnchantError(f"inventory type mask {rule.inventory_type_mask} names no slot")
    return slots, "kit" if item_id else "normal"


@dataclass(frozen=True)
class ClientTables:
    """The client rows `added_records` reads."""

    taught: Mapping[int, int]
    stats: Mapping[int, dict[str, float]]
    spell_names: Mapping[int, str]
    restrictions: Mapping[int, EquipRestriction]
    #: Spell id -> the items whose effects cast it (ItemXItemEffect).
    items_by_spell: Mapping[int, frozenset[int]]


def items_by_spell(
    item_effect_rows: Iterable[Mapping[str, str]], link_rows: Iterable[Mapping[str, str]]
) -> dict[int, frozenset[int]]:
    spell_by_effect = {int(row["ID"]): int(row["SpellID"]) for row in item_effect_rows}
    out: dict[int, set[int]] = {}
    for row in link_rows:
        spell = spell_by_effect.get(int(row["ItemEffectID"]))
        if spell:
            out.setdefault(spell, set()).add(int(row["ItemID"]))
    return {spell: frozenset(items) for spell, items in out.items()}


def added_records(
    curated: Iterable[CuratedEnchant],
    tables: ClientTables,
    icon_for: IconFor,
) -> list[EnchantRecord]:
    """The curated enchants the fork lacks, every column from the client.
    `icon_for(item_id, spell_id)` is `buffs.client_tables`' resolver."""
    out: list[EnchantRecord] = []
    for entry in curated:
        spell = entry.spell_id
        if spell not in tables.taught:
            raise ClientEnchantError(f"curated spell {spell} applies no enchant (SpellEffect 53)")
        if spell not in tables.restrictions:
            raise ClientEnchantError(f"curated spell {spell} has no SpellEquippedItems row")
        if entry.item_id and entry.item_id not in tables.items_by_spell.get(spell, frozenset()):
            raise ClientEnchantError(f"item {entry.item_id} does not cast curated spell {spell}")
        rule = tables.restrictions[spell]
        slots, shape = _slots_and_type(rule, entry.item_id)
        enchant_id = tables.taught[spell]
        out.append(
            EnchantRecord(
                id=enchant_id,
                name=tables.spell_names[spell],
                icon=icon_for(entry.item_id, spell),
                slots=slots,
                item_types=[shape],
                classes=[],
                stats=tables.stats[enchant_id],
                phase=0,
                spell_id=spell,
                item_id=entry.item_id,
                equipped_item_class=rule.item_class,
                equipped_item_subclass_mask=rule.subclass_mask,
                equipped_inventory_type_mask=rule.inventory_type_mask,
            )
        )
    return out


def with_client_enchants(
    fork_records: Iterable[EnchantRecord],
    curated: Iterable[CuratedEnchant],
    tables: ClientTables,
    icon_for: IconFor,
) -> list[EnchantRecord]:
    """The fork's rows rebound to the client's ids, plus the curated
    additions, in `build_enchants`' order. A curated spell the fork already
    carries is refused: the list is for what the fork lacks."""
    rebound = rebind_to_client(fork_records, tables.taught, tables.stats)
    known_spells = {record.spell_id for record in rebound}
    curated = list(curated)
    duplicate = sorted(entry.spell_id for entry in curated if entry.spell_id in known_spells)
    if duplicate:
        raise ClientEnchantError(f"curated spells {duplicate} are already in the fork's table")
    merged = [*rebound, *added_records(curated, tables, icon_for)]
    return sorted(merged, key=lambda record: (record.id, record.spell_id, record.item_id))


def build_enchant_table(
    build_dir: Path, fork: ForkDatabase, curated_dir: Path, icon_for: IconFor
) -> list[EnchantRecord]:
    """`enchants.json`: the fork's table over this build's own client rows
    (`with_client_enchants`), from `build_dir/raw`."""
    raw = build_dir / "raw"
    equipped_rows = read_csv(raw / "SpellEquippedItems.csv")
    restrictions = equip_restrictions(equipped_rows)
    effect_rows = read_csv(raw / "SpellEffect.csv")
    enchant_rows = read_csv(raw / "SpellItemEnchantment.csv")
    tables = ClientTables(
        taught=taught_enchants(effect_rows),
        stats=client_stats(
            enchant_rows,
            index_spell_effects(effect_rows),
            load_rating_factors(build_dir),
            frozenset(ranged_weapon_only_spells(equipped_rows)),
            frozenset(melee_weapon_only_spells(equipped_rows)),
        ),
        spell_names={int(row["ID"]): row["Name_lang"] for row in read_csv(raw / "SpellName.csv")},
        restrictions=restrictions,
        items_by_spell=items_by_spell(
            read_csv(raw / "ItemEffect.csv"), read_csv(raw / "ItemXItemEffect.csv")
        ),
    )
    displays = enchant_display_names(enchant_rows)
    merged = with_client_enchants(
        build_enchants(fork, restrictions, displays), load_curated(curated_dir), tables, icon_for
    )
    # A rebound or added record carries the client's id, so its text is read by that id.
    return [record.model_copy(update={"display": displays.get(record.id, "")}) for record in merged]
