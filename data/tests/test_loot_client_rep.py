"""Reputation sources from the client's own item requirement (sim-accuracy
round 9, items lane).

Darkspear Islands' rewards are reputation rewards on both sides. The client
states it on each item (ItemSparse.MinFactionID, MinReputation), and the
items/<class>.json rows carry it as `reputation`. Creeg Bothunk sells the
Darkspear Raiders set, but no vendor names the Theramore Expeditionary
Force set, so it had no source at all. Every such item now gets a `rep`
source, and its side goes into loot.json's `factions`, so each set stays
on its own side.
"""

import json

from pipeline.loot.client_rep import apply_client_reputations, load_item_reputations
from pipeline.models import ItemReputation, LootFile, LootSource

THERAMORE = ItemReputation(
    faction_id=2799,
    standing="exalted",
    faction_name="Theramore Expeditionary Force",
    side="alliance",
)
DARKSPEAR = ItemReputation(
    faction_id=2798, standing="exalted", faction_name="Darkspear Raiders", side="horde"
)


def document(*sources):
    return LootFile(sources=list(sources))


CREEG = LootSource(
    id="vendor:263570", kind="vendor", name="Creeg Bothunk", npc_id=263570, items=[272096]
)


def test_an_unsourced_reputation_reward_gets_its_rep_source_and_side():
    out, added = apply_client_reputations(document(CREEG), {271941: THERAMORE}, {271941, 272096})
    rep = {s.id: s for s in out.sources}["rep:theramore-expeditionary-force:exalted"]
    assert (rep.kind, rep.name, rep.faction_id, rep.standing, rep.items) == (
        "rep",
        "Theramore Expeditionary Force",
        2799,
        "exalted",
        [271941],
    )
    assert out.factions["271941"] == "alliance"
    assert added == 1


def test_a_vendor_item_gets_the_rep_source_beside_its_vendor_and_its_side():
    out, _ = apply_client_reputations(document(CREEG), {272096: DARKSPEAR}, {272096})
    ids = {s.id for s in out.sources}
    assert {"vendor:263570", "rep:darkspear-raiders:exalted"} <= ids
    assert out.factions["272096"] == "horde"


def test_an_item_a_rep_source_already_names_is_left_alone():
    existing = LootSource(
        id="rep:darkspear-raiders:honored",
        kind="rep",
        name="Darkspear Raiders",
        faction_id=2798,
        standing="honored",
        items=[272096],
    )
    out, added = apply_client_reputations(document(existing), {272096: DARKSPEAR}, {272096})
    assert [s.id for s in out.sources] == ["rep:darkspear-raiders:honored"]
    assert added == 0


def test_a_rep_source_of_the_same_id_takes_the_item():
    existing = LootSource(
        id="rep:theramore-expeditionary-force:exalted",
        kind="rep",
        name="Theramore Expeditionary Force",
        faction_id=2799,
        standing="exalted",
        items=[271940],
    )
    out, _ = apply_client_reputations(document(existing), {271941: THERAMORE}, {271940, 271941})
    assert [s.items for s in out.sources] == [[271940, 271941]]


def test_a_neutral_unnamed_or_absent_item_gets_nothing():
    neutral = THERAMORE.model_copy(update={"standing": "neutral"})
    unnamed = THERAMORE.model_copy(update={"faction_name": None})
    out, added = apply_client_reputations(
        document(), {1: neutral, 2: unnamed, 3: THERAMORE}, {1, 2}
    )
    assert out.sources == [] and out.factions == {} and added == 0


def test_a_both_sides_reputation_names_no_side():
    timbermaw = ItemReputation(faction_id=576, standing="honored", faction_name="Timbermaw Hold")
    out, _ = apply_client_reputations(document(), {5: timbermaw}, {5})
    assert [s.id for s in out.sources] == ["rep:timbermaw-hold:honored"]
    assert out.factions == {}


def test_load_item_reputations_reads_every_class_file(tmp_path):
    items = tmp_path / "items"
    items.mkdir()
    row = {"id": 271941, "reputation": THERAMORE.model_dump()}
    (items / "warrior.json").write_text(json.dumps({"items": [row, {"id": 7, "reputation": None}]}))
    (items / "paladin.json").write_text(json.dumps({"items": [row]}))
    assert load_item_reputations(tmp_path) == {271941: THERAMORE}
