import DesignSystem
import SwiftUI

// The site's shared controls drawn natively (spec §4.5, tenet 10). One definition each; a
// feature never draws its own variant.

/// The `.pill` / `.tag`: a 22 pt capsule, 11 pt bold uppercase, tinted by `color`.
struct Tag: View {
    let text: String
    let color: Color
    var systemImage: String?

    static let height: CGFloat = 22
    static let fillOpacity = 0.12
    static let lineOpacity = 0.35

    var body: some View {
        HStack(spacing: 4) {
            if let systemImage {
                Image(systemName: systemImage).font(.system(size: 9, weight: .bold)).accessibilityHidden(true)
            }
            Text(text.uppercased()).font(Theme.Typography.micro).tracking(1.2)
        }
        .foregroundStyle(color)
        .padding(.horizontal, Theme.Spacing.sm)
        .frame(minHeight: Self.height)
        .background(color.opacity(Self.fillOpacity), in: .rect(cornerRadius: Theme.Radius.pill))
        .overlay(RoundedRectangle(cornerRadius: Theme.Radius.pill).stroke(color.opacity(Self.lineOpacity)))
        .fixedSize()
    }
}

/// The secondary button: a 4.5% white fill and the gold 17% border, 48 tall, its label in
/// sentence case (the mock's `Back to {name}'s list`, `Try again`).
struct SecondaryButton: View {
    let title: String
    let action: () -> Void

    static let height: CGFloat = 48

    var body: some View {
        Button(action: action) {
            Text(title)
                .font(Theme.Typography.bodyStrong)
                .foregroundStyle(Theme.Palette.strong)
                .frame(maxWidth: .infinity, minHeight: Self.height)
        }
        .buttonStyle(RowPressStyle())
        .background(Theme.Palette.controlFill, in: .rect(cornerRadius: Theme.Radius.control))
        .overlay(RoundedRectangle(cornerRadius: Theme.Radius.control).stroke(Theme.Palette.lineCard))
    }
}

/// A list row's press: the 10% gold wash while held (spec §8).
struct RowPressStyle: ButtonStyle {
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .background(configuration.isPressed ? Theme.Palette.pressedWash : .clear)
    }
}

extension View {
    /// A chip or pill's ground: the capsule filled with `fill` and its 1 pt border inside the edge.
    func pillSurface(fill: Color, line: Color) -> some View {
        background { Capsule().fill(fill) }
            .overlay { Capsule().strokeBorder(line, lineWidth: 1) }
            .contentShape(.capsule)
    }
}
