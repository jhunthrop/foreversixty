"""data/curated/presets.json: the Phase 1 raid preset the BiS ranker runs level 60 under.

The Go side (`sim/request`'s LoadPresets) refuses an id the request layer cannot map;
this is the other direction, so an edit that drops a buff the owner asked for, or adds
a world buff or an Ahn'Qiraj consumable, shows up here too.
"""

import json
import re
from pathlib import Path

import pytest

from pipeline.curated import parse_sources

CURATED = Path("curated")
IDS = Path("../sim/request/IDS.md")

REQUIRED_BUFFS = {
    "arcane_brilliance",
    "gift_of_the_wild",
    "power_word_fortitude",
    "divine_spirit",
    "blessing_of_might",
    "blessing_of_wisdom",
    "battle_shout",
    "trueshot_aura",
    "leader_of_the_pack",
    "sanctity_aura",
    "strength_of_earth_totem",
    "grace_of_air_totem",
    "mana_spring_totem",
}
REQUIRED_DEBUFFS = {
    "curse_of_elements",
    "sunder_armor",
    "faerie_fire",
    "judgement_of_wisdom",
    "judgement_of_light",
    "hunters_mark",
    "curse_of_recklessness",
}
#: A buff the preset must never carry: Forever has no world buffs, Shadow Weaving is a
#: self-buff now, and Blessing of Kings is not in the live paladin trees.
FORBIDDEN = {"shadow_weaving", "blessing_of_kings", "songflower_serenade", "spirit_of_zandalar"}
#: Consumables from Ahn'Qiraj or Naxxramas, or a world buff by another name.
FORBIDDEN_CONSUMES = {
    "spirit_of_zanza",
    "sheen_of_zanza",
    "swiftness_of_zanza",
    "flask_of_the_titans",
}


@pytest.fixture(scope="module")
def raid() -> dict:
    document = json.loads((CURATED / "presets.json").read_text())
    parse_sources(document["sources"], "presets.json")
    return document["presets"]["raid"]


def _vocabulary(section: str) -> set[str]:
    """The ids one table of sim/request/IDS.md lists."""
    body = IDS.read_text().split(f"\n## {section}\n", 1)[1].split("\n## ", 1)[0]
    return set(re.findall(r"^\| `([^`]+)` \|", body, flags=re.MULTILINE))


def _consume_entries(raid: dict) -> list[dict]:
    consumes = raid["consumes"]
    entries = [e for group in consumes["groups"].values() for e in group["ids"]]
    entries += [e for group in consumes["groups"].values() for e in group.get("dual_wield_ids", [])]
    entries += [
        e
        for lists in (consumes["by_class"], consumes["by_spec"])
        for v in lists.values()
        for e in v
    ]
    return entries


def test_the_raid_preset_has_a_label_and_notes(raid):
    assert raid["label"] == "Raid-ready, Phase 1"
    assert raid["notes"].strip()


def test_the_raid_preset_names_every_buff_and_debuff_the_owner_asked_for(raid):
    assert REQUIRED_BUFFS <= {e["id"] for e in raid["buffs"]}
    assert REQUIRED_DEBUFFS <= {e["id"] for e in raid["debuffs"]}


def test_the_raid_preset_carries_nothing_forbidden(raid):
    ids = {e["id"] for e in raid["buffs"] + raid["debuffs"]}
    assert not ids & FORBIDDEN
    assert not {e["id"] for e in _consume_entries(raid)} & FORBIDDEN_CONSUMES


def test_every_buff_and_debuff_is_in_the_request_vocabulary(raid):
    buffs = _vocabulary("Buffs")
    for entry in raid["buffs"] + raid["debuffs"]:
        assert entry["id"] in buffs, entry["id"]


def test_every_consumable_is_in_the_request_vocabulary(raid):
    consumables = _vocabulary("Consumables")
    for entry in _consume_entries(raid):
        assert entry["id"] in consumables, entry["id"]


def test_every_entry_has_a_label_and_a_reason(raid):
    entries = raid["buffs"] + raid["debuffs"] + _consume_entries(raid)
    for entry in entries:
        assert entry["label"].strip(), entry["id"]
        assert entry["reason"].strip(), entry["id"]


def test_each_list_names_an_id_once(raid):
    for name in ("buffs", "debuffs"):
        ids = [e["id"] for e in raid[name]]
        assert len(ids) == len(set(ids)), name


def test_every_consumable_group_covers_a_role(raid):
    groups = raid["consumes"]["groups"]
    assert {"caster", "melee", "ranged"} <= set(groups)
    assert "hunter" in groups["ranged"]["classes"]
    assert "spell_power" in groups["caster"]["reference_stats"]
    assert "attack_power" in groups["melee"]["reference_stats"]
