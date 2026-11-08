import Foundation
import Models
import Networking

/// Who opened the item sheet (flows spec F5): the tooltip's request, the gain the row showed, and
/// the guild whose members may need the item (nil outside a guild).
struct ItemSheetRequest: Equatable, Identifiable, Sendable {
    let tooltip: ItemTooltipRequest
    let gain: HomeGain?
    let guildID: Int64?

    var id: Int { tooltip.itemID }
}

/// One of the client's tooltip lines, in its colour.
struct ItemSheetLine: Equatable, Hashable {
    enum Tone: Equatable, Hashable { case plain, green, gold }

    let text: String
    var tone: Tone = .plain
}

/// A row of the sheet's sections: a title, a line under it, and its lead image.
struct ItemSheetRow: Equatable, Identifiable {
    enum Lead: Equatable {
        case item(icon: String?, quality: Int?)
        case crest(className: String)
    }

    let id: String
    let title: String
    var titleQuality: Int?
    var titleClass: String?
    let sub: String?
    let lead: Lead
}

/// Who in the guild needs it, as far as it has loaded.
enum ItemSheetGuild: Equatable {
    /// Not in a guild: the section is left out.
    case none
    case loading
    case loaded(rows: [ItemSheetRow], note: String?)
    case failed
}

/// Everything the item sheet draws (mock ios-item-sheet): the item with its gain, the client's
/// tooltip lines (no item level: the client tooltip does not show it), where to get it with its
/// chance, the runner-up, who in the guild needs it or has it.
struct ItemSheetContent: Equatable {
    let look: HomeItemLook
    let slotLabel: String
    let gain: HomeGain?
    let lines: [ItemSheetLine]
    let source: ItemSheetRow?
    let runnerUp: ItemSheetRow?
    let guild: ItemSheetGuild

    func with(guild: ItemSheetGuild) -> ItemSheetContent {
        ItemSheetContent(look: look, slotLabel: slotLabel, gain: gain, lines: lines, source: source, runnerUp: runnerUp, guild: guild)
    }
}

/// Words the sheet from the build's files, the loot file and the guild's answer. Pure.
enum ItemSheetBuilder {
    static func content(
        _ request: ItemSheetRequest, sources: ItemTooltipSources, loot: LootFile?, guild: ItemSheetGuild
    ) -> ItemSheetContent {
        let tooltip = ItemTooltipModel.build(request.tooltip, sources: sources)
        let band = ItemTooltipModel.band(for: request.tooltip, in: sources.bis)
        let places = loot.map { ItemPlaces(loot: $0, level: nil) }
        let itemID = request.tooltip.itemID
        let listed = band.flatMap { listing(of: itemID, in: $0) }
        let item = sources.items[itemID]
        let slot = listed?.slot.slot ?? item?.slot
        return ItemSheetContent(
            look: tooltip.look,
            slotLabel: slot.map { ItemVocabulary.slotLabel($0, classSlug: request.tooltip.classSlug) } ?? tooltip.headerMeta,
            gain: request.gain,
            lines: lines(tooltip),
            source: source(itemID, look: tooltip.look, kind: listed?.kind, named: listed?.source, places: places, fallback: tooltip.source),
            runnerUp: runnerUp(listed, band: band, items: sources.items, places: places),
            guild: guild
        )
    }

    /// The client's lines after the name, in its order; no item level.
    static func lines(_ model: ItemTooltipModel) -> [ItemSheetLine] {
        var out: [String?] = [model.binds, model.unique ? Copy.tooltipUnique : nil]
        out.append([model.slotName, model.typeName].compactMap(\.self).joined(separator: separator).nilIfEmpty)
        out.append(model.armor)
        out.append([model.weaponDamage, model.weaponSpeed].compactMap(\.self).joined(separator: separator).nilIfEmpty)
        out.append(model.weaponDPS)
        var lines = out.compactMap(\.self).map { ItemSheetLine(text: $0) } + model.stats.map { ItemSheetLine(text: $0) }
        if let effect = model.effect { lines.append(ItemSheetLine(text: effect, tone: .green)) }
        if let requires = model.requiresLevel { lines.append(ItemSheetLine(text: requires)) }
        lines += model.equipLines.map { ItemSheetLine(text: $0, tone: .green) }
        if let set = model.setName {
            lines.append(ItemSheetLine(text: set, tone: .gold))
            lines += model.setBonuses.map { ItemSheetLine(text: $0) }
        }
        return lines
    }

    private static let separator = " · "

    /// Where the band lists the item: its slot, and its kind and source as a pick or a runner-up.
    struct Listing {
        let slot: BisSlot
        let isPick: Bool
        let kind: String?
        let source: String?
    }

    static func listing(of item: Int, in band: BisBand) -> Listing? {
        if let slot = band.slots.first(where: { $0.itemId == item }) {
            return Listing(slot: slot, isPick: true, kind: slot.sourceKind, source: slot.source)
        }
        for slot in band.slots {
            if let alt = slot.alternatives?.first(where: { $0.itemId == item }) {
                return Listing(slot: slot, isPick: false, kind: alt.sourceKind, source: alt.source)
            }
        }
        return nil
    }

