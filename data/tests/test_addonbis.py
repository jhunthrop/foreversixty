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


def test_chance_and_alternatives_pack_into_one_string_per_pick(tmp_path):
    """Fix round 2 (memory budget): one Lua string per pick, fields joined by
    `|` as item_id|kind|source|chance|alt1_id|alt1_kind|alt1_source|alt2_id|
    alt2_kind|alt2_source, trailing empty fields trimmed; Gear.pick decodes it."""
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(_DEADMINES_PICK))
    lua = _render_bis(build_bis(tmp_path, BUILD))
    assert (
        'main_hand = "7230|D|The Deadmines: Defias Overseer|0.14|'
        '270018|Q|Leaders of the Fang|1458|B|World drop",'
    ) in lua


def test_a_pick_with_no_chance_leaves_the_fourth_field_empty(tmp_path):
    quest_pick = dict(_DEADMINES_PICK, item_id=6087, source_kind="quest", source="The Defias Brotherhood")
    del quest_pick["chance"]
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(quest_pick))
    lua = _render_bis(build_bis(tmp_path, BUILD))
    assert (
        'main_hand = "6087|Q|The Defias Brotherhood||'
        '270018|Q|Leaders of the Fang|1458|B|World drop",'
    ) in lua


def test_a_pick_with_no_source_label_but_a_chance_keeps_chance_fourth(tmp_path):
    """The positions are fixed: an empty source is an empty field rather
    than letting chance slide into it."""
    pick = {"slot": "head", "item_id": 111, "source_kind": "dungeon", "chance": 0.05}
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(pick))
    lua = _render_bis(build_bis(tmp_path, BUILD))
    assert 'head = "111|D||0.05",' in lua


def test_a_band_written_before_chance_existed_has_nil_chance_and_no_alternatives(tmp_path):
    """bis/<spec>.json files leveling-bis wrote before it emitted `chance`
    and `runners_up` still load; their picks render exactly as before."""
    old_pick = {"slot": "chest", "item_id": 2317, "source_kind": "crafted", "source": "Leatherworking"}
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(old_pick))
    bis = build_bis(tmp_path, BUILD)
    pick = bis["warrior-arms"][0].factions["alliance"]["chest"]
    assert pick.chance is None
    assert pick.alternatives == []
    assert 'chest = "2317|C|Leatherworking",' in _render_bis(bis)


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
        "io.write(data.bis['warrior-arms'][20].alliance.main_hand)"
    )
    result = subprocess.run([lua_bin, "-e", script], capture_output=True, text=True)
    assert result.returncode == 0, result.stderr
    assert result.stdout == (
        "7230|D|The Deadmines: Defias Overseer|0.14|270018|Q|Leaders of the Fang|1458|B|World drop"
    )


def test_a_legacy_band_without_runners_up_takes_two_of_its_alternatives_and_no_chance(tmp_path):
    """bis files the nightly job wrote before leveling-bis emitted
    `runners_up` already carry `alternatives` (up to three, best first):
    the first two stand in until the next nightly run, chance stays nil."""
    legacy_pick = {
        "slot": "back",
        "item_id": 4706,
        "source_kind": "world_drop",
        "source": "World drop",
        "alternatives": [
            {"item_id": 279865, "item_name": "Grave Shroud", "source_kind": "quest", "source": "Abominable Creatures", "dps_delta": -0.019},
            {"item_id": 2316, "item_name": "Dark Leather Cloak", "source_kind": "crafted", "source": "Leatherworking", "dps_delta": -0.064},
            {"item_id": 279899, "item_name": "Catacomb Cloak", "source_kind": "quest", "source": "An Ancient Grudge", "dps_delta": -0.071},
        ],
    }
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(legacy_pick))
    bis = build_bis(tmp_path, BUILD)
    pick = bis["warrior-arms"][0].factions["alliance"]["back"]
    assert pick.chance is None
    assert [alt.item_id for alt in pick.alternatives] == [279865, 2316]
    assert (
        'back = "4706|B|World drop||'
        '279865|Q|Abominable Creatures|2316|C|Leatherworking",'
    ) in _render_bis(bis)


