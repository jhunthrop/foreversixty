@_exported import DesignTokens
import SwiftUI

/// The spec's type roles (design/specs/2026-10-30-mobile-apps.md §4.2), each scaling with
/// Dynamic Type through `relativeTo`. `FSFont` (generated) has fixed sizes only, so the faces
/// are named here once, with the same PostScript names lane T bundles under UIAppFonts.
/// JetBrains Mono is not bundled yet (apps/shared/fonts), so `mono` uses the system mono.
public enum ScaledType {
    private enum Face {
        static let heading = "BarlowCondensed-ExtraBold"
        static let body = "Barlow-Regular"
        static let bodySemibold = "Barlow-SemiBold"
        static let bodyBold = "Barlow-Bold"
        static let wordmark = "Cinzel-Bold"
    }

    /// Spec sizes the tokens do not carry.
    private enum Size {
        static let name: CGFloat = 17
        static let lead: CGFloat = 17
        static let tooltip: CGFloat = 16
        static let wordmark: CGFloat = 22
        static let button: CGFloat = 13
    }

    public static let title = Font.custom(Face.heading, size: FSTextSize.h1Phone, relativeTo: .largeTitle)
    public static let h2 = Font.custom(Face.heading, size: FSTextSize.h2, relativeTo: .title)
    public static let stat = Font.custom(Face.heading, size: FSTextSize.stat, relativeTo: .title)
    public static let sec = Font.custom(Face.heading, size: FSTextSize.cardTitle, relativeTo: .headline)
    public static let name = Font.custom(Face.bodyBold, size: Size.name, relativeTo: .headline)
    public static let lead = Font.custom(Face.bodySemibold, size: Size.lead, relativeTo: .body)
    public static let body = Font.custom(Face.body, size: FSTextSize.body, relativeTo: .subheadline)
    public static let bodyStrong = Font.custom(Face.bodySemibold, size: FSTextSize.body, relativeTo: .subheadline)
    public static let tooltip = Font.custom(Face.body, size: Size.tooltip, relativeTo: .body)
    public static let meta = Font.custom(Face.body, size: FSTextSize.meta, relativeTo: .footnote)
    public static let micro = Font.custom(Face.bodyBold, size: FSTextSize.micro, relativeTo: .caption2)
    public static let button = Font.custom(Face.bodyBold, size: Size.button, relativeTo: .subheadline)
    public static let wordmark = Font.custom(Face.wordmark, size: Size.wordmark, relativeTo: .title2)
    public static let mono = Font.system(size: FSTextSize.body, weight: .semibold, design: .monospaced)
}
