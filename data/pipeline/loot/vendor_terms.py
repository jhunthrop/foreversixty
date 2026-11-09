"""Where a vendor or a reputation quartermaster stands, and what a vendor
charges when it is not gold (feedback 2026-11-09: a BiS line read "Vendor:
Pix Xizzix" with no hint that the belt costs 100 Tarnished Undermine Reals).

The facts are curated in curated/vendor-terms.json, with provenance: no
client table names a vendor's subzone or its ItemExtendedCost, and wowhead's
Forever sold-by rows carry no price for these. `apply_vendor_terms` writes
them onto loot.json's own sources:

- `place` on a `vendor:<npc_id>` source and on every `rep` source of a
  curated faction;
- `item_costs` on a vendor source, priced by each item's gear slot. An item
  whose slot the file prices nothing for is left unpriced.
"""

from __future__ import annotations

import json
from collections.abc import Mapping
from pathlib import Path

from pydantic import BaseModel

from pipeline.curated import parse_sources
from pipeline.models import LootFile, LootSource, Source, VendorCost

VENDOR_TERMS = "vendor-terms.json"


class VendorTermsError(SystemExit):
    """curated/vendor-terms.json says something the pipeline will not apply."""


class VendorTerm(BaseModel):
    npc_id: int
    name: str
    place: str
    currency: str | None = None
    cost_by_slot: dict[str, int] = {}
    why: str

    model_config = {"extra": "forbid"}


class ReputationTerm(BaseModel):
    faction_id: int
    name: str
    place: str
    why: str

    model_config = {"extra": "forbid"}


class VendorTermsFile(BaseModel):
    sources: list[Source]
    notes: str
    vendors: list[VendorTerm] = []
    reputations: list[ReputationTerm] = []

    model_config = {"extra": "forbid"}


def _check(document: VendorTermsFile, where: str) -> None:
    for vendor in document.vendors:
        if vendor.cost_by_slot and not vendor.currency:
            raise VendorTermsError(f"{where} prices {vendor.name} but names no currency")
        if any(amount <= 0 for amount in vendor.cost_by_slot.values()):
            raise VendorTermsError(f"{where} gives {vendor.name} a price that is not positive")
    for terms, key in ((document.vendors, "npc_id"), (document.reputations, "faction_id")):
        ids = [getattr(term, key) for term in terms]
        if len(ids) != len(set(ids)):
            raise VendorTermsError(f"{where} lists one {key} twice")


def load_vendor_terms(curated_dir: Path) -> VendorTermsFile:
    path = curated_dir / VENDOR_TERMS
    if not path.exists():
        raise VendorTermsError(f"missing {path}")
    document = VendorTermsFile(**json.loads(path.read_text(encoding="utf-8")))
    parse_sources([source.model_dump() for source in document.sources], str(path))
    if not document.notes.strip():
        raise VendorTermsError(f"{path} has no notes; say what the terms are for")
    _check(document, str(path))
    return document


def _vendor_update(source: LootSource, term: VendorTerm, item_slots: Mapping[int, str]) -> dict:
    costs = {
        str(item_id): VendorCost(amount=term.cost_by_slot[slot], currency=term.currency or "")
        for item_id in source.items or []
        if (slot := item_slots.get(item_id)) in term.cost_by_slot
    }
    return {"place": term.place, **({"item_costs": costs} if costs else {})}


def apply_vendor_terms(
    document: LootFile, terms: VendorTermsFile, item_slots: Mapping[int, str]
) -> LootFile:
    """`document` with each curated vendor's and reputation's place, and each
    curated vendor's non-gold prices. `item_slots` is item id -> gear slot."""
    vendors = {term.npc_id: term for term in terms.vendors}
    reputations = {term.faction_id: term for term in terms.reputations}
    sources = []
    for source in document.sources:
        if source.kind == "vendor" and source.npc_id in vendors:
            source = source.model_copy(
                update=_vendor_update(source, vendors[source.npc_id], item_slots)
            )
        elif source.kind == "rep" and source.faction_id in reputations:
            source = source.model_copy(update={"place": reputations[source.faction_id].place})
        sources.append(source)
    return document.model_copy(update={"sources": sources})
