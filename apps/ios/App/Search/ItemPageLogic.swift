import Foundation
import Models

/// Where an item page was asked for: the item, and what the caller already knows of it (a BiS
/// row's or Home's look and class), drawn while the index loads.
struct ItemPageRoute: Hashable, Identifiable, Sendable {
    let itemID: Int
    var fallback: HomeItemLook?
    var classSlug: String?

    var id: Int { itemID }

    static func == (lhs: ItemPageRoute, rhs: ItemPageRoute) -> Bool { lhs.itemID == rhs.itemID }
    func hash(into hasher: inout Hasher) { hasher.combine(itemID) }
}

/// `Your Back pick` or `Runner-up 2 for Back` (spec F12.4), only for the current character's list.
struct YourPickCard: Equatable, Sendable {
    let title: String
    let meta: String
    let className: String?
}

/// One `Where it comes from` row.
struct ItemSourceRow: Equatable, Identifiable, Sendable {
    let id: Int
    let kind: String
    let title: String
    let meta: String

    var accessibilityLabel: String { meta.isEmpty ? title : Copy.itemPageA11ySource(title: title, meta: meta) }
}

/// `Pick` or `Runner-up {n}` on a `Best in slot for` row.
enum ItemSpecTag: Equatable, Sendable {
    case pick
    case runnerUp(Int)

    init(rank: Int) { self = rank <= 0 ? .pick : .runnerUp(rank) }

    var text: String {
        switch self {
        case .pick: Copy.itemPageTagPick
        case let .runnerUp(n): Copy.itemPageTagRunnerUp(n: n)
        }
    }
}

/// One `Best in slot for` row: the spec, its best placement, and the list its chevron opens.
struct ItemSpecRow: Equatable, Identifiable, Sendable {
    let use: ItemUseSpec
    let title: String
    let meta: String
    let icon: String?
    let tag: ItemSpecTag

    var id: String { use.spec }
    var accessibilityLabel: String { Copy.itemPageA11ySpec(spec: title, meta: meta, tag: tag.text) }
}

/// The item page's rules (spec F12.4, F12.5), pure.
enum ItemPageLogic {
    /// Six spec rows show, then `{n} more specs`.
    static let visibleSpecs = 6
    /// Boss drops are the kinds whose title is a boss and whose place is the instance.
    private static let bossKinds: Set<String> = ["raid", "dungeon"]

    /// `Rare · Back · item level 52`.
    static func meta(_ item: SearchItem) -> String {
        Copy.itemPageMeta(quality: ItemQuality(value: item.quality).name, slot: ItemVocabulary.slotLabel(item.slot), ilvl: item.itemLevel)
    }

    // MARK: Where it comes from

    static func sources(_ sources: [ItemUseSource]) -> [ItemSourceRow] {
        sources.enumerated().map { index, source in
            ItemSourceRow(id: index, kind: source.kind, title: source.title, meta: sourceMeta(source))
        }
    }

    /// `{place} · {kind}`, then for a boss drop `{n}% from {boss}` or `Drop chance not known`
    /// (never `0%`), and for a quest `Requires level {n}`.
    static func sourceMeta(_ source: ItemUseSource) -> String {
        var parts = [source.place, kindWord(source.kind)].compactMap { $0 }.filter { !$0.isEmpty && $0 != source.title }
        if bossKinds.contains(source.kind) {
            parts.append(source.chance.map { Copy.itemPageDropChance(chance: percent($0), boss: source.title) } ?? Copy.itemPageChanceUnknown)
        }
        if source.kind == "quest", let level = source.minLevel, level > 0 {
            parts.append(Copy.itemPageQuestLevel(n: level))
        }
        return parts.joined(separator: " · ")
    }

    /// The site's `dropPercentLabel`: a whole percent, `<1%` under one, so nothing reads `0%`.
    static func percent(_ chance: Double) -> String {
        chance < 1 ? "<1%" : "\(Int(chance.rounded()))%"
    }

    static func kindWord(_ kind: String) -> String? {
        switch kind {
        case "raid": Copy.itemPageKindRaid
        case "dungeon": Copy.itemPageKindDungeon
        case "world", "world_drop": Copy.itemPageKindWorld
        case "zone": Copy.itemPageKindZone
        case "vendor": Copy.itemPageKindVendor
        case "crafted": Copy.itemPageKindCrafted
        case "rep": Copy.itemPageKindRep
        case "pvp": Copy.itemPageKindPvp
        case "quest": Copy.itemPageKindQuest
        default: nil
        }
    }

    // MARK: Best in slot for

    /// One row per spec in the file's order (picks first, then class and spec).
    static func specs(_ uses: [ItemUseSpec], catalog: SpecCatalog) -> [ItemSpecRow] {
        uses.map { use in
            ItemSpecRow(
                use: use,
                title: SearchWords.specClass(spec: use.specName, classSlug: use.className),
                meta: use.preset == BisRange.raidPreset ? Copy.itemPageSpecMetaRaid(band: use.band) : Copy.itemPageSpecMetaBare(band: use.band),
                icon: catalog.spec(use.spec)?.icon,
                tag: ItemSpecTag(rank: use.rank)
            )
        }
    }

    /// The rows drawn, and how many `{n} more specs` hides.
    static func visible(_ rows: [ItemSpecRow], expanded: Bool) -> (rows: [ItemSpecRow], hidden: Int) {
        guard !expanded, rows.count > visibleSpecs else { return (rows, 0) }
        return (Array(rows.prefix(visibleSpecs)), rows.count - visibleSpecs)
    }

    /// The list a row's chevron opens: that spec's band and preset, its faction or the viewer's.
    static func selection(_ use: ItemUseSpec, viewerFaction: String?) -> BisSelection {
        let faction = use.faction ?? viewerFaction.flatMap { BisSelection.factions.contains($0) ? $0 : nil } ?? BisSelection.factions[1]
        return BisSelection(classSlug: use.className, spec: use.spec, range: BisRange(band: use.band, preset: use.preset), faction: faction)
    }

    // MARK: Your pick

    /// Where `itemID` sits on a band: the pick (0) or runner-up n of a slot.
    static func placement(of itemID: Int, in band: BisBand) -> (slot: String, rank: Int)? {
        for slot in band.slots {
            if slot.itemId == itemID { return (slot.slot, 0) }
            if let index = slot.alternatives?.firstIndex(where: { $0.itemId == itemID }) { return (slot.slot, index + 1) }
        }
        return nil
    }

    /// The card, only when the item is on the character's own list for its band, preset and faction.
    static func yourPick(itemID: Int, band: BisBand?, selection: BisSelection?, specName: String?, className: String?) -> YourPickCard? {
        guard let band, let selection, let placement = placement(of: itemID, in: band) else { return nil }
        let slot = ItemVocabulary.slotLabel(placement.slot)
        let title = placement.rank == 0 ? Copy.itemPagePick(slot: slot) : Copy.itemPageRunnerUp(n: placement.rank, slot: slot)
        let bandWord = BisLabels.band(selection.range)
        let faction = BisLabels.faction(selection.faction)
        let spec = specName ?? ""
        let meta = selection.range.preset == BisRange.raidPreset
            ? Copy.itemPagePickMeta(spec: spec, band: bandWord, preset: Copy.itemPagePresetRaid, faction: faction)
            : Copy.itemPagePickMetaBare(spec: spec, band: bandWord, faction: faction)
        return YourPickCard(title: title, meta: meta, className: className)
    }
}
