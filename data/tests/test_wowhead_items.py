# data/tests/test_wowhead_items.py
"""Wowhead's Forever gear planner as the source for items the client's ItemSparse lacks.

The fixture is six real-shaped records with the payload's own quirks (trailing commas,
a second setPageData call after the items, mixed-case icon names): a cloak the client
lacks, a one-hand sword with damage, a set id and a class mask, a quality-1 ring, a belt
the client does carry, a QA-named epic and a level-70 chest.
"""

from collections import Counter
from pathlib import Path

import httpx
import pytest

from pipeline import wowhead_items as wh
from pipeline.models import GearItem, Item

FIXTURE = Path(__file__).parent / "fixtures" / "wowhead-gear-planner.js"


def test_parses_the_item_object_past_trailing_commas_and_later_page_data() -> None:
    items = wh.load_items(FIXTURE)
    assert [item.id for item in items] == [11726, 264908, 271218, 279865, 777777, 888888]
    cloak = next(item for item in items if item.id == 279865)
    assert cloak.name == "Grave Shroud"
    assert cloak.icon == "inv_misc_cape_10"  # lowercased, the site's own vocabulary
    assert cloak.required_level == 16  # requiredLevel over stats.reqlevel
    assert (cloak.class_id, cloak.subclass_id) == (4, 1)  # wowhead's -6 is the client's cloth
    assert cloak.armor == 20 and cloak.class_mask is None


def test_refuses_a_payload_that_is_not_forever_data(tmp_path: Path) -> None:
    era = FIXTURE.read_text().replace('"versionNum":16001', '"versionNum":11300')
    path = tmp_path / "era.js"
    path.write_text(era)
    with pytest.raises(wh.WowheadPayloadError, match="11300"):
        wh.load_items(path)
    with pytest.raises(wh.WowheadPayloadError, match="no setPageData"):
        wh.parse_page_data('WH.setPageData("other", {})', wh.ITEM_PAGE_DATA)


def test_supplement_is_the_planner_gear_the_client_lacks() -> None:
    items = wh.load_items(FIXTURE)
    picked = wh.supplement(items, client_ids={11726, 264908})
    # The belt is the client's; the heirloom ring is quality 1; the QA robe is junk-named;
    # the hauberk needs level 70. The cloak and the sword remain.
    assert [item.id for item in picked] == [271218, 279865]


def test_stats_map_to_the_planner_vocabulary_and_count_what_it_does_not_track() -> None:
    sword = next(item for item in wh.load_items(FIXTURE) if item.id == 271218)
    untracked: Counter[str] = Counter()
    assert wh.planner_stats(sword, untracked) == {"agility": 2, "strength": 4, "crit": 3}
    assert untracked == {"hastertng": 1}


def test_gear_item_carries_weapon_damage_set_and_uniqueness() -> None:
    sword = next(item for item in wh.load_items(FIXTURE) if item.id == 271218)
    gear = wh.to_gear_item(sword)
    assert isinstance(gear, GearItem)
    assert (gear.slot, gear.damage_min, gear.damage_max, gear.speed, gear.dps) == (
        "main_hand",
        26,
        50,
        2.6,
        14.62,
    )
    assert gear.set_id == 9001 and gear.unique is True and gear.two_hand is False
    flat = wh.to_item(sword)
    assert isinstance(flat, Item)
    assert (flat.class_id, flat.subclass_id, flat.inventory_type) == (2, 7, 13)


def test_a_holdable_off_hand_item_wowhead_states_a_speed_for_gets_no_weapon_fields() -> None:
    """Antipodean Rod (2879) and Orb of Mistmantle (13031) are real items:
    Item.ClassID 4 (armour, not a weapon), InventoryType 23 HOLDABLE -- but
    wowhead's own scrape states a nonzero "speed" stat for both (1.6, no
    damage at all) anyway. to_gear_item used to copy every wowhead weapon
    field at face value; this is the wowhead-side twin of the client-row
    Father Flame defect (13371, also class 4, also InventoryType 23) that
    `is_weapon_row` fixes in normalize/gear.py -- the same gate belongs here
    too, so a non-weapon wowhead row never reports a swing speed.
    """
    rod = wh.WowheadItem(
        id=2879,
        name="Antipodean Rod",
        quality=3,
        item_level=22,
        required_level=17,
        class_id=4,
        subclass_id=0,
        inventory_type=23,
        icon="inv_wand_04",
        class_mask=None,
        stats={"speed": 1.6},
        set_id=None,
        unique=True,
        armor=0,
        damage_min=0,
        damage_max=0,
        speed=1.6,
        dps=0.0,
    )
    gear = wh.to_gear_item(rod)
    assert (gear.damage_min, gear.damage_max, gear.speed, gear.dps) == (0, 0, 0.0, 0.0)
    assert gear.two_hand is False


def test_class_allowed_reads_the_mask_and_the_proficiency_table() -> None:
    items = {item.id: item for item in wh.load_items(FIXTURE)}
    sword = items[271218]  # classMask 1029: warrior (1), hunter (4), rogue (1024)? bits 0, 2, 10
    assert wh.class_allowed(sword, 1) is True  # warrior
    assert wh.class_allowed(sword, 2) is False  # paladin: not in the mask
    cloak = items[279865]  # no mask: every class that can wear cloth
    assert wh.class_allowed(cloak, 8) is True  # mage
    hauberk = items[888888]  # mail
    assert wh.class_allowed(hauberk, 8) is False  # a mage cannot wear mail


def test_fetch_writes_the_payload_and_its_meta(tmp_path: Path) -> None:
    text = FIXTURE.read_text()

    def handler(request: httpx.Request) -> httpx.Response:
        assert request.url.params["dv"] == "100"
        return httpx.Response(200, text=text)

    client = httpx.Client(transport=httpx.MockTransport(handler))
    path = wh.fetch_wowhead("1.60.1.70009", root=tmp_path, client=client)
    assert path == tmp_path / "1.60.1.70009" / "raw" / wh.RAW_FILE
    assert path.read_text() == text
    meta = (path.parent / wh.META_FILE).read_text()
    assert '"items": 6' in meta and '"sha256"' in meta


def test_fetch_refuses_to_overwrite_with_era_data(tmp_path: Path) -> None:
    era = FIXTURE.read_text().replace('"versionNum":16001', '"versionNum":11300')
    client = httpx.Client(transport=httpx.MockTransport(lambda _: httpx.Response(200, text=era)))
    with pytest.raises(wh.WowheadPayloadError, match="11300"):
        wh.fetch_wowhead("1.60.1.70009", root=tmp_path, client=client)
    assert not (tmp_path / "1.60.1.70009").exists()
