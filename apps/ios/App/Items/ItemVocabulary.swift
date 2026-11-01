import DesignTokens
import SwiftUI

/// The game's words for an item, from the copy catalog: quality, slot, stat and standing names.
/// One table per word family, so the tooltip, Home and the gear list all say the same thing.
enum ItemVocabulary {
    /// The 17 character-frame slots, in the in-game frame's order (spec F2.7).
    static let frameSlots = [
        "head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist", "legs", "feet",
        "finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand", "ranged",
    ]

    static func slotName(_ slot: String) -> String? {
        switch slot {
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
        case "finger1": Copy.itemsSlotFinger1
        case "finger2": Copy.itemsSlotFinger2
        case "trinket1": Copy.itemsSlotTrinket1
        case "trinket2": Copy.itemsSlotTrinket2
        case "main_hand": Copy.itemsSlotMainHand
        case "off_hand": Copy.itemsSlotOffHand
        case "ranged": Copy.itemsSlotRanged
        case "finger": Copy.itemsSlotFinger
        case "trinket": Copy.itemsSlotTrinket
        default: nil
        }
    }

    /// A slot's label: two rings read `Finger 1` and `Finger 2`, two trinkets `Trinket 1` and
    /// `Trinket 2`, so the BiS rows can be told apart; an unknown key prints as published.
    static func slotLabel(_ slot: String) -> String { slotName(slot) ?? slot }

    /// The stats a tooltip lists, in the site's order (`STAT_KEYS`, armor excluded: it has its own line).
    static let statOrder: [(key: String, name: String)] = [
        ("strength", Copy.itemsStatStrength), ("agility", Copy.itemsStatAgility),
        ("stamina", Copy.itemsStatStamina), ("intellect", Copy.itemsStatIntellect),
        ("spirit", Copy.itemsStatSpirit), ("crit", Copy.itemsStatCrit), ("hit", Copy.itemsStatHit),
        ("spell_power", Copy.itemsStatSpellPower), ("healing", Copy.itemsStatHealing),
        ("attack_power", Copy.itemsStatAttackPower), ("ranged_attack_power", Copy.itemsStatRangedAttackPower),
        ("defense", Copy.itemsStatDefense), ("dodge", Copy.itemsStatDodge), ("parry", Copy.itemsStatParry),
        ("block", Copy.itemsStatBlock), ("block_value", Copy.itemsStatBlockValue), ("mp5", Copy.itemsStatMp5),
        ("spell_penetration", Copy.itemsStatSpellPenetration), ("fire_power", Copy.itemsStatFirePower),
        ("fire_res", Copy.itemsStatFireRes), ("frost_res", Copy.itemsStatFrostRes),
        ("nature_res", Copy.itemsStatNatureRes), ("shadow_res", Copy.itemsStatShadowRes),
        ("arcane_res", Copy.itemsStatArcaneRes),
    ]

    /// The stats the client prints as a green `Equip:` sentence rather than `+n Name`.
    static func equipLine(_ key: String, value: Int) -> String? {
        switch key {
        case "spell_power": Copy.itemsEquipSpellPower(n: value)
        case "healing": Copy.itemsEquipHealing(n: value)
        case "attack_power": Copy.itemsEquipAttackPower(n: value)
        case "ranged_attack_power": Copy.itemsEquipRangedAttackPower(n: value)
        case "mp5": Copy.itemsEquipMp5(n: value)
        case "spell_penetration": Copy.itemsEquipSpellPenetration(n: value)
        case "fire_power": Copy.itemsEquipFirePower(n: value)
        case "block_value": Copy.itemsEquipBlockValue(n: value)
        case "defense": Copy.itemsEquipDefense(n: value)
        default: nil
        }
    }

    /// A rating-family stat the data states as a literal percent (`percent_stats`): the client's
    /// `Equip:` sentence (the site's `PERCENT_STAT_LINES`); nil prints it as `+n Name`.
    static func percentLine(_ key: String, value: Int) -> String? {
        let percent = "\(value)%"
        return switch key {
        case "crit": Copy.itemsEquipCrit(percent: percent)
        case "hit": Copy.itemsEquipHit(percent: percent)
        case "dodge": Copy.itemsEquipDodge(percent: percent)
        case "parry": Copy.itemsEquipParry(percent: percent)
        case "block": Copy.itemsEquipBlock(percent: percent)
        default: nil
        }
    }

