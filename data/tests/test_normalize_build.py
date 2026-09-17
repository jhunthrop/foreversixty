import json
import shutil
from pathlib import Path

from pipeline.__main__ import main
from pipeline.models import ClassTalents
from pipeline.normalize import normalize_build
from pipeline.normalize.talents import flat_talents
from pipeline.wago import OPTIONAL_TABLES, TABLES

HERE = Path(__file__).parent
# classes.json and races.json are compared against the merged goldens below,
# because the orchestrator folds data/curated into them.
ENTITY_NAMES = ["zones", "dungeons", "items", "spells", "talents"]
#: Talent ids only the trait-table fixture (tests/fixtures/traits) can produce
#: -- the legacy Talent.csv fixture's own ids are all in the low thousands
#: (see fixtures/Talent.csv), so seeing one of these in the output proves
#: normalize_build took the trait branch of its reader switch rather than
#: quietly falling back to the legacy one.
TRAIT_ONLY_TALENT_ID = 900001


def prepare(tmp_path: Path) -> Path:
    raw = tmp_path / "1.0.0.1" / "raw"
    raw.mkdir(parents=True)
    for table in TABLES:
        fixture = HERE / "fixtures" / f"{table}.csv"
        if fixture.exists():
            shutil.copy(fixture, raw / f"{table}.csv")
        elif table in OPTIONAL_TABLES:
            # download_table writes this same stub for a 404'd optional table, so the
            # normalizer sees what a real Era build gives it. A required table with no
            # fixture is left absent on purpose: a normalizer that starts reading one
            # must bring its fixture, and this test fails loudly until it does.
            (raw / f"{table}.csv").write_text("ID\n", encoding="utf-8")
    meta = {"product": "test", "build": "1.0.0.1", "fetched_at": "t"}
    (raw / "_meta.json").write_text(json.dumps(meta))
    return tmp_path


def run(tmp_path: Path) -> Path:
    result = normalize_build(
        "1.0.0.1", root=prepare(tmp_path), curated_dir=HERE / "fixtures/curated"
    )
    assert result.skipped == ()
    return result.build_dir


def test_phase_0_entities_still_match_golden(tmp_path: Path):
    out = run(tmp_path)
    for name in ENTITY_NAMES:
        assert (out / f"{name}.json").read_text() == (HERE / "golden" / f"{name}.json").read_text()


def test_classes_and_races_are_merged_with_the_curated_facts(tmp_path: Path):
    out = run(tmp_path)
    assert (out / "classes.json").read_text() == (HERE / "golden/classes_merged.json").read_text()
    assert (out / "races.json").read_text() == (HERE / "golden/races_merged.json").read_text()


def test_phase_1_files_are_written(tmp_path: Path):
    out = run(tmp_path)
    assert (out / "talents" / "warrior.json").read_text() == (
        HERE / "golden/talents_warrior.json"
    ).read_text()
    assert (out / "items" / "warrior.json").read_text() == (
        HERE / "golden/items_warrior.json"
    ).read_text()
    assert (out / "sets.json").read_text() == (HERE / "golden/sets.json").read_text()
    assert (out / "combos.json").read_text() == (HERE / "golden/combos.json").read_text()


def test_manifest_lists_every_emitted_file(tmp_path: Path):
    out = run(tmp_path)
    files = set(json.loads((out / "manifest.json").read_text())["files"])
    assert {f"{name}.json" for name in ENTITY_NAMES} <= files
    assert {"classes.json", "races.json"} <= files
    assert {"sets.json", "combos.json", "talents/warrior.json", "items/warrior.json"} <= files


def test_a_rerun_is_byte_identical_and_clears_stale_class_files(tmp_path: Path):
    out = run(tmp_path)
    before = {p.relative_to(out).as_posix(): p.read_bytes() for p in sorted(out.rglob("*.json"))}
    (out / "talents" / "necromancer.json").write_text("{}\n")
    normalize_build("1.0.0.1", root=tmp_path, curated_dir=HERE / "fixtures/curated")
    after = {p.relative_to(out).as_posix(): p.read_bytes() for p in sorted(out.rglob("*.json"))}
    assert after == before


def break_the_item_table(root: Path) -> None:
    """Give one ItemSparse row a stat modifier id the normalizer will not guess at."""
    sparse = root / "1.0.0.1" / "raw" / "ItemSparse.csv"
    sparse.write_text(sparse.read_text().replace(",7,4,35,15", ",99,4,35,15"))


