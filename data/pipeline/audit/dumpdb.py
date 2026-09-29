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