    static func standingName(_ standing: String) -> String? {
        switch standing {
        case "hated": Copy.itemsStandingHated
        case "hostile": Copy.itemsStandingHostile
        case "unfriendly": Copy.itemsStandingUnfriendly
        case "neutral": Copy.itemsStandingNeutral
        case "friendly": Copy.itemsStandingFriendly
        case "honored": Copy.itemsStandingHonored
        case "revered": Copy.itemsStandingRevered
        case "exalted": Copy.itemsStandingExalted
        default: nil
        }
    }

    static func bindsLine(_ bonding: String?) -> String? {
        switch bonding {
        case "pickup": Copy.tooltipBindPickup
        case "equip": Copy.tooltipBindEquip
        case "use": Copy.tooltipBindUse
        case "quest": Copy.itemsBindQuest
        default: nil
        }
    }

    static func subclassName(_ subclass: String) -> String? {
        switch subclass {
        case "axe": Copy.itemsSubclassAxe
        case "bow": Copy.itemsSubclassBow
        case "gun": Copy.itemsSubclassGun
        case "mace": Copy.itemsSubclassMace
        case "polearm": Copy.itemsSubclassPolearm
        case "sword": Copy.itemsSubclassSword
        case "staff": Copy.itemsSubclassStaff
        case "fist_weapon": Copy.itemsSubclassFistWeapon
        case "dagger": Copy.itemsSubclassDagger
        case "thrown": Copy.itemsSubclassThrown
        case "crossbow": Copy.itemsSubclassCrossbow
        case "wand": Copy.itemsSubclassWand
        case "fishing_pole": Copy.itemsSubclassFishingPole
        default: nil
        }
    }

    static func weaponSlotName(_ slot: String) -> String? {
        switch slot {
        case "one_hand": Copy.itemsWeaponSlotOneHand
        case "two_hand": Copy.itemsWeaponSlotTwoHand
        case "main_hand": Copy.itemsWeaponSlotMainHand
        case "off_hand": Copy.itemsWeaponSlotOffHand
        case "ranged": Copy.itemsWeaponSlotRanged
        case "thrown": Copy.itemsWeaponSlotThrown
        default: nil
        }
    }

    static func armorTypeName(_ armorType: String?) -> String? {
        switch armorType {
        case "cloth": Copy.tooltipArmorTypeCloth
        case "leather": Copy.tooltipArmorTypeLeather
        case "mail": Copy.tooltipArmorTypeMail
        case "plate": Copy.tooltipArmorTypePlate
        case "shield": Copy.tooltipArmorTypeShield
        default: nil
        }
    }
}

/// An item's quality (0 poor to 5 legendary): its name, border colour and text colour.
struct ItemQuality: Equatable {
    let value: Int?

    var name: String {
        switch value {
        case 0: Copy.itemsQualityPoor
        case 1: Copy.itemsQualityCommon
        case 2: Copy.itemsQualityUncommon
        case 3: Copy.itemsQualityRare
        case 4: Copy.itemsQualityEpic
        case 5: Copy.itemsQualityLegendary
        default: ""
        }
    }

    /// The icon border (`--color-rarity-*`).
    var border: Color {
        switch value {
        case 0: FSColor.rarityPoor
        case 1: FSColor.rarityCommon
        case 2: FSColor.rarityUncommon
        case 3: FSColor.rarityRare
        case 4: FSColor.rarityEpic
        case 5: FSColor.rarityLegendary
        default: FSColor.line
        }
    }

    /// The name colour: the AA text variants for rare and epic, the border colour otherwise.
    var text: Color {
        switch value {
        case 3: FSColor.rarityRareText
        case 4: FSColor.rarityEpicText
        case nil: FSColor.muted
        default: border
        }
    }
}

/// Number formats the item surfaces share. Formatting only: every figure comes from the API or
/// the static data.
enum ItemNumbers {
    /// `+11.1` and `−0.8` (a true minus, as the server writes it).
    static func signedTenths(_ value: Double) -> String {
        let magnitude = String(format: "%.1f", abs(value))
        return value < 0 ? "\u{2212}\(magnitude)" : "+\(magnitude)"
    }

}
