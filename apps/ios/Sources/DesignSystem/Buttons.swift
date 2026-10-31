import SwiftUI

/// The site's buttons (spec §4.5): `PRIMARY_BUTTON_FIXED` (the gold gradient) and
/// `SECONDARY_BUTTON_FIXED` (4.5% white, the gold 17% border), plus the destructive and quiet
/// kinds, and the accent secondary (gold label and border, the reminder ask's Turn on). Full width
/// by default; `fullWidth: false` hugs the label. Pressed: 0.98 over 120 ms (§8), except under
/// Reduce Motion and on a disabled button, which never scale.
public struct BrandButtonStyle: ButtonStyle {
    public enum Kind: Sendable { case primary, secondary, accent, destructive, quiet }

    private static let horizontalPadding: CGFloat = 18
    private static let pressedScale: CGFloat = 0.98
    private static let pressDuration: Double = 0.12

    private let kind: Kind
    private let height: CGFloat
    private let fullWidth: Bool

    @Environment(\.accessibilityReduceMotion) private var reduceMotion
    @Environment(\.isEnabled) private var isEnabled

    public init(_ kind: Kind, height: CGFloat = 52, fullWidth: Bool = true) {
        self.kind = kind
        self.height = height
        self.fullWidth = fullWidth
    }

    public func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .font(ScaledType.button)
            // Spec §4.5: the primary button is uppercase; a secondary reads in sentence case.
            .textCase(kind == .secondary || kind == .accent ? nil : .uppercase)
            .tracking(0.9)
            .frame(maxWidth: fullWidth ? .infinity : nil, minHeight: height)
            .padding(.horizontal, Self.horizontalPadding)
            .foregroundStyle(foreground)
            .background(background, in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(border))
            .scaleEffect(Self.scale(pressed: configuration.isPressed, reduceMotion: reduceMotion, enabled: isEnabled))
            .animation(.easeOut(duration: Self.pressDuration), value: configuration.isPressed)
    }

    /// The press scale: 0.98 while held, 1 under Reduce Motion or when the button is disabled.
    static func scale(pressed: Bool, reduceMotion: Bool, enabled: Bool) -> CGFloat {
        pressed && !reduceMotion && enabled ? pressedScale : 1
    }

    private var foreground: Color {
        switch kind {
        case .primary: FSColor.bg
        case .secondary: FSColor.strong
        case .accent: FSColor.gold
        case .destructive: FSColor.wipe
        case .quiet: FSColor.nav
        }
    }

    private var background: AnyShapeStyle {
        switch kind {
        case .primary:
            AnyShapeStyle(LinearGradient(colors: [FSColor.goldHover, FSColor.gold], startPoint: .top, endPoint: .bottom))
        case .secondary, .accent, .destructive, .quiet:
            AnyShapeStyle(GameColors.controlFill)
        }
    }

    private var border: Color {
        switch kind {
        case .primary: .clear
        case .secondary, .quiet: FSColor.lineCard
        case .accent: FSColor.gold
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

/// The secondary button, 48 pt, at its label's width or the full width, with an optional
/// leading glyph.
public struct SecondaryButton: View {
    private let title: String
    private let systemImage: String?
    private let fullWidth: Bool
    private let action: () -> Void

    public init(title: String, systemImage: String? = nil, fullWidth: Bool = false, action: @escaping () -> Void) {
        self.title = title
        self.systemImage = systemImage
        self.fullWidth = fullWidth
        self.action = action
    }

    public var body: some View {
        Button(action: action) {
            HStack(spacing: FSSpacing.space2) {
                if let systemImage { Image(systemName: systemImage).accessibilityHidden(true) }
                Text(title)
            }
        }
        .buttonStyle(BrandButtonStyle(.secondary, height: 48, fullWidth: fullWidth))
    }
}
