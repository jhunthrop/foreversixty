"""Direct reads of the pinned cmangos/classic-db mysqldump for the handful
of tables `pipeline.classic_sources` does not itself expose (it only
persists per-item `ClassicDbSourceRecord` rows, not the raw tables it
parsed them from).

Optional: checks C, D and E ask for a boss npc's own spawn map
(`creature`/`creature_template`), a vendor item's `RequiredReputationFaction`
/`RequiredReputationRank` and a crafted item's recipe spell (`item_template`
+ `spell_template`'s `EffectItemType`), none of which the committed
`raw/classicdb/sources.json` cache carries. Pass `--classicdb-dump` a path
to the same dump (`.sql` or `.sql.gz`) `fetch-classic-sources` downloads --
the lane brief names a local copy for this run -- and this module reads it
directly, read-only, no network. Without one, the three checks that need it
say so in their own `CategoryResult.skipped` rather than guessing.
"""

from __future__ import annotations

import functools
import gzip
from collections import Counter, defaultdict
from pathlib import Path

from pipeline.sqldump import iter_table_records, unquote

#: cmangos' own SPELL_EFFECT_CREATE_ITEM id (Classic 1.12 effect list).
SPELL_EFFECT_CREATE_ITEM = 24

#: cmangos' own CONDITION_REPUTATION_RANK (matches
#: `pipeline.classic_sources._CONDITION_REPUTATION_RANK`; duplicated here
#: rather than imported since that name is private to that module).
CONDITION_REPUTATION_RANK = 5


def _read_text(path: Path) -> str:
    data = path.read_bytes()
    if path.suffix == ".gz":
        return gzip.decompress(data).decode("utf-8", errors="replace")
    return data.decode("utf-8", errors="replace")