def test_runners_up_win_over_alternatives_when_a_band_has_both(tmp_path):
    both = dict(
        _DEADMINES_PICK,
        alternatives=[{"item_id": 1, "source_kind": "quest", "source": "Not This One"}],
    )
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(both))
    pick = build_bis(tmp_path, BUILD)["warrior-arms"][0].factions["alliance"]["main_hand"]
    assert [alt.item_id for alt in pick.alternatives] == [270018, 1458]


def test_an_alternative_with_no_source_keeps_its_empty_field_in_the_middle(tmp_path):
    pick = dict(
        _DEADMINES_PICK,
        runners_up=[
            {"item_id": 5, "source_kind": "quest"},
            {"item_id": 1458, "source_kind": "world_drop", "source": "World drop"},
        ],
    )
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(pick))
    lua = _render_bis(build_bis(tmp_path, BUILD))
    assert 'main_hand = "7230|D|The Deadmines: Defias Overseer|0.14|5|Q||1458|B|World drop",' in lua


def test_a_source_with_the_field_separator_refuses_to_pack(tmp_path):
    from pipeline.addonbis import AddonBisError as RenderError

    pick = dict(_DEADMINES_PICK, source="The Deadmines | Defias Overseer")
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(pick))
    with pytest.raises(RenderError, match="separator"):
        _render_bis(build_bis(tmp_path, BUILD))


# --- Task L1.4: client-unconfirmed candidates never become picks; pick names and weapon dps ---

_VERDAN_SOURCE = "Wailing Caverns: Verdan the Everliving"

#: The Arms band-20 main-hand slot as leveling-bis wrote it in the real
#: build: its pick, The Axe of Severing (23171), is `client_unconfirmed` in
#: items.json, so the client never resolves it and the slot fell silent.
_VERDAN = {"item_id": 6631, "source_kind": "dungeon", "source": _VERDAN_SOURCE}

_SEVERING_SLOT = {
    "slot": "main_hand",
    "item_id": 23171,
    "item_name": "The Axe of Severing",
    "source_kind": "dungeon",
    "source": "Shadowfang Keep: Sever",
    "chance": 0.2,
    "runners_up": [
        _VERDAN,
        {"item_id": 2205, "source_kind": "dungeon", "source": "Shadowfang Keep: Son of Arugal"},
    ],
    "alternatives": [
        _VERDAN,
        {"item_id": 2205, "source_kind": "dungeon", "source": "Shadowfang Keep: Son of Arugal"},
        {"item_id": 7230, "source_kind": "dungeon", "source": "The Deadmines: Mr. Smite"},
    ],
}


def _item(item_id: int, name: str, *, unconfirmed: bool = False) -> dict:
    return {"id": item_id, "name": name, "client_unconfirmed": unconfirmed}


def _class_item(item_id: int, name: str, dps: float = 0.0) -> dict:
    return {"id": item_id, "name": name, "dps": dps}


def _write_items(root: Path, build: str, flat: list[dict], by_class: dict[str, list[dict]]) -> None:
    (root / build).mkdir(parents=True, exist_ok=True)
    (root / build / "items.json").write_text(json.dumps(flat), encoding="utf-8")
    (root / build / "items").mkdir(exist_ok=True)
    for class_slug, rows in by_class.items():
        (root / build / "items" / f"{class_slug}.json").write_text(
            json.dumps({"build": build, "class_slug": class_slug, "items": rows}), encoding="utf-8"
        )


def _severing_items(root: Path, *, all_unconfirmed: bool = False) -> None:
    _write_items(
        root,
        BUILD,
        [
            _item(23171, "The Axe of Severing", unconfirmed=True),
            _item(6631, "Living Root", unconfirmed=all_unconfirmed),
            _item(2205, "Duskbringer", unconfirmed=all_unconfirmed),
            _item(7230, "Smite's Mighty Hammer", unconfirmed=all_unconfirmed),
        ],
        {
            "warrior": [
                _class_item(23171, "The Axe of Severing", 21.0),
                _class_item(6631, "Living Root", 19.94),
                _class_item(2205, "Duskbringer", 20.05),
                _class_item(7230, "Smite's Mighty Hammer", 18.83),
            ]
        },
    )


