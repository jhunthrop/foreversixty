import DesignTokens
import SwiftUI

/// SwiftUI values from the generated `DesignTokens`. Views use `Theme`, never
/// raw hex, so a token regeneration reaches every screen through this file.
public enum Theme {
    public enum Palette {
        public static let background = Color(hex: DesignTokens.Color.bg)
        public static let raised = Color(hex: DesignTokens.Color.raised)
        public static let line = Color(hex: DesignTokens.Color.line)
        public static let text = Color(hex: DesignTokens.Color.text)
        public static let strong = Color(hex: DesignTokens.Color.strong)
        public static let muted = Color(hex: DesignTokens.Color.muted)
        public static let accent = Color(hex: DesignTokens.Color.gold)
    }

    public enum Spacing {
        public static let xs = CGFloat(DesignTokens.Spacing.xs)
        public static let sm = CGFloat(DesignTokens.Spacing.sm)
        public static let md = CGFloat(DesignTokens.Spacing.md)
        public static let lg = CGFloat(DesignTokens.Spacing.lg)
    }

    public enum Radius {
        public static let card = CGFloat(DesignTokens.Radius.card)
    }
}

public extension Color {
    /// An opaque sRGB colour from 0xRRGGBB.
    init(hex: UInt32) {
        let channel = { (shift: UInt32) in Double((hex >> shift) & 0xFF) / 255 }
        self.init(.sRGB, red: channel(16), green: channel(8), blue: channel(0), opacity: 1)
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
