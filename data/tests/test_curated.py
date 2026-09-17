import json
from pathlib import Path

import pytest

from pipeline.csvio import read_csv
from pipeline.curated import SOURCE_KINDS, CuratedError, merge_curated
from pipeline.models import PlayableClass, PlayableRace
from pipeline.normalize import write_json, write_records
from pipeline.normalize.classes import normalize_classes, normalize_races

HERE = Path(__file__).parent
#: The 1.60 client (Forever beta) is the first to export real Skyborne rows, so it is
#: the build the real curated/ directory is now validated against: it is a strict
#: superset of era's races and classes (every Classic race and class, plus the two
#: real Skyborne rows), where era has no Skyborne race at all. See the beta-data
#: report for why the Skyborne placeholder (a single synthetic race id 900) was
#: retired in favour of the client's own two per-faction rows rather than kept as a
#: build-conditional fallback.
BETA_BUILD = Path("builds/1.60.1.69893")

#: The Forever race and class pairs the research documents, as (race_id, class_id).
#: Skyborne is now the client's own two real per-faction rows: 95 (High Order
#: Skyborne, Alliance) gets the four default classes plus Mage, 96 (Windshaper
#: Skyborne, Horde) gets the four default classes plus Shaman.
SKYBORNE_DEFAULT_CLASSES = frozenset({1, 3, 4, 11})
SKYBORNE_COMBOS = {(95, c) for c in SKYBORNE_DEFAULT_CLASSES | {8}} | {
    (96, c) for c in SKYBORNE_DEFAULT_CLASSES | {7}
}
#: Six Deep Dive pairs:
#: Dwarf Shaman, Undead Paladin, Gnome Priest, Human Hunter, Orc Mage, Troll Warlock.
DEEP_DIVE_COMBOS = {(3, 7), (5, 2), (7, 5), (1, 3), (2, 8), (8, 9)}
NEW_COMBOS = DEEP_DIVE_COMBOS | SKYBORNE_COMBOS
#: How many pairs vanilla itself allows; the pairs above are the only additions.
CLASSIC_COMBO_COUNT = 40


def client_rows():
    return (
        normalize_classes(read_csv(HERE / "fixtures/ChrClasses.csv")),
        normalize_races(read_csv(HERE / "fixtures/ChrRaces.csv")),
    )


def merged():
    classes, races = client_rows()
    return merge_curated(classes, races, HERE / "fixtures/curated")


def test_forever_changes_are_attached_by_slug():
    classes, _races, _combos = merged()
    warrior = {c.slug: c for c in classes}["warrior"]
    assert [change.text for change in warrior.forever_changes] == ["Warriors gain a fourth tree."]
    assert warrior.forever_changes[0].sources[0].url == "https://example.test/forever-warrior"


def test_classes_without_curated_entries_keep_an_empty_list():
    classes, _races, _combos = merged()
    assert {c.slug: c for c in classes}["paladin"].forever_changes == []


def test_a_placeholder_race_is_appended_and_marked():
    _classes, races, _combos = merged()
    skyborne = {r.slug: r for r in races}["skyborne"]
    assert (skyborne.id, skyborne.faction, skyborne.placeholder) == (900, "neutral", True)
    assert {r.slug: r for r in races}["human"].placeholder is False


def test_a_client_race_with_no_combo_is_not_emitted():
    """ChrRaces carries races a player cannot pick; a race with zero legal classes
    would render as an empty picker. Only Human is in the fixture combos, so Orc and
    Pandaren drop out and the placeholder Skyborne stays without needing one."""
    _classes, races, _combos = merged()
    assert [race.slug for race in races] == ["human", "skyborne"]


def test_combos_are_emitted_sorted(tmp_path: Path):
    _classes, _races, combos = merged()
    out = tmp_path / "combos.json"
    write_records(combos, out)
    assert json.loads(out.read_text()) == [
        {"race_id": 1, "class_id": 1, "new_in_forever": False},
        {"race_id": 1, "class_id": 2, "new_in_forever": True},
    ]
    assert out.read_text() == (HERE / "golden/combos.json").read_text()


