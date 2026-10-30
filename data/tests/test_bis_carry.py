import json
from pathlib import Path

import pytest

from pipeline.bis_carry import NOT_IN_CLIENT_REASON, BisCarryError, carry_bis

OLD, NEW = "1.60.1.70658", "1.60.1.70685"


def _alt(item_id: int, name: str) -> dict:
    return {"item_id": item_id, "item_name": name, "source_kind": "drop", "source": "X"}


def _band(slots: list[dict]) -> dict:
    return {"spec": "mage-fire", "band": 40, "preset": "bare", "faction": "horde", "slots": slots}


def _write_build(root: Path, build: str, item_ids: list[int], bis: dict | None = None) -> None:
    build_dir = root / build
    build_dir.mkdir(parents=True)
    (build_dir / "items.json").write_text(
        json.dumps([{"id": i, "name": f"i{i}"} for i in item_ids])
    )
    if bis is not None:
        (build_dir / "bis").mkdir()
        (build_dir / "bis" / "mage-fire.json").write_text(json.dumps(bis, indent=2))
        (build_dir / "bis" / "mage-fire.md").write_text("# mage-fire\n")


def _payload(slots: list[dict]) -> dict:
    return {"spec": "mage-fire", "build": OLD, "engine_version": "abc", "bands": [_band(slots)]}


def test_a_carry_with_every_item_still_in_the_client_only_restamps_the_build(tmp_path: Path):
    slots = [{"slot": "head", "item_id": 1, "item_name": "Hood", "alternatives": [_alt(2, "Cap")]}]
    _write_build(tmp_path, OLD, [1, 2], _payload(slots))
    _write_build(tmp_path, NEW, [1, 2])

    assert carry_bis(tmp_path, OLD, NEW) == []

    carried = json.loads((tmp_path / NEW / "bis" / "mage-fire.json").read_text())
    assert carried["build"] == NEW
    assert carried["engine_version"] == "abc"
    assert carried["bands"][0]["slots"] == slots
    assert (tmp_path / NEW / "bis" / "mage-fire.md").read_text() == "# mage-fire\n"


def test_a_pick_the_new_client_lost_is_dropped_with_a_report_line(tmp_path: Path):
    """2026-10-29: 70685's items.json lost 320 items, 29 of them carried BiS picks, and
    `addon-data` stopped the move on the first one (bis item 9400)."""
    slots = [
        {
            "slot": "ranged",
            "item_id": 9400,
            "item_name": "Baelog's Shortbow",
            "source": "Baelog",
            "source_kind": "drop",
            "score": 12,
            "verified": True,
            "alternatives": [_alt(9401, "Bow"), _alt(9403, "Gone too")],
            "runners_up": [_alt(9403, "Gone too"), _alt(9401, "Bow")],
        },
        {"slot": "head", "item_id": 1, "item_name": "Hood", "verified": True},
    ]
    _write_build(tmp_path, OLD, [1, 9400, 9401, 9403], _payload(slots))
    _write_build(tmp_path, NEW, [1, 9401])

    lines = carry_bis(tmp_path, OLD, NEW)

    carried = json.loads((tmp_path / NEW / "bis" / "mage-fire.json").read_text())
    ranged, head = carried["bands"][0]["slots"]
    assert ranged == {
        "slot": "ranged",
        "verified": False,
        "alternatives": [_alt(9401, "Bow")],
        "runners_up": [_alt(9401, "Bow")],
        "empty_reason": NOT_IN_CLIENT_REASON,
    }
    assert head == slots[1]
    assert lines == [
        f"mage-fire band 40 horde bare ranged: pick 9400 Baelog's Shortbow is not in {NEW}; "
        "slot left empty until the ranking reruns",
        f"mage-fire band 40 horde bare ranged: alternative 9403 Gone too is not in {NEW}; dropped",
    ]
    # The old build's own file is untouched.
    assert (
        json.loads((tmp_path / OLD / "bis" / "mage-fire.json").read_text())["bands"][0]["slots"]
        == slots
    )


def test_an_alternative_list_emptied_by_the_carry_is_left_out(tmp_path: Path):
    slots = [{"slot": "head", "item_id": 1, "item_name": "Hood", "alternatives": [_alt(2, "Cap")]}]
    _write_build(tmp_path, OLD, [1, 2], _payload(slots))
    _write_build(tmp_path, NEW, [1])

    carry_bis(tmp_path, OLD, NEW)

    carried = json.loads((tmp_path / NEW / "bis" / "mage-fire.json").read_text())
    assert carried["bands"][0]["slots"] == [{"slot": "head", "item_id": 1, "item_name": "Hood"}]


def test_a_carry_never_overwrites_a_ranking_the_new_build_already_has(tmp_path: Path):
    _write_build(tmp_path, OLD, [1], _payload([]))
    _write_build(tmp_path, NEW, [1], _payload([]))
    with pytest.raises(BisCarryError, match="already has"):
        carry_bis(tmp_path, OLD, NEW)


def test_a_file_stamped_with_another_build_stops_the_carry(tmp_path: Path):
    _write_build(tmp_path, OLD, [1], {**_payload([]), "build": "1.60.1.70630"})
    _write_build(tmp_path, NEW, [1])
    with pytest.raises(BisCarryError, match="1.60.1.70630"):
        carry_bis(tmp_path, OLD, NEW)
    assert not (tmp_path / NEW / "bis").exists()