    /// Where to get it: an instance by its boss with the drop chance, a profession, a reputation
    /// with the standing it needs; any other source in the list's words.
    static func source(
        _ item: Int, look: HomeItemLook, kind: String?, named: String?, places: ItemPlaces?, fallback: String?
    ) -> ItemSheetRow? {
        let lead = ItemSheetRow.Lead.item(icon: look.icon, quality: look.quality)
        guard let place = places?.place(of: item, kind: kind, named: named) else {
            return fallback.map { ItemSheetRow(id: "source", title: $0, sub: nil, lead: lead) }
        }
        switch place.kind {
        case .instance:
            let title = place.boss.map { Copy.itemsSheetBossPlace(boss: $0, place: place.name) } ?? place.name
            let raid = place.sourceKind == ItemPlaces.raidKind
            let parts = [place.chance.map { Copy.gearSheetDropChance(chance: ItemNumbers.percent($0)) },
                         raid ? Copy.itemsSheetRaid : Copy.itemsSheetDungeon]
            return ItemSheetRow(id: "source", title: title, sub: parts.compactMap(\.self).joined(separator: separator), lead: lead)
        case .profession:
            return ItemSheetRow(id: "source", title: place.name, sub: GearPresenter.crafter(place.profession), lead: lead)
        case .reputation:
            let standing = place.standing.map { ItemVocabulary.standingName($0)?.lowercased() ?? $0 }
            return ItemSheetRow(id: "source", title: place.name, sub: standing.map { Copy.itemsSheetNeedsStanding(standing: $0) }, lead: lead)
        }
    }

    /// The slot's next pick after this one: `Onyxia · 8% · −1.8 DPS`.
    static func runnerUp(_ listed: Listing?, band: BisBand?, items: [Int: Item], places: ItemPlaces?) -> ItemSheetRow? {
        guard let listed, listed.isPick, let alt = listed.slot.alternatives?.first else { return nil }
        let unit = BisUnit(role: band?.role).word
        let item = items[alt.itemId]
        let place = places?.place(of: alt.itemId, kind: alt.sourceKind, named: alt.source)
        let delta = BisListBuilder.signedDelta(alt.dpsDelta).map { "\($0) \(unit)" } ?? Copy.itemsSameDps(unit: unit)
        let parts = [place?.boss ?? place?.name ?? alt.sourceLine ?? alt.source, (place?.chance ?? alt.dropChance).map(ItemNumbers.percent), delta]
        return ItemSheetRow(
            id: "runner-up", title: item?.name ?? alt.itemName, titleQuality: item?.quality ?? 1,
            sub: parts.compactMap(\.self).filter { !$0.isEmpty }.joined(separator: separator),
            lead: .item(icon: item?.icon, quality: item?.quality)
        )
    }

    /// Who needs it (picks first, as the server orders them), then who has it.
    static func guildRows(_ answer: ItemNeeds) -> ItemSheetGuild {
        let needs = answer.needs.map { need in
            let spec = specName(need.spec, need.className)
            return ItemSheetRow(
                id: "need-\(need.characterKey)", title: need.name, titleClass: need.className,
                sub: need.pick ? Copy.gearSheetNeedsPick(spec: spec) : Copy.gearSheetNeedsRunnerUp(spec: spec),
                lead: .crest(className: need.className.lowercased())
            )
        }
        let has = answer.has.map { holder in
            ItemSheetRow(
                id: "has-\(holder.characterKey)", title: holder.name, titleClass: holder.className,
                sub: Copy.gearSheetHas(spec: specName(holder.spec, holder.className)),
                lead: .crest(className: holder.className.lowercased())
            )
        }
        let rows = needs + has
        let note: String? = if rows.isEmpty {
            Copy.gearSheetNobody
        } else if answer.unknownGear > 0 {
            Copy.gearSheetUnknown(n: answer.unknownGear)
        } else {
            nil
        }
        return .loaded(rows: rows, note: note)
    }

    /// `Protection Warrior`.
    static func specName(_ spec: String, _ className: String) -> String {
        spec.isEmpty ? className : "\(spec) \(className)"
    }
}

/// Who in a guild needs an item (`GET /v1/guilds/{id}/items/{item}/needs`).
protocol ItemNeedsSource: Sendable {
    func needs(guild: Int64, item: Int) async throws(APIError) -> ItemNeeds
}

struct LiveItemNeedsSource: ItemNeedsSource {
    let client: APIClient

    func needs(guild: Int64, item: Int) async throws(APIError) -> ItemNeeds {
        try await client.send(.getGuildItemNeeds(id: guild, item: item))
    }
}

private extension String {
    var nilIfEmpty: String? { isEmpty ? nil : self }
}