class ClassicDbDump:
    """Read-only, lazily-indexed view of the dump's `creature_template`,
    `creature`, `item_template` and `spell_template` tables."""

    def __init__(self, path: Path) -> None:
        self.path = path

    @classmethod
    def from_text(cls, text: str) -> ClassicDbDump:
        """Build a view directly from already-read SQL text, for a caller
        that just downloaded it over the network and has no reason to write
        it to disk first only to read it straight back
        (`pipeline.classicdb_items.extract_records`, the same dump
        `pipeline.classic_sources.fetch_classic_db_sources` already
        downloads for its own tables)."""
        dump = cls.__new__(cls)
        dump.path = None
        dump.__dict__["_text"] = text
        return dump

    @functools.cached_property
    def _text(self) -> str:
        return _read_text(self.path)

    @functools.cached_property
    def creature_names(self) -> dict[int, str]:
        return {
            int(row["Entry"]): unquote(row["Name"]) or ""
            for row in iter_table_records(self._text, "creature_template")
        }

    @functools.cached_property
    def creature_spawn_map(self) -> dict[int, int]:
        """npc entry id -> the map id it spawns on most often -- same
        `Counter.most_common` rule `pipeline.classic_sources.
        _spawn_map_by_entry` uses, so an open-world npc that also has a
        rare instance-bound copy still resolves to its usual map."""
        by_entry: dict[int, Counter[int]] = defaultdict(Counter)
        for row in iter_table_records(self._text, "creature"):
            by_entry[int(row["id"])][int(row["map"])] += 1
        return {entry: counts.most_common(1)[0][0] for entry, counts in by_entry.items()}

    @functools.cached_property
    def item_template(self) -> dict[int, dict]:
        """item id -> `{required_reputation_faction, required_reputation_rank,
        spell_ids}` -- `spell_ids` is `spellid_1..5`, non-zero only, the
        candidate on-use spells a recipe item's own "Learn Recipe" spell
        would be one of."""
        out: dict[int, dict] = {}
        for row in iter_table_records(self._text, "item_template"):
            spell_ids = [
                int(row[f"spellid_{n}"]) for n in range(1, 6) if int(row[f"spellid_{n}"])
            ]
            out[int(row["entry"])] = {
                "required_reputation_faction": int(row["RequiredReputationFaction"]),
                "required_reputation_rank": int(row["RequiredReputationRank"]),
                "spell_ids": spell_ids,
            }
        return out

    @functools.cached_property
    def equippable_item_template_rows(self) -> list[dict[str, str]]:
        """Every `item_template` row that is real equippable gear: Item.ClassID
        2 (WEAPON) or 4 (ARMOR) with a real equip slot (`InventoryType != 0`)
        -- the raw dump's own equippable universe, kept as raw string rows
        (this dump's usual convention: a caller reads the columns it needs
        with its own `int()`/`unquote()` calls, same as every other
        row-shaped input in this pipeline).

        `pipeline.classicdb_items` is the one caller (catalogue-universe
        lane, 2026-09-30): it turns this into the committed
        `raw/classicdb/item_template.json` extract that fills the ids the
        client's `ItemSparse`/hotfix cache never carried at all -- 1,960 of
        them on build 1.60.1.70009, uncommon-or-better, including Hand of
        Justice (11815) and several ZG/AQ trinkets whose whole value is an
        on-equip/on-use spell effect, not a flat stat. No quality filter
        here: this is the raw universe
        `pipeline.classicdb_items.extract_records` narrows with the
        planner's own `PLANNER_QUALITIES` gate, the same way every other
        source's raw rows are narrowed downstream, not at the reader.
        """
        return [
            row
            for row in iter_table_records(self._text, "item_template")
            if int(row["class"]) in (2, 4) and int(row["InventoryType"]) != 0
        ]

    @functools.cached_property
    def spell_names(self) -> dict[int, str]:
        """spell id -> `spell_template.SpellName` (enUS), the internal
        cmangos label for the spell -- NOT player-facing tooltip text
        (compare "Increase Spell Dam 29" to the client's own resolved
        "Equip: Increases damage and healing done by magical spells and
        effects by up to 29"). Used only as `pipeline.classicdb_items`'
        last-resort fallback for a spell id the Forever client's own
        `Spell.csv` has no row for at all -- the client carries virtually
        every Classic-era spell, so this rarely fires; see that module's
        own doc.
        """
        return {
            int(row["Id"]): unquote(row["SpellName"]) or ""
            for row in iter_table_records(self._text, "spell_template")
        }

    @functools.cached_property
    def spell_effects(self) -> dict[int, dict]:
        """spell id -> `{proc_chance, effects: [{effect, aura, base_points,
        die_sides, misc_value, trigger_spell}, ...]}`, classic-db's own
        `spell_template` (1.12) columns, narrowed to the three effect slots
        that table carries (`Effect1..3`/`EffectApplyAuraName1..3`/
        `EffectBasePoints1..3`/`EffectDieSides1..3`/`EffectMiscValue1..3`/
        `EffectTriggerSpell1..3`).

        `pipeline.classicdb_items` is the one caller (classicdb-fidelity
        lane, 2026-09-30): a classic-db item's effect text and structured
        `stats` come from THIS table first, not the Forever client's own
        `Spell.csv` -- 1.12 spell ids are not stable across clients (see
        that module's `effect_text`/`equip_effects` own docs for the
        Devilsaur Eye/Hand of Justice defects this closes). Every effect's
        die-sides convention matches `pipeline.spelltext.effect_amount`'s
        own doc for Classic Era: the base amount is `EffectBasePoints + 1`
        when `EffectDieSides` is 1 (verified against Blackhand's Breadth's
        real +2% crit: EffectBasePoints1 1, EffectDieSides1 1).
        """
        out: dict[int, dict] = {}
        for row in iter_table_records(self._text, "spell_template"):
            effects = []
            for n in (1, 2, 3):
                effect = int(row[f"Effect{n}"])
                if effect == 0:
                    continue
                effects.append(
                    {
                        "effect": effect,
                        "aura": int(row[f"EffectApplyAuraName{n}"]),
                        "base_points": int(row[f"EffectBasePoints{n}"]),
                        "die_sides": int(row[f"EffectDieSides{n}"]),
                        "misc_value": int(row[f"EffectMiscValue{n}"]),
                        "trigger_spell": int(row[f"EffectTriggerSpell{n}"]),
                    }
                )
            out[int(row["Id"])] = {
                "proc_chance": int(row["ProcChance"]),
                "effects": effects,
            }
        return out

    @functools.cached_property
    def created_item_to_spells(self) -> dict[int, list[int]]:
        """The item a crafting spell's own `SPELL_EFFECT_CREATE_ITEM`
        effect produces -> every spell id that creates it (almost always
        one, but cmangos does carry a rare duplicate spell id for the
        same recipe on two ranks/patch variants)."""
        out: dict[int, list[int]] = defaultdict(list)
        for row in iter_table_records(self._text, "spell_template"):
            spell_id = int(row["Id"])
            for n in (1, 2, 3):
                if int(row[f"Effect{n}"]) != SPELL_EFFECT_CREATE_ITEM:
                    continue
                created = int(row[f"EffectItemType{n}"])
                if created:
                    out[created].append(spell_id)
        return dict(out)


def load_dump(path: Path | None) -> ClassicDbDump | None:
    if path is None or not path.exists():
        return None
    return ClassicDbDump(path)
