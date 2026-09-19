from pathlib import Path

import pytest

from pipeline.forkdb import (
    CLASS_SLUGS,
    ENCHANT_TYPES,
    ITEM_TYPE_SLOTS,
    PROFESSIONS,
    REP_LEVELS,
    ForkDbError,
    decode,
    load_fork_database,
)

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"


def test_a_missing_checkout_says_where_it_looked(tmp_path):
    with pytest.raises(ForkDbError, match="assets/database/db.json"):
        load_fork_database(tmp_path)


def test_the_fixture_database_loads_every_table():
    fork = load_fork_database(ENGINE)
    assert len(fork.items) == 13
    assert len(fork.enchants) == 4
    assert len(fork.random_suffixes) == 2
    assert fork.zones == {2717: "Molten Core", 1581: "The Deadmines"}
    assert fork.npcs == {
        900: "Big Boss",
        902: "Dungeon Boss",
        903: "Azuregos",
        906: "Re-itemised Boss",
    }
    assert fork.factions == {529: "Argent Dawn"}
    assert fork.item_icons[2304] == "inv_misc_armorkit_17"
    assert fork.spell_icons[7420] == "spell_holy_chest"
    assert len(fork.spell_icon_rows) == 5
    assert len(fork.item_icon_rows) == 1


def test_the_tables_are_immutable():
    fork = load_fork_database(ENGINE)
    # `items` is a `tuple[dict, ...]` (see ForkDatabase); a plain tuple has
    # no `.append` at all, so the interpreter raises AttributeError here,
    # not TypeError.
    with pytest.raises(AttributeError):
        fork.items.append({})  # type: ignore[attr-defined]


def test_decode_names_the_value_it_could_not_read():
    assert decode(PROFESSIONS, 2, "profession") == "blacksmithing"
    with pytest.raises(ForkDbError, match="profession 7"):
        decode(PROFESSIONS, 7, "profession")


def test_the_enum_tables_match_the_forks_protos():
    """Transcribed from wowsims-forever proto/common.proto and
    proto/ui.proto; a mismatch here is a fork change to pick up."""
    assert PROFESSIONS[11] == "tailoring"
    assert REP_LEVELS[8] == "exalted"
    assert ITEM_TYPE_SLOTS[13] == ("main_hand", "off_hand")
    assert ITEM_TYPE_SLOTS[4] == ("back",)
    assert ENCHANT_TYPES[3] == "kit"
    assert CLASS_SLUGS[9] == "warrior"
    assert sorted(CLASS_SLUGS) == list(range(1, 10))
