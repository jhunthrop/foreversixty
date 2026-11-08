import Auth
import Foundation
import Networking

/// The App Group the app and its extensions belong to (project.yml entitlements).
enum AppGroup {
    static let id = "group.gg.foreversixty.app"

    /// The shared container; nil when the process lacks the entitlement (a unit-test host
    /// built without signing), where the caller falls back to its own Application Support.
    static var containerURL: URL? {
        FileManager.default.containerURL(forSecurityApplicationGroupIdentifier: id)
    }

    /// Settings every process reads (the Live surfaces switches).
    static var defaults: UserDefaults { UserDefaults(suiteName: id) ?? .standard }
}

extension APIClient {
    /// An extension's own client, from its Info.plist (the same FSAPI keys as the app) and the
    /// token the app shares through the App Group keychain. Fixture builds answer from fixtures.
    static func appGroup(_ bundle: Bundle = .main) -> APIClient? {
        guard let configuration = try? APIConfiguration(infoDictionary: bundle.infoDictionary ?? [:]) else { return nil }
        let tokens: any TokenStore = switch configuration.mode {
        case .live: KeychainTokenStore(accessGroup: AppGroup.id)
        case .fixture: InMemoryTokenStore(token: "fsd_fixture")
        }
        let identity = ClientIdentity.bundled(bundle, installID: InstallID.current(in: AppGroup.defaults))
        return .make(configuration, tokens: tokens, identity: identity, bundle: bundle)
    }
}