def test_an_unconfirmed_pick_gives_way_to_the_next_confirmed_candidate(tmp_path):
    """23171 is client_unconfirmed: the first runner-up becomes the pick
    (with its own source, and no chance -- the chance was 23171's), and the
    alternatives are the next two confirmed candidates."""
    _severing_items(tmp_path)
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(_SEVERING_SLOT))
    pick = build_bis(tmp_path, BUILD)["warrior-arms"][0].factions["alliance"]["main_hand"]
    assert (pick.item_id, pick.source, pick.chance) == (
        6631, "Wailing Caverns: Verdan the Everliving", None
    )
    assert [alt.item_id for alt in pick.alternatives] == [2205, 7230]


def test_an_unconfirmed_runner_up_is_never_an_alternative(tmp_path):
    _severing_items(tmp_path)
    slot = dict(_SEVERING_SLOT, item_id=7230, source="The Deadmines: Mr. Smite")
    slot["runners_up"] = [
        {"item_id": 23171, "source_kind": "dungeon", "source": "Shadowfang Keep: Sever"},
        _VERDAN,
    ]
    slot["alternatives"] = slot["runners_up"] + [
        {"item_id": 2205, "source_kind": "dungeon", "source": "Shadowfang Keep: Son of Arugal"}
    ]
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(slot))
    pick = build_bis(tmp_path, BUILD)["warrior-arms"][0].factions["alliance"]["main_hand"]
    assert pick.item_id == 7230 and pick.chance == 0.2
    assert [alt.item_id for alt in pick.alternatives] == [6631, 2205]


def test_a_slot_with_only_unconfirmed_candidates_packs_nothing(tmp_path):
    _severing_items(tmp_path, all_unconfirmed=True)
    _write_bis(tmp_path, BUILD, "warrior-arms", _deadmines_band(_SEVERING_SLOT))
    band = build_bis(tmp_path, BUILD)["warrior-arms"][0]
    assert band.factions["alliance"] == {}
    assert 'main_hand = "' not in _render_bis(build_bis(tmp_path, BUILD))


def test_the_skips_are_counted_per_build(tmp_path):
    from pipeline.addonbis import bis_unconfirmed_skips

    _severing_items(tmp_path)
    _write_bis(
        tmp_path,
        BUILD,
        "warrior-arms",
        [
            _band(20, "alliance", [_SEVERING_SLOT]),
            _band(20, "horde", [_SEVERING_SLOT]),
        ],
    )
    _write_bis(tmp_path, BUILD, "warrior-fury", [_band(20, "alliance", [_SEVERING_SLOT])])
    assert bis_unconfirmed_skips(tmp_path, BUILD) == 3


def test_the_real_arms_band_20_main_hand_is_never_the_axe_of_severing():
    bis = build_bis(Path("builds"), BUILD)
    band = next(b for b in bis["warrior-arms"] if b.level == 20)
    for faction, slots in band.factions.items():
        pick = slots["main_hand"]
        ids = [pick.item_id] + [alt.item_id for alt in pick.alternatives]
        assert 23171 not in ids, faction


def test_no_real_pick_or_alternative_is_client_unconfirmed():
    flat = json.loads((Path("builds") / BUILD / "items.json").read_text(encoding="utf-8"))
    unconfirmed = {row["id"] for row in flat if row.get("client_unconfirmed")}
    for spec, bands in build_bis(Path("builds"), BUILD).items():
        for band in bands:
            for faction, slots in band.factions.items():
                for slot, pick in slots.items():
                    ids = {pick.item_id} | {alt.item_id for alt in pick.alternatives}
                    assert not ids & unconfirmed, (spec, band.level, faction, slot)


