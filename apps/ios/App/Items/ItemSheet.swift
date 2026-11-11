import Foundation
import Models
import Networking
import Routing

/// Who opened the item sheet (flows spec F5): the tooltip's request, the gain the row showed, and
/// the guild whose members may need the item (nil outside a guild).
struct ItemSheetRequest: Equatable, Identifiable, Sendable {
    let tooltip: ItemTooltipRequest
    let gain: HomeGain?
    let guildID: Int64?

    var id: Int { tooltip.itemID }
}

/// A row of the sheet's sections: a title, a line under it, and its lead image.
struct ItemSheetRow: Equatable, Identifiable {
    enum Lead: Equatable {
        case item(icon: String?, quality: Int?)
        case crest(className: String)
        /// No image: the source row is words alone, so it never reads as an item (round 8, row 1).
        case none
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

/// The guild's next raid to the zone that drops the item (needs.next_raid), opening that raid.
struct ItemSheetNextRaid: Equatable {
    let line: String
    let link: DeepLink
}

/// Everything the item sheet draws (round 5, mocks ios-item-sheet and m-ios-item): the client's
/// tooltip with the gain closing it (no item level: the client tooltip does not show it), then one
/// list: where it drops with its chance (and the guild's next raid there), what it replaces, the
/// runner-up, who in the guild needs it or has it.
struct ItemSheetContent: Equatable {
    let tooltip: ItemTooltipModel
    let gain: HomeGain?
    let source: ItemSheetRow?
    /// `Replaces Ironwood Blade`, Equipped: what the character wears in the slot today.
    let replaces: ItemSheetRow?
    let runnerUp: ItemSheetRow?
    let guild: ItemSheetGuild
    /// The raid instance the item drops in, by name, for the next raid line; nil for any other source.
    var raidZone: String?
    var nextRaid: ItemSheetNextRaid?

    var look: HomeItemLook { tooltip.look }

    func with(guild: ItemSheetGuild, nextRaid: ItemSheetNextRaid? = nil) -> ItemSheetContent {
        ItemSheetContent(
            tooltip: tooltip, gain: gain, source: source, replaces: replaces, runnerUp: runnerUp, guild: guild,
            raidZone: raidZone, nextRaid: nextRaid
        )
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
        let place = places?.place(of: itemID, kind: listed?.kind, named: listed?.source)
        // The client's tooltip has no item level (spec 2026-11-08, round 4 cut), so the sheet drops it.
        return ItemSheetContent(
            tooltip: tooltip.withoutItemLevel,
            gain: request.gain,
            // Unplaced by the loot file, the list's own source in the shared words (`Vendor: Lady Palanseer`).
            source: source(
                itemID, kind: listed?.kind, named: listed?.source, places: places,
                fallback: listed?.source.map { BisListBuilder.sourceLine($0, kind: listed?.kind) } ?? tooltip.source
            ),
            replaces: replaces(request.tooltip.wornID, item: itemID, items: sources.items),
            runnerUp: runnerUp(listed, band: band, items: sources.items, places: places, pickPlace: place),
            guild: guild,
            raidZone: place.flatMap { $0.kind == .instance && $0.sourceKind == ItemPlaces.raidKind ? $0.name : nil }
        )
    }

    /// `Next raid: Onyxia's Lair, Wed 8:00 PM`: the API sends next_raid only for the zone that drops
    /// the item; the day and time read in the device's locale and time zone.
    static func nextRaid(
        _ next: ItemNextRaid?, guild: Int64, zone: String?, locale: Locale = .current, timeZone: TimeZone = .current
    ) -> ItemSheetNextRaid? {
        guard let next, let zone else { return nil }
        var style = Date.FormatStyle(date: .omitted, time: .shortened, locale: locale, timeZone: timeZone)
        style = style.weekday(.abbreviated)
        return ItemSheetNextRaid(
            // ICU puts a narrow no-break space before AM/PM; the copy's lines use plain spaces.
            line: Copy.itemsSheetNextRaid(zone: zone, when: next.startsAt.formatted(style).replacingOccurrences(of: "\u{202F}", with: " ")),
            link: .raid(guild: guild, raid: next.raidId)
        )
    }

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

    /// Where to get it, one self-describing row with no image and no caption (round 8, row 1; Android's
    /// anatomy): an instance by its boss with the drop chance (`Onyxia, Onyxia's Lair, 8%`), a profession
    /// with the skill its recipe needs (`Crafted: Blacksmithing (245)`, WoW's Requires shape), a
    /// reputation with its standing (`Stormpike Guard, revered`); any other source in the list's words
    /// (`World drop`).
    static func source(_ item: Int, kind: String?, named: String?, places: ItemPlaces?, fallback: String?) -> ItemSheetRow? {
        guard let place = places?.place(of: item, kind: kind, named: named) else {
            return fallback.map { ItemSheetRow(id: "source", title: withoutFaction($0), sub: nil, lead: .none) }
        }
        let parts: [String?] = switch place.kind {
        case .instance: [place.boss, place.name, place.chance.map(ItemNumbers.percent)]
        case .profession: [professionWords(place)]
        case .reputation: [place.name, standingWord(place)]
        }
        return ItemSheetRow(id: "source", title: parts.compactMap(\.self).joined(separator: listSeparator), sub: nil, lead: .none)
    }

