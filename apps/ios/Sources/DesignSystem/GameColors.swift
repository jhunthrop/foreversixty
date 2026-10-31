import SwiftUI

/// The class colours by the API's class word (`warrior`, `shaman`, ...).
public enum ClassColor {
    public static func of(_ className: String?) -> Color {
        switch className?.lowercased() {
        case "warrior": FSColor.classWarrior
        case "paladin": FSColor.classPaladin
        case "hunter": FSColor.classHunter
        case "rogue": FSColor.classRogue
        case "priest": FSColor.classPriest
        case "shaman": FSColor.classShaman
        case "mage": FSColor.classMage
        case "warlock": FSColor.classWarlock
        case "druid": FSColor.classDruid
        default: FSColor.text
        }
    }
}

/// The game's colours beyond the class colours (item quality lives with the items, `ItemQuality`),
/// plus the spec-only values with no site token yet. Views pass the data's own key, never a hex.
public enum GameColors {
    public static func faction(_ slug: String) -> Color {
        slug == "horde" ? FSColor.horde : FSColor.alliance
    }

    // Spec-only values (spec §4.1, §5.0, F2.6, F3.2); lane T folds them into tokens.css when the
    // site gains them.
    /// The offline strip's ground.
    public static let offlineGround = Color(tokenHex: 0x11151F)
    /// A picker sheet's ground.
    public static let sheet = Color(tokenHex: 0x10141D)
    /// The tooltip's inner panel, and the 1 pt line around it and around a spec icon.
    public static let tooltipPanel = Color(tokenHex: 0x0A0D14)
    public static let iconLine = Color(tokenHex: 0x3A4256)
    /// A chip or secondary control's 4.5% white fill, and the pressed 10% gold wash.
    public static let controlFill = Color.white.opacity(0.045)
    public static let pressedWash = FSColor.gold.opacity(0.10)
}
