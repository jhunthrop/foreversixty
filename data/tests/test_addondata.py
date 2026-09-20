import json
import subprocess
import sys
from pathlib import Path

import pytest

from pipeline.addondata import AddonDataError, build_addon_data, check_addon_data, write_addon_data
from pipeline.addonlua import render_lua, write_lua

BUILD = "1.60.1.69893"


def test_every_class_has_three_tabs_in_position_order():
    data = build_addon_data(BUILD)
    paladin = data.classes["paladin"]
    assert [tab.name for tab in paladin.tabs] == ["Holy", "Protection", "Retribution"]


def test_talents_keep_the_site_array_order_and_are_one_based():
    """The export encodes ranks in this order, so it is the contract, not a detail."""
    source = json.loads(
        Path(f"builds/{BUILD}/talents/paladin.json").read_text(encoding="utf-8")
    )
    holy = next(tree for tree in source["trees"] if tree["position"] == 0)
    tab = build_addon_data(BUILD).classes["paladin"].tabs[0]
    assert [talent.name for talent in tab.talents] == [t["name"] for t in holy["talents"]]
    assert tab.talents[0].tier == holy["talents"][0]["tier"] + 1
    assert tab.talents[0].column == holy["talents"][0]["column"] + 1
    assert tab.talents[0].max_rank == holy["talents"][0]["max_rank"]


def test_no_two_talents_in_a_tab_share_a_tier_and_column():
    """The addon keys the client's talents by tier:column; a collision would alias them."""
    data = build_addon_data(BUILD)
    for class_slug, entry in data.classes.items():
        for tab in entry.tabs:
            cells = [(talent.tier, talent.column) for talent in tab.talents]
            assert len(cells) == len(set(cells)), f"{class_slug} {tab.name}"


def test_the_weights_travel_with_the_layout():
    data = build_addon_data(BUILD)
    assert data.weights["paladin-holy"]["spell_power"] == 1.0
    assert len(data.weights) == 27


def test_the_build_id_is_the_directory():
    assert build_addon_data(BUILD).build == BUILD


def test_a_talent_file_from_another_build_is_refused(tmp_path):
    (tmp_path / BUILD / "talents").mkdir(parents=True)
    (tmp_path / BUILD / "talents" / "paladin.json").write_text(
        json.dumps({"build": "1.15.9.69722", "class_id": 2, "class_slug": "paladin", "trees": []}),
        encoding="utf-8",
    )
    (tmp_path / BUILD / "classes.json").write_text(
        json.dumps([{"id": 2, "name": "Paladin", "slug": "paladin", "color": "#f58cba"}]),
        encoding="utf-8",
    )
    with pytest.raises(AddonDataError, match="1.15.9.69722"):
        build_addon_data(BUILD, root=tmp_path)


def test_write_addon_data_round_trips(tmp_path):
    path = write_addon_data(BUILD, out_root=tmp_path)
    assert json.loads(path.read_text(encoding="utf-8"))["build"] == BUILD


def test_the_rendered_lua_parses_and_matches_the_golden_paladin_tab():
    lua = render_lua(build_addon_data(BUILD))
    golden = Path("tests/golden/Data.paladin.lua").read_text(encoding="utf-8")
    start = lua.index('["paladin"] = {')
    assert lua[start : start + len(golden)] == golden


def test_the_rendered_lua_is_loadable_by_a_real_lua():
    """A generated chunk that does not parse is worse than no chunk at all."""
    lua = render_lua(build_addon_data(BUILD))
    result = subprocess.run(
        ["lua", "-e", f"local f, err = load([==[{lua}]==], 'Data.lua'); assert(f, err)"],
        capture_output=True,
        text=True,
    )
    assert result.returncode == 0, result.stderr


def test_a_name_with_a_quote_is_escaped():
    from pipeline.models import AddonClass, AddonData, AddonTab, AddonTalent

    data = AddonData(
        build="x",
        classes={
            "rogue": AddonClass(
                tabs=[
                    AddonTab(
                        name='He said "hi"',
                        talents=[AddonTalent(name="A\\B", tier=1, column=1, max_rank=1)],
                    )
                ]
            )
        },
        weights={},
    )
    lua = render_lua(data)
    assert '"He said \\"hi\\""' in lua
    assert '"A\\\\B"' in lua


def test_write_lua_and_check_agree(tmp_path):
    from pipeline.addonlua import lua_has_drifted

    path = write_lua(BUILD, lua_path=tmp_path / "Data.lua")
    assert not lua_has_drifted(BUILD, lua_path=path)
    path.write_text("-- edited by hand\n", encoding="utf-8")
    assert lua_has_drifted(BUILD, lua_path=path)


def test_check_addon_data_agrees_with_a_deliberately_drifted_copy(tmp_path):
    """The Lua half has both directions covered by test_write_lua_and_check_agree;
    this is the same shape for the JSON half."""
    (tmp_path / BUILD / "talents").mkdir(parents=True)
    (tmp_path / BUILD / "talents" / "paladin.json").write_text(
        json.dumps({"build": BUILD, "class_id": 2, "class_slug": "paladin", "trees": []}),
        encoding="utf-8",
    )
    (tmp_path / BUILD / "classes.json").write_text(
        json.dumps([{"id": 2, "name": "Paladin", "slug": "paladin", "color": "#f58cba"}]),
        encoding="utf-8",
    )
    write_addon_data(BUILD, root=tmp_path)
    assert not check_addon_data(BUILD, root=tmp_path)
    (tmp_path / BUILD / "addon-data.json").write_text('{"build": "drifted"}', encoding="utf-8")
    assert check_addon_data(BUILD, root=tmp_path)


def test_the_cli_check_passes_on_the_committed_file():
    result = subprocess.run(
        [sys.executable, "-m", "pipeline", "addon-data", "--build", BUILD, "--check"],
        capture_output=True,
        text=True,
    )
    assert result.returncode == 0, result.stderr
