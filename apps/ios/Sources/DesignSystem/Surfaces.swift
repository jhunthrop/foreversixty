import SwiftUI

/// The site's uppercase section label (`micro`, .14em).
public struct MicroLabel: View {
    private let text: String
    private let color: Color

    public init(_ text: String, color: Color = FSColor.muted) {
        self.text = text
        self.color = color
    }

    public var body: some View {
        Text(text.uppercased())
            .font(ScaledType.micro)
            .tracking(1.5)
            .foregroundStyle(color)
            .accessibilityLabel(text)
    }
}

/// A tag (`.pill`): 22 pt, 11 pt 700 uppercase, tinted ground and border, an optional glyph first.
public struct TagView: View {
    private let text: String
    private let color: Color
    private let systemImage: String?

    public init(text: String, color: Color, systemImage: String? = nil) {
        self.text = text
        self.color = color
        self.systemImage = systemImage
    }

    public var body: some View {
        HStack(spacing: 4) {
            if let systemImage {
                Image(systemName: systemImage).font(.system(size: 9, weight: .bold)).accessibilityHidden(true)
            }
            Text(text.uppercased()).font(ScaledType.micro).tracking(1)
        }
        .foregroundStyle(color)
        .padding(.horizontal, FSSpacing.space2)
        .frame(minHeight: 22)
        .background(color.opacity(0.12), in: .rect(cornerRadius: FSRadius.pill))
        .overlay(RoundedRectangle(cornerRadius: FSRadius.pill).stroke(color.opacity(0.35)))
        .fixedSize()
    }
}

/// The offline strip (§5.0): 36 pt, `#11151f`, wifi-off glyph, then the line
/// (`offline.strip` with its stamp).
public struct OfflineStrip: View {
    private let text: String

    public init(text: String) {
        self.text = text
    }

    public var body: some View {
        HStack(spacing: FSSpacing.space2) {
            Image(systemName: "wifi.slash").font(.system(size: 14)).accessibilityHidden(true)
            Text(text).font(ScaledType.meta)
            Spacer(minLength: 0)
        }
        .foregroundStyle(FSColor.nav)
        .padding(.horizontal, FSSpacing.space4)
        .frame(minHeight: 36)
        .background(GameColors.offlineGround)
        .overlay(alignment: .bottom) { FSColor.line.frame(height: 1) }
        .accessibilityElement(children: .combine)
        .onAppear { AccessibilityNotification.Announcement(text).post() }
    }
}

/// The primary card (`--surface-card`): the warm gradient and its gold border.
struct PrimaryCard: ViewModifier {
    func body(content: Content) -> some View {
        content
            .background(
                LinearGradient(colors: [Color(tokenHex: 0x16130E), FSColor.raised], startPoint: .top, endPoint: .bottom),
                in: .rect(cornerRadius: FSRadius.panel)
            )
            .overlay(RoundedRectangle(cornerRadius: FSRadius.panel).stroke(FSColor.lineCard))
    }
}

/// A quiet list card: raised ground, soft line.
struct QuietCard: ViewModifier {
    func body(content: Content) -> some View {
        content
            .background(FSColor.raised, in: .rect(cornerRadius: FSRadius.panel))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.panel).stroke(FSColor.lineSoft))
    }
}

/// The hero band (`--surface-band`): the ember radial over `#14100b`, a soft bottom line.
struct HeroBand: ViewModifier {
    func body(content: Content) -> some View {
        content
            .background(
                ZStack {
                    Color(tokenHex: 0x14100B)
                    RadialGradient(colors: [FSColor.ember.opacity(0.22), .clear], center: .top, startRadius: 0, endRadius: 320)
                }
            )
            .overlay(alignment: .bottom) { FSColor.lineSoft.frame(height: 1) }
    }
}

public extension View {
    func primaryCard() -> some View { modifier(PrimaryCard()) }
    func quietCard() -> some View { modifier(QuietCard()) }
    func heroBand() -> some View { modifier(HeroBand()) }
}

/// A list row's press: the 10% gold wash while held (spec §8).
public struct RowPressStyle: ButtonStyle {
    public init() {}

    public func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .background(configuration.isPressed ? GameColors.pressedWash : .clear)
    }
}

public extension View {
    /// A chip or pill's ground: the capsule filled with `fill` and its 1 pt border inside the edge.
    func pillSurface(fill: Color, line: Color) -> some View {
        background { Capsule().fill(fill) }
            .overlay { Capsule().strokeBorder(line, lineWidth: 1) }
            .contentShape(.capsule)
    }
}
