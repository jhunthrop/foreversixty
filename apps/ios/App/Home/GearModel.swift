import Foundation
import Models

/// One of the 17 rows of Your gear (spec F2.7): what is worn and the band's pick.
struct GearRow: Equatable, Identifiable {
    let slot: String
    let slotName: String
    /// Nil for an empty slot.
    let worn: HomeItemLook?
    /// The worn item's name, or the server's words for one our data does not know.
    let wornName: String
    let wornLevel: String?
    /// The pick when it differs from what is worn.
    let pick: HomeItemLook?
    let isBestInSlot: Bool

    var id: String { slot }
}

/// Builds the 17 rows from the home answer: an upgrade row names worn and pick, an
/// already-best row names the one item, and any other slot shows the export's worn item looked
/// up in the build's items file. Nothing is ranked or compared here.
enum GearModel {
    static func rows(home: CharacterHome, gear: [String: Int], items: [Int: Item]) -> [GearRow] {
        let upgrades = Dictionary(home.upgrades.map { ($0.slot, $0) }, uniquingKeysWith: { first, _ in first })
        let best = Dictionary(home.alreadyBest.map { ($0.slot, $0) }, uniquingKeysWith: { first, _ in first })
        return ItemVocabulary.frameSlots.map { slot in
            let name = ItemVocabulary.slotName(slot) ?? slot
            if let upgrade = upgrades[slot] { return upgradeRow(upgrade, slotName: name) }
            if let item = best[slot]?.item { return bestRow(item, slot: slot, slotName: name) }
            return exportRow(gear[slot], slot: slot, slotName: name, items: items)
        }
    }

    private static func upgradeRow(_ upgrade: HomeUpgrade, slotName: String) -> GearRow {
        let pick = HomeItemLook(id: upgrade.pick.id, name: upgrade.pick.name, quality: upgrade.pick.quality,
                                icon: upgrade.pick.icon, itemLevel: upgrade.pick.itemLevel)
        guard let worn = upgrade.worn else {
            return GearRow(slot: upgrade.slot, slotName: slotName, worn: nil, wornName: Copy.gearEmptySlot,
                           wornLevel: nil, pick: pick, isBestInSlot: false)
        }
        let look = HomeItemLook(id: worn.id, name: worn.name, quality: worn.quality, icon: worn.icon, itemLevel: worn.itemLevel)
        return GearRow(slot: upgrade.slot, slotName: slotName, worn: look, wornName: worn.known ? (worn.name ?? worn.label) : worn.label,
                       wornLevel: worn.itemLevel.map { Copy.itemsIlvl(n: $0) }, pick: pick, isBestInSlot: false)
    }

    private static func bestRow(_ item: HomeItem, slot: String, slotName: String) -> GearRow {
        let look = HomeItemLook(id: item.id, name: item.name, quality: item.quality, icon: item.icon, itemLevel: item.itemLevel)
        return GearRow(slot: slot, slotName: slotName, worn: look, wornName: item.name ?? Copy.itemsNotInData,
                       wornLevel: item.itemLevel.map { Copy.itemsIlvl(n: $0) }, pick: nil, isBestInSlot: true)
    }

    private static func exportRow(_ id: Int?, slot: String, slotName: String, items: [Int: Item]) -> GearRow {
        guard let id else {
            return GearRow(slot: slot, slotName: slotName, worn: nil, wornName: Copy.gearEmptySlot,
                           wornLevel: nil, pick: nil, isBestInSlot: false)
        }
        guard let item = items[id] else {
            return GearRow(slot: slot, slotName: slotName, worn: HomeItemLook(id: id, name: nil, quality: nil, icon: nil, itemLevel: nil),
                           wornName: Copy.itemsNotInData, wornLevel: nil, pick: nil, isBestInSlot: false)
        }
        let look = HomeItemLook(id: id, name: item.name, quality: item.quality, icon: item.icon, itemLevel: item.itemLevel)
        return GearRow(slot: slot, slotName: slotName, worn: look, wornName: item.name,
                       wornLevel: Copy.itemsIlvl(n: item.itemLevel), pick: nil, isBestInSlot: false)
    }
}
