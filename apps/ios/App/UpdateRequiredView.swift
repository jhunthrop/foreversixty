import DesignSystem
import Networking
import Observation
import SwiftUI

/// Flips once any request answers the API's minimum-version gate, and then stays flipped:
/// the build is too old for the API, so nothing else in the app can work until it updates.
@MainActor
@Observable
final class UpdateGate: UpdateRequiredSink {
    /// Where the button goes when the gate sent no store link. The App Store id does not
    /// exist until Apple enrollment clears; the site links to the store from then on.
    static let fallbackStoreURL = URL(string: "https://foreversixty.gg")!

    private(set) var isRequired = false
    private(set) var storeURL: URL = UpdateGate.fallbackStoreURL

    init(isRequired: Bool = false) {
        self.isRequired = isRequired
    }

    nonisolated func updateRequired(storeURL: URL?) {
        Task { @MainActor in self.require(storeURL: storeURL) }
    }

    func require(storeURL: URL?) {
        if let storeURL { self.storeURL = storeURL }
        isRequired = true
    }
}

/// The blocking screen the gate shows: no tabs, no way past it but the store.
struct UpdateRequiredView: View {
    let storeURL: URL
    @Environment(\.openURL) private var openURL

    var body: some View {
        VStack(spacing: Theme.Spacing.md) {
            Spacer()
            Text("Update Forever Sixty")
                .font(Theme.Typography.title)
                .foregroundStyle(Theme.Palette.strong)
                .multilineTextAlignment(.center)
            Text("This version no longer works with Forever Sixty. Update to keep your character, raids and lists in sync.")
                .font(Theme.Typography.body)
                .foregroundStyle(Theme.Palette.text)
                .multilineTextAlignment(.center)
            Button {
                openURL(storeURL)
            } label: {
                Text("Open the App Store")
                    .font(Theme.Typography.meta)
                    .foregroundStyle(Theme.Palette.background)
                    .frame(maxWidth: .infinity)
            }
            .buttonStyle(.borderedProminent)
            .tint(Theme.Palette.accent)
            .controlSize(.large)
            Spacer()
        }
        .padding(Theme.Spacing.lg)
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(Theme.Palette.background)
    }
}

#Preview("Update required") {
    UpdateRequiredView(storeURL: UpdateGate.fallbackStoreURL).preferredColorScheme(.dark)
}
