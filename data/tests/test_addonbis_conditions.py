"""The Gear tab's Conditions row and hit-cap line (design/specs/2026-11-13-addon-hub.md 4.3 and 13):
the raid preset's list (bisRaid), what it was ranked under (bisConditions) and each band's distance
to the hit cap (hitToCap), as pipeline.addonbis reads them from bis/<spec>.json and
pipeline.addonlua renders them into Data.lua."""

import json
import shutil
import subprocess
from pathlib import Path

import pytest

from pipeline.addonbis import AddonBisError, build_bis, build_bis_items, build_bis_tables
from pipeline.addondata import build_addon_data
from pipeline.addonlua import render_lua
from pipeline.models import AddonData

BUILD = "1.60.1.70374"
REPO = Path(__file__).resolve().parents[2]
ACTIVE = json.loads((REPO / "web" / "src" / "data" / "active-build.json").read_text())["build"]

RAID = {
    "label": "Raid-ready, Phase 1",
    "buffs": [
        {"id": "battle_shout", "label": "Battle Shout"},
        {"id": "gotw", "label": "Mark of the Wild"},
    ],
    "debuffs": [{"id": "sunder_armor", "label": "Sunder Armor"}],
    "consumes": [{"id": "juju_might", "label": "Juju Might"}],
    "notes": "prose the addon never shows",
}


def _band(level: int, faction: str, preset: str, item_id: int, hit_to_cap: dict | None) -> dict:
    band = {
        "band": level,
        "faction": faction,
        "preset": preset,
        "slots": [{"slot": "head", "item_id": item_id, "source_kind": "quest", "source": "Q"}],
        "weights": [],
    }
    if hit_to_cap is not None:
        band["hit_to_cap"] = hit_to_cap
    return band


def _write(root: Path, spec: str, bands: list[dict], presets: dict | None = None) -> None:
    bis_dir = root / BUILD / "bis"
    bis_dir.mkdir(parents=True, exist_ok=True)
    payload = {"spec": spec, "build": BUILD, "bands": bands}
    if presets is not None:
        payload["presets"] = presets
    (bis_dir / f"{spec}.json").write_text(json.dumps(payload), encoding="utf-8")


#: A dual-wield melee spec at 60: 6.1% hit from the list, the engine's own float noise kept; under
#: raid buffs 9.4%, past the special cap (the distance stops at 0).
FURY_BARE = {"baseline": 6.1000000000000005, "specials": 2.8999999999999995, "white": 21.9}
FURY_RAID = {"baseline": 9.4, "specials": 0, "white": 18.6}


def _fury(tmp_path: Path) -> None:
    _write(
        tmp_path,
        "warrior-fury",
        [
            _band(60, "alliance", "bare", 111, FURY_BARE),
            _band(60, "horde", "bare", 111, FURY_BARE),
            _band(60, "alliance", "raid", 222, FURY_RAID),
            _band(60, "horde", "raid", 222, FURY_RAID),
        ],
        {"raid": RAID},
    )


def test_the_raid_preset_is_its_own_list_beside_the_bare_one(tmp_path):
    _fury(tmp_path)
    tables = build_bis_tables(tmp_path, BUILD)
    assert tables.bare["warrior-fury"][0].factions["horde"]["head"].item_id == 111
    assert tables.raid["warrior-fury"][0].factions["horde"]["head"].item_id == 222
    assert [band.level for band in tables.raid["warrior-fury"]] == [60]
    # build_bis is still the bare list alone.
    assert build_bis(tmp_path, BUILD)["warrior-fury"][0].factions["alliance"]["head"].item_id == 111


def test_a_spec_without_a_raid_preset_has_no_raid_list_or_conditions(tmp_path):
    _write(tmp_path, "warrior-arms", [_band(20, "alliance", "bare", 111, None)])
    tables = build_bis_tables(tmp_path, BUILD)
    assert tables.raid == {}
    assert tables.conditions == {}
    assert tables.bare["warrior-arms"][0].hit_to_cap == {}


