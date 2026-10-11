import json
from pathlib import Path

import pytest

from pipeline.addonbis import (
    SOURCE_KIND_CODES,
    AddonBisError,
    build_bis,
)

BUILD = "1.60.1.70291"


def _write_bis(root: Path, build: str, spec: str, bands: list[dict]) -> None:
    bis_dir = root / build / "bis"
    bis_dir.mkdir(parents=True, exist_ok=True)
    (bis_dir / f"{spec}.json").write_text(
        json.dumps({"spec": spec, "build": build, "bands": bands}), encoding="utf-8"
    )


def _band(
    level: int,
    faction: str,
    slots: list[dict],
    new_at_band: list[str] | None = None,
    weights: list[dict] | None = None,
) -> dict:
    return {
        "spec": "hunter-marksmanship",
        "band": level,
        "faction": faction,
        "slots": slots,
        "new_at_band": new_at_band or [],
        "weights": weights or [],
    }


def test_missing_bis_directory_answers_empty():
    assert build_bis(Path("/nonexistent-root-for-a-test"), BUILD) == {}


def test_a_spec_with_no_bis_file_is_simply_absent(tmp_path):
    (tmp_path / BUILD).mkdir()
    (tmp_path / BUILD / "bis").mkdir()
    assert build_bis(tmp_path, BUILD) == {}


def test_two_factions_at_one_band_merge_into_one_entry(tmp_path):
    _write_bis(
        tmp_path,
        BUILD,
        "hunter-marksmanship",
        [
            _band(20, "alliance", [{"slot": "head", "item_id": 111, "source_kind": "quest"}]),
            _band(20, "horde", [{"slot": "head", "item_id": 222, "source_kind": "dungeon"}]),
        ],
    )
    bis = build_bis(tmp_path, BUILD)
    band = bis["hunter-marksmanship"][0]
    assert band.level == 20
    assert band.factions["alliance"]["head"].item_id == 111
    assert band.factions["horde"]["head"].item_id == 222
    assert band.factions["horde"]["head"].source_kind == "dungeon"


def test_a_slot_with_no_item_id_is_left_out(tmp_path):
    """leveling-bis writes a bare {"slot": ..., "verified": false} row when
    nothing eligible was found -- no pick, no line in the addon table."""
    _write_bis(
        tmp_path,
        BUILD,
        "hunter-marksmanship",
        [_band(35, "alliance", [{"slot": "off_hand", "verified": False}])],
    )
    bis = build_bis(tmp_path, BUILD)
    assert bis["hunter-marksmanship"][0].factions["alliance"] == {}


def test_new_at_band_resolves_slot_names_to_the_bands_own_item_ids(tmp_path):
    _write_bis(
        tmp_path,
        BUILD,
        "hunter-marksmanship",
        [
            _band(
                20,
                "alliance",
                [
                    {"slot": "head", "item_id": 111, "source_kind": "quest"},
                    {"slot": "neck", "item_id": 333, "source_kind": "quest"},
                ],
                new_at_band=["head: Helm of the Pathfinder"],
            )
        ],
    )
    bis = build_bis(tmp_path, BUILD)
    assert bis["hunter-marksmanship"][0].new_at_band["alliance"] == [111]


def test_a_new_at_band_slot_with_no_matching_pick_is_skipped(tmp_path):
    _write_bis(
        tmp_path,
        BUILD,
        "hunter-marksmanship",
        [_band(20, "alliance", [], new_at_band=["head: Something Removed"])],
    )
    bis = build_bis(tmp_path, BUILD)
    assert bis["hunter-marksmanship"][0].new_at_band["alliance"] == []


def test_bands_come_out_sorted_by_level(tmp_path):
    _write_bis(
        tmp_path,
        BUILD,
        "hunter-marksmanship",
        [
            _band(30, "alliance", []),
            _band(10, "alliance", []),
            _band(20, "alliance", []),
        ],
    )
    bis = build_bis(tmp_path, BUILD)
    assert [band.level for band in bis["hunter-marksmanship"]] == [10, 20, 30]


def test_an_unknown_faction_refuses(tmp_path):
    _write_bis(tmp_path, BUILD, "hunter-marksmanship", [_band(20, "scourge", [])])
    with pytest.raises(AddonBisError, match="scourge"):
        build_bis(tmp_path, BUILD)


def test_an_unknown_slot_refuses(tmp_path):
    _write_bis(
        tmp_path,
        BUILD,
        "hunter-marksmanship",
        [_band(20, "alliance", [{"slot": "cape", "item_id": 1, "source_kind": "quest"}])],
    )
    with pytest.raises(AddonBisError, match="cape"):
        build_bis(tmp_path, BUILD)


