"""What a trait talent's effects are worth at each of its ranks.

The 1.60 client has no per-rank spells. A talent is one `TraitDefinition`
pointing at one spell, and `TraitDefinitionEffectPoints` says which curve
drives each of that spell's effects; the curve's points are the ranks, with
`Pos_0` the rank and `Pos_1` the value. Measured on build 1.60.1.69893: 640
rows, `OperationType` 0 in every one, `Pos_0` always 1..5, every `Pos_1` a
whole number, and every multi-rank talent covered.

`pipeline/normalize/traits.py` hands the values for one rank to
`SpellText.describe` as effect overrides, which is how a rank gets its own
sentence out of the one description the client stores.
"""

from __future__ import annotations

from dataclasses import dataclass


class CurveDataError(ValueError):
    """The curve tables do not have the shape per-rank values need."""


@dataclass(frozen=True)
class RankPoints:
    """definition id -> rank -> 0-based effect index -> the effect's value."""

    by_definition: dict[int, dict[int, dict[int, int]]]

    def for_rank(self, definition_id: int, rank: int) -> dict[int, int]:
        """The effect values at `rank`, or {} when this definition has no curve."""
        return self.by_definition.get(definition_id, {}).get(rank, {})

    def highest_rank(self, definition_id: int) -> int:
        """The last rank any of this definition's curves names; 0 when it has none."""
        ranks = self.by_definition.get(definition_id)
        return max(ranks) if ranks else 0


def load_rank_points(
    effect_point_rows: list[dict[str, str]], curve_point_rows: list[dict[str, str]]
) -> RankPoints:
    curves: dict[int, dict[int, int]] = {}
    for row in curve_point_rows:
        rank = round(float(row["Pos_0"]))
        curves.setdefault(int(row["CurveID"]), {})[rank] = round(float(row["Pos_1"]))
    by_definition: dict[int, dict[int, dict[int, int]]] = {}
    for row in effect_point_rows:
        # 0 is "this curve gives the effect's value". Anything else would mean
        # the value is computed from the spell's own base points in a way this
        # module does not implement, and quietly emitting the base value would
        # put a wrong number in a tooltip.
        if row["OperationType"] != "0":
            raise CurveDataError(
                f"TraitDefinitionEffectPoints {row['ID']} has OperationType "
                f"{row['OperationType']}; only 0 (set the value) is understood"
            )
        curve_id = int(row["CurveID"])
        curve = curves.get(curve_id)
        if not curve:
            raise CurveDataError(
                f"TraitDefinitionEffectPoints {row['ID']} names curve {curve_id}, "
                "which has no points"
            )
        index = int(row["EffectIndex"])
        ranks = by_definition.setdefault(int(row["TraitDefinitionID"]), {})
        for rank, value in curve.items():
            ranks.setdefault(rank, {})[index] = value
    return RankPoints(by_definition)
