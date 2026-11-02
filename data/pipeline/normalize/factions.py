"""The client's reputation factions: each one's name, and the side that earns it.

An item's reputation requirement (ItemSparse.MinFactionID) names a faction by
id only. The Faction table gives the name, and its race masks give the side.
Up to four entries each pair a 64-bit ReputationRaceMasks<i> (wago exports it
as two signed 32-bit halves, _0 low and _1 high) with a ReputationBase_<i>.
The side whose races start at Hated (-42000), and can never move, does not
earn it. Darkspear Raiders 2798 is earned by the Horde races, and Theramore
Expeditionary Force 2799 by the Alliance races. A reputation that both sides
earn, such as Timbermaw Hold, has no side.
"""

from __future__ import annotations

from collections.abc import Iterable, Mapping
from dataclasses import dataclass
from typing import Literal

Side = Literal["alliance", "horde"]

#: Faction.ReputationBase for a race that is Hated and stays Hated.
PERMANENTLY_HATED = -42000

#: The Faction table's race-mask entries.
RACE_MASK_ENTRIES = range(4)

#: ChrRaces.Alliance: 0 for an Alliance race, 1 for a Horde race (2 neutral).
SIDE_BY_CHR_RACES_ALLIANCE: dict[int, Side] = {0: "alliance", 1: "horde"}

_HALF = 32
_HALF_MASK = (1 << _HALF) - 1


@dataclass(frozen=True)
class FactionInfo:
    name: str
    side: Side | None


def _race_mask(row: Mapping[str, str], entry: int) -> int:
    low = int(row.get(f"ReputationRaceMasks{entry}_0") or 0) & _HALF_MASK
    high = int(row.get(f"ReputationRaceMasks{entry}_1") or 0) & _HALF_MASK
    return high << _HALF | low


def race_bit_sides(race_rows: Iterable[Mapping[str, str]]) -> dict[int, Side]:
    """A playable race's mask bit (ChrRaces.PlayableRaceBit, not its id: Dark
    Iron Dwarf 34 is bit 11) -> its side. A race with no bit is not playable."""
    out: dict[int, Side] = {}
    for row in race_rows:
        bit = int(row.get("PlayableRaceBit") or -1)
        side = SIDE_BY_CHR_RACES_ALLIANCE.get(int(row.get("Alliance") or -1))
        if bit >= 0 and side is not None:
            out[bit] = side
    return out


def _earning_sides(row: Mapping[str, str], bit_sides: Mapping[int, Side]) -> set[Side]:
    sides: set[Side] = set()
    for entry in RACE_MASK_ENTRIES:
        mask = _race_mask(row, entry)
        if not mask or int(row.get(f"ReputationBase_{entry}") or 0) <= PERMANENTLY_HATED:
            continue
        sides |= {side for bit, side in bit_sides.items() if mask >> bit & 1}
    return sides


def load_factions(
    faction_rows: Iterable[Mapping[str, str]] | None, race_rows: Iterable[Mapping[str, str]]
) -> dict[int, FactionInfo]:
    """Faction id -> its name and side, or `{}` for a build with no Faction table."""
    if faction_rows is None:
        return {}
    bit_sides = race_bit_sides(race_rows)
    out: dict[int, FactionInfo] = {}
    for row in faction_rows:
        sides = _earning_sides(row, bit_sides)
        side = next(iter(sides)) if len(sides) == 1 else None
        out[int(row["ID"])] = FactionInfo(name=row["Name_lang"], side=side)
    return out
