import SwiftUI

/// The site's buttons (spec §4.5): `PRIMARY_BUTTON_FIXED` (the gold gradient) and
/// `SECONDARY_BUTTON_FIXED` (4.5% white, the gold 17% border), plus the destructive and quiet
/// kinds. Full width by default; `fullWidth: false` hugs the label. Pressed: 0.98 over 120 ms (§8).
public struct BrandButtonStyle: ButtonStyle {
    public enum Kind: Sendable { case primary, secondary, destructive, quiet }

    private static let horizontalPadding: CGFloat = 18
    private static let pressedScale: CGFloat = 0.98
    private static let pressDuration: Double = 0.12

    private let kind: Kind
    private let height: CGFloat
    private let fullWidth: Bool

    public init(_ kind: Kind, height: CGFloat = 52, fullWidth: Bool = true) {
        self.kind = kind
        self.height = height
        self.fullWidth = fullWidth
    }

    public func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .font(ScaledType.button)
            .textCase(.uppercase)
            .tracking(0.9)
            .frame(maxWidth: fullWidth ? .infinity : nil, minHeight: height)
            .padding(.horizontal, Self.horizontalPadding)
            .foregroundStyle(foreground)
            .background(background, in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(border))
            .scaleEffect(configuration.isPressed ? Self.pressedScale : 1)
            .animation(.easeOut(duration: Self.pressDuration), value: configuration.isPressed)
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
        case .secondary, .quiet: FSColor.lineCard
        case .destructive: FSColor.wipe.opacity(0.45)
        }
    }
}

/// The primary button at its label's width, 48 pt. One per screen at most.
public struct PrimaryButton: View {
    private let title: String
    private let action: () -> Void

    public init(title: String, action: @escaping () -> Void) {
        self.title = title
        self.action = action
    }

    public var body: some View {
        Button(title, action: action).buttonStyle(BrandButtonStyle(.primary, height: 48, fullWidth: false))
    }
}

/// The secondary button at its label's width, 48 pt, with an optional leading glyph.
public struct SecondaryButton: View {
    private let title: String
    private let systemImage: String?
    private let action: () -> Void

    public init(title: String, systemImage: String? = nil, action: @escaping () -> Void) {
        self.title = title
        self.systemImage = systemImage
        self.action = action
    }

    public var body: some View {
        Button(action: action) {
            HStack(spacing: FSSpacing.space2) {
                if let systemImage { Image(systemName: systemImage).accessibilityHidden(true) }
                Text(title)
            }
        }
        .buttonStyle(BrandButtonStyle(.secondary, height: 48, fullWidth: false))
    }
}
