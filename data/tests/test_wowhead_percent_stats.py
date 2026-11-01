"""Wowhead's flat-percentage stats and defense reach the planner and the sim.

On the wowhead supplement (items the client's ItemSparse lacks), wowhead
states an on-equip "+1% crit" as `critstrkpct` and "+1% hit" as `hitpct`
(likewise `dodgepct`, `parrypct`, `blockpct`), while `critstrkrtng` and
`hitrtng` are rating points. The percent keys used to vanish (15 planner
items carry critstrkpct, 23 hitpct, among them Icemail Jerkin and Might of
Hakkar). They now count in `stats` and are recorded in `percent_stats`, the
client path's own contract (`GearItem.percent_stats`), so no consumer
divides them by a rating factor. `def` is the defense an item shows
(Greatflank Flayer: wowhead 3, the client row's own `defense` 3), so it maps
to `defense` like `defrtng`. `armorbonus` was already added to `armor` on
load. `blockamount` is a shield's intrinsic block value, which no client
row carries either, so it stays unmapped.
"""

from collections import Counter

from pipeline import wowhead_items as wh
from pipeline.simdb.items import build_wowhead_sim_items
from pipeline.simproto import pb

RATING_FACTORS = {
    "hit": 10.0,
    "crit": 14.0,
    "dodge": 12.0,
    "parry": 15.0,
    "block": 5.0,
    "defense": 1.5,
}


def _ring(stats: dict[str, float]) -> wh.WowheadItem:
    return wh.WowheadItem(
        id=990001,
        name="Trial Band",
        quality=3,
        item_level=60,
        required_level=55,
        class_id=4,
        subclass_id=0,
        inventory_type=11,
        icon="inv_jewelry_ring_01",
        class_mask=None,
        stats=stats,
        set_id=None,
        unique=False,
        armor=0,
        damage_min=0,
        damage_max=0,
        speed=0.0,
        dps=0.0,
    )


def test_percent_keys_count_in_stats_and_are_recorded_as_percent() -> None:
    ring = _ring({"critstrkpct": 1.0, "hitpct": 2.0, "dodgepct": 1.0, "def": 3.0, "sta": 5.0})
    untracked: Counter[str] = Counter()
    assert wh.planner_stats(ring, untracked) == {
        "crit": 1,
        "hit": 2,
        "dodge": 1,
        "defense": 3,
        "stamina": 5,
    }
    assert untracked == {}
    assert wh.planner_percent_stats(ring) == {"crit": 1, "hit": 2, "dodge": 1}
    assert wh.to_gear_item(ring).percent_stats == {"crit": 1, "hit": 2, "dodge": 1}


def test_a_percent_stat_reaches_the_engine_undivided_and_a_rating_divided() -> None:
    ring = _ring({"critstrkpct": 1.0, "hitrtng": 10.0, "def": 3.0})
    stats = list(build_wowhead_sim_items([ring], RATING_FACTORS)[0].stats)
    assert stats[pb.Stat.Value("StatCrit")] == 1.0
    # hitrtng 10 at 10 a percent is 1%; the unified StatHit carries it.
    assert stats[pb.Stat.Value("StatHit")] == 1.0
    assert stats[pb.Stat.Value("StatDefense")] == 2.0
