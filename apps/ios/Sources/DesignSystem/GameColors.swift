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
