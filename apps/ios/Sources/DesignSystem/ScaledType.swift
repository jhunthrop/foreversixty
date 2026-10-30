@_exported import DesignTokens
import SwiftUI

/// The spec's type roles (design/specs/2026-10-30-mobile-apps.md §4.2), each scaling with
/// Dynamic Type through `relativeTo`. `FSFont` (generated) has fixed sizes only, so the faces
/// are named here once, with the same PostScript names lane T bundles under UIAppFonts.
public enum ScaledType {
    private enum Face {
        static let heading = "BarlowCondensed-ExtraBold"
        static let body = "Barlow-Regular"
        static let bodySemibold = "Barlow-SemiBold"
        static let bodyBold = "Barlow-Bold"
        static let wordmark = "Cinzel-Bold"
    }

    /// Spec sizes the tokens do not carry: the selector and switcher name, the segments.
    private enum Size {
        static let name: CGFloat = 17
        static let wordmark: CGFloat = 22
        static let button: CGFloat = 13
    }

    public static let title = Font.custom(Face.heading, size: FSTextSize.h1Phone, relativeTo: .largeTitle)
    public static let h2 = Font.custom(Face.heading, size: FSTextSize.h2, relativeTo: .title)
    public static let sec = Font.custom(Face.heading, size: FSTextSize.cardTitle, relativeTo: .headline)
    public static let name = Font.custom(Face.bodyBold, size: Size.name, relativeTo: .headline)
    public static let body = Font.custom(Face.body, size: FSTextSize.body, relativeTo: .subheadline)
    public static let bodyStrong = Font.custom(Face.bodySemibold, size: FSTextSize.body, relativeTo: .subheadline)
    public static let meta = Font.custom(Face.body, size: FSTextSize.meta, relativeTo: .footnote)
    public static let micro = Font.custom(Face.bodyBold, size: FSTextSize.micro, relativeTo: .caption2)
    public static let button = Font.custom(Face.bodyBold, size: Size.button, relativeTo: .subheadline)
    public static let wordmark = Font.custom(Face.wordmark, size: Size.wordmark, relativeTo: .title2)
}

/// The site's primary and secondary buttons (spec §4.5), full width.
public struct BrandButtonStyle: ButtonStyle {
    public enum Kind: Sendable { case primary, secondary, destructive, quiet }

    private let kind: Kind
    private let height: CGFloat

    public init(_ kind: Kind, height: CGFloat = 52) {
        self.kind = kind
        self.height = height
    }

    public func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .font(ScaledType.button)
            .textCase(.uppercase)
            .tracking(0.9)
            .frame(maxWidth: .infinity, minHeight: height)
            .padding(.horizontal, FSSpacing.space4)
            .foregroundStyle(foreground)
            .background(background, in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(border))
            .opacity(configuration.isPressed ? 0.8 : 1)
    }

    private var foreground: Color {
        switch kind {
        case .primary: FSColor.bg
        case .secondary: FSColor.strong
        case .destructive: FSColor.wipe
        case .quiet: FSColor.nav
        }
    }

    private var background: AnyShapeStyle {
        switch kind {
        case .primary:
            AnyShapeStyle(LinearGradient(colors: [FSColor.goldHover, FSColor.gold], startPoint: .top, endPoint: .bottom))
        case .secondary, .destructive, .quiet:
            AnyShapeStyle(Color.white.opacity(0.045))
        }
    }

    private var border: Color {
        switch kind {
        case .primary: .clear
        case .secondary, .quiet: FSColor.gold.opacity(0.17)
        case .destructive: FSColor.wipe.opacity(0.45)
        }
    }
}

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
