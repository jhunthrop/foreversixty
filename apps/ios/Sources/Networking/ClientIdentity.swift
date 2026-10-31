import Foundation

/// Who is calling, sent on every API request (docs/contracts/app-v1.md "Client identity"):
/// `X-App: ios/<version> (<build>)` and `X-App-Install: <install id>`. The install id is a
/// random UUID made on first launch. It is tied to no user and dies with the app, so a
/// reinstall starts a new one.
public struct ClientIdentity: Sendable, Equatable {
    public static let appHeader = "X-App"
    public static let installHeader = "X-App-Install"
    public static let platform = "ios"

    public let version: String
    public let build: String
    public let installID: String

    public init(version: String, build: String, installID: String) {
        self.version = version
        self.build = build
        self.installID = installID
    }

    /// `ios/1.2.0 (42)`.
    public var appHeaderValue: String {
        "\(Self.platform)/\(version) (\(build))"
    }

    /// The bundle's own version and build (apps/VERSION, through Version.xcconfig).
    public static func bundled(_ bundle: Bundle = .main, installID: String) -> ClientIdentity {
        let info = bundle.infoDictionary ?? [:]
        return ClientIdentity(
            version: info["CFBundleShortVersionString"] as? String ?? "0.0.0",
            build: info["CFBundleVersion"] as? String ?? "0",
            installID: installID
        )
    }
}

/// The install id, kept in UserDefaults on purpose: UserDefaults is deleted with the app,
/// where the Keychain would survive a reinstall and make the id a durable device identifier.
public enum InstallID {
    public static let defaultsKey = "gg.foreversixty.install-id"

    /// The stored id, or a new random one stored on first call.
    public static func current(in defaults: UserDefaults = .standard) -> String {
        if let existing = defaults.string(forKey: defaultsKey), UUID(uuidString: existing) != nil {
            return existing
        }
        let fresh = UUID().uuidString.lowercased()
        defaults.set(fresh, forKey: defaultsKey)
        return fresh
    }
}
