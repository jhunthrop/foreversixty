import Foundation
import Models

/// Where a list item comes from, as the web's page model resolves it (`lib/bis/page-model.ts`
/// `BisSourceInfo`: the place, its kind and a reputation's standing): an instance with the boss
/// and its chance, a profession, or a reputation at a standing.
struct ItemPlace: Equatable, Sendable {
    /// Where to go tonight's groups, in their order.
    enum Kind: Int, Comparable, Sendable, CaseIterable {
        case instance, profession, reputation

        static func < (lhs: Kind, rhs: Kind) -> Bool { lhs.rawValue < rhs.rawValue }
    }

    let kind: Kind
    /// The loot file's kind: `dungeon`, `raid`, `crafted` or `rep`.
    let sourceKind: String
    /// The instance, the profession (`Blacksmithing`) or the reputation faction.
    let name: String
    var boss: String?
    /// Percent, when the loot file names one.
    var chance: Double?
    /// A reputation's standing, lower case (`exalted`).
    var standing: String?
    /// A profession's key (`blacksmithing`).
    var profession: String?
    /// The profession skill the recipe needs (`245`), when the loot file knows it.
    var skill: Int?
}

/// Every item of the loot file by where it comes from, for the kinds Where to go groups by.
/// An instance above the character's reach (its lowest level more than `levelReach` above) is left
/// out; of the instances that drop an item, the one with the best chance wins (the web's
/// `findBestChanceSource`).
struct ItemPlaces: Sendable {
    static let dungeonKind = "dungeon"
    static let raidKind = "raid"
    static let craftedKind = "crafted"
    static let repKind = "rep"
    /// An instance whose lowest level is this far above the character's is not tonight's run.
    static let levelReach = 3

    private var instances: [String: [Int: ItemPlace]] = [:]
    private var crafted: [Int: ItemPlace] = [:]
    private var reputations: [Int: [ItemPlace]] = [:]

    init(loot: LootFile, level: Int?) {
        for source in loot.sources {
            switch source.kind {
            case Self.dungeonKind, Self.raidKind:
                if let level, let min = source.levelMin, min > level + Self.levelReach { continue }
                addInstance(source)
            case Self.craftedKind:
                for item in source.items ?? [] where crafted[item] == nil {
                    crafted[item] = ItemPlace(
                        kind: .profession, sourceKind: source.kind, name: source.name, profession: source.profession,
                        skill: source.requiredSkill?[String(item)]
                    )
                }
            case Self.repKind:
                for item in source.items ?? [] {
                    reputations[item, default: []].append(ItemPlace(
                        kind: .reputation, sourceKind: source.kind, name: source.name, standing: source.standing?.lowercased()
                    ))
                }
            default:
                continue
            }
        }
    }

    private mutating func addInstance(_ source: LootSource) {
        var found: [(item: Int, boss: String?, chance: Double?)] = []
        for boss in source.bosses ?? [] {
            let items = boss.items + (boss.itemChances?.keys.compactMap { Int($0) } ?? [])
            found += items.map { ($0, boss.name.isEmpty ? nil : boss.name, boss.itemChances?[String($0)]) }
        }
        found += ((source.items ?? []) + (source.trash ?? [])).map { ($0, nil, source.itemChances?[String($0)]) }
        for entry in found {
            let place = ItemPlace(
                kind: .instance, sourceKind: source.kind, name: source.name, boss: entry.boss, chance: entry.chance
            )
            let current = instances[source.kind]?[entry.item]
            if current == nil || (place.chance ?? 0) > (current?.chance ?? 0) {
                instances[source.kind, default: [:]][entry.item] = place
            }
        }
    }

    /// Where `item` comes from under the list's kind; a reputation the list names wins (an item
    /// both PvP factions sell has a source per faction).
    func place(of item: Int, kind: String?, named: String? = nil) -> ItemPlace? {
        switch kind {
        case Self.dungeonKind, Self.raidKind: instances[kind ?? ""]?[item]
        case Self.craftedKind: crafted[item]
        case Self.repKind:
            reputations[item]?.first { $0.name == named } ?? reputations[item]?.first
        default: nil
        }
    }
}

