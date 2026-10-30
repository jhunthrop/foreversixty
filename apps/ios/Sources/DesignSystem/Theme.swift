import DesignTokens
import SwiftUI

/// The app's semantic roles over the generated tokens (`FSColor`, `FSFont`, `FSTextSize`,
/// `FSSpacing`, `FSRadius`; Shared/DesignTokens.swift, from web/src/styles/tokens.css).
/// Views use `Theme`, never raw hex or sizes, so a token regeneration reaches every screen
/// through this file.
public enum Theme {
    public enum Palette {
        public static let background = FSColor.bg
        public static let raised = FSColor.raised
        public static let line = FSColor.line
        public static let text = FSColor.text
        public static let strong = FSColor.strong
        public static let muted = FSColor.muted
        public static let accent = FSColor.gold
        public static let accentHover = FSColor.goldHover
        public static let nav = FSColor.nav
        public static let lineSoft = FSColor.lineSoft
        /// The primary card's 17% gold border.
        public static let lineCard = FSColor.lineCard
        public static let cardTop = FSColor.cardTop
        /// A good state: checked, `Matches`, `You wear it`.
        public static let kill = FSColor.kill
        /// A failure line.
        public static let wipe = FSColor.wipe
        /// Spec-only values with no site token yet (spec §4.1, §5.0, F2.6); lane T folds them
        /// into tokens.css when the site gains them.
        public static let offlineGround = Color(tokenHex: 0x11151F)
        public static let sheet = Color(tokenHex: 0x10141D)
        public static let tooltipPanel = Color(tokenHex: 0x0A0D14)
        public static let iconLine = Color(tokenHex: 0x3A4256)
        /// The game's green for Equip and Use lines, and its gold for Item Level.
        public static let gameGreen = FSColor.parseGreen
        public static let itemLevel = FSColor.parseGold
        /// A chip or secondary control's 4.5% white fill, and the pressed 10% gold wash.
        public static let controlFill = Color.white.opacity(0.045)
        public static let pressedWash = FSColor.gold.opacity(0.10)
    }

    /// The site's type: Cinzel for the wordmark, Barlow Condensed for headings, Barlow for
    /// body copy. The faces are bundled from apps/shared/fonts and listed under UIAppFonts.
    public enum Typography {
        public static let title = FSFont.headingBold(FSTextSize.h2)
        public static let cardTitle = FSFont.headingBold(FSTextSize.cardTitle)
        public static let body = FSFont.bodyRegular(FSTextSize.body)
        public static let meta = FSFont.bodySemibold(FSTextSize.meta)
        /// The spec's type roles (§4.2), each scaling with the system text size.
        public static let screenTitle = Font.custom("BarlowCondensed-ExtraBold", size: FSTextSize.h1Phone, relativeTo: .largeTitle)
        public static let section = Font.custom("BarlowCondensed-ExtraBold", size: FSTextSize.cardTitle, relativeTo: .headline)
        public static let name = Font.custom("Barlow-Bold", size: 17, relativeTo: .headline)
        public static let bodyRole = Font.custom("Barlow-Regular", size: FSTextSize.body, relativeTo: .subheadline)
        public static let bodyStrong = Font.custom("Barlow-SemiBold", size: FSTextSize.body, relativeTo: .subheadline)
        public static let metaRole = Font.custom("Barlow-Regular", size: FSTextSize.meta, relativeTo: .footnote)
        public static let micro = Font.custom("Barlow-Bold", size: FSTextSize.micro, relativeTo: .caption2)
        /// Figures, counts and times: the system monospaced face at body size.
        public static let figure = Font.system(.subheadline, design: .monospaced, weight: .semibold)
    }

    public enum Spacing {
        public static let xs = FSSpacing.space1
        public static let sm = FSSpacing.space2
        public static let md = FSSpacing.space4
        public static let lg = FSSpacing.space6
    }

    public enum Radius {
        public static let card = FSRadius.control
        public static let panel = FSRadius.panel
        public static let control = FSRadius.control
        public static let pill = FSRadius.pill
    }
}

/// The scaffold's card surface: raised background, hairline border.
public struct CardSurface: ViewModifier {
    public init() {}

    public func body(content: Content) -> some View {
        content
            .padding(Theme.Spacing.md)
            .frame(maxWidth: .infinity, alignment: .leading)
            .background(Theme.Palette.raised, in: .rect(cornerRadius: Theme.Radius.card))
            .overlay(RoundedRectangle(cornerRadius: Theme.Radius.card).stroke(Theme.Palette.line))
    }
}

public extension View {
    func cardSurface() -> some View { modifier(CardSurface()) }
}
