import DesignTokens
import SwiftUI
import UIKit

/// The app's semantic roles over the generated tokens (`FSColor`, `FSFont`, `FSType`, `FSSpace`,
/// `FSRadius`; Shared/DesignTokens.swift, from design/tokens/tokens.json), in round 5's look
/// (design/specs/2026-11-09-directions.md, Mobile): the painted brown ground under every screen,
/// the painted blue band behind the bars, brown-black panels with the warm ring, the teal primary
/// button. Views use `Theme`, never raw hex or sizes, so a token regeneration reaches every screen
/// through this file.
public enum Theme {
    public enum Palette {
        /// The deepest brown-black well: a screen that paints nothing of its own. The painted
        /// ground is `ScreenGround` (`.screenGround()`), taken screen by screen as each tab lane
        /// restyles it.
        public static let background = FSColor.bg
        public static let well = FSColor.bg
        /// The earth under the ground tile.
        public static let ground = FSColor.ground
        /// A panel's brown-black (their card and footer).
        public static let raised = FSColor.raised
        public static let line = FSColor.line
        public static let text = FSColor.text
        public static let strong = FSColor.strong
        public static let muted = FSColor.muted
        /// Muted ink straight on the ground: white at .5 fails there (tokens `mutedGround`).
        public static let mutedOnGround = FSColor.mutedGround
        /// Titles: the cream display ink.
        public static let title = FSColor.cream
        /// Links and the accent: their gold.
        public static let accent = FSColor.gold
        /// The bars: the band under the navigation bar, the tab bar's ground.
        public static let band = FSColor.band
        public static let bar = FSColor.navBar
        public static let tabSelected = FSColor.tabOn
        public static let tabUnselected = FSColor.tabOff
    }

    /// Open Sans for body, labels and figures; Marcellus for titles (apps/shared/fonts, UIAppFonts).
    public enum Typography {
        public static let title = Font.custom(FSFont.title, size: FSTextSize.h2)
        public static let cardTitle = Font.custom(FSFont.title, size: FSTextSize.cardTitle)
        public static let body = Font.custom(FSFont.body, size: FSType.body.size)
        public static let meta = Font.custom(FSFont.bodySemibold, size: FSTextSize.meta)
        /// The navigation bar's inline title (Open Sans 600 17, the mock's `.navb .t`).
        public static var barTitle: UIFont? { UIFont(name: FSFont.bodySemibold, size: 17) }
        /// The large title (Marcellus, the mock's `.bigt h1`).
        public static var largeTitle: UIFont? { UIFont(name: FSFont.title, size: FSType.title.size) }
        /// A tab's label (11 pt 600, the mock's `.tab`).
        public static var tab: UIFont? { UIFont(name: FSFont.bodySemibold, size: 11) }
    }

    public enum Spacing {
        public static let xs = FSSpace.space1
        public static let sm = FSSpace.space2
        public static let md = FSSpace.space4
        public static let lg = FSSpace.space6
    }

    /// Round 5's one radius scale: panels, controls, item icons.
    public enum Radius {
        public static let card = FSRadius.plate
        public static let panel = FSRadius.plate
        public static let control = FSRadius.field
        public static let icon = FSRadius.pill
    }

    /// The tab bar (spec "Tab-bar icons"): the owner's icons at 24 pt.
    public enum TabBar {
        public static let iconSize: CGFloat = 24
    }
}

/// The scaffold's card surface: a plain round-5 panel with the content's padding.
public struct CardSurface: ViewModifier {
    public init() {}

    public func body(content: Content) -> some View {
        content
            .padding(Theme.Spacing.md)
            .frame(maxWidth: .infinity, alignment: .leading)
            .panel()
    }
}

public extension View {
    func cardSurface() -> some View { modifier(CardSurface()) }
}
