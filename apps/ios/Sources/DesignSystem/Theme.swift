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
    }

    /// The site's type: Cinzel for the wordmark, Barlow Condensed for headings, Barlow for
    /// body copy. The faces are bundled from apps/shared/fonts and listed under UIAppFonts.
    public enum Typography {
        public static let title = FSFont.headingBold(FSTextSize.h2)
        public static let cardTitle = FSFont.headingBold(FSTextSize.cardTitle)
        public static let body = FSFont.bodyRegular(FSTextSize.body)
        public static let meta = FSFont.bodySemibold(FSTextSize.meta)
    }

    public enum Spacing {
        public static let xs = FSSpacing.space1
        public static let sm = FSSpacing.space2
        public static let md = FSSpacing.space4
        public static let lg = FSSpacing.space6
    }

    public enum Radius {
        public static let card = FSRadius.control
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
