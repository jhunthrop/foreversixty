import json
from pathlib import Path

import pytest
from pydantic import ValidationError

from pipeline.forkdb import load_fork_database
from pipeline.loot.buffs import (
    BuffError,
    build_simbuffs,
    client_tables,
    fork_tables,
    ids_md_ids,
    load_overrides,
)

HERE = Path(__file__).parent
ENGINE = HERE / "fixtures/loot"
CURATED = Path("curated")


def ids():
    return ids_md_ids((ENGINE / "IDS.md").read_text(encoding="utf-8"))


def tables_and_icons():
    client, icon_for = client_tables(ENGINE, ENGINE)
    return fork_tables(load_fork_database(ENGINE)) + client, icon_for


def built():
    tables, icon_for = tables_and_icons()
    return build_simbuffs(
        ids(), tables, load_overrides(ENGINE, name="curated-simbuffs.json"), icon_for
    ).entries


def test_ids_md_yields_the_buff_and_consumable_ids_once_each_not_professions():
    assert ids() == [
        "arcane_brilliance",
        "curse_of_elements",
        "elixir_of_fire_power",
        "food_grilled_squid",
        "hunters_mark",
        "innervates",
    ]
    assert "alchemy" not in ids()


def test_every_id_resolves_to_a_name_and_an_icon():
    entries = built()
    assert set(entries) == set(ids())
    assert all(entry.name and entry.icon for entry in entries.values())


def test_an_id_matching_a_client_spell_takes_the_clients_icon():
    entries = built()
    assert entries["arcane_brilliance"].name == "Arcane Brilliance"
    assert entries["arcane_brilliance"].icon == "spell_holy_arcaneintellect"


def test_an_apostrophe_and_a_plural_both_resolve():
    entries = built()
    assert entries["hunters_mark"].name == "Hunter's Mark"
    assert entries["innervates"].name == "Innervate"


def test_an_enum_name_prefix_is_stripped_the_way_ids_md_says():
    entries = built()
    assert entries["food_grilled_squid"].name == "Grilled Squid"


def test_a_curated_override_wins_over_every_join():
    entries = built()
    assert entries["curse_of_elements"].name == "Curse of the Elements"
    # spell 1459's icon in the fixture: the override named a client row,
    # so the art came from the build and not from the curated file.
    assert entries["curse_of_elements"].icon == "spell_holy_arcaneintellect"


def test_an_id_nothing_resolves_stops_the_build_and_names_it():
    tables, icon_for = tables_and_icons()
    with pytest.raises(BuffError, match="mystery_buff"):
        build_simbuffs(["mystery_buff"], tables, {}, icon_for)


def test_an_override_naming_both_or_neither_id_is_refused(tmp_path):
    for entry in ({"name": "X"}, {"name": "X", "item_id": 1, "spell_id": 2}):
        (tmp_path / "simbuffs.json").write_text(
            json.dumps(
                {
                    "sources": [
                        {"label": "l", "url": "https://x.invalid", "kind": "site"}
                    ],
                    "notes": "x",
                    "entries": {"curse_of_elements": entry},
                }
            ),
            encoding="utf-8",
        )
        with pytest.raises(BuffError, match="exactly one"):
            load_overrides(tmp_path)


def test_an_override_with_an_unknown_key_is_refused(tmp_path):
    (tmp_path / "simbuffs.json").write_text(
        json.dumps(
            {
                "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
                "notes": "x",
                "entries": {
                    "curse_of_elements": {
                        "name": "X",
                        "item_id": 1,
                        "icon": "sneaky_icon",
                    }
                },
            }
        ),
        encoding="utf-8",
    )
    with pytest.raises(ValidationError, match="icon"):
        load_overrides(tmp_path)


def test_the_real_override_file_is_sourced_and_covers_only_real_ids():
    document = json.loads((CURATED / "simbuffs.json").read_text(encoding="utf-8"))
    assert document["sources"] and document["notes"].strip()
    real = set(ids_md_ids(Path("../sim/request/IDS.md").read_text(encoding="utf-8")))
    assert set(document["entries"]) <= real
    assert len(document["entries"]) == 14