/// Where to go tonight (flows spec F5): every upgrade on the list grouped by where it comes from.
/// An upgrade is a slot's pick when the slot gains, or a runner-up that still gains over what is
/// worn (the slot's gain plus its delta). Computed from the cached list and loot files.
struct WhereToGo: Equatable, Sendable {
    struct Entry: Equatable, Sendable {
        let itemID: Int
        let name: String
        let slot: String
        let gain: Double
        let place: ItemPlace
        /// The boss's chance, else the list's own.
        let chance: Double?
    }

    struct Group: Equatable, Sendable, Identifiable {
        let place: ItemPlace
        /// The best upgrade per slot, the largest gain first.
        let entries: [Entry]
        /// A reputation's highest standing among its entries.
        let standing: String?

        var id: String { "\(place.kind.rawValue):\(place.name)" }
        var gain: Double { entries.reduce(0) { $0 + $1.gain } }
    }

    /// By kind (instance, profession, reputation), then the largest total gain.
    let groups: [Group]

    init(slots: [BisSlot], gains: [String: Double], places: ItemPlaces) {
        let entries = slots.flatMap { Self.entries($0, gains: gains, places: places) }
        let grouped = Dictionary(grouping: entries) { "\($0.place.kind.rawValue):\($0.place.name)" }
        groups = grouped.values.compactMap(Self.group).sorted {
            $0.place.kind != $1.place.kind ? $0.place.kind < $1.place.kind : $0.gain > $1.gain
        }
    }

    /// Tonight: the best place of each kind.
    var tonight: [Group] {
        ItemPlace.Kind.allCases.compactMap { kind in groups.first { $0.place.kind == kind } }
    }

    /// The upgrade with the largest gain from a source kind (`dungeon`: Home's Best run).
    func best(sourceKind: String) -> Entry? {
        groups.flatMap(\.entries).filter { $0.place.sourceKind == sourceKind }.max { $0.gain < $1.gain }
    }

    private static func entries(_ slot: BisSlot, gains: [String: Double], places: ItemPlaces) -> [Entry] {
        guard let slotGain = gains[slot.slot], slotGain > 0 else { return [] }
        var out: [Entry] = []
        if let id = slot.itemId, let name = slot.itemName,
           let place = places.place(of: id, kind: slot.sourceKind, named: slot.source) {
            out.append(Entry(itemID: id, name: name, slot: slot.slot, gain: slotGain, place: place, chance: place.chance ?? slot.dropChance))
        }
        for alt in slot.alternatives ?? [] where slotGain + alt.dpsDelta > 0 {
            guard let place = places.place(of: alt.itemId, kind: alt.sourceKind, named: alt.source) else { continue }
            out.append(Entry(
                itemID: alt.itemId, name: alt.itemName, slot: slot.slot, gain: slotGain + alt.dpsDelta, place: place,
                chance: place.chance ?? alt.dropChance
            ))
        }
        return out
    }

    private static func group(_ entries: [Entry]) -> Group? {
        let bestPerSlot = Dictionary(grouping: entries, by: \.slot).values.compactMap { $0.max { $0.gain < $1.gain } }
        let sorted = bestPerSlot.sorted { $0.gain != $1.gain ? $0.gain > $1.gain : $0.name < $1.name }
        guard let first = sorted.first else { return nil }
        let standing = sorted.compactMap(\.place.standing).max { standingRank($0) < standingRank($1) }
        return Group(place: first.place, entries: sorted, standing: standing)
    }

    /// The reputation ladder, lowest first.
    private static let standings = ["hated", "hostile", "unfriendly", "neutral", "friendly", "honored", "revered", "exalted"]

    private static func standingRank(_ standing: String) -> Int {
        standings.firstIndex(of: standing) ?? -1
    }
}
