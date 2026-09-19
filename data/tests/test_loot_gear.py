import json
import shutil
from pathlib import Path

from pipeline.forkdb import load_fork_database
from pipeline.loot.gear import (
    apply_fork_columns,
    build_suffixes,
    faction_restrictions,
    suffix_options,
)

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"


def fork():
    return load_fork_database(ENGINE)


def test_every_suffix_is_emitted_with_its_stats_named():
    records = build_suffixes(fork())
    assert [(r.id, r.name) for r in records] == [(5, "of Intellect"), (6, "of Strength")]
    assert records[0].stats == {"intellect": 4.0}
    assert records[1].stats == {"strength": 7.0}


def test_suffix_options_index_only_the_items_that_roll_one():
    assert suffix_options(fork()) == {110: [5, 6]}


def test_faction_restrictions_index_only_the_restricted_items():
    assert faction_restrictions(fork()) == {100: "alliance_only"}


def prepared(tmp_path):
    build_dir = tmp_path / "1.60.1.69893"
    build_dir.mkdir()
    shutil.copy(ENGINE / "items.json", build_dir / "items.json")
    return build_dir


def test_apply_fork_columns_fills_both_and_leaves_the_rest_empty(tmp_path):
    build_dir = prepared(tmp_path)
    with_suffixes, restricted = apply_fork_columns(
        build_dir, suffix_options(fork()), faction_restrictions(fork())
    )
    assert (with_suffixes, restricted) == (1, 1)
    rows = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    by_id = {row["id"]: row for row in rows}
    assert by_id[110]["suffixes"] == [5, 6]
    assert by_id[100]["suffixes"] == []
    assert by_id[100]["faction_restriction"] == "alliance_only"
    assert by_id[110]["faction_restriction"] == ""
    assert len(rows) == 11


def test_apply_fork_columns_keeps_the_key_order_and_appends_the_new_ones(tmp_path):
    build_dir = prepared(tmp_path)
    apply_fork_columns(build_dir, suffix_options(fork()), faction_restrictions(fork()))
    rows = json.loads((build_dir / "items.json").read_text(encoding="utf-8"))
    assert list(rows[0]) == [
        "id",
        "name",
        "quality",
        "item_level",
        "required_level",
        "class_id",
        "subclass_id",
        "inventory_type",
        "suffixes",
        "faction_restriction",
    ]


def test_apply_fork_columns_is_idempotent(tmp_path):
    build_dir = prepared(tmp_path)
    options, restrictions = suffix_options(fork()), faction_restrictions(fork())
    apply_fork_columns(build_dir, options, restrictions)
    once = (build_dir / "items.json").read_bytes()
    apply_fork_columns(build_dir, options, restrictions)
    assert (build_dir / "items.json").read_bytes() == once
