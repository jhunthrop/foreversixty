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
from pathlib import Path

from pipeline.loot.client_rep import apply_client_reputations, load_item_reputations
from pipeline.loot.content_phases import apply_content_phases, load_content_phases
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


def test_a_vendor_item_keeps_its_vendor_as_its_way_in_and_gets_its_side():
    """Round 10: a rep source beside the vendor that sells the item is a
    second way in that no gate on the vendor closes. The vendor is where the
    item comes from; the ranker holds it to the client standing
    (sim/cmd/leveling-bis reputation.go), so only the side is added."""
    out, added = apply_client_reputations(document(CREEG), {272096: DARKSPEAR}, {272096})
    assert [s.id for s in out.sources] == ["vendor:263570"]
    assert out.factions["272096"] == "horde"
    assert added == 0


SEA_GIANTS_TOE_RING = 274746
BOOTY_BAY = ItemReputation(faction_id=21, standing="friendly", faction_name="Booty Bay")
GEZZY = LootSource(
    id="vendor:265575",
    kind="vendor",
    name="Gezzy Gunkgear",
    npc_id=265575,
    items=[SEA_GIANTS_TOE_RING],
)


def test_a_gated_vendors_reputation_stock_stays_gated():
    """Round 10: Gezzy Gunkgear (forever-new-vendors, gated 'later') sells
    Sea Giant's Toe Ring at Booty Bay Friendly (ItemSparse 70707
    MinFactionID 21, MinReputation 4). The client reputation source must not
    open what the vendor's gate closes."""
    out, _ = apply_client_reputations(
        document(GEZZY), {SEA_GIANTS_TOE_RING: BOOTY_BAY}, {SEA_GIANTS_TOE_RING}
    )
    gated, _ = apply_content_phases(out, load_content_phases(Path("curated")))
    open_ways = [
        source.id
        for source in gated.sources
        if SEA_GIANTS_TOE_RING in (source.items or []) and not source.opens
    ]
    assert open_ways == []


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
