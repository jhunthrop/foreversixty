"""levels.json: per-level base stats, race offsets and spell crit per intellect,
from wowhead's Forever gear planner (docs/superpowers/specs/
2026-09-27-level-aware-sim-design.md).

tests/fixtures/wowhead-gear-planner.js (also test_wowhead_items.py's fixture) carries
a baseStats/critSpell pair for two classes -- warrior (1, no spell-crit curve) and
paladin (2, one) -- and two races -- human (1, every offset zero) and a Skyborne-style
row (95, nonzero offsets) -- with warrior's level 60 pinned to the engine's own
ClassBaseStats: 1689 health, 110 stamina.
"""

import json
from pathlib import Path

import pytest

from pipeline import levels
from pipeline.models import LevelsSource
from pipeline.wowhead_items import parse_page_data

FIXTURE = Path(__file__).parent / "fixtures" / "wowhead-gear-planner.js"
CLASS_ROWS = [
    {"id": 1, "name": "Warrior", "slug": "warrior"},
    {"id": 2, "name": "Paladin", "slug": "paladin"},
]
RACE_ROWS = [
    {"id": 1, "name": "Human", "slug": "human"},
    {"id": 95, "name": "High-Order Skyborne", "slug": "high-order-skyborne"},
]
SOURCE = LevelsSource(url="https://nether.wowhead.com/forever/data/gear-planner", fetched_at="t")


def _payload() -> tuple[dict, dict]:
    text = FIXTURE.read_text(encoding="utf-8")
    return (
        parse_page_data(text, levels.BASESTATS_KEY),
        parse_page_data(text, levels.CRITSPELL_KEY),
    )


def _build() -> levels.LevelsFile:
    base_stats, crit_spell = _payload()
    return levels.build_levels_file(
        "9.9.9.9", base_stats, crit_spell, SOURCE, CLASS_ROWS, RACE_ROWS
    )


def test_every_class_has_sixty_levels_numbered_one_through_sixty():
    result = _build()
    assert set(result.classes) == {"warrior", "paladin"}
    for class_levels in result.classes.values():
        assert [row.level for row in class_levels.levels] == list(range(1, 61))


def test_warrior_level_sixty_matches_the_engines_class_base_stats():
    warrior = _build().classes["warrior"].levels[-1]
    assert warrior.level == 60
    assert warrior.health == 1689
    assert warrior.stamina == 110


def test_spell_crit_per_int_is_null_only_for_a_class_with_no_curve():
    result = _build()
    assert all(row.spell_crit_per_int is None for row in result.classes["warrior"].levels)
    paladin = result.classes["paladin"].levels
    assert paladin[0].spell_crit_per_int == pytest.approx(0.075)
    assert paladin[-1].spell_crit_per_int == pytest.approx(0.019)
    assert all(row.spell_crit_per_int is not None for row in paladin)


def test_race_offsets_are_keyed_by_slug_not_id():
    result = _build()
    assert set(result.race_offsets) == {"human", "high-order-skyborne"}
    assert result.race_offsets["human"].model_dump() == {
        "agility": 0,
        "strength": 0,
        "intellect": 0,
        "spirit": 0,
        "stamina": 0,
    }
    assert result.race_offsets["high-order-skyborne"].model_dump() == {
        "agility": 1,
        "strength": 0,
        "intellect": 2,
        "spirit": 0,
        "stamina": 1,
    }


def test_an_unmappable_class_id_raises():
    base_stats, crit_spell = _payload()
    base_stats = {**base_stats, "stats": {**base_stats["stats"], "999": base_stats["stats"]["1"]}}
    with pytest.raises(levels.UnmappableIdError, match="999"):
        levels.build_levels_file("9.9.9.9", base_stats, crit_spell, SOURCE, CLASS_ROWS, RACE_ROWS)


def test_an_unmappable_race_id_raises():
    base_stats, crit_spell = _payload()
    base_stats = {
        **base_stats,
        "raceOffsets": {**base_stats["raceOffsets"], "999": base_stats["raceOffsets"]["1"]},
    }
    with pytest.raises(levels.UnmappableIdError, match="999"):
        levels.build_levels_file("9.9.9.9", base_stats, crit_spell, SOURCE, CLASS_ROWS, RACE_ROWS)


def test_output_is_deterministic():
    first, second = _build().model_dump(), _build().model_dump()
    assert first == second
    # classes and race_offsets are both keyed by slug, sorted -- not by
    # whatever order the wowhead payload's own dict happened to state.
    assert list(first["classes"]) == ["paladin", "warrior"]


def test_load_levels_is_none_without_a_wowhead_payload(tmp_path: Path, caplog) -> None:
    build_dir = tmp_path / "builds" / "1.0.0.1"
    build_dir.mkdir(parents=True)
    with caplog.at_level("INFO"):
        assert levels.load_levels("1.0.0.1", root=tmp_path / "builds") is None
    assert "no wowhead payload" in caplog.text


def _write_build_dir(build_dir: Path) -> None:
    build_dir.mkdir(parents=True)
    raw = build_dir / "raw"
    raw.mkdir()
    (raw / "wowhead-gear-planner.js").write_text(FIXTURE.read_text(encoding="utf-8"))
    meta = {"url": "https://example.test/gear-planner", "fetched_at": "2026-09-27T00:00:00Z"}
    (raw / "wowhead-gear-planner.meta.json").write_text(json.dumps(meta))
    (build_dir / "classes.json").write_text(json.dumps(CLASS_ROWS))
    (build_dir / "races.json").write_text(json.dumps(RACE_ROWS))
    manifest = {
        "build": "1.0.0.1",
        "product": "test",
        "fetched_at": "2026-09-27T00:00:00Z",
        "files": {},
    }
    (build_dir / "manifest.json").write_text(json.dumps(manifest))


def test_write_levels_writes_the_file_and_carries_the_metas_provenance(tmp_path: Path) -> None:
    build_dir = tmp_path / "builds" / "1.0.0.1"
    _write_build_dir(build_dir)
    path = levels.write_levels("1.0.0.1", root=tmp_path / "builds")
    assert path == build_dir / "levels.json"
    written = json.loads(path.read_text(encoding="utf-8"))
    assert written["build"] == "1.0.0.1"
    assert written["source"] == {
        "url": "https://example.test/gear-planner",
        "fetched_at": "2026-09-27T00:00:00Z",
    }
    assert written["classes"]["warrior"]["levels"][-1]["stamina"] == 110


def test_write_levels_raises_without_the_meta_file(tmp_path: Path) -> None:
    build_dir = tmp_path / "builds" / "1.0.0.1"
    _write_build_dir(build_dir)
    (build_dir / "raw" / "wowhead-gear-planner.meta.json").unlink()
    with pytest.raises(SystemExit, match="meta.json"):
        levels.write_levels("1.0.0.1", root=tmp_path / "builds")
