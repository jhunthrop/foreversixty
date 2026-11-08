import Foundation
import Models

/// Home's Best run line (flows spec F5): the dungeon pick with the largest gain over what the
/// character wears, its instance and its chance, computed from the cached list and loot files
/// (`/data/{build}/bis/{spec}.json`, `loot.json`) as the web's `lib/bis/page-model` groups them.
/// No endpoint: the gains are the home answer's, a runner-up adds its delta to its slot's gain.
struct BestRun: Equatable, Sendable {
    let place: String
    let item: String
    let gain: Double
    /// Percent, when the loot file names one.
    let chance: Double?

    static let dungeonKind = "dungeon"
    /// A dungeon whose lowest level is this far above the character's is not tonight's run.
    static let levelReach = 3

    /// The slots' modelled gains from the home answer, by slot.
    static func gains(_ upgrades: [HomeUpgrade]) -> [String: Double] {
        Dictionary(upgrades.compactMap { upgrade in
            guard upgrade.modelled, let gain = upgrade.gain, gain > 0 else { return nil }
            return (upgrade.slot, gain)
        }, uniquingKeysWith: max)
    }

    static func best(slots: [BisSlot], gains: [String: Double], loot: LootFile, level: Int?) -> BestRun? {
        let places = DungeonIndex(loot: loot, level: level)
        var best: BestRun?
        for slot in slots {
            guard let slotGain = gains[slot.slot] else { continue }
            for candidate in candidates(slot, slotGain: slotGain) {
                guard let place = places.place(of: candidate.id) else { continue }
                let run = BestRun(
                    place: place.name, item: candidate.name, gain: candidate.gain,
                    chance: place.chance ?? candidate.chance
                )
                if run.gain > (best?.gain ?? 0) { best = run }
            }
        }
        return best
    }

    private struct Candidate {
        let id: Int
        let name: String
        let gain: Double
        let chance: Double?
    }

    /// The pick when it drops in a dungeon, and every dungeon runner-up that still gains.
    private static func candidates(_ slot: BisSlot, slotGain: Double) -> [Candidate] {
        var out: [Candidate] = []
        if slot.sourceKind == dungeonKind, let id = slot.itemId, let name = slot.itemName {
            out.append(Candidate(id: id, name: name, gain: slotGain, chance: slot.dropChance))
        }
        for alt in slot.alternatives ?? [] where alt.sourceKind == dungeonKind && slotGain + alt.dpsDelta > 0 {
            out.append(Candidate(id: alt.itemId, name: alt.itemName, gain: slotGain + alt.dpsDelta, chance: alt.dropChance))
        }
        return out
    }
}

/// Where each dungeon item drops: the instance, and the boss's chance for it.
private struct DungeonIndex {
    struct Place {
        let name: String
        let chance: Double?
    }

    private var places: [Int: Place] = [:]

    init(loot: LootFile, level: Int?) {
        for source in loot.sources where source.kind == BestRun.dungeonKind {
            if let level, let min = source.levelMin, min > level + BestRun.levelReach { continue }
            for boss in source.bosses ?? [] {
                for item in boss.items + (boss.itemChances?.keys.compactMap { Int($0) } ?? []) where places[item] == nil {
                    places[item] = Place(name: source.name, chance: boss.itemChances?[String(item)])
                }
            }
            for item in (source.items ?? []) + (source.trash ?? []) where places[item] == nil {
                places[item] = Place(name: source.name, chance: source.itemChances?[String(item)])
            }
        }
    }

    func place(of item: Int) -> Place? { places[item] }
}
