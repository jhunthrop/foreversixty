import Foundation
import Models

/// Who asked for a tooltip and what they already know: the item, the build its static data
/// lives in, and the band the pick belongs to. Built by Home and the gear list.
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
    /// The server's source words for this item, when the caller has them (a Home pick).
    let sourceLine: String?
    /// The item worn in the same slot, so the next best can say "You wear this".
    let wornID: Int?

    var id: Int { itemID }
}

/// An item's name, quality, icon and level as the API sends them; every field but id may be null.
struct HomeItemLook: Equatable, Sendable {
    let id: Int
    let name: String?
    let quality: Int?
    let icon: String?
    let itemLevel: Int?
}

/// The static files a tooltip reads, already loaded.
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

/// Everything the tooltip sheet draws, in the game's order (spec F2.6, the site's
/// `lib/items/tooltip.ts`). Pure: the sheet renders it and the tests read it.
struct ItemTooltipModel: Equatable {
    let look: HomeItemLook
    let headerMeta: String
    /// False when the build's file does not list the item; the panel then says so.
    let known: Bool
    let binds: String?
    let unique: Bool
    let slotName: String?
    let typeName: String?
    let armor: String?
    let weaponDamage: String?
    let weaponSpeed: String?
    let weaponDPS: String?
    let stats: [String]
    let effect: String?
    let setName: String?
    let setBonuses: [String]
    let requiresLevel: String?
    let itemLevel: String?
    let source: String?
    let alternatives: [ItemTooltipAlternative]
    let pickFooter: String?

    static func build(_ request: ItemTooltipRequest, sources: ItemTooltipSources) -> ItemTooltipModel {
        guard let item = sources.items[request.itemID] else { return unknown(request) }
        let quality = ItemQuality(value: item.quality)
        let slotName = ItemVocabulary.slotName(item.slot)
        let band = Self.band(for: request, in: sources.bis)
        let pickSlot = band?.slots.first { $0.itemId == item.id }
        let weapon = Self.isWeapon(item)
        let set = item.setId.flatMap { id in sources.sets.first { $0.id == id } }
        return ItemTooltipModel(
            look: HomeItemLook(id: item.id, name: item.name, quality: item.quality, icon: item.icon, itemLevel: item.itemLevel),
            headerMeta: Copy.itemsHeader(quality: quality.name, slot: slotName ?? ""),
            known: true,
            binds: ItemVocabulary.bindsLine(item.bonding),
            unique: item.unique,
            slotName: slotName,
            typeName: Self.typeName(item, isWeapon: weapon),
            armor: item.armor > 0 ? Copy.tooltipArmor(n: item.armor) : nil,
            weaponDamage: weapon ? Copy.itemsWeaponDamage(min: Int(item.damageMin.rounded()), max: Int(item.damageMax.rounded())) : nil,
            weaponSpeed: weapon ? Copy.itemsWeaponSpeed(speed: String(format: "%.2f", item.speed)) : nil,
            weaponDPS: weapon ? Copy.itemsWeaponDps(dps: String(format: "%.1f", item.dps)) : nil,
            stats: Self.statLines(item),
            effect: Self.effect(item),
            setName: set?.name,
            setBonuses: set?.bonuses.map { Copy.itemsSetBonus(n: $0.pieces, text: $0.description) } ?? [],
            requiresLevel: item.requiredLevel > 0 ? Copy.tooltipRequires(n: item.requiredLevel) : nil,
            itemLevel: Copy.tooltipItemLevel(n: item.itemLevel),
            source: Self.source(item, request: request, band: band),
            alternatives: Self.alternatives(pickSlot, request: request, items: sources.items),
            pickFooter: pickSlot.flatMap { _ in Self.footer(slotName: slotName, request: request, bis: sources.bis) }
        )
    }

    private static func unknown(_ request: ItemTooltipRequest) -> ItemTooltipModel {
        let quality = ItemQuality(value: request.fallback.quality)
        return ItemTooltipModel(
            look: request.fallback, headerMeta: quality.name, known: false, binds: nil, unique: false,
            slotName: nil, typeName: nil, armor: nil, weaponDamage: nil, weaponSpeed: nil, weaponDPS: nil,
            stats: [], effect: nil, setName: nil, setBonuses: [], requiresLevel: nil,
            itemLevel: request.fallback.itemLevel.map { Copy.tooltipItemLevel(n: $0) },
            source: request.sourceLine, alternatives: [], pickFooter: nil
        )
    }

