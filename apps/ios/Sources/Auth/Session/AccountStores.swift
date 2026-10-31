import Foundation
import Synchronization

/// One small string the app keeps between launches on this device. Neither value here is a
/// secret (a character key, a device id), so UserDefaults holds them; the token stays in the
/// Keychain (`TokenStore`).
public protocol StoredValue: Sendable {
    func load() -> String?
    func save(_ value: String?)
}

public struct DefaultsValue: StoredValue {
    private let key: String
    // UserDefaults is documented thread-safe.
    private nonisolated(unsafe) let defaults: UserDefaults

    public init(key: String, defaults: UserDefaults = .standard) {
        self.key = key
        self.defaults = defaults
    }

    public func load() -> String? { defaults.string(forKey: key) }

    public func save(_ value: String?) {
        if let value { defaults.set(value, forKey: key) } else { defaults.removeObject(forKey: key) }
    }
}

public final class MemoryValue: StoredValue {
    private let stored: Mutex<String?>

    public init(_ value: String? = nil) {
        stored = Mutex(value)
    }

    public func load() -> String? { stored.withLock { $0 } }

    public func save(_ value: String?) { stored.withLock { $0 = value } }
}

public extension DefaultsValue {
    /// The character last picked in the switcher (spec §3.4: it persists per device).
    static let pickedCharacter = DefaultsValue(key: "gg.foreversixty.picked-character")
    /// This install's device id from the sign-in exchange, so Devices can mark "This iPhone".
    static let deviceID = DefaultsValue(key: "gg.foreversixty.device-id")
}
