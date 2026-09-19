import json
import warnings
from pathlib import Path

import pytest

from pipeline.curated import CuratedError
from pipeline.loot.overlay import OverlayError, apply_overlays, load_overlays
from pipeline.models import LootBoss, LootFile, LootSource

HERE = Path(__file__).parent
FIXTURE = HERE / "fixtures/loot/curated"
CURATED = Path("curated/loot")


def base() -> LootFile:
    return LootFile(
        sources=[
            LootSource(
                id="raid:molten-core",
                kind="raid",
                name="Molten Core",
                zone_id=2717,
                bosses=[LootBoss(id="raid:molten-core:900", name="Big Boss", npc_id=900,
                                 items=[100])],
                trash=[101],
            ),
            LootSource(
                id="dungeon:the-deadmines",
                kind="dungeon",
                name="The Deadmines",
                zone_id=1581,
                bosses=[LootBoss(id="dungeon:the-deadmines:902", name="Dungeon Boss",
                                 npc_id=902, items=[103])],
            ),
        ]
    )


def write(tmp_path: Path, name: str, payload: dict) -> Path:
    (tmp_path / name).write_text(json.dumps(payload), encoding="utf-8")
    return tmp_path


def test_a_missing_overlay_directory_is_not_an_error(tmp_path):
    assert load_overlays(tmp_path / "nope") == []


def test_the_fixture_overlays_add_patch_and_remove():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    by_id = {source.id: source for source in result.sources}
    assert set(by_id) == {"raid:molten-core", "raid:barrow-deeps"}
    assert by_id["raid:molten-core"].opens == "later"
    assert by_id["raid:barrow-deeps"].items == [100]


def test_a_replace_only_touches_the_keys_it_names():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert molten.name == "Molten Core"
    assert molten.zone_id == 2717
    assert molten.trash == [101]
    assert [b.items for b in molten.bosses] == [[100]]


def test_the_result_stays_sorted_by_kind_then_id():
    result = apply_overlays(base(), load_overlays(FIXTURE))
    assert [s.id for s in result.sources] == ["raid:barrow-deeps", "raid:molten-core"]


def test_an_overlay_with_no_source_is_refused(tmp_path):
    write(tmp_path, "a.json", {"notes": "x", "add": []})
    with pytest.raises(CuratedError):
        load_overlays(tmp_path)


def test_an_overlay_source_kind_outside_the_vocabulary_is_refused(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "u", "kind": "hearsay"}], "notes": "x"})
    with pytest.raises(CuratedError, match="hearsay"):
        load_overlays(tmp_path)


def test_adding_a_source_that_already_exists_is_refused(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "add": [{"id": "raid:molten-core", "kind": "raid", "name": "Again"}],
    })
    with pytest.raises(OverlayError, match="already"):
        apply_overlays(base(), load_overlays(tmp_path))


def test_replacing_or_removing_an_unknown_source_is_refused(tmp_path):
    for payload in (
        {"replace": [{"id": "raid:nope", "opens": "later"}]},
        {"remove": ["raid:nope"]},
    ):
        write(tmp_path, "a.json", {
            "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
            "notes": "x",
        } | payload)
        with pytest.raises(OverlayError, match="raid:nope"):
            apply_overlays(base(), load_overlays(tmp_path))


def test_a_replace_may_restate_a_sources_item_lists(tmp_path):
    # Contract 10.4 has no per-item removal, so this is how a wrong item
    # is corrected: restate the list it is in.
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "replace": [{"id": "raid:molten-core", "trash": []}],
    })
    result = apply_overlays(base(), load_overlays(tmp_path))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert molten.trash == []
    assert [b.items for b in molten.bosses] == [[100]]


def test_a_replace_may_not_change_a_sources_kind(tmp_path):
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "replace": [{"id": "raid:molten-core", "kind": "dungeon"}],
    })
    with pytest.raises(OverlayError, match="kind"):
        apply_overlays(base(), load_overlays(tmp_path))


def test_a_replace_touching_bosses_yields_real_loot_boss_instances(tmp_path):
    # model_copy(update=...) assigns a patch's dumped dict verbatim, without
    # re-validating; a `replace` that names `bosses` must still come back
    # out as `LootBoss` instances, not raw dicts, since `LootSourcePatch`
    # explicitly advertises `bosses` as a replaceable field.
    write(tmp_path, "a.json", {
        "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
        "notes": "x",
        "replace": [{
            "id": "raid:molten-core",
            "bosses": [
                {"id": "raid:molten-core:900", "name": "New Boss", "npc_id": 900,
                 "items": [100, 200]}
            ],
        }],
    })
    result = apply_overlays(base(), load_overlays(tmp_path))
    molten = next(s for s in result.sources if s.id == "raid:molten-core")
    assert len(molten.bosses) == 1
    boss = molten.bosses[0]
    assert isinstance(boss, LootBoss)
    assert boss.name == "New Boss"
    assert boss.npc_id == 900
    assert boss.items == [100, 200]
    with warnings.catch_warnings():
        warnings.simplefilter("error")
        molten.model_dump()


def test_the_real_curated_overlays_all_load_and_apply_cleanly():
    """The committed overlays must be loadable on their own terms; that
    they match the committed loot.json is tests/test_loot_build.py's job."""
    for path, document in load_overlays(CURATED):
        assert document.sources, path
        assert document.notes.strip(), path
