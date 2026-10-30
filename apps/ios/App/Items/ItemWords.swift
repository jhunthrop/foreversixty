import Foundation

/// The words every item surface shares: slot labels and quality names (spec §4.3). One place,
/// so the BiS list, the tooltip sheet and Home name an item the same way.
enum ItemWords {
    /// The paperdoll label of a BiS or item slot key. Both rings read `Ring` and both trinkets
    /// `Trinket`, as the site's list does (`slot-display-labels.ts`).
    static func slot(_ key: String) -> String {
        switch key {
        case "head": Copy.bisSlotHead
        case "neck": Copy.bisSlotNeck
        case "shoulder": Copy.bisSlotShoulder
        case "back": Copy.bisSlotBack
        case "chest": Copy.bisSlotChest
        case "wrist": Copy.bisSlotWrist
        case "hands": Copy.bisSlotHands
        case "waist": Copy.bisSlotWaist
        case "legs": Copy.bisSlotLegs
        case "feet": Copy.bisSlotFeet
        case "finger", "finger1", "finger2": Copy.bisSlotFinger
        case "trinket", "trinket1", "trinket2": Copy.bisSlotTrinket
        case "main_hand": Copy.bisSlotMainHand
        case "off_hand": Copy.bisSlotOffHand
        case "ranged": Copy.bisSlotRanged
        default: key
        }
    }

    /// `quality` is 0 (poor) to 5 (legendary).
    static func quality(_ quality: Int) -> String {
        switch quality {
        case 0: Copy.bisQualityPoor
        case 1: Copy.bisQualityCommon
        case 2: Copy.bisQualityUncommon
        case 3: Copy.bisQualityRare
        case 4: Copy.bisQualityEpic
        default: Copy.bisQualityLegendary
        }
    }
}