def test_items_are_skipped_when_the_item_table_is_unreadable(tmp_path: Path, caplog):
    root = prepare(tmp_path)
    break_the_item_table(root)
    result = normalize_build("1.0.0.1", root=root, curated_dir=HERE / "fixtures/curated")
    assert not (result.build_dir / "items").exists()
    assert (result.build_dir / "sets.json").exists()
    assert "unknown stat modifier id 99" in caplog.text


def test_a_skipped_output_is_reported_in_the_result(tmp_path: Path):
    """normalize_build stays soft, but it says what it could not emit."""
    root = prepare(tmp_path)
    break_the_item_table(root)
    result = normalize_build("1.0.0.1", root=root, curated_dir=HERE / "fixtures/curated")
    assert len(result.skipped) == 1
    assert result.skipped[0].startswith("items/: ")


def cli_workspace(tmp_path: Path, monkeypatch) -> Path:
    """A working directory the CLI's default `builds/` and `curated/` paths resolve in."""
    root = prepare(tmp_path / "builds")
    shutil.copytree(HERE / "fixtures/curated", tmp_path / "curated")
    monkeypatch.chdir(tmp_path)
    return root / "1.0.0.1"


def test_the_cli_exits_non_zero_when_the_items_files_were_skipped(tmp_path: Path, monkeypatch):
    """A build directory with no items/ must never reach the commit step of CI."""
    build_dir = cli_workspace(tmp_path, monkeypatch)
    break_the_item_table(build_dir.parent)
    assert main(["normalize", "--build", "1.0.0.1"]) == 1
    assert not (build_dir / "items").exists()
    # Everything else is still emitted: the run is soft, only its exit code is not.
    assert (build_dir / "sets.json").exists()
    assert (build_dir / "talents" / "warrior.json").exists()


def test_the_cli_exits_zero_when_every_output_was_emitted(tmp_path: Path, monkeypatch):
    build_dir = cli_workspace(tmp_path, monkeypatch)
    assert main(["normalize", "--build", "1.0.0.1"]) == 0
    assert (build_dir / "items" / "warrior.json").exists()


def trait_switch_workspace(tmp_path: Path) -> Path:
    """The usual fixture raw dir, with its talent/trait tables swapped for
    tests/fixtures/traits' (see traits_fixture.py and
    test_normalize_trait_trees.py), so SkillLineXTraitTree is populated and
    has_trait_trees(...) is true. ChrClasses stays the default fixture's
    (Warrior/Paladin/Mage): the trait fixture's own tables only ever name
    class 1 (Warrior, ClassMask 1), which the default ChrClasses already
    has, and swapping ChrClasses too would strand the curated fixture's
    combos, which name classes the trait fixture doesn't carry.
    """
    root = prepare(tmp_path)
    raw = root / "1.0.0.1" / "raw"
    for fixture in (HERE / "fixtures/traits").glob("*.csv"):
        if fixture.name != "ChrClasses.csv":
            shutil.copy(fixture, raw / fixture.name)
    return root


def test_normalize_build_picks_the_trait_reader_when_the_client_has_trait_trees(tmp_path: Path):
    """The entire deliverable of Task 6: normalize_build must read a build's
    talent trees from the trait tables, not the legacy Talent table, whenever
    the client's SkillLineXTraitTree names one. Before this test, nothing
    exercised this branch of the switch -- every other normalize_build test
    fixture leaves SkillLineXTraitTree empty and takes the legacy branch."""
    out = normalize_build(
        "1.0.0.1", root=trait_switch_workspace(tmp_path), curated_dir=HERE / "fixtures/curated"
    ).build_dir
    warrior = ClassTalents.model_validate_json(
        (out / "talents" / "warrior.json").read_text(encoding="utf-8")
    )
    talent_ids = {t.id for tree in warrior.trees for t in tree.talents}
    assert TRAIT_ONLY_TALENT_ID in talent_ids

    # talents.json is written by write_json, which sorts by id; flat_talents
    # itself preserves tree/talent order, so the expectation is sorted the
    # same way before comparing.
    flat = json.loads((out / "talents.json").read_text(encoding="utf-8"))
    expected = sorted(
        (node.model_dump() for node in flat_talents([warrior])), key=lambda n: n["id"]
    )
    assert flat == expected
