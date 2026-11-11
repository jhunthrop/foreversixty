import SwiftUI

/// The one load error (spec 3.4, `ui/LoadError`): a Plate with the sentence in wipe, read as an
/// alert, and a secondary Try again. Wipe ink reads on the panel's dark, never loose on the painted
/// ground, so every failing region draws this rather than its own red line. Android: `LoadError`.
public struct LoadError: View {
    private static let gap: CGFloat = 12

    private let text: String
    private let retryLabel: String
    private let retry: () -> Void

    public init(_ text: String, retryLabel: String, retry: @escaping () -> Void) {
        self.text = text
        self.retryLabel = retryLabel
        self.retry = retry
    }

    public var body: some View {
        VStack(alignment: .leading, spacing: Self.gap) {
            Text(text).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                .fixedSize(horizontal: false, vertical: true)
                .accessibilityAddTraits(.isStaticText)
            SecondaryButton(title: retryLabel, systemImage: "arrow.clockwise", action: retry)
        }
        .padding(FSSpace.space4)
        .frame(maxWidth: .infinity, alignment: .leading)
        .panel()
    }
}
