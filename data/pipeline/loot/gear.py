"""The two gear tables Top Gear needs, and the column they add to items.json.

Parity contract 6.2 and 6.3. Both come from the engine fork's database
rather than from DB2:

* **Suffixes.** `ItemRandomSuffix` 404s on build 1.60.1.69893 -- which is
  why `pipeline/simdb/__init__.py` already emits `SimDatabase.random_suffixes`
  empty -- so the fork's `randomSuffixes` and its per-item
  `randomSuffixOptions` are the only statement of what "of the Bear" is
  worth and which items roll it.
* **Enchants.** The client's `SpellItemEnchantment` has 2,216 rows and no
  idea which of them a player can actually apply; the fork's `UIEnchant` is
  the curated 173 its own UI offers, with the slots, classes and phase the
  picker needs.

`stats` on both records (via `stat_keys`) uses the planner's own stat
vocabulary -- `pipeline/simdb/statmap.py`'s `PROTO_STAT_ALIASES` keys, the
same ones `items/<class-slug>.json` uses -- not contract 10.1 A7's
`reference_stat` vocabulary (`pipeline/simdb/statmap.py`'s `stat_id`, the
fork's own `proto.Stat` enum names in lower snake case). The two disagree on
spelling for at least `healing` (A7: `healing_power`) and `arcane_res` (A7:
`arcane_resistance`); a consumer that maps `enchants.json`'s keys onto
`proto.Stat` by name needs to translate through `PROTO_STAT_ALIASES`, not
assume A7's spelling.
"""

from __future__ import annotations

import json
import re
from collections.abc import Iterable, Mapping
from dataclasses import dataclass
from pathlib import Path

from pipeline.forkdb import (
    CLASS_SLUGS,
    ENCHANT_TYPES,
    FACTION_RESTRICTIONS,
    ITEM_TYPE_SLOTS,
    ForkDatabase,
    decode,
)
from pipeline.models import EnchantRecord, Item, SuffixRecord
from pipeline.normalize import write_json
from pipeline.simdb.statmap import stat_keys


def build_suffixes(fork: ForkDatabase) -> list[SuffixRecord]:
    return sorted(
        (
            SuffixRecord(
                id=int(row["id"]),
                name=row["name"],
                stats=stat_keys(row.get("stats", [])),
            )
            for row in fork.random_suffixes
        ),
        key=lambda record: record.id,
    )


def suffix_options(fork: ForkDatabase) -> dict[int, list[int]]:
    """Item id -> the suffix ids it rolls, for the items that roll any."""
    return {
        int(row["id"]): sorted(int(suffix) for suffix in row["randomSuffixOptions"])
        for row in fork.items
        if row.get("randomSuffixOptions")
    }


def faction_restrictions(fork: ForkDatabase) -> dict[int, str]:
    """Item id -> "alliance_only" or "horde_only", for the items so marked.

    The client cannot answer this on build 1.60.1.69893: 19,066 of its
    19,171 `ItemSparse` rows carry `AllowableRace` -1/-1, and every one of
    the 819 items the fork marks restricted is among them.
    """
    return {
        int(row["id"]): decode(
            FACTION_RESTRICTIONS, int(row["factionRestriction"]), "faction restriction"
        )
        for row in fork.items
        if row.get("factionRestriction")
    }


def apply_fork_columns(
    build_dir: Path,
    options: dict[int, list[int]],
    restrictions: dict[int, str],
) -> tuple[int, int]:
    """Fill `items.json`'s two fork-derived columns.

    Returns (rows with a suffix list, rows with a faction restriction).

    Read-modify-write through the `Item` model rather than through the raw
    dicts, so the file comes back out in exactly the key order and
    serialization `normalize` would have written -- re-running this on an
    already-filled build rewrites the same bytes.
    """
    path = build_dir / "items.json"
    if not path.exists():
        raise SystemExit(f"no {path}; run `python -m pipeline normalize` for this build first")
    rows = json.loads(path.read_text(encoding="utf-8"))
    items = [
        Item(**row).model_copy(
            update={
                "suffixes": options.get(int(row["id"]), []),
                "faction_restriction": restrictions.get(int(row["id"]), ""),
            }
        )
        for row in rows
    ]
    write_json(items, path)
    return (
        sum(1 for item in items if item.suffixes),
        sum(1 for item in items if item.faction_restriction),
    )


def _enchant_icon(fork: ForkDatabase, row: dict) -> str:
    """The enchant's art: its own item's icon where it has one, else its
    spell's. Every row on the pinned fork resolves through one or the
    other; a row that resolved through neither would reach the site as
    `icons/.webp`, so it is refused rather than emitted blank."""
    item_id, spell_id = int(row.get("itemId", 0)), int(row.get("spellId", 0))
    icon = fork.item_icons.get(item_id) or fork.spell_icons.get(spell_id)
    if not icon:
        raise SystemExit(
            f"enchant {row['effectId']} ({row.get('name', '?')}) has no icon in the "
            f"fork's itemIcons or spellIcons"
        )
    return icon