def test_every_source_kind_code_is_a_single_uppercase_letter():
    """leveling-bis's own vocabulary (sim/cmd/leveling-bis/band.go's
    sourceKindPriority); a code for a kind that vocabulary drops would be
    dead, and a new kind with no code here is what
    pipeline.addonlua._source_kind_code refuses at render time."""
    assert set(SOURCE_KIND_CODES) == {
        "quest",
        "vendor",
        "dungeon",
        "crafted",
        "rep",
        "pvp",
        "world",
        "raid",
        "world_drop",
    }
    assert all(len(code) == 1 and code.isupper() for code in SOURCE_KIND_CODES.values())
    assert len(set(SOURCE_KIND_CODES.values())) == len(SOURCE_KIND_CODES)


def test_multiple_specs_each_get_their_own_table(tmp_path):
    _write_bis(tmp_path, BUILD, "hunter-marksmanship", [_band(20, "alliance", [])])
    _write_bis(tmp_path, BUILD, "warrior-fury", [_band(20, "alliance", [])])
    bis = build_bis(tmp_path, BUILD)
    assert set(bis) == {"hunter-marksmanship", "warrior-fury"}


def test_a_pick_carries_its_source_label(tmp_path):
    """lane addon-tooltip-polish item 1: the tooltip's muted source line
    reads this, in preference to a bare source kind."""
    _write_bis(
        tmp_path,
        BUILD,
        "hunter-marksmanship",
        [
            _band(
                20,
                "alliance",
                [
                    {
                        "slot": "chest",
                        "item_id": 2317,
                        "source": "Leatherworking",
                        "source_kind": "crafted",
                    }
                ],
            )
        ],
    )
    bis = build_bis(tmp_path, BUILD)
    assert bis["hunter-marksmanship"][0].factions["alliance"]["chest"].source == "Leatherworking"


def test_a_pick_with_no_source_label_is_simply_empty(tmp_path):
    _write_bis(
        tmp_path,
        BUILD,
        "hunter-marksmanship",
        [_band(20, "alliance", [{"slot": "chest", "item_id": 2317, "source_kind": "crafted"}])],
    )
    bis = build_bis(tmp_path, BUILD)
    assert bis["hunter-marksmanship"][0].factions["alliance"]["chest"].source == ""


def test_a_bands_weights_keep_only_the_significant_ones(tmp_path):
    """lane addon-tooltip-polish item 4: the tooltip's verdict scores a
    hovered item against these, never an insignificant one leveling-bis
    itself flagged as noise."""
    _write_bis(
        tmp_path,
        BUILD,
        "hunter-marksmanship",
        [
            _band(
                20,
                "alliance",
                [],
                weights=[
                    {"stat": "agility", "weight": 2.0452764470665192, "insignificant": False},
                    {"stat": "hit", "weight": 5.226522612412781, "insignificant": True},
                ],
            )
        ],
    )
    bis = build_bis(tmp_path, BUILD)
    assert bis["hunter-marksmanship"][0].weights == {"agility": 2.045}


def test_a_bands_weights_are_the_same_for_every_faction(tmp_path):
    """Only the first faction's own `weights` list is kept (sorted, so
    always the same one for a given band) -- see AddonBisBand.weights'
    docstring on why factions are not kept apart for this field."""
    _write_bis(
        tmp_path,
        BUILD,
        "hunter-marksmanship",
        [
            _band(20, "horde", [], weights=[{"stat": "agility", "weight": 1.0}]),
            _band(20, "alliance", [], weights=[{"stat": "agility", "weight": 2.0}]),
        ],
    )
    bis = build_bis(tmp_path, BUILD)
    assert bis["hunter-marksmanship"][0].weights == {"agility": 2.0}


def test_a_band_with_no_weights_measured_yet_is_simply_empty(tmp_path):
    _write_bis(tmp_path, BUILD, "hunter-marksmanship", [_band(20, "alliance", [])])
    bis = build_bis(tmp_path, BUILD)
    assert bis["hunter-marksmanship"][0].weights == {}


def test_a_raid_preset_entry_never_replaces_the_bare_one(tmp_path):
    """Level 60 carries a bare and a raid entry per faction; the addon table is the bare one."""
    bare = _band(60, "alliance", [{"slot": "head", "item_id": 111, "source_kind": "quest"}])
    bare["preset"] = "bare"
    raid = _band(60, "alliance", [{"slot": "head", "item_id": 999, "source_kind": "quest"}])
    raid["preset"] = "raid"
    _write_bis(tmp_path, BUILD, "hunter-marksmanship", [bare, raid])
    band = build_bis(tmp_path, BUILD)["hunter-marksmanship"][0]
    assert band.factions["alliance"]["head"].item_id == 111


def _deadmines_band(slot: dict) -> list[dict]:
    return [_band(20, "alliance", [slot])]


def _render_bis(bis: dict) -> str:
    from pipeline.addonlua import render_lua
    from pipeline.models import AddonData

    return render_lua(AddonData(build="x", classes={}, weights={}, bis=bis))