def test_every_pick_and_alternative_has_a_name_and_only_weapons_have_dps(tmp_path):
    from pipeline.addonbis import build_bis_items

    _severing_items(tmp_path)
    _write_items(
        tmp_path,
        BUILD,
        json.loads((tmp_path / BUILD / "items.json").read_text(encoding="utf-8"))
        + [_item(111, "Cap of Testing")],
        {
            "warrior": [
                _class_item(6631, "Living Root", 19.94),
                _class_item(2205, "Duskbringer", 20.05),
                _class_item(7230, "Smite's Mighty Hammer", 18.83),
                _class_item(111, "Cap of Testing"),
            ]
        },
    )
    head = {"slot": "head", "item_id": 111, "source_kind": "quest", "source": "Q"}
    _write_bis(tmp_path, BUILD, "warrior-arms", [_band(20, "alliance", [_SEVERING_SLOT, head])])
    names, dps = build_bis_items(tmp_path, BUILD, build_bis(tmp_path, BUILD))
    assert names == {
        111: "Cap of Testing",
        2205: "Duskbringer",
        6631: "Living Root",
        7230: "Smite's Mighty Hammer",
    }
    # One decimal; the cap is not a weapon, and 23171 was never packed.
    assert dps == {2205: 20.1, 6631: 19.9, 7230: 18.8}


def test_a_name_with_the_field_separator_becomes_a_slash(tmp_path):
    from pipeline.addonbis import build_bis_items

    _write_items(tmp_path, BUILD, [_item(111, "Cap | of Testing")], {"warrior": []})
    head = {"slot": "head", "item_id": 111, "source_kind": "quest", "source": "Q"}
    _write_bis(tmp_path, BUILD, "warrior-arms", [_band(20, "alliance", [head])])
    names, _ = build_bis_items(tmp_path, BUILD, build_bis(tmp_path, BUILD))
    assert names == {111: "Cap / of Testing"}


def test_a_pick_the_items_file_does_not_name_refuses(tmp_path):
    from pipeline.addonbis import build_bis_items

    _write_items(tmp_path, BUILD, [], {"warrior": []})
    head = {"slot": "head", "item_id": 111, "source_kind": "quest", "source": "Q"}
    _write_bis(tmp_path, BUILD, "warrior-arms", [_band(20, "alliance", [head])])
    with pytest.raises(AddonBisError, match="111"):
        build_bis_items(tmp_path, BUILD, build_bis(tmp_path, BUILD))


def test_the_name_and_dps_tables_render_beside_an_unchanged_pick_string():
    from pipeline.addonlua import render_lua
    from pipeline.models import AddonBisBand, AddonBisItem, AddonData

    pick = AddonBisItem(item_id=7230, source_kind="dungeon", source="The Deadmines: Mr. Smite")
    data = AddonData(
        build="x",
        classes={},
        weights={},
        bis={"warrior-arms": [AddonBisBand(level=20, factions={"alliance": {"main_hand": pick}})]},
        bis_names={7230: "Smite's Mighty Hammer", 111: 'Cap "Quoted"'},
        bis_dps={7230: 18.8, 2205: 20.0},
    )
    lua = render_lua(data)
    assert 'main_hand = "7230|D|The Deadmines: Mr. Smite",' in lua
    assert (
        "\tbisNames = {\n"
        '\t\t[111] = "Cap \\"Quoted\\"",\n'
        "\t\t[7230] = \"Smite's Mighty Hammer\",\n"
        "\t},\n"
    ) in lua
    assert "\tbisDps = {\n\t\t[2205] = 20,\n\t\t[7230] = 18.8,\n\t},\n" in lua


def test_the_audit_line_names_the_build_the_skips_and_the_distinct_items(tmp_path):
    from pipeline.addonbis import unconfirmed_skips_line

    _severing_items(tmp_path)
    _write_bis(
        tmp_path,
        BUILD,
        "warrior-arms",
        [_band(20, "alliance", [_SEVERING_SLOT]), _band(20, "horde", [_SEVERING_SLOT])],
    )
    assert unconfirmed_skips_line(tmp_path, BUILD) == (
        f"bis {BUILD}: skipped 2 client_unconfirmed candidates (1 distinct items)"
    )
