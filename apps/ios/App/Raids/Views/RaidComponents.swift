import DesignSystem
import SwiftUI
import UIKit

/// Where the game icons come from: `/data/{build}/icons/{stem}.webp` on the site (app-v1.md).
struct RaidIcons: Sendable {
    let build: String?

    func url(_ stem: String?) -> URL? {
        guard let build, let stem, !stem.isEmpty else { return nil }
        return StaticSite.url.appending(path: "data/\(build)/icons/\(stem).webp")
    }
}

enum StaticSite {
    static let url = URL(string: "https://foreversixty.gg")!
}

private struct RaidIconsKey: EnvironmentKey {
    static let defaultValue = RaidIcons(build: nil)
}

extension EnvironmentValues {
    var raidIcons: RaidIcons {
        get { self[RaidIconsKey.self] }
        set { self[RaidIconsKey.self] = newValue }
    }
}

/// The ringed HD class crest (spec 4.3), bundled from web/public/icons/hd/crests.
struct CrestView: View {
    let className: String
    let size: CGFloat

    var body: some View {
        let color = Theme.Palette.classColor(className)
        ZStack {
            Circle().fill(Theme.Palette.raised)
            if let image = Self.image(className) {
                Image(uiImage: image).resizable().scaledToFill().clipShape(Circle())
            }
        }
        .frame(width: size, height: size)
        .overlay(Circle().stroke(color.opacity(0.6), lineWidth: 2))
        .shadow(color: color.opacity(0.35), radius: size > 32 ? 7 : 4)
        .accessibilityHidden(true)
    }

    static func image(_ className: String) -> UIImage? {
        guard let folder = Bundle.main.url(forResource: "crests", withExtension: nil) else { return nil }
        return UIImage(contentsOfFile: folder.appending(path: "\(className.lowercased()).webp").path)
    }
}

/// A square game icon: the spec icon (border `#3a4256`) or an item icon (quality border).
struct GameIcon: View {
    let stem: String?
    let size: CGFloat
    var border: Color = Color(red: 0x3A / 255, green: 0x42 / 255, blue: 0x56 / 255)
    var borderWidth: CGFloat = 1
    @Environment(\.raidIcons) private var icons

    var body: some View {
        RoundedRectangle(cornerRadius: Theme.Radius.pill)
            .fill(Theme.Palette.cardTop)
            .overlay {
                if let url = icons.url(stem) {
                    AsyncImage(url: url, transaction: Transaction(animation: .easeIn(duration: 0.12))) { phase in
                        if let image = phase.image { image.resizable().scaledToFill() } else { Color.clear }
                    }
                }
            }
            .clipShape(RoundedRectangle(cornerRadius: Theme.Radius.pill))
            .overlay(RoundedRectangle(cornerRadius: Theme.Radius.pill).stroke(border, lineWidth: borderWidth))
            .frame(width: size, height: size)
            .accessibilityHidden(true)
    }
}

/// `micro` uppercase .14em section label; at accessibility sizes the tracking drops and the label
/// shrinks rather than breaking.
struct SectionLabel: View {
    let text: String
    var color: Color = Theme.Palette.muted
    var trailing: String?
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        HStack(spacing: 6) {
            Text(text.uppercased()).font(Theme.Typography.micro).tracking(typeSize.isAccessibilitySize ? 0 : 1.5)
                .lineLimit(1).minimumScaleFactor(0.5).foregroundStyle(color)
            if let trailing {
                Text(trailing).font(Theme.Typography.monoMeta).foregroundStyle(color)
            }
        }
        .accessibilityAddTraits(.isHeader)
    }
}

/// The status tag: `Planned` gold, `Locked` muted, `Cancelled` wipe, `Ended` muted, `Done` kill.
struct StatusTag: View {
    let text: String
    let color: Color

    init(text: String, color: Color) {
        self.text = text
        self.color = color
    }

    init(chip: RaidChip) {
        text = chip.label
        color = switch chip {
        case .planned: Theme.Palette.accent
        case .cancelled: Theme.Palette.wipe
        case .done: Theme.Palette.kill
        case .locked, .ended: Theme.Palette.muted
        }
    }

    var body: some View {
        Text(text.uppercased())
            .font(Theme.Typography.micro)
            .tracking(0.8)
            .lineLimit(1)
            .minimumScaleFactor(0.5)
            .foregroundStyle(color)
            .padding(.horizontal, 7)
            .frame(minHeight: 22)
            .background(color.opacity(0.12), in: .rect(cornerRadius: Theme.Radius.pill))
            .overlay(RoundedRectangle(cornerRadius: Theme.Radius.pill).stroke(color.opacity(0.45)))
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
            Text(parts.weekday).font(Theme.Typography.micro).foregroundStyle(Theme.Palette.muted)
            Text(parts.day).font(Theme.Typography.monoDay).foregroundStyle(Theme.Palette.strong)
            Text(parts.month).font(Theme.Typography.micro).foregroundStyle(Theme.Palette.muted)
        }
        // The block keeps its width at every size; its words shrink, never break.
        .lineLimit(1)
        .minimumScaleFactor(0.4)
        .padding(.vertical, 4)
        .padding(.horizontal, 2)
        .frame(width: width)
        .frame(minHeight: height)
        .background(Theme.Palette.background.opacity(0.5), in: .rect(cornerRadius: Theme.Radius.control))
        .overlay(RoundedRectangle(cornerRadius: Theme.Radius.control).stroke(Theme.Palette.line))
        .accessibilityHidden(true)
    }
}

