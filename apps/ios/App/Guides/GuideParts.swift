import DesignSystem
import Models
import SwiftUI

/// F9's sizes (spec F9, §4).
enum GuidesLayout {
    static let gutter: CGFloat = 16
    static let sectionGap: CGFloat = 22
    static let cardPadding: CGFloat = 14
    static let headIcon: CGFloat = 56
    static let rowIcon: CGFloat = 40
    static let tierIcon: CGFloat = 36
    static let crest: CGFloat = 44
    static let chip: CGFloat = 36
    static let chipDisc: CGFloat = 22
    static let yourCard: CGFloat = 80
    static let tierRow: CGFloat = 68
    /// The skeleton's stand-in for a row's figure column.
    static let tierFigure: CGFloat = 56
    static let yoursEdge: CGFloat = 3
    static let dateBlock = CGSize(width: 56, height: 48)
    static let factionMark: CGFloat = 24
    static let glyph: CGFloat = 18
    /// The guide's section strip: 36 pt pills with their padding.
    static let stripHeight: CGFloat = 52
    /// A section counts as current once its top passes under the strip.
    static let currentThreshold: CGFloat = 80
    static let tabBarClearance: CGFloat = 96
    static let chipNumberSize: CGFloat = 12
}

/// A group's head over the ground (Overview, All classes, Rotation, What changed): `SectionTitle`, as
/// Android's, with the spec's 22 pt above it.
struct GuideSectionLabel: View {
    let text: String

    var body: some View {
        SectionTitle(text)
            .padding(.top, GuidesLayout.sectionGap)
            .padding(.bottom, 8)
    }
}

/// A numbered stat chip (spec §4): 36 pt pill, a 22 pt number disc in mono 12, the stat. The
/// disc and its number grow with the text; a stat that wraps keeps the chip's 18 pt corners, so
/// no letter is cut by the curve.
struct StatChipView: View {
    let chip: StatChip
    @ScaledMetric(relativeTo: .caption) private var disc = GuidesLayout.chipDisc
    @ScaledMetric(relativeTo: .caption) private var number = GuidesLayout.chipNumberSize

    var body: some View {
        let shape = RoundedRectangle(cornerRadius: GuidesLayout.chip / 2)
        HStack(spacing: 8) {
            Text("\(chip.number)")
                .font(Font.custom(FSFont.bodySemibold, size: number).monospacedDigit())
                .foregroundStyle(FSColor.strong)
                .frame(width: disc, height: disc)
                .background(FSColor.cardTop, in: .circle)
                .overlay(Circle().stroke(FSColor.line))
            Text(chip.stat)
                .font(ScaledType.bodyStrong)
                .foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
        }
        .padding(.leading, 7)
        .padding(.trailing, 14)
        .padding(.vertical, 4)
        .frame(minHeight: GuidesLayout.chip)
        .background(GameColors.controlFill, in: shape)
        .overlay(shape.strokeBorder(FSColor.lineSoft, lineWidth: 1))
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(chip.a11y)
    }
}

/// The stat priority as wrapping numbered chips; never cut.
struct StatChips: View {
    let chips: [StatChip]

    var body: some View {
        WrapLayout(spacing: 8) {
            ForEach(chips) { StatChipView(chip: $0) }
        }
    }
}

/// Bulleted lines of a guide inside a panel (strengths, weaknesses, a raid's notes): a Label12 head,
/// never a display-face heading inside a Plate, then each line after its dot.
struct GuideLines: View {
    let title: String
    let lines: [String]

    static let dot: CGFloat = 5

    var body: some View {
        if !lines.isEmpty {
            VStack(alignment: .leading, spacing: 6) {
                Label12(title).accessibilityAddTraits(.isHeader)
                ForEach(lines, id: \.self) { line in
                    HStack(alignment: .firstTextBaseline, spacing: FSSpace.space2) {
                        Circle().fill(FSColor.muted).frame(width: Self.dot, height: Self.dot)
                            .alignmentGuide(.firstTextBaseline) { $0[.bottom] + Self.dot }
                            .accessibilityHidden(true)
                        Text(line).font(ScaledType.body).foregroundStyle(FSColor.text)
                            .fixedSize(horizontal: false, vertical: true)
                    }
                }
            }
        }
    }
}

/// One link row: a line in gold and its chevron, 48 pt.
struct GuideLinkRow: View {
    let text: String
    var systemImage = "chevron.right"
    let action: () -> Void

    var body: some View {
        Button(action: action) {
            HStack(spacing: 12) {
                Text(text).font(ScaledType.bodyStrong).foregroundStyle(FSColor.gold)
                    .fixedSize(horizontal: false, vertical: true)
                    .frame(maxWidth: .infinity, alignment: .leading)
                Image(systemName: systemImage).font(.system(size: 13, weight: .semibold)).foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
            }
            .padding(.horizontal, GuidesLayout.cardPadding)
            .frame(minHeight: 48)
            .contentShape(Rectangle())
        }
        .buttonStyle(RowPressStyle())
    }
}

/// `Sources`: one quiet disclosure at a guide's foot; each label opens its page.
struct GuideSources: View {
    let sources: [GuideSourceLink]
    @State private var isOpen = false
    @Environment(\.openURL) private var openURL

    var body: some View {
        if !sources.isEmpty {
            DisclosureGroup(isExpanded: $isOpen) {
                VStack(alignment: .leading, spacing: 0) {
                    ForEach(sources, id: \.url) { source in
                        if let url = URL(string: source.url) {
                            GuideLinkRow(text: source.label, systemImage: "arrow.up.right") { openURL(url) }
                        }
                    }
                }
            } label: {
                Text(Copy.guidesSources).font(ScaledType.bodyStrong).foregroundStyle(FSColor.nav)
            }
            .tint(FSColor.muted)
            .padding(GuidesLayout.cardPadding)
            .quietCard()
        }
    }
}

/// The not-saved sentence (spec F9.6, §5.7).
struct GuidesNotCached: View {
    let text: String

    var body: some View {
        HStack(alignment: .top, spacing: 12) {
            Image(systemName: "icloud.and.arrow.down").font(.system(size: GuidesLayout.glyph)).foregroundStyle(FSColor.muted)
                .accessibilityHidden(true)
            Text(text).font(ScaledType.lead).foregroundStyle(FSColor.text).fixedSize(horizontal: false, vertical: true)
        }
        .padding(16)
        .frame(maxWidth: .infinity, alignment: .leading)
        .quietCard()
    }
}

/// The offline strip over a saved file, when it is one.
struct GuidesOfflineStrip: View {
    let since: Date?
    let now: Date

    var body: some View {
        if let since { OfflineStrip(text: Copy.offlineStrip(time: OfflineStamp.format(since, now: now))) }
    }
}

/// A row of a quiet card list, with the card's soft divider before every row but the first.
struct GuideCardRows<Item: Identifiable, Row: View>: View {
    let items: [Item]
    @ViewBuilder let row: (Item) -> Row

    var body: some View {
        VStack(spacing: 0) {
            ForEach(Array(items.enumerated()), id: \.element.id) { index, item in
                if index > 0 { Divider().overlay(FSColor.lineSoft) }
                row(item)
            }
        }
        .quietCard()
    }
}
