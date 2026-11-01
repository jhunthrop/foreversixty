import json
import os
import shutil
import subprocess
import sys
from pathlib import Path

import pytest

from pipeline.addondata import AddonDataError, build_addon_data, check_addon_data, write_addon_data
from pipeline.addonlua import render_lua, write_lua

# The active build, so this test follows every client move instead of pinning one.
BUILD = json.loads(
    (Path(__file__).resolve().parents[2] / "web" / "src" / "data" / "active-build.json").read_text()
)["build"]


def test_every_class_has_three_tabs_in_position_order():
    data = build_addon_data(BUILD)
    paladin = data.classes["paladin"]
    assert [tab.name for tab in paladin.tabs] == ["Holy", "Protection", "Retribution"]


def test_talents_keep_the_site_array_order_and_are_one_based():
    """The export encodes ranks in this order, so it is the contract, not a detail."""
    source = json.loads(Path(f"builds/{BUILD}/talents/paladin.json").read_text(encoding="utf-8"))
    holy = next(tree for tree in source["trees"] if tree["position"] == 0)
    tab = build_addon_data(BUILD).classes["paladin"].tabs[0]
    assert [talent.name for talent in tab.talents] == [t["name"] for t in holy["talents"]]
    assert tab.talents[0].tier == holy["talents"][0]["tier"] + 1
    assert tab.talents[0].column == holy["talents"][0]["column"] + 1
    assert tab.talents[0].max_rank == holy["talents"][0]["max_rank"]
    # The trait node id is what C_Traits.GetNodeInfo keys on in the 1.60 client.
    assert tab.talents[0].node == holy["talents"][0]["id"]


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
    assert len(data.weights) >= 27


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


