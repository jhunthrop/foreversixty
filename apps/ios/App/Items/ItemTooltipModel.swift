import Foundation
import Models

// MERGE NOTE (lanes F2 and F3): F2's ItemTooltipModel.swift replaces this file at integration.
// The request, the look and the sources below carry F2's names and fields exactly, so BiS's call
// site (`ItemTooltipSheet(request:catalog:)`) compiles against either. F2's model must word the
// next best in the band's unit (`tooltip.delta` and `items.sameDps` take `{unit}`) and drop the
// preset from the footer (`tooltip.pickOf` has no `{preset}`).

/// Who asked for a tooltip and what they already know: the item, the build its static data
/// lives in, and the band the pick belongs to. Built by Home, the gear list and BiS.
struct ItemTooltipRequest: Equatable, Identifiable, Sendable {
    let itemID: Int
    /// The name and look to draw while the static file loads, or when it lacks the item.
    let fallback: HomeItemLook
    let build: String
    let classSlug: String
    let spec: String?
    let specName: String?
    let band: Int?
    let bandLabel: String?
    let preset: String?
    let faction: String?
    /// The caller's source words for this item, when it has them.
    let sourceLine: String?
    /// The item worn in the same slot, so the next best can say "You wear this".
    let wornID: Int?

    var id: Int { itemID }
}

/// An item's name, quality, icon and level; every field but id may be null.
struct HomeItemLook: Equatable, Sendable {
    let id: Int
    let name: String?
    let quality: Int?
    let icon: String?
    let itemLevel: Int?
}

/// The build's static files a tooltip reads, already loaded.
struct ItemTooltipSources: Sendable {
    let items: [Int: Item]
    let sets: [ItemSet]
    let bis: BisFile?
}

/// One runner-up under "The next best".
struct ItemTooltipAlternative: Equatable, Identifiable {
    let look: HomeItemLook
    let meta: String
    var id: Int { look.id }
}

/// What F3's minimal sheet draws: the header, the panel's level lines, where it comes from, the
/// next best in the band's unit, and the pick footer.
struct ItemTooltipModel: Equatable {
    let look: HomeItemLook
    let item: Item?
    let slotKey: String?
    let source: String?
    let alternatives: [ItemTooltipAlternative]
    let pickFooter: String?

    /// Spec F2.6: up to two alternatives.
    static let maxAlternatives = 2

    static func build(_ request: ItemTooltipRequest, sources: ItemTooltipSources) -> ItemTooltipModel {
        let item = sources.items[request.itemID]
        let band = band(for: request, in: sources.bis)
        let slot = band?.slots.first { $0.itemId == request.itemID }
        let unit = BisUnit(role: band?.role)
        let look = HomeItemLook(
            id: request.itemID,
            name: item?.name ?? request.fallback.name,
            quality: item?.quality ?? request.fallback.quality,
            icon: item?.icon ?? request.fallback.icon,
            itemLevel: item?.itemLevel ?? request.fallback.itemLevel
        )
        return ItemTooltipModel(
            look: look,
            item: item,
            slotKey: slot?.slot ?? item?.slot,
            source: request.sourceLine,
            alternatives: alternatives(slot, request: request, items: sources.items, unit: unit),
            pickFooter: footer(slot: slot, request: request)
        )
    }

    /// The band the request names, by the site's selection rule.
    static func band(for request: ItemTooltipRequest, in bis: BisFile?) -> BisBand? {
        guard let bis, let band = request.band, let preset = request.preset, let faction = request.faction else { return nil }
        return bis.band(for: BisRange(band: band, preset: preset), faction: faction)
    }

    private static func alternatives(
        _ slot: BisSlot?,
        request: ItemTooltipRequest,
        items: [Int: Item],
        unit: BisUnit
    ) -> [ItemTooltipAlternative] {
        (slot?.alternatives ?? []).prefix(maxAlternatives).map { alt in
            let row = items[alt.itemId]
            let look = HomeItemLook(id: alt.itemId, name: row?.name ?? alt.itemName, quality: row?.quality, icon: row?.icon, itemLevel: row?.itemLevel)
            let delta = BisListBuilder.signedDelta(alt.dpsDelta).map { Copy.tooltipDelta(delta: $0, unit: unit.word) }
                ?? Copy.itemsSameDps(unit: unit.word)
            let worn = alt.itemId == request.wornID ? Copy.tooltipYouWear : ""
            return ItemTooltipAlternative(look: look, meta: worn + delta)
        }
    }

    /// `Best in slot for Head · Fury Warrior 60 · Phase 1 raid`, only when the item is the pick.
    /// The band label carries the phase; no preset id is ever printed.
    private static func footer(slot: BisSlot?, request: ItemTooltipRequest) -> String? {
        guard let slot, let specName = request.specName, let bandLabel = request.bandLabel else { return nil }
        return Copy.tooltipPickOf(slot: ItemWords.slot(slot.slot), spec: specName, band: bandLabel)
    }
}