def test_a_change_with_no_source_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text(
        json.dumps(
            [{"slug": "warrior", "forever_changes": [{"text": "Trust me.", "sources": []}]}]
        )
    )
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text("[]")
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="at least one source"):
        merge_curated(classes, races, tmp_path)


def test_an_unknown_source_kind_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text(
        json.dumps(
            [
                {
                    "slug": "warrior",
                    "forever_changes": [
                        {
                            "text": "x",
                            "sources": [{"label": "l", "url": "https://e.test", "kind": "rumour"}],
                        }
                    ],
                }
            ]
        )
    )
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text("[]")
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="rumour"):
        merge_curated(classes, races, tmp_path)


def test_an_unknown_slug_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text(
        json.dumps([{"slug": "necromancer", "forever_changes": []}])
    )
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text("[]")
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="necromancer"):
        merge_curated(classes, races, tmp_path)


def test_a_combo_naming_an_unknown_id_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text("[]")
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text(
        json.dumps([{"race_id": 1, "class_id": 99, "new_in_forever": False}])
    )
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="class_id 99"):
        merge_curated(classes, races, tmp_path)


def test_a_new_combo_without_a_source_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text("[]")
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text(
        json.dumps([{"race_id": 1, "class_id": 1, "new_in_forever": True}])
    )
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="at least one source"):
        merge_curated(classes, races, tmp_path)


def test_a_placeholder_race_may_not_reuse_a_client_race_id(tmp_path: Path):
    (tmp_path / "classes.json").write_text("[]")
    (tmp_path / "races.json").write_text(
        json.dumps(
            [
                {
                    "id": 1,
                    "name": "Skyborne",
                    "slug": "skyborne",
                    "faction": "neutral",
                    "placeholder": True,
                    "forever_changes": [],
                }
            ]
        )
    )
    (tmp_path / "combos.json").write_text("[]")
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="race id 1"):
        merge_curated(classes, races, tmp_path)


def test_merged_classes_and_races_match_golden(tmp_path: Path):
    classes, races, _combos = merged()
    write_json(classes, tmp_path / "classes.json")
    write_json(races, tmp_path / "races.json")
    assert (tmp_path / "classes.json").read_text() == (
        HERE / "golden/classes_merged.json"
    ).read_text()
    assert (tmp_path / "races.json").read_text() == (
        HERE / "golden/races_merged.json"
    ).read_text()


def real_merged():
    """data/curated merged onto the committed rows of the build it is now curated
    against: 1.60 (Forever beta), the first client to export real Skyborne rows.

    The committed rows already carry the previous merge's `forever_changes`;
    merge_curated replaces them, so merging them again is what the pipeline does.
    era's own committed build (builds/1.15.9.69722) is a frozen historical artifact:
    its client has no Skyborne race at all, and combos.json's two real Skyborne race
    ids (95, 96) would make merge_curated raise if merged onto it. Beta is a strict
    superset of era's races and classes, so this is the only build the shared curated/
    directory can validate against now that Skyborne is real data, not a placeholder.
    """
    classes_json = json.loads((BETA_BUILD / "classes.json").read_text())
    races_json = json.loads((BETA_BUILD / "races.json").read_text())
    classes = [PlayableClass(**row) for row in classes_json]
    races = [PlayableRace(**row) for row in races_json]
    return merge_curated(classes, races, Path("curated"))


def test_the_real_curated_directory_is_valid():
    """data/curated must merge cleanly onto the committed beta rows."""
    _classes, merged_races, combos = real_merged()
    assert {"high-order-skyborne", "windshaper-skyborne"} <= {r.slug for r in merged_races}
    assert len(combos) == CLASSIC_COMBO_COUNT + len(NEW_COMBOS)