def test_the_rendered_lua_is_loadable_by_a_real_lua(tmp_path):
    """A generated chunk that does not parse is worse than no chunk at all.

    Skipped when no interpreter is installed, the way test_simdb_build.py skips
    its Go round-trip -- a developer without Lua still gets a green suite. CI is
    not allowed that excuse: .github/workflows/data.yml's `test` job installs
    lua5.4 before pytest precisely so this gate runs there.

    The chunk goes through a file, not an argument: Data.lua passed 128 KB once
    the rotation table joined it, and Linux refuses a single argument that long
    ("Argument list too long") while macOS does not -- CI red, local green.
    """
    lua_bin = shutil.which("lua") or shutil.which("lua5.4")
    if not lua_bin:
        pytest.skip("no lua interpreter is available")
    chunk = tmp_path / "Data.lua"
    chunk.write_text(render_lua(build_addon_data(BUILD)), encoding="utf-8")
    result = subprocess.run(
        [lua_bin, "-e", f"local f, err = loadfile({str(chunk)!r}); assert(f, err)"],
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
                        talents=[AddonTalent(name="A\\B", tier=1, column=1, max_rank=1, node=7)],
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
    # build_rotations reads root/<build>/spellranks.json; an empty-but-well-formed
    # one is enough (every rotation line resolves to nothing, which is not an error).
    (tmp_path / BUILD / "spellranks.json").write_text(
        json.dumps({"build": BUILD, "classes": {}}), encoding="utf-8"
    )
    write_addon_data(BUILD, root=tmp_path)
    assert not check_addon_data(BUILD, root=tmp_path)
    (tmp_path / BUILD / "addon-data.json").write_text('{"build": "drifted"}', encoding="utf-8")
    assert check_addon_data(BUILD, root=tmp_path)


def test_bis_lua_carries_the_source_label_when_present():
    from pipeline.models import AddonBisBand, AddonBisItem, AddonData

    data = AddonData(
        build="x",
        classes={},
        weights={},
        bis={
            "hunter-marksmanship": [
                AddonBisBand(
                    level=20,
                    factions={
                        "alliance": {
                            "chest": AddonBisItem(
                                item_id=2317,
                                source_kind="crafted",
                                source="Leatherworking",
                            )
                        }
                    },
                )
            ]
        },
    )
    lua = render_lua(data)
    assert 'chest = "2317|C|Leatherworking",' in lua


def test_bis_lua_leaves_the_source_element_out_when_there_is_none():
    from pipeline.models import AddonBisBand, AddonBisItem, AddonData

    data = AddonData(
        build="x",
        classes={},
        weights={},
        bis={
            "hunter-marksmanship": [
                AddonBisBand(
                    level=20,
                    factions={
                        "alliance": {"chest": AddonBisItem(item_id=2317, source_kind="crafted")}
                    },
                )
            ]
        },
    )
    lua = render_lua(data)
    assert 'chest = "2317|C",' in lua


def test_bis_weights_lua_carries_a_measured_bands_weights():
    from pipeline.models import AddonBisBand, AddonData

    data = AddonData(
        build="x",
        classes={},
        weights={},
        bis={
            "hunter-beast-mastery": [
                AddonBisBand(level=20, factions={}, weights={"agility": 2.045}),
            ]
        },
    )
    lua = render_lua(data)
    start = lua.index("bis_weights = {")
    assert '["hunter-beast-mastery"] = {\n\t\t\t[20] = { ["agility"] = 2.045 },' in lua[start:]


def test_bis_weights_lua_skips_a_spec_with_no_measured_band():
    from pipeline.models import AddonBisBand, AddonData

    data = AddonData(
        build="x",
        classes={},
        weights={},
        bis={"hunter-beast-mastery": [AddonBisBand(level=20, factions={})]},
    )
    lua = render_lua(data)
    start = lua.index("bis_weights = {")
    assert "hunter-beast-mastery" not in lua[start:]


def test_the_cli_check_passes_on_a_fresh_regeneration(regenerated_addon_build, tmp_path):
    """The CLI's `--check` round trip: run where `builds/`, `curated/` and
    `../addon/ForeverSixty/Data.lua` are a scratch copy whose generated files
    were just regenerated, it reports no drift. It runs against a copy, never
    the committed files: those trail a curated edit until data.yml's
    `addon-data` job (on the push) or the nightly bis.yml regenerates them,
    and this suite runs before either."""
    fresh = regenerated_addon_build
    data_dir = tmp_path / "data"
    shutil.copytree(fresh.root / fresh.build, data_dir / "builds" / fresh.build)
    (data_dir / "curated").symlink_to(Path("curated").resolve())
    lua_path = tmp_path / "addon" / "ForeverSixty" / "Data.lua"
    lua_path.parent.mkdir(parents=True)
    shutil.copyfile(fresh.lua_path, lua_path)
    result = subprocess.run(
        [sys.executable, "-m", "pipeline", "addon-data", "--build", fresh.build, "--check"],
        capture_output=True,
        text=True,
        cwd=data_dir,
        env={**os.environ, "PYTHONPATH": str(Path.cwd())},
    )
    assert result.returncode == 0, result.stderr


def test_every_weighted_spec_has_a_51_point_leveling_and_raid_build():
    """data/builds/<build>/addon-builds.json (sim/cmd/addon-builds) is what
    the addon loads on a character's first login; a spec without one would
    start with no build at all."""
    data = build_addon_data(BUILD)
    for spec in data.weights:
        builds = data.default_builds.get(spec)
        assert builds is not None, f"{spec} has no default builds"
        for kind in ("leveling", "raid"):
            assert len(builds[kind].order) == 51, f"{spec} {kind}"
    arms = data.default_builds["warrior-arms"]
    assert arms["leveling"].name == "Arms leveling"
    assert arms["raid"].name == "Arms raid"
    # The codec's own triple, 1-based: Arms spends its first point in its own tab's top row.
    assert arms["leveling"].order[0][:2] == (1, 1)


def test_default_builds_are_absent_when_the_build_has_no_addon_builds_file(tmp_path):
    (tmp_path / BUILD).mkdir(parents=True)
    (tmp_path / BUILD / "classes.json").write_text("[]", encoding="utf-8")
    (tmp_path / BUILD / "spellranks.json").write_text(
        json.dumps({"build": BUILD, "classes": {}}), encoding="utf-8"
    )
    assert build_addon_data(BUILD, root=tmp_path).default_builds == {}


def test_default_builds_lua_is_triples_per_point():
    from pipeline.models import AddonData, AddonDefaultBuild

    data = AddonData(
        build="x",
        classes={},
        weights={},
        default_builds={
            "warrior-arms": {
                "leveling": AddonDefaultBuild(name="Arms leveling", order=[(1, 1, 2), (1, 2, 3)]),
                "raid": AddonDefaultBuild(name="Arms raid", order=[(1, 1, 3)]),
            }
        },
    )
    lua = render_lua(data)
    start = lua.index("defaultBuilds = {")
    assert (
        '\t\t["warrior-arms"] = {\n'
        '\t\t\tleveling = { name = "Arms leveling", order = { {1,1,2}, {1,2,3} } },\n'
        '\t\t\traid = { name = "Arms raid", order = { {1,1,3} } },\n'
        "\t\t},"
    ) in lua[start:]


def test_the_rendered_default_builds_load_in_a_real_lua(tmp_path):
    """Skipped without an interpreter, like test_the_rendered_lua_is_loadable_by_a_real_lua."""
    lua_bin = shutil.which("lua") or shutil.which("lua5.4")
    if not lua_bin:
        pytest.skip("no lua interpreter is available")
    chunk = tmp_path / "Data.lua"
    chunk.write_text(render_lua(build_addon_data(BUILD)), encoding="utf-8")
    script = (
        f"local data = assert(loadfile({str(chunk)!r}))() "
        "local build = data.defaultBuilds['warrior-arms'].leveling "
        "io.write(build.name, '|', #build.order, '|', build.order[1][1], ',', build.order[1][2])"
    )
    result = subprocess.run([lua_bin, "-e", script], capture_output=True, text=True)
    assert result.returncode == 0, result.stderr
    assert result.stdout == "Arms leveling|51|1,1"


# --- Task L1.3: dpsPerPoint, encounterNeeds, abilityLevels (contract 1) ---

#: Lucifron, Molten Core's first boss (loot.json's raid:molten-core source).
LUCIFRON = 12118
#: Execute's resolved ids in the Arms rotation: rank 1 (learned at 24) up to rank 5.
EXECUTE_RANKS = (5308, 20658, 20660, 20661, 20662)


def test_the_rendered_data_lua_carries_dps_per_point_ability_levels_and_encounter_needs(tmp_path):
    """Task L1.3 step 1, read back through a real Lua the way the addon reads it."""
    lua_bin = shutil.which("lua") or shutil.which("lua5.4") or shutil.which("lua5.1")
    if not lua_bin:
        pytest.skip("no lua interpreter is available")
    chunk = tmp_path / "Data.lua"
    chunk.write_text(render_lua(build_addon_data(BUILD)), encoding="utf-8")
    script = (
        f"local data = assert(loadfile({chunk.as_posix()!r}))() "
        "local arms = data.dpsPerPoint['warrior-arms'] "
        f"local lucifron = data.encounterNeeds[{LUCIFRON}] "
        "io.write(arms.reference, '|', tostring(arms.dps > 0), '|', tostring(arms.healing), '|', "
        f"data.abilityLevels[{EXECUTE_RANKS[-1]}], '|', #lucifron.debuffs, '|', lucifron.resist)"
    )
    result = subprocess.run([lua_bin, "-e", script], capture_output=True, text=True)
    assert result.returncode == 0, result.stderr
    reference, positive, healing, execute, debuffs, resist = result.stdout.split("|")
    assert (reference, positive, healing, execute, resist) == (
        "attack_power", "true", "false", "24", "fire"
    )
    assert int(debuffs) > 0


def test_dps_per_point_is_read_from_the_builds_sim_file():
    data = build_addon_data(BUILD)
    arms = data.dps_per_point["warrior-arms"]
    assert arms.reference == "attack_power" and arms.dps > 0 and arms.healing is False
    holy = data.dps_per_point["priest-holy"]
    assert holy.reference == "healing_power" and holy.healing is True
    # A tank's reference stat (stamina) does no damage: no entry, so the
    # addon falls back to "score" units for it.
    assert "warrior-protection" not in data.dps_per_point


def test_dps_per_point_is_absent_when_the_build_has_no_sim_file(tmp_path):
    (tmp_path / BUILD).mkdir(parents=True)
    (tmp_path / BUILD / "classes.json").write_text("[]", encoding="utf-8")
    (tmp_path / BUILD / "spellranks.json").write_text(
        json.dumps({"build": BUILD, "classes": {}}), encoding="utf-8"
    )
    data = build_addon_data(BUILD, root=tmp_path)
    assert data.dps_per_point == {}
    # No loot.json either: no raid record to check npc ids against, so none ship.
    assert data.encounter_needs == {}


def test_ability_levels_are_the_first_ranks_learn_level_for_every_rotation_id():
    data = build_addon_data(BUILD)
    for spell_id in EXECUTE_RANKS:
        assert data.ability_levels[spell_id] == 24, spell_id
    for spec, bands in data.rotations.items():
        for band in bands:
            for line in band.lines:
                level = data.ability_levels.get(line.spell_id)
                assert level is not None, f"{spec} {band.level} {line.name} has no ability level"
                # A band only shows abilities learned by its level.
                assert level <= band.level, f"{spec} {band.level} {line.name} learned at {level}"


def test_encounter_needs_cover_molten_core_and_onyxia():
    needs = build_addon_data(BUILD).encounter_needs
    molten_core = [12118, 11982, 12259, 12057, 12056, 12264, 12098, 11988, 12018, 11502]
    for npc_id in molten_core:
        assert needs[npc_id].resist == "fire", npc_id
        assert needs[npc_id].debuffs and needs[npc_id].buffs, npc_id
    assert needs[10184].resist is None  # Onyxia
    assert 11597 in needs[LUCIFRON].debuffs  # Sunder Armor


def test_the_three_additions_render_as_contract_one_shapes():
    from pipeline.models import AddonData, AddonDpsPerPoint, AddonEncounterNeeds

    data = AddonData(
        build="x",
        classes={},
        weights={},
        dps_per_point={
            "priest-holy": AddonDpsPerPoint(reference="healing_power", dps=0.305, healing=True),
            "warrior-arms": AddonDpsPerPoint(reference="attack_power", dps=0.2342, healing=False),
        },
        encounter_needs={
            10184: AddonEncounterNeeds(buffs=[11551], debuffs=[11597, 9907]),
            12118: AddonEncounterNeeds(buffs=[11551, 9885], debuffs=[11597], resist="fire"),
        },
        ability_levels={20662: 24, 5308: 24},
    )
    lua = render_lua(data)
    assert (
        "\tdpsPerPoint = {\n"
        '\t\t["priest-holy"] = { reference = "healing_power", dps = 0.305, healing = true },\n'
        '\t\t["warrior-arms"] = { reference = "attack_power", dps = 0.2342, healing = false },\n'
        "\t},\n"
    ) in lua
    assert (
        "\tencounterNeeds = {\n"
        "\t\t[10184] = { buffs = { 11551 }, debuffs = { 11597, 9907 } },\n"
        '\t\t[12118] = { buffs = { 11551, 9885 }, debuffs = { 11597 }, resist = "fire" },\n'
        "\t},\n"
    ) in lua
    assert "\tabilityLevels = {\n\t\t[5308] = 24,\n\t\t[20662] = 24,\n\t},\n" in lua


def test_the_real_data_lua_names_every_pick_and_prices_its_weapons(tmp_path):
    """Task L1.4: bisNames/bisDps read back through a real Lua. The Arms
    band-20 main hand is a weapon, so it has both a name and a dps; the
    head pick is not a weapon, so it has a name and no dps."""
    lua_bin = shutil.which("lua") or shutil.which("lua5.4") or shutil.which("lua5.1")
    if not lua_bin:
        pytest.skip("no lua interpreter is available")
    chunk = tmp_path / "Data.lua"
    chunk.write_text(render_lua(build_addon_data(BUILD)), encoding="utf-8")
    script = (
        f"local data = assert(loadfile({chunk.as_posix()!r}))() "
        "local band = data.bis['warrior-arms'][20].alliance "
        "local weapon = tonumber(band.main_hand:match('^(%d+)')) "
        "local head = tonumber(band.head:match('^(%d+)')) "
        "io.write(data.bisNames[weapon], '|', tostring(data.bisDps[weapon]), '|', "
        "data.bisNames[head], '|', tostring(data.bisDps[head]))"
    )
    result = subprocess.run([lua_bin, "-e", script], capture_output=True, text=True)
    assert result.returncode == 0, result.stderr
    weapon_name, weapon_dps, head_name, head_dps = result.stdout.split("|")
    assert weapon_name and weapon_name != "The Axe of Severing"
    assert float(weapon_dps) > 0
    assert head_name
    assert head_dps == "nil"


def test_every_real_pick_and_alternative_is_named():
    data = build_addon_data(BUILD)
    for spec, bands in data.bis.items():
        for band in bands:
            for slots in band.factions.values():
                for pick in slots.values():
                    for item_id in [pick.item_id] + [alt.item_id for alt in pick.alternatives]:
                        assert data.bis_names.get(item_id), (spec, band.level, item_id)
    assert set(data.bis_dps) <= set(data.bis_names)
    assert all(value > 0 and round(value, 1) == value for value in data.bis_dps.values())


def _wsl_path(path: Path) -> str:
    """C:\\x\\y -> /mnt/c/x/y, for a file WSL's lua5.1 must open."""
    resolved = path.resolve()
    drive = resolved.drive.rstrip(":").lower()
    return f"/mnt/{drive}" + resolved.as_posix()[len(resolved.drive) :]


def _lua51() -> tuple[list[str], object] | None:
    """The command prefix for a Lua 5.1 (the client's own) and a function
    mapping a local path to the path that Lua opens: `lua5.1` on PATH, else
    WSL's `lua5.1`. None when neither is available."""
    local = shutil.which("lua5.1")
    if local:
        return [local], lambda path: path.as_posix()
    wsl = shutil.which("wsl")
    if wsl:
        probe = subprocess.run(
            [wsl, "-e", "lua5.1", "-v"], capture_output=True, text=True, timeout=60
        )
        if probe.returncode == 0:
            return [wsl, "-e", "lua5.1"], _wsl_path
    return None


def test_the_real_data_lua_stays_inside_the_two_megabyte_idle_budget(tmp_path):
    """Ruling 3 (contract 1's 2 MB idle budget), measured the way L1.2/L1.3
    did, under Lua 5.1 only (fix round 1, item 3: Lua 5.4's heap is smaller
    and would pass a file the client cannot hold): the heap delta of loading
    and calling Data.lua, after full collections either side."""
    lua51 = _lua51()
    if lua51 is None:
        pytest.skip("no Lua 5.1 (neither lua5.1 on PATH nor WSL's lua5.1): the budget is 5.1's")
    command, path_for = lua51
    chunk = tmp_path / "Data.lua"
    chunk.write_text(render_lua(build_addon_data(BUILD)), encoding="utf-8")
    script = tmp_path / "heap.lua"
    script.write_text(
        "collectgarbage() collectgarbage() local before = collectgarbage('count') "
        "local data = assert(loadfile(arg[1]))() "
        "collectgarbage() collectgarbage() "
        "io.write(_VERSION, '|', string.format('%.1f', collectgarbage('count') - before))",
        encoding="utf-8",
    )
    result = subprocess.run(
        command + [path_for(script), path_for(chunk)], capture_output=True, text=True, timeout=120
    )
    assert result.returncode == 0, result.stderr
    version, heap = result.stdout.split("|")
    assert version == "Lua 5.1"
    assert float(heap) < 2048


def test_the_cli_logs_the_unconfirmed_skips_audit_line():
    result = subprocess.run(
        [sys.executable, "-m", "pipeline", "addon-data", "--build", BUILD, "--check"],
        capture_output=True,
        text=True,
    )
    assert f"bis {BUILD}: skipped " in result.stderr
    assert " client_unconfirmed candidates (" in result.stderr


def test_one_cli_run_parses_each_bis_file_once(monkeypatch):
    import pipeline.addonbis as addonbis
    from pipeline.__main__ import main

    parsed: list[str] = []
    real = addonbis._spec_bis

    def counting(path, *args, **kwargs):
        parsed.append(path.stem)
        return real(path, *args, **kwargs)

    monkeypatch.setattr(addonbis, "_spec_bis", counting)
    main(["addon-data", "--build", BUILD, "--check"])
    specs = sorted(path.stem for path in (Path("builds") / BUILD / "bis").glob("*.json"))
    assert sorted(parsed) == specs


# --- Task L1.5: dpsPerPoint per spec and band, with a measured weapon term ---

#: The figures are bis.yml's nightly re-measure (sim/cmd/dps-per-point), so the
#: tests read them from the committed dps-per-point.json, the generator's own
#: input, and assert they are carried through, never a pinned measurement: the
#: 2026-11-01 re-measure moved the Arms band-20 weapon term 2.1% and failed the
#: pin (memory: data contracts are floors).
#: The per-band model's threshold (task L1.5, ruling 2).
BAND_DIFFERS = 0.25


def _measured_dps_per_point(spec: str) -> dict:
    path = Path("builds") / BUILD / "dps-per-point.json"
    return json.loads(path.read_text(encoding="utf-8"))[spec]


def test_dps_per_point_carries_each_bands_figure_and_weapon_term():
    data = build_addon_data(BUILD)
    arms = data.dps_per_point["warrior-arms"]
    measured = _measured_dps_per_point("warrior-arms")
    assert arms.dps == pytest.approx(measured["dps"])
    assert arms.weapon_dps == pytest.approx(measured["weapon_dps"])
    assert sorted(arms.bands) == [20, 30, 40, 50, 60]
    for level, band in arms.bands.items():
        assert band.dps == pytest.approx(measured["bands"][str(level)]["dps"])
        assert band.weapon_dps == pytest.approx(measured["bands"][str(level)]["weapon_dps"])
        assert band.dps > 0 and band.weapon_dps > 0
    band = arms.bands[20]
    assert band.reference == "attack_power"
    # A higher band's character gains more from a point than a lower band's.
    assert arms.bands[30].dps > band.dps
    assert arms.bands[30].weapon_dps > band.weapon_dps
    # The reason for the per-band model: the level-60 figure is far off at 20.
    assert abs(band.dps - arms.dps) / arms.dps > BAND_DIFFERS


def test_a_casters_and_a_healers_weapon_term_is_zero_in_their_reference_stat():
    data = build_addon_data(BUILD)
    frost, holy = data.dps_per_point["mage-frost"], data.dps_per_point["priest-holy"]
    assert frost.reference == "spell_power" and holy.reference == "healing_power"
    assert holy.healing is True
    for entry in (frost, holy):
        assert entry.weapon_dps == 0
        assert entry.bands, entry.reference
        for band in entry.bands.values():
            assert band.reference == entry.reference and band.weapon_dps == 0 and band.dps > 0


def test_dps_per_point_bands_render_beside_the_spec_figure():
    from pipeline.models import AddonData, AddonDpsPerPoint, AddonDpsPerPointBand

    data = AddonData(
        build="x",
        classes={},
        weights={},
        dps_per_point={
            "warrior-arms": AddonDpsPerPoint(
                reference="attack_power",
                dps=0.2222,
                healing=False,
                weapon_dps=3.1536,
                bands={
                    30: AddonDpsPerPointBand(
                        reference="attack_power", dps=0.0793, weapon_dps=0.9831
                    ),
                    20: AddonDpsPerPointBand(
                        reference="attack_power", dps=0.0385, weapon_dps=0.5432
                    ),
                },
            ),
        },
    )
    assert (
        "\tdpsPerPoint = {\n"
        '\t\t["warrior-arms"] = { reference = "attack_power", dps = 0.2222, healing = false, '
        "weaponDps = 3.1536,\n"
        '\t\t\t[20] = { reference = "attack_power", dps = 0.0385, weaponDps = 0.5432 },\n'
        '\t\t\t[30] = { reference = "attack_power", dps = 0.0793, weaponDps = 0.9831 },\n'
        "\t\t},\n"
        "\t},\n"
    ) in render_lua(data)


def test_the_real_data_lua_reads_dps_per_point_by_spec_and_band(tmp_path):
    lua_bin = shutil.which("lua") or shutil.which("lua5.4") or shutil.which("lua5.1")
    if not lua_bin:
        pytest.skip("no lua interpreter is available")
    chunk = tmp_path / "Data.lua"
    chunk.write_text(render_lua(build_addon_data(BUILD)), encoding="utf-8")
    script = (
        f"local data = assert(loadfile({chunk.as_posix()!r}))() "
        "local arms = data.dpsPerPoint['warrior-arms'] "
        "local frost = data.dpsPerPoint['mage-frost'][20] "
        "io.write(arms.reference, '|', arms[20].reference, '|', arms[20].dps, '|', "
        "arms[20].weaponDps, '|', arms.dps, '|', frost.weaponDps)"
    )
    result = subprocess.run([lua_bin, "-e", script], capture_output=True, text=True)
    assert result.returncode == 0, result.stderr
    reference, band_reference, dps, weapon, spec_dps, frost_weapon = result.stdout.split("|")
    assert (reference, band_reference, float(frost_weapon)) == ("attack_power", "attack_power", 0)
    measured = _measured_dps_per_point("warrior-arms")
    # Lua prints numbers rounded to its own precision.
    assert float(dps) == pytest.approx(measured["bands"]["20"]["dps"], rel=1e-3)
    assert float(weapon) == pytest.approx(measured["bands"]["20"]["weapon_dps"], rel=1e-3)
    assert float(spec_dps) == pytest.approx(measured["dps"], rel=1e-3)
