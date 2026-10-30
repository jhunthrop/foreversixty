import DesignTokens
import SwiftUI

/// The game's colours and the spec's type roles (design/specs/2026-10-30-mobile-apps.md 4.1, 4.2),
/// over the generated tokens. Added by lane F4 beside `Theme`; every feature reads them here.
public extension Theme.Palette {
    static let lineSoft = FSColor.lineSoft
    static let lineWarm = FSColor.lineWarm
    static let lineCard = FSColor.lineCard
    static let cardTop = FSColor.cardTop
    static let nav = FSColor.nav
    static let goldHover = FSColor.goldHover
    static let kill = FSColor.kill
    static let wipe = FSColor.wipe
    /// The offline strip's ground (spec 5.0).
    static let strip = Color(red: 0x11 / 255, green: 0x15 / 255, blue: 0x1F / 255)
    /// Sheets: `--color-raised` mixed 4% white (spec 4.1).
    static let sheet = Color(red: 0x10 / 255, green: 0x14 / 255, blue: 0x1D / 255)
    /// The viewer's own roster rows and a pressed row: gold at 5% and 10%.
    static let rowWash = FSColor.gold.opacity(0.05)
    static let pressed = FSColor.gold.opacity(0.10)
    static let horde = FSColor.horde
    static let alliance = FSColor.alliance

    /// A class's colour (`--color-class-*`), muted text for a class the tokens do not name.
    static func classColor(_ slug: String) -> Color {
        switch slug.lowercased() {
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

    /// An item's quality border (0 poor to 5 legendary).
    static func qualityBorder(_ quality: Int) -> Color {
        switch quality {
        case 0: FSColor.rarityPoor
        case 1: FSColor.rarityCommon
        case 2: FSColor.rarityUncommon
        case 3: FSColor.rarityRare
        case 4: FSColor.rarityEpic
        default: FSColor.rarityLegendary
        }
    }

    /// An item name's quality text colour (the AA variants for rare and epic).
    static func qualityText(_ quality: Int) -> Color {
        switch quality {
        case 3: FSColor.rarityRareText
        case 4: FSColor.rarityEpicText
        default: qualityBorder(quality)
        }
    }

    /// A parse percentile's colour (`--color-parse-*`).
    static func parse(_ percentile: Int) -> Color {
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

/// The spec's type roles (4.2), each scaling with Dynamic Type relative to its text style.
public extension Theme.Typography {
    static let pageTitle = Font.custom("BarlowCondensed-ExtraBold", size: FSTextSize.h1Phone, relativeTo: .largeTitle)
    static let h2 = Font.custom("BarlowCondensed-ExtraBold", size: FSTextSize.h2, relativeTo: .title)
    static let stat = Font.custom("BarlowCondensed-ExtraBold", size: FSTextSize.stat, relativeTo: .title)
    static let sec = Font.custom("BarlowCondensed-ExtraBold", size: FSTextSize.cardTitle, relativeTo: .headline)
    static let cardHead = Font.custom("BarlowCondensed-ExtraBold", size: 22, relativeTo: .title2)
    static let name = Font.custom("Barlow-Bold", size: 17, relativeTo: .headline)
    static let bodyText = Font.custom("Barlow-Regular", size: FSTextSize.body, relativeTo: .subheadline)
    static let bodyStrong = Font.custom("Barlow-SemiBold", size: FSTextSize.body, relativeTo: .subheadline)
    static let metaText = Font.custom("Barlow-Regular", size: FSTextSize.meta, relativeTo: .footnote)
    static let micro = Font.custom("Barlow-Bold", size: FSTextSize.micro, relativeTo: .caption2)
    static let segment = Font.custom("Barlow-SemiBold", size: 14, relativeTo: .subheadline)
    static let button = Font.custom("Barlow-Bold", size: 13, relativeTo: .subheadline)
    /// Times, counts and countdowns (JetBrains Mono is not bundled yet; the system mono stands in).
    static let mono = Font.system(.subheadline, design: .monospaced)
    static let monoMeta = Font.system(.footnote, design: .monospaced)
    static let monoDay = Font.system(.title3, design: .monospaced).weight(.semibold)
}

public extension Theme.Radius {
    static let panel = FSRadius.panel
    static let control = FSRadius.control
    static let pill = FSRadius.pill
}