@dataclass(frozen=True)
class EquipRestriction:
    """A spell's SpellEquippedItems row: the items it can be cast on."""

    item_class: int
    subclass_mask: int
    inventory_type_mask: int


#: SpellEquippedItems' "no restriction" class.
ANY_ITEM_CLASS = -1


def equip_restrictions(rows: Iterable[Mapping[str, str]]) -> dict[int, EquipRestriction]:
    """SpellEquippedItems, by spell id. A spell with no row has no restriction."""
    return {
        int(row["SpellID"]): EquipRestriction(
            item_class=int(row["EquippedItemClass"]),
            subclass_mask=int(row["EquippedItemSubclass"]),
            inventory_type_mask=int(row["EquippedItemInvTypes"]),
        )
        for row in rows
    }


def with_equip_restrictions(
    records: Iterable[EnchantRecord], restrictions: Mapping[int, EquipRestriction]
) -> list[EnchantRecord]:
    """Each enchant with its spell's client restriction, so a page offers an
    enchant only for an item the client lets it go on (a weapon enchant never
    for a held-in-off-hand item)."""
    unrestricted = EquipRestriction(ANY_ITEM_CLASS, 0, 0)
    out: list[EnchantRecord] = []
    for record in records:
        rule = restrictions.get(record.spell_id, unrestricted)
        out.append(
            record.model_copy(
                update={
                    "equipped_item_class": rule.item_class,
                    "equipped_item_subclass_mask": rule.subclass_mask,
                    "equipped_inventory_type_mask": rule.inventory_type_mask,
                }
            )
        )
    return out


#: `$k<n>` in an enchant's name: effect n's points (1-based), as the client fills it.
_AMOUNT_TOKEN = re.compile(r"\$[kK]([1-3])")


def _fill_amounts(row: Mapping[str, str]) -> str | None:
    """The name with every `$k<n>` filled from `EffectPointsMin_<n-1>`, or None when a
    token has no points behind it (never print a raw token)."""
    missing = False

    def amount(match: re.Match[str]) -> str:
        nonlocal missing
        points = row.get(f"EffectPointsMin_{int(match.group(1)) - 1}", "")
        if points in ("", "0"):
            missing = True
            return match.group(0)
        return points

    filled = _AMOUNT_TOKEN.sub(amount, row["Name_lang"])
    return None if missing else filled


def enchant_display_names(rows: Iterable[Mapping[str, str]]) -> dict[int, str]:
    """`SpellItemEnchantment.Name_lang` by id: the enchant's own text, the line the
    game prints on the item ("+7 Strength"), not the recipe's name. Amount tokens are
    filled the way the client fills them; a name that cannot be filled is left out."""
    out: dict[int, str] = {}
    for row in rows:
        text = _fill_amounts(row)
        if text is not None:
            out[int(row["ID"])] = text
    return out


def build_enchants(
    fork: ForkDatabase,
    restrictions: Mapping[int, EquipRestriction],
    display_names: Mapping[int, str] | None = None,
) -> list[EnchantRecord]:
    """`restrictions` is `equip_restrictions` over the build's SpellEquippedItems;
    `display_names` is `enchant_display_names` over its SpellItemEnchantment."""
    displays = display_names or {}
    records = [
        EnchantRecord(
            id=int(row["effectId"]),
            name=row["name"],
            display=displays.get(int(row["effectId"]), ""),
            icon=_enchant_icon(fork, row),
            slots=sorted(
                {
                    slot
                    for item_type in [row["type"], *row.get("extraTypes", [])]
                    for slot in decode(ITEM_TYPE_SLOTS, int(item_type), "enchant item type")
                }
            ),
            item_types=[decode(ENCHANT_TYPES, int(row.get("enchantType", 0)), "enchant type")],
            classes=sorted(
                decode(CLASS_SLUGS, int(class_id), "enchant class")
                for class_id in row.get("classAllowlist", [])
            ),
            stats=stat_keys(row.get("stats", [])),
            phase=int(row.get("phase", 0)),
            spell_id=int(row.get("spellId", 0)),
            item_id=int(row.get("itemId", 0)),
        )
        for row in fork.enchants
    ]
    # Not `write_json`'s sort: `id` alone is not unique here, so the tie is
    # broken by the spell and then the item, which together are.
    ordered = sorted(records, key=lambda record: (record.id, record.spell_id, record.item_id))
    return with_equip_restrictions(ordered, restrictions)