def test_the_documented_new_combos_are_the_ones_marked_new():
    """Exactly the pairs the research documents carry new_in_forever."""
    _classes, _races, combos = real_merged()
    assert {(c.race_id, c.class_id) for c in combos if c.new_in_forever} == NEW_COMBOS


def test_the_skyborne_rows_have_their_combos():
    """High Order Skyborne (95, Alliance) and Windshaper Skyborne (96, Horde) each get
    the four default classes plus their own faction-exclusive one."""
    _classes, _races, combos = real_merged()
    by_race: dict[int, set[int]] = {}
    for combo in combos:
        if combo.race_id in (95, 96):
            by_race.setdefault(combo.race_id, set()).add(combo.class_id)
    assert by_race == {
        95: SKYBORNE_DEFAULT_CLASSES | {8},
        96: SKYBORNE_DEFAULT_CLASSES | {7},
    }


def test_every_class_and_race_documents_a_forever_change():
    classes, races, _combos = real_merged()
    assert [record.slug for record in classes if not record.forever_changes] == []
    assert [record.slug for record in races if not record.forever_changes] == []


def test_every_forever_change_carries_a_usable_source():
    classes, races, _combos = real_merged()
    for record in [*classes, *races]:
        for change in record.forever_changes:
            where = f"{record.slug}: {change.text}"
            assert change.text.endswith("."), where
            assert change.sources, where
            for source in change.sources:
                assert source.kind in SOURCE_KINDS, where
                assert source.url.startswith("https://"), where
                assert source.label.strip(), where


def test_every_new_combo_names_a_source():
    """merge_curated rejects an unsourced new combo, but the sources are dropped
    from the emitted row, so the curated file is where they can be asserted."""
    for row in json.loads(Path("curated/combos.json").read_text()):
        if row["new_in_forever"]:
            assert row["sources"], row
            assert {source["kind"] for source in row["sources"]} <= SOURCE_KINDS


def test_the_committed_build_matches_the_curated_facts():
    """A curated edit that was never normalized into builds/<build>/ fails here."""
    classes, races, combos = real_merged()
    emitted = ((classes, "classes.json"), (races, "races.json"), (combos, "combos.json"))
    for records, name in emitted:
        committed = json.loads((BETA_BUILD / name).read_text())
        assert [record.model_dump() for record in records] == committed, name


def test_a_duplicate_class_slug_fails(tmp_path: Path):
    sourced = [{"label": "l", "url": "https://e.test", "kind": "blizzard"}]
    (tmp_path / "classes.json").write_text(
        json.dumps(
            [
                {"slug": "warrior", "forever_changes": [{"text": "First.", "sources": sourced}]},
                {"slug": "warrior", "forever_changes": [{"text": "Second.", "sources": sourced}]},
            ]
        )
    )
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text("[]")
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="class slug 'warrior' twice"):
        merge_curated(classes, races, tmp_path)


def test_a_duplicate_race_slug_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text("[]")
    (tmp_path / "races.json").write_text(
        json.dumps(
            [{"slug": "human", "forever_changes": []}, {"slug": "human", "forever_changes": []}]
        )
    )
    (tmp_path / "combos.json").write_text("[]")
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="race slug 'human' twice"):
        merge_curated(classes, races, tmp_path)


def test_a_duplicate_combo_pair_fails(tmp_path: Path):
    (tmp_path / "classes.json").write_text("[]")
    (tmp_path / "races.json").write_text("[]")
    (tmp_path / "combos.json").write_text(
        json.dumps(
            [
                {"race_id": 1, "class_id": 1, "new_in_forever": False},
                {"race_id": 1, "class_id": 1, "new_in_forever": False},
            ]
        )
    )
    classes, races = client_rows()
    with pytest.raises(CuratedError, match="race_id 1 class_id 1 twice"):
        merge_curated(classes, races, tmp_path)