def test_the_hit_cap_is_each_parts_distance_and_the_baseline_plus_it_at_one_decimal(tmp_path):
    _fury(tmp_path)
    tables = build_bis_tables(tmp_path, BUILD)
    bare = tables.bare["warrior-fury"][0].hit_to_cap["horde"]
    assert (bare.hit, bare.white, bare.spell, bare.school) == ((2.9, 9.0), (21.9, 28.0), None, None)
    raid = tables.raid["warrior-fury"][0].hit_to_cap["alliance"]
    assert (raid.hit, raid.white) == ((0.0, 9.0), (18.6, 28.0))


def test_a_caster_has_the_spell_cap_and_a_talented_schools_own_distance(tmp_path):
    shadow = {
        "kind": "spell",
        "baseline": 0,
        "spell": 16,
        "school": {"names": ["Shadow"], "hit": 6.5, "to_cap": 9.5},
    }
    _write(tmp_path, "priest-shadow", [_band(40, "alliance", "bare", 111, shadow)])
    hit = build_bis_tables(tmp_path, BUILD).bare["priest-shadow"][0].hit_to_cap["alliance"]
    assert (hit.hit, hit.white, hit.spell, hit.school) == (
        None,
        None,
        (16.0, 16.0),
        (9.5, "Shadow"),
    )


def test_a_melee_spec_that_casts_keeps_its_spell_cap_too(tmp_path):
    enhancement = {"baseline": 4, "specials": 5, "spell": 12}
    _write(tmp_path, "shaman-enhancement", [_band(60, "horde", "bare", 111, enhancement)])
    hit = build_bis_tables(tmp_path, BUILD).bare["shaman-enhancement"][0].hit_to_cap["horde"]
    assert (hit.hit, hit.spell) == ((5.0, 9.0), (12.0, 16.0))


@pytest.mark.parametrize(
    "raw",
    [
        {"specials": 2},
        {"kind": "spell", "baseline": 1},
        {"kind": "melee", "baseline": 1, "specials": 8},
        {"baseline": 1},
        {"baseline": 1, "specials": 9},
        {"baseline": 0, "specials": 9, "white": 20},
        {"kind": "spell", "baseline": 2, "spell": 13},
        {
            "kind": "spell",
            "baseline": 0,
            "spell": 16,
            "school": {"names": ["Fire"], "hit": 4, "to_cap": 11},
        },
    ],
)
def test_a_hit_cap_the_addon_cannot_word_refuses(tmp_path, raw):
    _write(tmp_path, "warrior-arms", [_band(20, "alliance", "bare", 111, raw)])
    with pytest.raises(AddonBisError, match="hit_to_cap"):
        build_bis_tables(tmp_path, BUILD)


def test_the_conditions_join_each_groups_labels_in_the_files_order(tmp_path):
    _fury(tmp_path)
    conditions = build_bis_tables(tmp_path, BUILD).conditions["warrior-fury"]
    assert conditions.model_dump() == {
        "label": "Raid-ready, Phase 1",
        "buffs": "Battle Shout, Mark of the Wild",
        "debuffs": "Sunder Armor",
        "consumes": "Juju Might",
    }


@pytest.mark.parametrize(
    ("raid", "message"),
    [
        ({key: value for key, value in RAID.items() if key != "consumes"}, "consumes"),
        ({**RAID, "debuffs": [{"id": "x", "label": "Sunder | Armor"}]}, r"'\|'"),
    ],
)
def test_conditions_the_tooltip_cannot_show_refuse(tmp_path, raid, message):
    _write(tmp_path, "warrior-fury", [_band(60, "horde", "bare", 111, None)], {"raid": raid})
    with pytest.raises(AddonBisError, match=message):
        build_bis_tables(tmp_path, BUILD)


def test_a_raid_pick_is_named_like_any_other(tmp_path):
    _fury(tmp_path)
    items = [{"id": 111, "name": "Bare Helm"}, {"id": 222, "name": "Raid Helm"}]
    (tmp_path / BUILD / "items.json").write_text(json.dumps(items), encoding="utf-8")
    (tmp_path / BUILD / "items").mkdir()
    (tmp_path / BUILD / "items" / "warrior.json").write_text(
        json.dumps({"items": []}), encoding="utf-8"
    )
    tables = build_bis_tables(tmp_path, BUILD)
    names, _ = build_bis_items(tmp_path, BUILD, tables.bare, tables.raid)
    assert names == {111: "Bare Helm", 222: "Raid Helm"}


