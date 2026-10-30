import Foundation

/// The words every item surface shares: slot labels and quality names (spec §4.3). One place,
/// so the BiS list, the tooltip sheet and Home name an item the same way.
enum ItemWords {
    /// The paperdoll label of a BiS or item slot key. Both rings read `Ring` and both trinkets
    /// `Trinket`, as the site's BiS list does (`slot-display-labels.ts`); the rest are F2's `items.slot.*`.
    static func slot(_ key: String) -> String {
        switch key {
        case "head": Copy.itemsSlotHead
        case "neck": Copy.itemsSlotNeck
        case "shoulder": Copy.itemsSlotShoulder
        case "back": Copy.itemsSlotBack
        case "chest": Copy.itemsSlotChest
        case "wrist": Copy.itemsSlotWrist
        case "hands": Copy.itemsSlotHands
        case "waist": Copy.itemsSlotWaist
        case "legs": Copy.itemsSlotLegs
        case "feet": Copy.itemsSlotFeet
        case "finger", "finger1", "finger2": Copy.bisSlotRing
        case "trinket", "trinket1", "trinket2": Copy.itemsSlotTrinket
        case "main_hand": Copy.itemsSlotMainHand
        case "off_hand": Copy.itemsSlotOffHand
        case "ranged": Copy.itemsSlotRanged
        default: key
        }
    }

    /// `quality` is 0 (poor) to 5 (legendary).
    static func quality(_ quality: Int) -> String {
        switch quality {
        case 0: Copy.itemsQualityPoor
        case 1: Copy.itemsQualityCommon
        case 2: Copy.itemsQualityUncommon
        case 3: Copy.itemsQualityRare
        case 4: Copy.itemsQualityEpic
        default: Copy.itemsQualityLegendary
        }
    }
}
