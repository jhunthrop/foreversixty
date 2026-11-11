import SwiftUI

/// The site's buttons in round 5 (web lib/ui/button.ts): primary their flat teal with the teal edge
/// (Buy Now), secondary their blue (Upgrade Now), the accent gold outline (the reminder ask's Turn
/// on), destructive the wipe outline and quiet the rule outline, each a white or tone label in
/// sentence case on the 6 pt control radius. Full width by default; `fullWidth: false` hugs the
/// label. Pressed: 0.98 over 120 ms (spec section 8), except under Reduce Motion and on a disabled
/// button, which never scale. Disabled: the whole button at 50%, as on Android and the web.
public struct BrandButtonStyle: ButtonStyle {
    public enum Kind: Sendable { case primary, secondary, accent, destructive, quiet }

    private static let horizontalPadding: CGFloat = 18
    private static let pressedScale: CGFloat = 0.98
    private static let pressDuration: Double = 0.12
    /// A disabled button's opacity.
    public static let disabledOpacity: Double = 0.5

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
            .frame(maxWidth: fullWidth ? .infinity : nil, minHeight: height)
            .padding(.horizontal, Self.horizontalPadding)
            .foregroundStyle(foreground)
            .background(background, in: .rect(cornerRadius: Theme.Radius.control))
            .overlay(RoundedRectangle(cornerRadius: Theme.Radius.control).strokeBorder(border, lineWidth: FSStroke.hairline))
            .opacity(isEnabled ? 1 : Self.disabledOpacity)
            .scaleEffect(Self.scale(pressed: configuration.isPressed, reduceMotion: reduceMotion, enabled: isEnabled))
            .animation(.easeOut(duration: Self.pressDuration), value: configuration.isPressed)
    }

    /// The press scale: 0.98 while held, 1 under Reduce Motion or when the button is disabled.
    public nonisolated static func scale(pressed: Bool, reduceMotion: Bool, enabled: Bool) -> CGFloat {
        pressed && !reduceMotion && enabled ? pressedScale : 1
    }

    private var foreground: Color {
        switch kind {
        case .primary, .secondary: FSColor.strong
        case .quiet: FSColor.cream
        case .accent: FSColor.gold
        case .destructive: FSColor.wipe
        }
    }

    private var background: Color {
        switch kind {
        case .primary: FSColor.teal
        case .secondary: FSColor.blue
        case .accent, .destructive, .quiet: .clear
        }
    }

    private var border: Color {
        switch kind {
        case .primary: FSColor.tealEdge
        case .secondary: FSColor.blueEdge
        // The cream frame line: the rule's slate disappears on the painted ground.
        case .quiet: FSColor.frameLine
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
            HStack(spacing: FSSpace.space2) {
                if let systemImage { Image(systemName: systemImage).accessibilityHidden(true) }
                Text(title)
            }
        }
        .buttonStyle(BrandButtonStyle(.secondary, height: 48, fullWidth: fullWidth))
    }
}
