import Foundation
import Models

/// Home's Best run line (flows spec F5): the dungeon upgrade with the largest gain over what the
/// character wears, its instance and its chance. It is Where to go's best dungeon entry
/// (`WhereToGo`), so Home and Gear read one index of the cached list and loot files.
struct BestRun: Equatable, Sendable {
    let place: String
    let item: String
    let gain: Double
    /// Percent, when the loot file names one.
    let chance: Double?
    /// The drop's item id, for its icon on Home's Best run row.
    var itemID: Int?

    /// The slots' modelled gains from the home answer, by slot.
    static func gains(_ upgrades: [HomeUpgrade]) -> [String: Double] {
        Dictionary(upgrades.compactMap { upgrade in
            guard upgrade.modelled, let gain = upgrade.gain, gain > 0 else { return nil }
            return (upgrade.slot, gain)
        }, uniquingKeysWith: max)
    }

    static func best(slots: [BisSlot], gains: [String: Double], loot: LootFile, level: Int?) -> BestRun? {
        let places = WhereToGo(slots: slots, gains: gains, places: ItemPlaces(loot: loot, level: level))
        guard let entry = places.best(sourceKind: ItemPlaces.dungeonKind) else { return nil }
        return BestRun(place: entry.place.name, item: entry.name, gain: entry.gain, chance: entry.chance, itemID: entry.itemID)
    }
}