_DEADMINES_PICK = {
    "slot": "main_hand",
    "item_id": 7230,
    "source_kind": "dungeon",
    "source": "The Deadmines: Defias Overseer",
    "chance": 0.14,
    "runners_up": [
        {"item_id": 270018, "source_kind": "quest", "source": "Leaders of the Fang"},
        {"item_id": 1458, "source_kind": "world_drop", "source": "World drop"},
    ],
}


def test_a_pick_carries_its_drop_chance_and_two_runners_up_as_alternatives(tmp_path):
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(_DEADMINES_PICK))
    pick = build_bis(tmp_path, BUILD)["warrior-arms"][0].factions["alliance"]["main_hand"]
    assert pick.chance == 0.14
    assert [(alt.item_id, alt.source_kind, alt.source) for alt in pick.alternatives] == [
        (270018, "quest", "Leaders of the Fang"),
        (1458, "world_drop", "World drop"),
    ]


def test_chance_and_alternatives_render_as_the_fourth_and_fifth_positional_fields(tmp_path):
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(_DEADMINES_PICK))
    lua = _render_bis(build_bis(tmp_path, BUILD))
    assert (
        'main_hand = { 7230, "D", "The Deadmines: Defias Overseer", 0.14, '
        '{ { 270018, "Q", "Leaders of the Fang" }, { 1458, "B", "World drop" } } },'
    ) in lua


def test_a_pick_with_no_chance_renders_nil_in_the_fourth_field(tmp_path):
    quest_pick = dict(_DEADMINES_PICK, item_id=6087, source_kind="quest", source="The Defias Brotherhood")
    del quest_pick["chance"]
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(quest_pick))
    lua = _render_bis(build_bis(tmp_path, BUILD))
    assert (
        'main_hand = { 6087, "Q", "The Defias Brotherhood", nil, '
        '{ { 270018, "Q", "Leaders of the Fang" }, { 1458, "B", "World drop" } } },'
    ) in lua


def test_a_pick_with_no_source_label_but_a_chance_keeps_chance_fourth(tmp_path):
    """The positions are fixed: an empty source becomes nil in the third
    field rather than letting chance slide into it."""
    pick = {"slot": "head", "item_id": 111, "source_kind": "dungeon", "chance": 0.05}
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(pick))
    lua = _render_bis(build_bis(tmp_path, BUILD))
    assert 'head = { 111, "D", nil, 0.05 },' in lua


def test_a_band_written_before_chance_existed_has_nil_chance_and_no_alternatives(tmp_path):
    """bis/<spec>.json files leveling-bis wrote before it emitted `chance`
    and `runners_up` still load; their picks render exactly as before."""
    old_pick = {"slot": "chest", "item_id": 2317, "source_kind": "crafted", "source": "Leatherworking"}
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(old_pick))
    bis = build_bis(tmp_path, BUILD)
    pick = bis["warrior-arms"][0].factions["alliance"]["chest"]
    assert pick.chance is None
    assert pick.alternatives == []
    assert 'chest = { 2317, "C", "Leatherworking" },' in _render_bis(bis)


def test_a_chance_above_one_is_a_percent_and_refuses(tmp_path):
    """loot.json states percents (0-100); the addon contract is a fraction."""
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(dict(_DEADMINES_PICK, chance=14)))
    with pytest.raises(AddonBisError, match="chance"):
        build_bis(tmp_path, BUILD)


def test_more_than_two_runners_up_refuses_the_chance_contract(tmp_path):
    three = dict(
        _DEADMINES_PICK,
        runners_up=_DEADMINES_PICK["runners_up"] + [{"item_id": 5, "source_kind": "quest", "source": "Q"}],
    )
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(three))
    with pytest.raises(AddonBisError, match="runners_up"):
        build_bis(tmp_path, BUILD)


def test_the_chance_pick_table_reads_back_in_a_real_lua(tmp_path):
    """Skipped without an interpreter, like test_addondata's real-Lua gates."""
    import shutil
    import subprocess

    lua_bin = shutil.which("lua") or shutil.which("lua5.4")
    if not lua_bin:
        pytest.skip("no lua interpreter is available")
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(_DEADMINES_PICK))
    chunk = tmp_path / "Data.lua"
    chunk.write_text(_render_bis(build_bis(tmp_path, BUILD)), encoding="utf-8")
    script = (
        f"local data = assert(loadfile({str(chunk)!r}))() "
        "local pick = data.bis['warrior-arms'][20].alliance.main_hand "
        "io.write(pick[1], '|', pick[2], '|', pick[3], '|', pick[4], '|', #pick[5], '|', "
        "pick[5][1][1], ',', pick[5][1][2], ',', pick[5][1][3], '|', pick[5][2][1])"
    )
    result = subprocess.run([lua_bin, "-e", script], capture_output=True, text=True)
    assert result.returncode == 0, result.stderr
    assert result.stdout == "7230|D|The Deadmines: Defias Overseer|0.14|2|270018,Q,Leaders of the Fang|1458"
