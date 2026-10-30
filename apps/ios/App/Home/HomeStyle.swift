import DesignTokens
import SwiftUI

/// The spec's type roles (§4.2), each scaling with Dynamic Type through `relativeTo`.
/// JetBrains Mono is not bundled yet (apps/shared/fonts), so mono figures use the system mono.
enum HomeType {
    static let stat = Font.custom("BarlowCondensed-ExtraBold", size: FSTextSize.stat, relativeTo: .title)
    static let sectionTitle = Font.custom("BarlowCondensed-ExtraBold", size: FSTextSize.cardTitle, relativeTo: .headline)
    static let pageTitle = Font.custom("BarlowCondensed-ExtraBold", size: FSTextSize.h1Phone, relativeTo: .largeTitle)
    static let lead = Font.custom("Barlow-SemiBold", size: 17, relativeTo: .body)
    static let body = Font.custom("Barlow-Regular", size: FSTextSize.body, relativeTo: .subheadline)
    static let bodyStrong = Font.custom("Barlow-SemiBold", size: FSTextSize.body, relativeTo: .subheadline)
    static let meta = Font.custom("Barlow-Regular", size: FSTextSize.meta, relativeTo: .footnote)
    static let micro = Font.custom("Barlow-Bold", size: FSTextSize.micro, relativeTo: .caption2)
    static let button = Font.custom("Barlow-Bold", size: 13, relativeTo: .subheadline)
    static let tooltip = Font.custom("Barlow-Regular", size: 16, relativeTo: .body)
    static let mono = Font.system(size: FSTextSize.body, weight: .semibold, design: .monospaced)
}

/// The site's uppercase section label (`micro`, .14em).
struct MicroLabel: View {
    let text: String
    var color: Color = FSColor.muted

    init(_ text: String, color: Color = FSColor.muted) {
        self.text = text
        self.color = color
    }

    var body: some View {
        Text(text.uppercased())
            .font(HomeType.micro)
            .tracking(1.5)
            .foregroundStyle(color)
            .accessibilityLabel(text)
    }
}

/// `PRIMARY_BUTTON_FIXED`: the gold gradient, 48 pt, uppercase label. One per screen at most.
struct PrimaryButton: View {
    let title: String
    let action: () -> Void

    var body: some View {
        Button(action: action) {
            Text(title.uppercased())
                .font(HomeType.button).tracking(0.9)
                .foregroundStyle(FSColor.bg)
                .padding(.horizontal, 18)
                .frame(minHeight: 48)
                .background(
                    LinearGradient(colors: [FSColor.goldHover, FSColor.gold], startPoint: .top, endPoint: .bottom),
                    in: .rect(cornerRadius: FSRadius.control)
                )
        }
        .buttonStyle(PressScale())
    }
}

/// `SECONDARY_BUTTON_FIXED`: 4.5% white, the gold 17% border.
struct SecondaryButton: View {
    let title: String
    var systemImage: String?
    let action: () -> Void

    var body: some View {
        Button(action: action) {
            HStack(spacing: 8) {
                if let systemImage { Image(systemName: systemImage).accessibilityHidden(true) }
                Text(title.uppercased()).font(HomeType.button).tracking(0.9)
            }
            .foregroundStyle(FSColor.strong)
            .padding(.horizontal, 18)
            .frame(minHeight: 48)
            .background(Color.white.opacity(0.045), in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.lineCard))
        }
        .buttonStyle(PressScale())
    }
}

/// The §8 pressed look: 0.98 scale over 120 ms.
struct PressScale: ButtonStyle {
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .scaleEffect(configuration.isPressed ? 0.98 : 1)
            .animation(.easeOut(duration: 0.12), value: configuration.isPressed)
    }
}

/// The primary card (`--surface-card`): the warm gradient, its gold border and inset highlight.
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

/// The hero band (`--surface-band`): the ember radial over `#14100b`.
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

extension View {
    func primaryCard() -> some View { modifier(PrimaryCard()) }
    func quietCard() -> some View { modifier(QuietCard()) }
    func heroBand() -> some View { modifier(HeroBand()) }
}

/// A tag (`.pill`): 22 pt capsule, 11 pt 700 uppercase.
struct TagView: View {
    let text: String
    let color: Color

    var body: some View {
        Text(text.uppercased())
            .font(HomeType.micro).tracking(1)
            .foregroundStyle(color)
            .padding(.horizontal, 8)
            .frame(minHeight: 22)
            .background(color.opacity(0.12), in: .rect(cornerRadius: FSRadius.pill))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.pill).stroke(color.opacity(0.35)))
    }
}

/// The offline strip (§5.0): 36 pt, `#11151f`, wifi-off glyph, the stamp.
struct OfflineStrip: View {
    let stamp: String

    var body: some View {
        HStack(spacing: 8) {
            Image(systemName: "wifi.slash").font(.system(size: 14)).accessibilityHidden(true)
            Text(Copy.offlineStrip(time: stamp)).font(HomeType.meta)
            Spacer(minLength: 0)
        }
        .foregroundStyle(FSColor.nav)
        .padding(.horizontal, 16)
        .frame(minHeight: 36)
        .background(Color(tokenHex: 0x11151F))
        .overlay(alignment: .bottom) { FSColor.line.frame(height: 1) }
        .accessibilityElement(children: .combine)
    }
}
