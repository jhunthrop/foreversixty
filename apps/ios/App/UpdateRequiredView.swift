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

/// The blocking screen the gate shows (mock ios-update-required): the seal, the title, why, and
/// the store button; no tabs, no way past it but the store.
struct UpdateRequiredView: View {
    let storeURL: URL
    @Environment(\.openURL) private var openURL

    /// `.seal`: 84 pt, 40 over and 8 under; `.stack .btn`: 48 high, 20 under the line.
    static let sealSize: CGFloat = 84
    static let sealTop: CGFloat = 40
    static let sealBottom: CGFloat = FSSpace.space2
    static let buttonHeight: CGFloat = 48

    var body: some View {
        ScrollView {
            VStack(spacing: 0) {
                Seal(size: Self.sealSize).padding(.top, Self.sealTop).padding(.bottom, Self.sealBottom)
                PageHeader(title: Copy.updateRequiredTitle, alignment: .center)
                Text(Copy.updateRequiredBody)
                    .type(FSType.small).foregroundStyle(FSColor.muted)
                    .multilineTextAlignment(.center)
                    .fixedSize(horizontal: false, vertical: true)
                    .frame(maxWidth: 300)
                Button(Copy.updateRequiredAction) { openURL(storeURL) }
                    .buttonStyle(FSButtonStyle(.primary))
                    .frame(minHeight: Self.buttonHeight)
                    .padding(.top, FSSpace.space5)
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .parityContent()
        }
        .background(FSColor.bg)
    }
}

#Preview("Update required") {
    UpdateRequiredView(storeURL: UpdateGate.fallbackStoreURL).preferredColorScheme(.dark)
}
