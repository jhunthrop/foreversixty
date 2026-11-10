@_exported import DesignTokens
import SwiftUI

/// The spec's type roles (design/specs/2026-10-30-mobile-apps.md §4.2), each scaling with
/// Dynamic Type through `relativeTo`. `FSFont` (generated) has fixed sizes only, so the faces
/// are the generated `FSFont` PostScript names. Round 5 (design/specs/2026-11-09-directions.md,
/// Mobile): headings and the wordmark are Marcellus 400, as on the site; body, labels and figures
/// are Open Sans. JetBrains Mono is not bundled, so `mono` uses the system mono.
public enum ScaledType {
    private enum Face {
        static let heading = FSFont.title
        static let body = FSFont.body
        static let bodySemibold = FSFont.bodySemibold
        static let bodyBold = FSFont.bodyBold
        static let wordmark = FSFont.title
    }

    /// Spec sizes the tokens do not carry.
    private enum Size {
        static let name: CGFloat = 17
        static let lead: CGFloat = 17
        static let socketName: CGFloat = 12
        static let tooltip: CGFloat = 16
        static let wordmark: CGFloat = 22
        static let button: CGFloat = 13
        static let cardHead: CGFloat = 22
        static let segment: CGFloat = 14
    }

    public static let title = Font.custom(Face.heading, size: FSTextSize.h1Phone, relativeTo: .largeTitle)
    public static let h2 = Font.custom(Face.heading, size: FSTextSize.h2, relativeTo: .title)
    public static let stat = Font.custom(Face.bodyBold, size: FSTextSize.stat, relativeTo: .title)
    public static let sec = Font.custom(Face.heading, size: FSTextSize.cardTitle, relativeTo: .headline)
    /// A raid card's title (spec F4.1).
    public static let cardHead = Font.custom(Face.heading, size: Size.cardHead, relativeTo: .title2)
    public static let name = Font.custom(Face.bodyBold, size: Size.name, relativeTo: .headline)
    public static let lead = Font.custom(Face.bodySemibold, size: Size.lead, relativeTo: .body)
    public static let body = Font.custom(Face.body, size: FSType.body.size, relativeTo: .subheadline)
    public static let bodyStrong = Font.custom(Face.bodySemibold, size: FSType.body.size, relativeTo: .subheadline)
    public static let tooltip = Font.custom(Face.body, size: Size.tooltip, relativeTo: .body)
    public static let meta = Font.custom(Face.body, size: FSTextSize.meta, relativeTo: .footnote)
    /// The 12 pt label (`FSType.label`): 700, sentence case on native (ruling N1).
    public static let label = Font.custom(Face.bodyBold, size: FSType.label.size, relativeTo: .caption)
    /// A section's title over the ground (round 5, the mocks' `.sh h2`): the display face at 21.
    public static let sectionTitle = Font.custom(Face.heading, size: 21, relativeTo: .title3)
    public static let button = Font.custom(Face.bodyBold, size: Size.button, relativeTo: .subheadline)
    /// A segmented control's label (the sign-up segments, spec F4.3).
    public static let segment = Font.custom(Face.bodySemibold, size: Size.segment, relativeTo: .subheadline)
    public static let wordmark = Font.custom(Face.wordmark, size: Size.wordmark, relativeTo: .title2)
    /// A Paperdoll item name beside its socket (paperdoll spec §3.3 phone: 12 pt 600).
    public static let socketName = Font.custom(Face.bodySemibold, size: Size.socketName, relativeTo: .caption)
    /// The Forever Score figure: Open Sans 700 at 36, tabular (`.largeTitle` is the scaling role nearest it).
    public static let scoreFigure = Font.custom(Face.bodyBold, size: 36, relativeTo: .largeTitle).monospacedDigit()
    /// Figures, counts and times beside body text: Open Sans 600, tabular, as the site sets its numbers.
    public static let tabular = Font.custom(Face.bodySemibold, size: FSType.body.size, relativeTo: .subheadline).monospacedDigit()
    /// Clock digits and counts beside meta text (13 pt, `.footnote`), tabular.
    public static let tabularMeta = Font.custom(Face.body, size: FSTextSize.meta, relativeTo: .footnote).monospacedDigit()
    /// A date block's day number, a level (`.title3`), tabular.
    public static let tabularDay = Font.custom(Face.bodySemibold, size: 20, relativeTo: .title3).monospacedDigit()
}
