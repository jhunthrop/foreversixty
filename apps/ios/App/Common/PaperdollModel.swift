import Foundation
import Models

/// One socket of the Paperdoll (design/specs/2026-10-31-paperdoll.md, spec §4 Paperdoll).
struct PaperdollSlot: Equatable, Identifiable, Sendable {
    enum State: Equatable, Sendable {
        /// A worn item; quality and icon nil when the items file does not have it but the API named it.
        case item(HomeItemLook)
        /// A worn item our data does not know: its id only.
        case unknown(Int)
        case empty
        /// The empty off hand under a two-handed main hand, with its words.
        case twoHandOff(String)
    }

    let slot: String
    let state: State
    /// The gold dot: the character's BiS list has a better item for this slot.
    let better: Bool

    var id: String { slot }
    var label: String { ItemVocabulary.slotLabel(slot) }

    /// The worn item's look, for the tooltip.
    var look: HomeItemLook? {
        if case let .item(look) = state { return look }
        return nil
    }
}

/// The frame: two mirrored columns (the client's left and right slots) and the weapons row.
struct PaperdollFrame: Equatable, Sendable {
    let left: [PaperdollSlot]
    let right: [PaperdollSlot]
    let weapons: [PaperdollSlot]

    /// All 17 in the character frame's order: the largest-text single column.
    var all: [PaperdollSlot] {
        let bySlot = Dictionary((left + right + weapons).map { ($0.slot, $0) }, uniquingKeysWith: { first, _ in first })
        return ItemVocabulary.frameSlots.compactMap { bySlot[$0] }
    }
}

/// The Paperdoll's pure model: where each slot sits, what it holds, and its gold dot. One model
/// for the character page (F11.5) and Home's Your gear (F2.7).
enum PaperdollModel {
    static let leftSlots = ["head", "neck", "shoulder", "back", "chest", "wrist"]
    static let rightSlots = ["hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2"]
    static let weaponSlots = ["main_hand", "off_hand", "ranged"]

    private static let mainHand = "main_hand"
    private static let offHand = "off_hand"
    /// The slot pairs whose items are matched either way round (bis `wearsListedBest`).
    private static let twins = ["finger1": "finger2", "finger2": "finger1", "trinket1": "trinket2", "trinket2": "trinket1"]

    /// The frame from any slots; a slot the data lacks is drawn empty.
    static func frame(_ slots: [PaperdollSlot]) -> PaperdollFrame {
        let bySlot = Dictionary(slots.map { ($0.slot, $0) }, uniquingKeysWith: { first, _ in first })
        func pick(_ keys: [String]) -> [PaperdollSlot] {
            keys.map { bySlot[$0] ?? PaperdollSlot(slot: $0, state: .empty, better: false) }
        }
        return PaperdollFrame(left: pick(leftSlots), right: pick(rightSlots), weapons: pick(weaponSlots))
    }

    /// The character page's sockets from `CharacterProfile.gear`, looked up in the class's items file.
    static func character(gear: [CharacterGearSlot], items: [Int: Item]) -> [PaperdollSlot] {
        let bySlot = Dictionary(gear.map { ($0.slot, $0) }, uniquingKeysWith: { first, _ in first })
        let wornBySlot = bySlot.compactMapValues { $0.worn?.id }
        let twoHander = wornBySlot[mainHand].flatMap { items[$0] }.flatMap { $0.twoHand ? $0.name : nil }
        return ItemVocabulary.frameSlots.map { slot in
            guard let entry = bySlot[slot], let worn = entry.worn else {
                if slot == offHand, let twoHander {
                    return PaperdollSlot(slot: slot, state: .twoHandOff(Copy.itemsOffHandTwoHand(item: twoHander)), better: false)
                }
                return PaperdollSlot(slot: slot, state: .empty, better: bySlot[slot]?.pick != nil)
            }
            return PaperdollSlot(slot: slot, state: state(worn, items: items), better: isBetter(entry, wornBySlot: wornBySlot))
        }
    }

    private static func state(_ worn: CharacterGearSlotWorn, items: [Int: Item]) -> PaperdollSlot.State {
        if let item = items[worn.id] {
            return .item(HomeItemLook(id: item.id, name: item.name, quality: item.quality, icon: item.icon, itemLevel: item.itemLevel))
        }
        guard let name = worn.name else { return .unknown(worn.id) }
        return .item(HomeItemLook(id: worn.id, name: name, quality: nil, icon: nil, itemLevel: worn.itemLevel))
    }

    /// The gold dot (paperdoll spec §3.4): a pick for a worn item our data does not know, or a
    /// different pick with a gain. Never when the pick is worn, in this slot or its twin.
    static func isBetter(_ entry: CharacterGearSlot, wornBySlot: [String: Int]) -> Bool {
        guard let pick = entry.pick else { return false }
        guard let worn = entry.worn else { return true }
        if worn.name == nil { return true }
        if pick.id == worn.id { return false }
        if let twin = twins[entry.slot], wornBySlot[twin] == pick.id { return false }
        return (entry.gap ?? 0) > 0
    }

    /// Home's Your gear (F2.7): the same 17 rows `GearModel` builds, a dot on every upgrade.
    static func home(_ rows: [GearRow]) -> [PaperdollSlot] {
        rows.map { row in
            let state: PaperdollSlot.State = if let worn = row.worn {
                worn.name == nil && worn.icon == nil ? .unknown(worn.id) : .item(worn)
            } else if row.blockedByTwoHand {
                .twoHandOff(row.wornName)
            } else {
                .empty
            }
            return PaperdollSlot(slot: row.slot, state: state, better: row.pick != nil)
        }
    }

    /// The socket's VoiceOver name (spec §6 Paperdoll socket).
    static func a11y(_ slot: PaperdollSlot) -> String {
        let base: String = switch slot.state {
        case let .item(look):
            Copy.paperdollA11ySocket(
                slot: slot.label, item: look.name ?? Copy.itemsNotInData, quality: ItemQuality(value: look.quality).name,
                level: look.itemLevel.map(String.init) ?? ""
            )
        case .unknown:
            "\(slot.label), \(Copy.itemsNotInData)."
        case .empty:
            Copy.paperdollA11yEmpty(slot: slot.label)
        case let .twoHandOff(words):
            "\(slot.label), \(words)"
        }
        return slot.better ? "\(base) \(Copy.paperdollA11yBetter)" : base
    }
}