    /// The band the pick belongs to: the request's band and preset, the character's faction when known.
    static func band(for request: ItemTooltipRequest, in bis: BisFile?) -> BisBand? {
        guard let bis, let level = request.band else { return nil }
        let candidates = bis.bands.filter { $0.band == level && $0.preset == (request.preset ?? $0.preset) }
        return candidates.first { $0.faction == request.faction } ?? candidates.first
    }

    private static let weaponSlots: Set<String> = ["main_hand", "off_hand", "ranged"]

    private static func isWeapon(_ item: Item) -> Bool {
        weaponSlots.contains(item.slot) && item.damageMax > 0
    }

    /// The right-hand word on the slot line: the armor type from the client, else a weapon's
    /// handedness, which is exact from the published row (the site's `typeLabel`).
    private static func typeName(_ item: Item, isWeapon: Bool) -> String? {
        if let armor = ItemVocabulary.armorTypeName(item.armorType) { return armor }
        guard isWeapon else { return nil }
        if item.twoHand { return Copy.itemsWeaponTwoHand }
        return item.slot == "ranged" ? Copy.itemsWeaponRanged : Copy.itemsWeaponOneHand
    }

    /// `+22 Intellect`, `-15 Parry`: the client's sign convention, in the site's stat order.
    private static func statLines(_ item: Item) -> [String] {
        ItemVocabulary.statOrder.compactMap { stat in
            guard let value = item.stats[stat.key], value != 0 else { return nil }
            return "\(value < 0 ? "-" : "+")\(abs(value)) \(stat.name)"
        }
    }

    private static func effect(_ item: Item) -> String? {
        let text = item.effectText.split(whereSeparator: \.isWhitespace).joined(separator: " ")
        return text.isEmpty ? nil : text
    }

    /// Where it comes from: the BiS row's words for the item, else the caller's, with the
    /// reputation standing the client requires (`Warsong Outriders (Exalted)`).
    private static func source(_ item: Item, request: ItemTooltipRequest, band: BisBand?) -> String? {
        let (words, kind) = bisSource(of: item.id, in: band) ?? (request.sourceLine, nil)
        guard let words else { return nil }
        guard let standing = item.minStanding.flatMap(ItemVocabulary.standingName) else { return words }
        if kind == "rep" { return Copy.tooltipRepSource(faction: words, standing: standing) }
        return words + Copy.tooltipRequiresStanding(standing: standing)
    }

    private static func bisSource(of id: Int, in band: BisBand?) -> (String?, String?)? {
        guard let band else { return nil }
        if let slot = band.slots.first(where: { $0.itemId == id }), let words = slot.source {
            return (words, slot.sourceKind)
        }
        for slot in band.slots {
            if let alt = slot.alternatives?.first(where: { $0.itemId == id }) { return (alt.source, alt.sourceKind) }
        }
        return nil
    }

    /// Up to two runners-up from the BiS row; the worn one leads with "You wear this · ".
    private static func alternatives(_ slot: BisSlot?, request: ItemTooltipRequest, items: [Int: Item]) -> [ItemTooltipAlternative] {
        (slot?.alternatives ?? []).prefix(2).map { alt in
            let row = items[alt.itemId]
            let look = HomeItemLook(id: alt.itemId, name: row?.name ?? alt.itemName, quality: row?.quality, icon: row?.icon, itemLevel: row?.itemLevel)
            let delta = alt.dpsDelta == 0 ? Copy.itemsSameDps : Copy.tooltipDelta(delta: ItemNumbers.delta(alt.dpsDelta))
            let worn = alt.itemId == request.wornID ? Copy.tooltipYouWear : ""
            return ItemTooltipAlternative(look: look, meta: worn + delta)
        }
    }

    private static func footer(slotName: String?, request: ItemTooltipRequest, bis: BisFile?) -> String? {
        guard let slotName, let specName = request.specName, let bandLabel = request.bandLabel,
              let preset = request.preset, let label = bis?.presets[preset]?.label
        else { return nil }
        return Copy.tooltipPickOf(slot: slotName, spec: specName, band: bandLabel, preset: label)
    }
}