def _render(tmp_path: Path) -> str:
    tables = build_bis_tables(tmp_path, BUILD)
    return render_lua(
        AddonData(
            build="x",
            classes={},
            weights={},
            bis=tables.bare,
            bis_raid=tables.raid,
            bis_conditions=tables.conditions,
        )
    )


def test_the_raid_list_conditions_and_pooled_hit_caps_render(tmp_path):
    _fury(tmp_path)
    lua = _render(tmp_path)
    assert "\tbisRaid = {\n" in lua
    assert '\t\t\t\t\thead = "222|Q|Q",\n' in lua
    assert (
        '\t\t["warrior-fury"] = { label = "Raid-ready, Phase 1", '
        'buffs = "Battle Shout, Mark of the Wild", '
        'debuffs = "Sunder Armor", consumes = "Juju Might" },\n'
    ) in lua
    # One pooled table per distinct cap; both factions of a band point at the same one.
    assert (
        "local hitCaps = {\n\t{ hit = { 2.9, 9 }, white = { 21.9, 28 } },\n"
        "\t{ hit = { 0, 9 }, white = { 18.6, 28 } },\n}"
    ) in lua
    assert (
        "\t\t\tbare = {\n\t\t\t\t[60] = { alliance = hitCaps[1], horde = hitCaps[1] },\n\t\t\t},\n"
        in lua
    )
    assert (
        "\t\t\traid = {\n\t\t\t\t[60] = { alliance = hitCaps[2], horde = hitCaps[2] },\n\t\t\t},\n"
        in lua
    )


def test_a_build_with_no_hit_caps_or_raid_list_renders_neither(tmp_path):
    _write(tmp_path, "warrior-arms", [_band(20, "alliance", "bare", 111, None)])
    lua = _render(tmp_path)
    for name in ("bisRaid", "bisConditions", "hitToCap", "hitCaps"):
        assert name not in lua


def test_the_new_tables_read_back_in_a_real_lua(tmp_path):
    """Skipped without an interpreter, like test_addonbis' real-Lua gates."""
    lua_bin = shutil.which("lua") or shutil.which("lua5.1")
    if not lua_bin:
        pytest.skip("no lua interpreter is available")
    _fury(tmp_path)
    chunk = tmp_path / "Data.lua"
    chunk.write_text(_render(tmp_path), encoding="utf-8")
    script = (
        f"local data = assert(loadfile({str(chunk)!r}))() "
        "local cap = data.hitToCap['warrior-fury'].bare[60].horde "
        "io.write(data.bisRaid['warrior-fury'][60].alliance.head, ';', "
        "cap.hit[1], '/', cap.hit[2], ';', cap.white[1], '/', cap.white[2], ';', "
        "data.bisConditions['warrior-fury'].debuffs)"
    )
    result = subprocess.run([lua_bin, "-e", script], capture_output=True, text=True)
    assert result.returncode == 0, result.stderr
    assert result.stdout == "222|Q|Q;2.9/9;21.9/28;Sunder Armor"


def test_every_real_spec_with_a_raid_preset_carries_its_list_conditions_and_caps():
    """The active build: each raid band has both factions' lists, and every band's hit cap reads
    (the generator refuses one that does not land on its cap)."""
    data = build_addon_data(ACTIVE)
    assert data.bis_raid, "the active build has no raid preset at all"
    assert set(data.bis_conditions) == set(data.bis_raid)
    for spec, bands in data.bis_raid.items():
        assert [band.level for band in bands] == [60], spec
        assert set(bands[0].factions) == {"alliance", "horde"}, spec
    for table in (data.bis, data.bis_raid):
        for spec, bands in table.items():
            for band in bands:
                for faction, cap in band.hit_to_cap.items():
                    where = (spec, band.level, faction)
                    parts = (cap.hit, cap.white, cap.spell)
                    assert any(part is not None for part in parts), where
                    assert all(part is None or 0 <= part[0] <= part[1] for part in parts), where


def test_every_real_raid_pick_has_a_name():
    data = build_addon_data(ACTIVE)
    for spec, bands in data.bis_raid.items():
        for band in bands:
            for faction, slots in band.factions.items():
                for slot, pick in slots.items():
                    assert pick.item_id in data.bis_names, (spec, faction, slot)
