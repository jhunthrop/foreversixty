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

/// One runner-up under "The next best": its delta against the pick and where it comes from.
struct ItemTooltipAlternative: Equatable, Identifiable {
    let look: HomeItemLook
    let meta: String
    let source: String?
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
    /// White `+n Name` lines.
    let stats: [String]
    /// Green `Equip:` sentences, as the client prints spell power, healing, attack power, mp5…
    let equipLines: [String]
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
            slotName: weapon ? Self.weaponSlotName(item) : slotName,
            typeName: weapon ? item.subclass.flatMap(ItemVocabulary.subclassName) : ItemVocabulary.armorTypeName(item.armorType),
            armor: item.armor > 0 ? Copy.tooltipArmor(n: item.armor) : nil,
            weaponDamage: weapon ? Copy.itemsWeaponDamage(min: Int(item.damageMin.rounded()), max: Int(item.damageMax.rounded())) : nil,
            weaponSpeed: weapon ? Copy.itemsWeaponSpeed(speed: String(format: "%.2f", item.speed)) : nil,
            weaponDPS: weapon ? Copy.itemsWeaponDps(dps: String(format: "%.1f", item.dps)) : nil,
            stats: Self.statLines(item),
            equipLines: Self.equipLines(item),
            effect: Self.effect(item),
            setName: set?.name,
            setBonuses: set?.bonuses.map { Copy.itemsSetBonus(n: $0.pieces, text: $0.description) } ?? [],
            requiresLevel: item.requiredLevel > 0 ? Copy.tooltipRequires(n: item.requiredLevel) : nil,
            itemLevel: Copy.tooltipItemLevel(n: item.itemLevel),
            source: Self.source(item, request: request, band: band),
            alternatives: Self.alternatives(pickSlot, request: request, items: sources.items, unit: BisUnit(role: band?.role)),
            pickFooter: pickSlot.flatMap { Self.footer(slot: $0.slot, request: request) }
        )
    }

    private static func unknown(_ request: ItemTooltipRequest) -> ItemTooltipModel {
        let quality = ItemQuality(value: request.fallback.quality)
        return ItemTooltipModel(
            look: request.fallback, headerMeta: quality.name, known: false, binds: nil, unique: false,
            slotName: nil, typeName: nil, armor: nil, weaponDamage: nil, weaponSpeed: nil, weaponDPS: nil,
            stats: [], equipLines: [], effect: nil, setName: nil, setBonuses: [], requiresLevel: nil,
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

    /// The left word of a weapon's slot line as the client prints it (`Two-Hand`, `Main Hand`):
    /// the client's InventoryType when the data carries it, else what `two_hand` and the slot say.
    private static func weaponSlotName(_ item: Item) -> String? {
        if let word = item.weaponSlot.flatMap(ItemVocabulary.weaponSlotName) { return word }
        if item.twoHand { return Copy.itemsWeaponSlotTwoHand }
        return item.slot == "ranged" ? Copy.itemsWeaponSlotRanged : Copy.itemsWeaponSlotOneHand
    }

    /// `+22 Intellect`, `-15 Parry`: the client's sign convention, in the site's stat order (its
    /// `statLines`). A stat the client prints as an `Equip:` sentence (spell power, healing, attack
    /// power, mp5 and the rest, `ItemVocabulary.equipLine`) or a rating stat `percent_stats`
    /// states as a percent is not a white line.
    private static func statLines(_ item: Item) -> [String] {
        ItemVocabulary.statOrder.compactMap { stat in
            guard let value = item.stats[stat.key], value != 0, percentLine(item, key: stat.key) == nil,
                  ItemVocabulary.equipLine(stat.key, value: value) == nil else { return nil }
            return "\(value < 0 ? "-" : "+")\(abs(value)) \(stat.name)"
        }
    }

    /// The client's green `Equip:` sentences: the percent-stated rating stats, then one per positive
    /// Equip stat share. One spell that gives equal spell power and healing is one sentence (the
    /// site's `equipShares`). A sentence the item's own effect text already says is not repeated.
    private static func equipLines(_ item: Item) -> [String] {
        let shares = equipShares(item)
        let percent = ItemVocabulary.statOrder.compactMap { percentLine(item, key: $0.key) }
        let spells = ItemVocabulary.statOrder.compactMap { stat -> String? in
            guard let share = shares[stat.key], share > 0 else { return nil }
            return ItemVocabulary.equipLine(stat.key, value: share)
        }
        let effect = Self.effect(item) ?? ""
        return (percent + spells).filter { !effect.localizedCaseInsensitiveContains($0.dropFirst(equipPrefix.count)) }
    }

    private static let equipPrefix = "Equip: "

    /// The amount each Equip sentence states: one "Increase Spell Dam" spell gives equal spell power
    /// and healing and is one sentence, so healing states only what it adds beyond spell power.
    static func equipShares(_ item: Item) -> [String: Int] {
        var shares = item.stats
        let spellPower = shares["spell_power"] ?? 0
        shares["healing"] = max(0, (shares["healing"] ?? 0) - spellPower)
        return shares
    }

    private static func percentLine(_ item: Item, key: String) -> String? {
        guard let value = item.stats[key], value != 0, item.percentStats[key] == value else { return nil }
        return ItemVocabulary.percentLine(key, value: value)
    }

    private static func effect(_ item: Item) -> String? {
        let text = item.effectText.split(whereSeparator: \.isWhitespace).joined(separator: " ")
        return text.isEmpty ? nil : text
    }

    /// Where it comes from: the site's formatted line (the BiS row's, else the caller's, which the
    /// API formats the same way), plus the standing a vendor requires (` · requires Friendly`).
    /// A reputation reward's line already names its standing.
    private static func source(_ item: Item, request: ItemTooltipRequest, band: BisBand?) -> String? {
        let (line, kind) = bisSource(of: item.id, in: band) ?? (request.sourceLine, nil)
        guard let line else { return nil }
        guard kind == "vendor", let standing = item.minStanding.flatMap(ItemVocabulary.standingName) else { return line }
        return line + Copy.tooltipRequiresStanding(standing: standing)
    }

    private static func bisSource(of id: Int, in band: BisBand?) -> (String?, String?)? {
        guard let band else { return nil }
        if let slot = band.slots.first(where: { $0.itemId == id }), let line = slot.sourceLine ?? slot.source {
            return (line, slot.sourceKind)
        }
        for slot in band.slots {
            if let alt = slot.alternatives?.first(where: { $0.itemId == id }) { return (alt.sourceLine ?? alt.source, alt.sourceKind) }
        }
        return nil
    }

    /// Spec F2.6: up to two runners-up.
    private static let maxAlternatives = 2

    /// Up to two runners-up from the BiS row, each in the band's unit (DPS, HPS or score); the
    /// worn one leads with "You wear this · ".
    private static func alternatives(
        _ slot: BisSlot?, request: ItemTooltipRequest, items: [Int: Item], unit: BisUnit
    ) -> [ItemTooltipAlternative] {
        (slot?.alternatives ?? []).prefix(maxAlternatives).map { alt in
            let row = items[alt.itemId]
            let look = HomeItemLook(id: alt.itemId, name: row?.name ?? alt.itemName, quality: row?.quality, icon: row?.icon, itemLevel: row?.itemLevel)
            let delta = BisListBuilder.signedDelta(alt.dpsDelta).map { Copy.tooltipDelta(delta: $0, unit: unit.word) }
                ?? Copy.itemsSameDps(unit: unit.word)
            let worn = alt.itemId == request.wornID ? Copy.tooltipYouWear : ""
            return ItemTooltipAlternative(look: look, meta: worn + delta, source: alt.sourceLine ?? alt.source)
        }
    }

    /// `Best in slot for Finger 1 · Fury Warrior 60 · Phase 1 raid`, named by the BiS row's slot; the
    /// band label carries the phase, never a preset id.
    private static func footer(slot: String, request: ItemTooltipRequest) -> String? {
        guard let specName = request.specName, let bandLabel = request.bandLabel else { return nil }
        return Copy.tooltipPickOf(slot: ItemVocabulary.slotLabel(slot), spec: specName, band: bandLabel)
    }
}
