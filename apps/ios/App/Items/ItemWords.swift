import Foundation

/// The words every item surface shares: slot labels and quality names (spec §4.3). One place,
/// so the BiS list, the tooltip sheet and Home name an item the same way.
enum ItemWords {
    /// The paperdoll label of a BiS or item slot key. Two rings read `Ring 1` and
    /// `Ring 2`, two trinkets `Trinket 1` and `Trinket 2`, so the rows can be told apart.
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
        case "finger": Copy.itemsSlotFinger
        case "finger1": Copy.itemsSlotFinger1
        case "finger2": Copy.itemsSlotFinger2
        case "trinket": Copy.itemsSlotTrinket
        case "trinket1": Copy.itemsSlotTrinket1
        case "trinket2": Copy.itemsSlotTrinket2
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
