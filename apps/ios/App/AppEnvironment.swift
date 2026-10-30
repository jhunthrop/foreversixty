import Auth
import Foundation
import Networking

/// The app's dependencies, built once from Info.plist. Feature lanes read the
/// client and token store from here; tests and previews build their own.
struct AppEnvironment {
    let configuration: APIConfiguration
    let tokens: any TokenStore
    let client: APIClient
    /// Flips when the API says this build is below its minimum; RootView then blocks.
    let updateGate: UpdateGate

    @MainActor
    init(
        configuration: APIConfiguration,
        tokens: any TokenStore,
        identity: ClientIdentity,
        updateGate: UpdateGate = UpdateGate(),
        bundle: Bundle = .main
    ) {
        self.configuration = configuration
        self.tokens = tokens
        self.updateGate = updateGate
        self.client = APIClient.make(
            configuration, tokens: tokens, identity: identity, updateGate: updateGate, bundle: bundle
        )
    }

    /// Fixture mode keeps its token in memory so a fixture session never reaches the Keychain.
    @MainActor
    static func fromBundle(_ bundle: Bundle = .main) -> AppEnvironment {
        let configuration: APIConfiguration
        do {
            configuration = try APIConfiguration(infoDictionary: bundle.infoDictionary ?? [:])
        } catch {
            preconditionFailure("Build is misconfigured: \(error)")
        }
        let tokens: any TokenStore = switch configuration.mode {
        case .live: KeychainTokenStore()
        case .fixture: InMemoryTokenStore(token: "fsd_fixture")
        }
        let identity = ClientIdentity.bundled(bundle, installID: InstallID.current())
        return AppEnvironment(configuration: configuration, tokens: tokens, identity: identity, bundle: bundle)
    }

    // TODO(lane A2, F7): register the APNs device token with the push_tokens
    // endpoint after sign-in and when the system rotates it; unregister on sign-out.
}