    /// `Crafted: Blacksmithing (245)`, or without the skill when the loot file has none.
    private static func professionWords(_ place: ItemPlace) -> String {
        place.skill.map { Copy.bisSourceCraftedSkill(name: place.name, skill: $0) } ?? Copy.bisSourceCrafted(name: place.name)
    }

    private static func standingWord(_ place: ItemPlace) -> String? {
        place.standing.map { ItemVocabulary.standingName($0)?.lowercased() ?? $0 }
    }

    /// A PvP source without its faction word: the rank title already belongs to one faction
    /// (`PvP rank 11 · Blood Guard`, round 9, row 3).
    static func withoutFaction(_ line: String) -> String {
        let factions: Set = [Copy.bisFactionHorde, Copy.bisFactionAlliance]
        var parts = line.components(separatedBy: separator)
        if parts.count > 1, let last = parts.last, factions.contains(last) { parts.removeLast() }
        return parts.joined(separator: separator)
    }

    /// Whether the runner-up comes from the pick's own place (Android's round-9 rule): placed by the loot
    /// file, the same kind, name and boss as the pick's place; else its list words equal the pick's.
    static func isPickPlace(_ runnerUp: ItemPlace?, words: String?, pick: ItemPlace?, pickWords: String?) -> Bool {
        if let runnerUp {
            return runnerUp.sourceKind == pick?.sourceKind && runnerUp.name == pick?.name && runnerUp.boss == pick?.boss
        }
        return words != nil && words == pickWords
    }

    /// The runner-up's source in the source row's shared words, the boss in its short shape
    /// (`Ancient Treasure · 1%`), the list's words (`bis.source.*`) when the loot file has no place; nil
    /// when nothing is known.
    static func runnerUpSource(_ place: ItemPlace?, chance: Double?, line: String?, kind: String?) -> String? {
        guard let place else { return line.map { withoutFaction(BisListBuilder.sourceLine($0, kind: kind)) } }
        switch place.kind {
        case .instance:
            return [place.boss ?? place.name, chance.map(ItemNumbers.percent)].compactMap(\.self).joined(separator: separator)
        case .profession:
            return professionWords(place)
        case .reputation:
            return [place.name, standingWord(place)].compactMap(\.self).joined(separator: listSeparator)
        }
    }

    private static let listSeparator = ", "
    private static let separator = " · "

    /// `Replaces Ironwood Blade`, Equipped, with its icon: what the slot holds today; nil when the
    /// item is what it holds, or nothing is known.
    static func replaces(_ wornID: Int?, item: Int, items: [Int: Item]) -> ItemSheetRow? {
        guard let wornID, wornID != item, let worn = items[wornID] else { return nil }
        return ItemSheetRow(
            id: "replaces", title: Copy.gearSheetReplaces(item: worn.name), sub: Copy.gearWorn,
            lead: .item(icon: worn.icon, quality: worn.quality)
        )
    }

    /// The slot's next pick after this one: `Runner-up · Onyxia · 8% · −1.8 DPS`; its source only when it
    /// differs from the pick's (round 9, row 2: `Runner-up · −5.0 DPS` under the same crafter), and `Same
    /// DPS` when nothing separates them (row 1).
    static func runnerUp(
        _ listed: Listing?, band: BisBand?, items: [Int: Item], places: ItemPlaces?, pickPlace: ItemPlace? = nil
    ) -> ItemSheetRow? {
        guard let listed, listed.isPick, let alt = listed.slot.alternatives?.first else { return nil }
        let unit = BisUnit(role: band?.role).word
        let item = items[alt.itemId]
        let place = places?.place(of: alt.itemId, kind: alt.sourceKind, named: alt.source)
        let line = alt.sourceLine ?? alt.source
        let own = isPickPlace(place, words: line, pick: pickPlace, pickWords: listed.slot.sourceLine ?? listed.source)
        let delta = BisListBuilder.signedDelta(alt.dpsDelta).map { "\($0) \(unit)" } ?? Copy.itemsSameDps(unit: unit)
        let parts = [
            Copy.gearSheetRunnerUp,
            own ? nil : runnerUpSource(place, chance: place?.chance ?? alt.dropChance, line: line, kind: alt.sourceKind),
            delta,
        ]
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

extension ItemTooltipModel {
    /// The model without its item level line.
    var withoutItemLevel: ItemTooltipModel {
        ItemTooltipModel(
            look: look, headerMeta: headerMeta, known: known, binds: binds, unique: unique, slotName: slotName, typeName: typeName,
            armor: armor, weaponDamage: weaponDamage, weaponSpeed: weaponSpeed, weaponDPS: weaponDPS, stats: stats,
            equipLines: equipLines, effect: effect, setName: setName, setBonuses: setBonuses, requiresLevel: requiresLevel,
            itemLevel: nil, source: source, alternatives: alternatives, pickFooter: pickFooter
        )
    }
}
