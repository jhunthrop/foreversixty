"""pipeline.icons_fix: the standalone `python -m pipeline icons --build <build>
[--engine <path>]` pass that rewrites an already-committed items/<class>.json's
placeholder `icon` fields, without needing a fresh raw ItemSparse export.
"""

import json
from pathlib import Path

from pipeline.icons import PLACEHOLDER_ICON
from pipeline.icons_fix import (
    fix_icons_for_build,
    fix_placeholder_icons,
    load_fork_icons,
    load_wowhead_icons,
)

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"
WOWHEAD_FIXTURE = HERE / "fixtures" / "wowhead-gear-planner.js"


def _gear_item(item_id: int, icon: str) -> dict:
    return {
        "id": item_id,
        "name": f"Item {item_id}",
        "icon": icon,
        "slot": "head",
        "quality": 3,
        "required_level": 1,
        "item_level": 1,
        "armor": 0,
        "stats": {},
        "set_id": None,
        "unique": False,
    }


def write_class_items(build_dir: Path, class_slug: str, items: list[dict]) -> Path:
    path = build_dir / "items" / f"{class_slug}.json"
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(
        json.dumps({"build": "1.0.0.1", "class_slug": class_slug, "items": items}),
        encoding="utf-8",
    )
    return path


def test_load_fork_icons_reads_the_engine_checkout():
    assert load_fork_icons(ENGINE)[100] == "inv_a"


def test_load_wowhead_icons_reads_every_payload_id(tmp_path: Path):
    build_dir = tmp_path / "1.0.0.1"
    (build_dir / "raw").mkdir(parents=True)
    (build_dir / "raw" / "wowhead-gear-planner.js").write_text(
        WOWHEAD_FIXTURE.read_text(encoding="utf-8"), encoding="utf-8"
    )
    icons = load_wowhead_icons(build_dir)
    assert icons[11726] == "inv_belt_03"
    assert icons[279865] == "inv_misc_cape_10"


def test_load_wowhead_icons_is_empty_without_a_payload(tmp_path: Path):
    build_dir = tmp_path / "1.0.0.1"
    build_dir.mkdir()
    assert load_wowhead_icons(build_dir) == {}


def test_fix_placeholder_icons_rewrites_only_placeholder_rows(tmp_path: Path):
    build_dir = tmp_path / "1.0.0.1"
    path = write_class_items(
        build_dir,
        "warrior",
        [_gear_item(1, "inv_sword_04"), _gear_item(2, PLACEHOLDER_ICON)],
    )
    results = fix_placeholder_icons(build_dir, fork_icons={}, wowhead_icons={2: "inv_helm_09"})
    assert len(results) == 1
    result = results[0]
    assert (result.total_items, result.before_placeholder, result.after_placeholder) == (2, 1, 0)
    assert (result.fixed_by_fork, result.fixed_by_wowhead) == (0, 1)
    written = json.loads(path.read_text(encoding="utf-8"))
    icons = {item["id"]: item["icon"] for item in written["items"]}
    assert icons == {1: "inv_sword_04", 2: "inv_helm_09"}


def test_fix_placeholder_icons_prefers_the_fork_source(tmp_path: Path):
    build_dir = tmp_path / "1.0.0.1"
    write_class_items(build_dir, "warrior", [_gear_item(1, PLACEHOLDER_ICON)])
    results = fix_placeholder_icons(
        build_dir, fork_icons={1: "inv_fork"}, wowhead_icons={1: "inv_wowhead"}
    )
    assert (results[0].fixed_by_fork, results[0].fixed_by_wowhead) == (1, 0)


def test_fix_placeholder_icons_leaves_an_unresolvable_item_on_the_placeholder(tmp_path: Path):
    build_dir = tmp_path / "1.0.0.1"
    path = write_class_items(build_dir, "warrior", [_gear_item(1, PLACEHOLDER_ICON)])
    before_text = path.read_text(encoding="utf-8")
    results = fix_placeholder_icons(build_dir, fork_icons={}, wowhead_icons={})
    assert (results[0].before_placeholder, results[0].after_placeholder) == (1, 1)
    # Nothing changed, so the file is left exactly as it was written.
    assert path.read_text(encoding="utf-8") == before_text


def test_fix_placeholder_icons_does_not_rewrite_a_class_with_no_placeholders(tmp_path: Path):
    build_dir = tmp_path / "1.0.0.1"
    path = write_class_items(build_dir, "warrior", [_gear_item(1, "inv_sword_04")])
    before_stat = path.stat().st_mtime_ns
    results = fix_placeholder_icons(build_dir, fork_icons={}, wowhead_icons={1: "inv_other"})
    assert results[0].before_placeholder == 0
    assert path.stat().st_mtime_ns == before_stat


def test_fix_icons_for_build_orchestrates_fork_and_wowhead_sources(tmp_path: Path):
    build_dir = tmp_path / "1.0.0.1"
    write_class_items(
        build_dir,
        "warrior",
        [_gear_item(100, PLACEHOLDER_ICON), _gear_item(11726, PLACEHOLDER_ICON)],
    )
    (build_dir / "raw").mkdir(parents=True)
    (build_dir / "raw" / "wowhead-gear-planner.js").write_text(
        WOWHEAD_FIXTURE.read_text(encoding="utf-8"), encoding="utf-8"
    )
    results = fix_icons_for_build("1.0.0.1", root=tmp_path, engine=ENGINE)
    result = results[0]
    # 100 comes from the fork db (`assets/database/db.json`), 11726 only from wowhead.
    assert (result.fixed_by_fork, result.fixed_by_wowhead) == (1, 1)
    assert result.after_placeholder == 0


def test_fix_icons_for_build_without_an_engine_only_tries_wowhead(tmp_path: Path):
    build_dir = tmp_path / "1.0.0.1"
    write_class_items(build_dir, "warrior", [_gear_item(11726, PLACEHOLDER_ICON)])
    (build_dir / "raw").mkdir(parents=True)
    (build_dir / "raw" / "wowhead-gear-planner.js").write_text(
        WOWHEAD_FIXTURE.read_text(encoding="utf-8"), encoding="utf-8"
    )
    results = fix_icons_for_build("1.0.0.1", root=tmp_path, engine=None)
    assert (results[0].fixed_by_fork, results[0].fixed_by_wowhead) == (0, 1)


def test_fix_icons_for_build_warns_with_neither_source(tmp_path: Path, caplog):
    build_dir = tmp_path / "1.0.0.1"
    write_class_items(build_dir, "warrior", [_gear_item(1, PLACEHOLDER_ICON)])
    with caplog.at_level("WARNING"):
        results = fix_icons_for_build("1.0.0.1", root=tmp_path, engine=None)
    assert results[0].after_placeholder == 1
    assert "no placeholder icon names can be fixed" in caplog.text
