import DesignTokens
import SwiftUI

/// The game's colours over the generated tokens: item quality (spec §4.1, §4.3), class and
/// faction. Views pass the data's own key (a quality 0 to 5, a class slug) and never a hex.
public enum GamePalette {
    /// The quality border: the game's own colour. `quality` is 0 (poor) to 5 (legendary).
    public static func qualityBorder(_ quality: Int) -> Color {
        switch quality {
        case 0: FSColor.rarityPoor
        case 1: FSColor.rarityCommon
        case 2: FSColor.rarityUncommon
        case 3: FSColor.rarityRare
        case 4: FSColor.rarityEpic
        default: FSColor.rarityLegendary
        }
    }

    /// An item name: the site's AA variants for rare and epic on the dark ground.
    public static func qualityText(_ quality: Int) -> Color {
        switch quality {
        case 3: FSColor.rarityRareText
        case 4: FSColor.rarityEpicText
        default: qualityBorder(quality)
        }
    }

    /// A class's text colour by slug (`shaman` is the site's AA blue); muted when unknown.
    public static func classColor(_ slug: String) -> Color {
        switch slug {
        case "warrior": FSColor.classWarrior
        case "paladin": FSColor.classPaladin
        case "hunter": FSColor.classHunter
        case "rogue": FSColor.classRogue
        case "priest": FSColor.classPriest
        case "shaman": FSColor.classShaman
        case "mage": FSColor.classMage
        case "warlock": FSColor.classWarlock
        case "druid": FSColor.classDruid
        default: FSColor.muted
        }
    }

    public static func faction(_ slug: String) -> Color {
        slug == "horde" ? FSColor.horde : FSColor.alliance
    }
}
