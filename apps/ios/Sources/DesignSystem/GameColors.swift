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

    // The surfaces the tokens name by another role, in round 5's palette (lane M0).
    /// The offline strip's ground: the deepest brown-black well.
    public static let offlineGround = FSColor.bg
    /// A picker sheet's ground: the panel brown-black.
    public static let sheet = FSColor.raised
    /// The tooltip's inner panel (the client's navy), and the 1 pt line around it and around a
    /// spec icon (their rule).
    public static let tooltipPanel = FSColor.tooltip
    public static let iconLine = FSColor.rule
    /// A chip or secondary control's 4.5% white fill, and the pressed 10% gold wash.
    public static let controlFill = Color.white.opacity(0.045)
    public static let pressedWash = FSColor.gold.opacity(0.10)
    /// The viewer's own row (a roster's own line): gold at 5%.
    public static let rowWash = FSColor.gold.opacity(0.05)

    /// A parse percentile's colour (`--color-parse-*`).
    public static func parse(_ percentile: Int) -> Color {
        switch percentile {
        case 100...: FSColor.parseGold
        case 99: FSColor.parsePink
        case 95...: FSColor.parseOrange
        case 75...: FSColor.parsePurple
        case 50...: FSColor.parseBlue
        case 25...: FSColor.parseGreen
        default: FSColor.parseGreyText
        }
    }
}