/// The 8 pt progress bar: line-soft track, gold fill. A nil width takes the full row.
struct FillBar: View {
    let fill: Double
    let width: CGFloat?
    var muted = false

    var body: some View {
        Capsule().fill(Theme.Palette.lineSoft)
            .overlay(alignment: .leading) {
                GeometryReader { proxy in
                    Capsule().fill(muted ? Theme.Palette.muted : Theme.Palette.accent).frame(width: proxy.size.width * fill)
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
                Text(facts.countLine).font(Theme.Typography.bodyStrong).foregroundStyle(Theme.Palette.strong)
                if let fill = facts.fill { FillBar(fill: fill, width: nil, muted: muted) }
            } else {
                HStack(spacing: 12) {
                    if let fill = facts.fill { FillBar(fill: fill, width: 120, muted: muted) }
                    Text(facts.countLine).font(Theme.Typography.bodyStrong).foregroundStyle(Theme.Palette.strong)
                }
            }
            Text(facts.breakdown).font(Theme.Typography.metaText).foregroundStyle(Theme.Palette.muted)
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(Copy.raidsA11yCount(count: facts.countLine, breakdown: facts.breakdown.replacingOccurrences(of: " · ", with: ", ")))
    }
}

/// The offline strip (spec 5.0): 36 pt, wifi-off glyph, `Offline. Showing what you had at {time}.`
struct OfflineStrip: View {
    let text: String

    var body: some View {
        HStack(spacing: 8) {
            Image(systemName: "wifi.slash").font(.system(size: 14)).accessibilityHidden(true)
            Text(text).font(Theme.Typography.metaText)
        }
        .foregroundStyle(Theme.Palette.nav)
        .padding(.horizontal, 16)
        .frame(maxWidth: .infinity, minHeight: 36, alignment: .leading)
        .background(Theme.Palette.strip)
        .overlay(alignment: .bottom) { Rectangle().fill(Theme.Palette.line).frame(height: 1) }
        .accessibilityElement(children: .combine)
    }
}

/// The primary card surface (`--surface-card`): the warm gradient, its gold border and highlight.
struct PrimaryCard: ViewModifier {
    func body(content: Content) -> some View {
        content
            .padding(16)
            .frame(maxWidth: .infinity, alignment: .leading)
            .background(
                LinearGradient(colors: [Color(red: 0x1A / 255, green: 0x17 / 255, blue: 0x12 / 255), Theme.Palette.raised], startPoint: .top, endPoint: .bottom),
                in: .rect(cornerRadius: Theme.Radius.panel)
            )
            .overlay(RoundedRectangle(cornerRadius: Theme.Radius.panel).stroke(Theme.Palette.lineCard))
    }
}

/// A quiet card: raised, line-soft border.
struct QuietCard: ViewModifier {
    var padding: CGFloat = 0

    func body(content: Content) -> some View {
        content
            .padding(padding)
            .frame(maxWidth: .infinity, alignment: .leading)
            .background(Theme.Palette.raised, in: .rect(cornerRadius: Theme.Radius.panel))
            .overlay(RoundedRectangle(cornerRadius: Theme.Radius.panel).stroke(Theme.Palette.lineSoft))
    }
}

extension View {
    func primaryCard() -> some View { modifier(PrimaryCard()) }
    func quietCard(padding: CGFloat = 0) -> some View { modifier(QuietCard(padding: padding)) }
}

/// The secondary button (4.5% white, gold 17% border, 48 tall, 13 pt 700 uppercase).
struct SecondaryButtonStyle: ButtonStyle {
    var gold = false

    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .font(Theme.Typography.button)
            .tracking(0.9)
            .textCase(.uppercase)
            .lineLimit(1)
            .minimumScaleFactor(0.6)
            .foregroundStyle(gold ? Theme.Palette.accent : Theme.Palette.strong)
            .padding(.horizontal, 14)
            .frame(minHeight: 48)
            .background(configuration.isPressed ? Theme.Palette.pressed : Color.white.opacity(0.045), in: .rect(cornerRadius: Theme.Radius.control))
            .overlay(RoundedRectangle(cornerRadius: Theme.Radius.control).stroke(gold ? Theme.Palette.accent : Theme.Palette.lineCard))
    }
}

/// The gold text link (`.alink`), 44 pt hit.
struct GoldLinkStyle: ButtonStyle {
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .font(Theme.Typography.bodyStrong)
            .lineLimit(1)
            .minimumScaleFactor(0.6)
            .foregroundStyle(configuration.isPressed ? Theme.Palette.goldHover : Theme.Palette.accent)
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
            Text(message).font(Theme.Typography.bodyText).foregroundStyle(Theme.Palette.wipe)
                .accessibilityAddTraits(.isStaticText)
            Button(Copy.commonTryAgain, action: retry).buttonStyle(SecondaryButtonStyle())
        }
        .quietCard(padding: 16)
    }
}
