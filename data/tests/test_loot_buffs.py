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
    merge_simbuffs,
)
from pipeline.models import BuffOverride, SimBuffEntry

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
                    "sources": [{"label": "l", "url": "https://x.invalid", "kind": "site"}],
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
    # A weapon imbue's override is keyed by its second half ("windfury"), which
    # covers the main and off hand, imbue and stone ids together.
    real |= {buff_id.split(":", 1)[1] for buff_id in real if ":" in buff_id}
    assert set(document["entries"]) <= real
    assert len(document["entries"]) >= 13


# Round 10 (platform): IDS.md's weapon imbues are qualified ids, and the row
# pattern stopped at the colon, so not one of the 96 had a name or an icon and
# the /sim panel's weapon section drew humanised ids with no art.
IMBUE_IDS_MD = """## Buffs

| id | lands in |
| --- | --- |
| `blessing_of_might` | IndividualBuffs |
| `blessing_of_might:improved` | IndividualBuffs |

## Consumables

| id | sets |
| --- | --- |
| `main_hand_imbue:wizard_oil` | Consumes.main_hand_imbue |
| `off_hand_stone:wizard_oil` | Consumes.off_hand_stone |
| `main_hand_imbue:windfury_weapon` | Consumes.main_hand_imbue |
| `main_hand_imbue:windfury` | Consumes.main_hand_imbue |
"""
SPELLS = {
    "wizard_oil": ("Wizard Oil", "temp"),
    "windfury_weapon": ("Windfury Weapon", "spell_nature_cyclone"),
    "blessing_of_might": ("Blessing of Might", "spell_holy_fistofjustice"),
}
ITEMS = {"wizard_oil": ("Wizard Oil", "inv_potion_104")}


def test_ids_md_yields_weapon_imbue_ids_but_not_improved_grades():
    assert ids_md_ids(IMBUE_IDS_MD) == [
        "blessing_of_might",
        "main_hand_imbue:windfury",
        "main_hand_imbue:windfury_weapon",
        "main_hand_imbue:wizard_oil",
        "off_hand_stone:wizard_oil",
    ]


def test_a_weapon_imbue_takes_its_items_name_and_icon_before_the_enchant_spells():
    windfury = BuffOverride(name="Windfury Totem", spell_id=8512)
    entries = build_simbuffs(
        ids_md_ids(IMBUE_IDS_MD),
        [SPELLS, ITEMS],
        {"windfury": windfury},
        lambda item_id, spell_id: f"icon-{item_id}-{spell_id}",
        item_tables=[ITEMS],
    ).entries
    assert entries["main_hand_imbue:wizard_oil"] == SimBuffEntry(
        name="Wizard Oil", icon="inv_potion_104"
    )
    assert entries["off_hand_stone:wizard_oil"] == entries["main_hand_imbue:wizard_oil"]
    assert entries["main_hand_imbue:windfury_weapon"].icon == "spell_nature_cyclone"
    assert entries["main_hand_imbue:windfury"] == SimBuffEntry(
        name="Windfury Totem", icon="icon-0-8512"
    )
    # An unqualified id still joins spells first, as before.
    assert entries["blessing_of_might"].icon == "spell_holy_fistofjustice"


def test_the_merge_night_names_a_new_weapon_imbue_from_the_item_too():
    out = merge_simbuffs(
        ["main_hand_imbue:wizard_oil"], [SPELLS, ITEMS], {}, {}, item_tables=[ITEMS]
    )
    assert out.entries["main_hand_imbue:wizard_oil"].icon == "inv_potion_104"


# The same audit: 26 committed entries carried a placeholder from a fork spell row, a food's
# recipe spell (Seal of Kings art on Tender Wolf Steak) or the client's 'temp' (every
# protection potion), where the item itself has real art.
def test_a_food_takes_its_items_icon_before_its_recipe_spells():
    spells = {"tender_wolf_steak": ("Tender Wolf Steak", "spell_shadow_sealofkings")}
    items = {"tender_wolf_steak": ("Tender Wolf Steak", "inv_misc_food_47")}
    entries = build_simbuffs(
        ["food_tender_wolf_steak"], [spells, items], {}, lambda *_: "", item_tables=[items]
    ).entries
    assert entries["food_tender_wolf_steak"].icon == "inv_misc_food_47"


def test_a_temp_placeholder_loses_to_real_art_and_stands_only_alone():
    spells = {"mana_potion": ("Mana Potion", "temp"), "lonely": ("Lonely", "temp")}
    items = {"mana_potion": ("Mana Potion", "inv_potion_76")}
    entries = build_simbuffs(["lonely", "mana_potion"], [spells, items], {}, lambda *_: "").entries
    assert entries["mana_potion"].icon == "inv_potion_76"
    assert entries["lonely"].icon == "temp"
