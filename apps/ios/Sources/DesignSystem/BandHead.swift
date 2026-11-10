import SwiftUI
import UIKit

// One bar system for every pushed screen (spec ruling N3, design/specs/2026-11-07-mobile-redesign.md
// 0A): the navigation bar is clear and carries no title at rest, over the painted band; the screen's
// head continues that band from under the status bar and the bar to below the head, with the title
// in the display face. Once the head scrolls under it, the bar takes the band and its title fades in.

/// The band head's measures.
public enum BandReach {
    /// How far a head's band art reaches above it: past the navigation bar, the status bar and the island.
    public static let underBars: CGFloat = 240
    /// The head's side and bottom padding, and the gap a screen leaves between the band and its content.
    public static let gap: CGFloat = 16
}

public extension View {
    /// The band art behind a head, running up under the status bar and the navigation bar, with the
    /// band's black bottom edge. Secondary lines on it take `mutedGround`. `bleed`: how far the band
    /// reaches past the head's sides, for a head inside a padded column, so it still meets the edges.
    func bandUnderBars(bleed: CGFloat = 0) -> some View {
        background(alignment: .bottom) { BandArt().padding(.top, -BandReach.underBars).padding(.horizontal, -bleed) }
            .overlay(alignment: .bottom) { Color.black.frame(height: FSStroke.hairline).padding(.horizontal, -bleed) }
            .environment(\.subInk, FSColor.mutedGround)
    }

    /// A pushed screen's bar (N3): `title` names the screen for VoiceOver and the next screen's back
    /// button, while the bar shows only the back chevron and the actions until the head scrolls under.
    func pushedBandBar(_ title: String) -> some View {
        navigationTitle(title)
            .navigationBarTitleDisplayMode(.inline)
            .background { BandBar(hidesTitleAtRest: true).frame(width: 0, height: 0).accessibilityHidden(true) }
    }
}

/// A pushed screen's head on the band (the guide reader's and the character's pattern, shared): an
/// optional lead (a crest or an icon), the title in the display face at 32, strong or a class colour,
/// its facts under it, and an optional trailing control. The screen sets its content `BandReach.gap`
/// under the head.
public struct BandHead<Lead: View, Facts: View, Trailing: View>: View {
    /// The space between the bar and the title.
    static var top: CGFloat { FSSpace.space2 }
    static var leadGap: CGFloat { 14 }
    static var factsGap: CGFloat { 4 }

    private let title: String
    private let titleColor: Color
    private let lead: Lead
    private let facts: Facts
    private let trailing: Trailing

    public init(
        title: String,
        titleColor: Color = FSColor.strong,
        @ViewBuilder lead: () -> Lead,
        @ViewBuilder facts: () -> Facts,
        @ViewBuilder trailing: () -> Trailing
    ) {
        self.title = title
        self.titleColor = titleColor
        self.lead = lead()
        self.facts = facts()
        self.trailing = trailing()
    }

    public var body: some View {
        HStack(alignment: .center, spacing: Self.leadGap) {
            lead
            VStack(alignment: .leading, spacing: Self.factsGap) {
                Text(title)
                    .type(FSType.title)
                    .foregroundStyle(titleColor)
                    .fixedSize(horizontal: false, vertical: true)
                    .accessibilityAddTraits(.isHeader)
                facts
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            trailing
        }
        .padding(.horizontal, BandReach.gap)
        .padding(.top, Self.top)
        .padding(.bottom, BandReach.gap)
        .frame(maxWidth: .infinity, alignment: .leading)
        .bandUnderBars()
    }
}

public extension BandHead where Trailing == EmptyView {
    init(title: String, titleColor: Color = FSColor.strong, @ViewBuilder lead: () -> Lead, @ViewBuilder facts: () -> Facts) {
        self.init(title: title, titleColor: titleColor, lead: lead, facts: facts, trailing: { EmptyView() })
    }
}

public extension BandHead where Lead == EmptyView, Trailing == EmptyView {
    init(title: String, titleColor: Color = FSColor.strong, @ViewBuilder facts: () -> Facts) {
        self.init(title: title, titleColor: titleColor, lead: { EmptyView() }, facts: facts, trailing: { EmptyView() })
    }
}

public extension BandHead where Lead == EmptyView, Facts == EmptyView, Trailing == EmptyView {
    init(title: String, titleColor: Color = FSColor.strong) {
        self.init(title: title, titleColor: titleColor, lead: { EmptyView() }, facts: { EmptyView() }, trailing: { EmptyView() })
    }
}

/// A head's fact line on the band: small, in the band's secondary ink.
public struct BandFact: View {
    private let text: String

    public init(_ text: String) {
        self.text = text
    }

    public var body: some View {
        Text(text).font(ScaledType.meta).foregroundStyle(FSColor.mutedGround).fixedSize(horizontal: false, vertical: true)
    }
}
