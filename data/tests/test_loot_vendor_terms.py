"""pipeline.loot.vendor_terms: a vendor's place and non-gold prices, and a
reputation quartermaster's place (feedback 2026-11-09)."""

import json
from pathlib import Path

import pytest

from pipeline.loot.vendor_terms import (
    VendorTermsError,
    apply_vendor_terms,
    load_vendor_terms,
)
from pipeline.models import ItemReputation, LootFile, LootSource, VendorCost

CURATED = Path("curated")
VALIANTS_WAISTGUARD, RING_OF_AWARENESS, SENTINELS_LIBRAM = 272403, 272409, 272434
THERAMORE_GLOVES = 271948
INSURGENTS_BAND, RAIDER_GLOVES = 272064, 272099
PROVENANCE = [{"label": "l", "url": "https://x.invalid", "kind": "site"}]


def _document() -> LootFile:
    return LootFile(
        sources=[
            LootSource(
                id="vendor:227853", kind="vendor", name="Pix Xizzix", npc_id=227853,
                items=[VALIANTS_WAISTGUARD, RING_OF_AWARENESS, SENTINELS_LIBRAM],
            ),
            LootSource(id="vendor:1", kind="vendor", name="Other", npc_id=1, items=[5]),
            LootSource(
                id="vendor:263570", kind="vendor", name="Creeg Bothunk", npc_id=263570,
                items=[INSURGENTS_BAND, RAIDER_GLOVES],
            ),
            LootSource(
                id="rep:theramore-expeditionary-force:exalted", kind="rep",
                name="Theramore Expeditionary Force", faction_id=2799, standing="exalted",
                items=[THERAMORE_GLOVES],
            ),
        ]
    )  # fmt: skip


SLOTS = {
    VALIANTS_WAISTGUARD: "waist",
    RING_OF_AWARENESS: "finger",
    SENTINELS_LIBRAM: "ranged",
    THERAMORE_GLOVES: "hands",
    5: "waist",
}


def _darkspear(standing: str) -> ItemReputation:
    return ItemReputation(
        faction_id=2798, standing=standing, faction_name="Darkspear Raiders", side="horde"
    )


REPUTATIONS = {INSURGENTS_BAND: _darkspear("honored"), RAIDER_GLOVES: _darkspear("exalted")}


def test_the_curated_terms_price_pix_xizzix_and_place_theramore():
    out = apply_vendor_terms(_document(), load_vendor_terms(CURATED), SLOTS, REPUTATIONS)
    by_id = {source.id: source for source in out.sources}
    pix = by_id["vendor:227853"]
    assert pix.place == "Booty Bay"
    reals = "Tarnished Undermine Reals"
    assert pix.item_costs == {
        str(VALIANTS_WAISTGUARD): VendorCost(amount=100, currency=reals),
        str(RING_OF_AWARENESS): VendorCost(amount=75, currency=reals),
    }  # the libram's price was never recorded, so it stays unpriced
    assert by_id["rep:theramore-expeditionary-force:exalted"].place == "The Barrens"
    other = by_id["vendor:1"]
    assert other.place is None and other.item_costs is None and other.reputation is None


def test_a_quartermaster_reads_as_its_reputation_like_the_other_sides_rep_source():
    out = apply_vendor_terms(_document(), load_vendor_terms(CURATED), SLOTS, REPUTATIONS)
    creeg = {source.id: source for source in out.sources}["vendor:263570"]
    assert creeg.place == "The Barrens"
    assert creeg.reputation == "Darkspear Raiders"
    assert creeg.faction_id == 2798
    assert creeg.item_standings == {str(INSURGENTS_BAND): "honored", str(RAIDER_GLOVES): "exalted"}


def test_a_vendor_with_any_unreputed_item_is_no_quartermaster():
    reputations = {INSURGENTS_BAND: _darkspear("honored")}
    out = apply_vendor_terms(_document(), load_vendor_terms(CURATED), SLOTS, reputations)
    creeg = {source.id: source for source in out.sources}["vendor:263570"]
    assert creeg.reputation is None and creeg.item_standings is None


def _write(tmp_path: Path, vendors: list[dict]) -> Path:
    document = {"sources": PROVENANCE, "notes": "why", "vendors": vendors}
    (tmp_path / "vendor-terms.json").write_text(json.dumps(document), encoding="utf-8")
    return tmp_path


@pytest.mark.parametrize(
    ("vendor", "message"),
    [
        ({"cost_by_slot": {"waist": 100}}, "names no currency"),
        ({"currency": "Reals", "cost_by_slot": {"waist": 0}}, "not positive"),
    ],
)
def test_a_price_without_a_currency_or_a_real_amount_fails_fast(tmp_path, vendor, message):
    row = {"npc_id": 1, "name": "V", "place": "P", "why": "w", **vendor}
    with pytest.raises(VendorTermsError, match=message):
        load_vendor_terms(_write(tmp_path, [row]))
