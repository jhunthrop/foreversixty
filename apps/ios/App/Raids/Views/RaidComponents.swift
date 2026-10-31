import DesignSystem
import SwiftUI

/// The site, for the raid screens' web links.
enum StaticSite {
    static let url = URL(string: "https://foreversixty.gg")!
}

extension EnvironmentValues {
    /// The active data build the raid screens' game icons load from (`GameIconImage`).
    @Entry var raidIconBuild: String?
}

/// A square game icon on the shared icon loader: the spec icon (the `#3a4256` line) or a buff.
/// Until the icon loads, or when it fails, the tile shows [fallback] (`MotW`, `Bless`).
/// A [dimmed] tile (a missing buff) greys and fades the art only; its label stays readable.
struct GameIcon: View {
    let stem: String?
    let size: CGFloat
    var border: Color = GameColors.iconLine
    var borderWidth: CGFloat = 1
    var fallback: String?
    var dimmed = false
    @Environment(\.raidIconBuild) private var build

    var body: some View {
        let shape = RoundedRectangle(cornerRadius: FSRadius.pill)
        GameIconImage(icon: stem, build: build, fallback: fallback, dimmed: dimmed)
            .id("\(build ?? "")/\(stem ?? "")")
            .frame(width: size, height: size)
            .clipShape(shape)
            .overlay(shape.stroke(border, lineWidth: borderWidth))
            .accessibilityHidden(true)
    }
}

/// `micro` uppercase .14em section label; at accessibility sizes the tracking drops and the label
/// shrinks rather than breaking.
struct SectionLabel: View {
    let text: String
    var color: Color = FSColor.muted
    var trailing: String?
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        HStack(spacing: 6) {
            Text(text.uppercased()).font(ScaledType.micro).tracking(typeSize.isAccessibilitySize ? 0 : 1.5)
                .lineLimit(1).minimumScaleFactor(0.5).foregroundStyle(color)
            if let trailing {
                Text(trailing).font(ScaledType.monoMeta).foregroundStyle(color)
            }
        }
        .accessibilityAddTraits(.isHeader)
    }
}

extension TagView {
    /// The raid status tag: `Planned` gold, `Locked` muted, `Cancelled` wipe, `Ended` muted, `Done` kill.
    init(chip: RaidChip) {
        let color = switch chip {
        case .planned: FSColor.gold
        case .cancelled: FSColor.wipe
        case .done: FSColor.kill
        case .locked, .ended: FSColor.muted
        }
        self.init(text: chip.label, color: color)
    }
}

/// The date block, `THU` / `10` / `DEC` (hidden from VoiceOver; the title row reads the date).
struct DateBlock: View {
    let date: Date
    let clock: RaidClock
    var width: CGFloat = 52
    var height: CGFloat = 56

    var body: some View {
        let parts = clock.dateBlock(date)
        VStack(spacing: 0) {
            Text(parts.weekday).font(ScaledType.micro).foregroundStyle(FSColor.muted)
            Text(parts.day).font(ScaledType.monoDay).foregroundStyle(FSColor.strong)
            Text(parts.month).font(ScaledType.micro).foregroundStyle(FSColor.muted)
        }
        // The block keeps its width at every size; its words shrink, never break.
        .lineLimit(1)
        .minimumScaleFactor(0.4)
        .padding(.vertical, 4)
        .padding(.horizontal, 2)
        .frame(width: width)
        .frame(minHeight: height)
        .background(FSColor.bg.opacity(0.5), in: .rect(cornerRadius: FSRadius.control))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line))
        .accessibilityHidden(true)
    }
}

/// The 8 pt progress bar: line-soft track, gold fill. A nil width takes the full row.
struct FillBar: View {
    let fill: Double
    let width: CGFloat?
    var muted = false

    var body: some View {
        Capsule().fill(FSColor.lineSoft)
            .overlay(alignment: .leading) {
                GeometryReader { proxy in
                    Capsule().fill(muted ? FSColor.muted : FSColor.gold).frame(width: proxy.size.width * fill)
                }
            }
            .frame(width: width, height: 8)
            .frame(maxWidth: width == nil ? .infinity : nil, alignment: .leading)
            .accessibilityHidden(true)
    }
}

/// A one-word name: one line at every size, shrinking before it would break inside the word.
struct NameText: View {
    let text: String
    let font: Font
    let color: Color

    var body: some View {
        Text(text).font(font).foregroundStyle(color).lineLimit(1).minimumScaleFactor(0.5)
    }
}

/// The count block (spec F4.1): the bar, `24 of 40 signed up`, and the role breakdown.
struct CountBlock: View {
    let facts: RaidFacts
    var muted = false
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        VStack(alignment: .leading, spacing: 6) {
            if typeSize.isAccessibilitySize {
                // Large text: the count above a full-width bar (spec 10).
                Text(facts.countLine).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                if let fill = facts.fill { FillBar(fill: fill, width: nil, muted: muted) }
            } else {
                HStack(spacing: 12) {
                    if let fill = facts.fill { FillBar(fill: fill, width: 120, muted: muted) }
                    Text(facts.countLine).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                }
            }
            Text(facts.breakdown).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.raidsA11yCount(count: facts.countLine, breakdown: facts.breakdown.replacingOccurrences(of: " · ", with: ", ")))
    }
}

extension View {
    /// A raid card's box inside its DS surface: full width, leading, padded.
    func raidCardBox(_ padding: CGFloat) -> some View {
        self.padding(padding).frame(maxWidth: .infinity, alignment: .leading)
    }
}

/// The gold text link (`.alink`), 44 pt hit.
struct GoldLinkStyle: ButtonStyle {
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .font(ScaledType.bodyStrong)
            .lineLimit(1)
            .minimumScaleFactor(0.6)
            .foregroundStyle(configuration.isPressed ? FSColor.goldHover : FSColor.gold)
            .frame(minHeight: 44)
            .contentShape(Rectangle())
    }
}

/// An error region (spec 5.0): the sentence in wipe red, read as an alert, and Try again.
struct ErrorRegion: View {
    let message: String
    let retry: () -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            Text(message).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                .accessibilityAddTraits(.isStaticText)
            SecondaryButton(title: Copy.commonTryAgain, action: retry)
        }
        .raidCardBox(16).quietCard()
    }
}
