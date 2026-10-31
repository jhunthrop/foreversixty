import Foundation
import Networking
import Security

/// Holds the device bearer token between launches. Conforms to Networking's
/// `BearerTokenSource` so the API client reads the token straight from it.
public protocol TokenStore: BearerTokenSource {
    func load() throws(TokenStoreError) -> String?
    func save(_ token: String) throws(TokenStoreError)
    func clear() throws(TokenStoreError)
}

public extension TokenStore {
    func bearerToken() -> String? { try? load() }

    /// A dead token is cleared; a Keychain failure here leaves it to the next 401.
    func discardBearerToken() { try? clear() }
}

public enum TokenStoreError: Error, Equatable {
    /// A Keychain call failed with this OSStatus.
    case keychain(OSStatus)
    case malformed
}

/// The production store: one generic-password item, readable only on this
/// device after its first unlock (so background refresh works), never synced.
public struct KeychainTokenStore: TokenStore {
    public static let defaultService = "gg.foreversixty.app.device-token"
    static let account = "device"

    private let service: String
    /// The keychain access group a new item is written to: the App Group id, so the widget
    /// extension's sign-up intent reads the same token (lane F14). Reads, updates and deletes
    /// search every group the process holds, so an item written before the group still loads.
    private let accessGroup: String?

    public init(service: String = Self.defaultService, accessGroup: String? = nil) {
        self.service = service
        self.accessGroup = accessGroup
    }

    public func load() throws(TokenStoreError) -> String? {
        var query = baseQuery
        query[kSecReturnData as String] = true
        query[kSecMatchLimit as String] = kSecMatchLimitOne
        var result: CFTypeRef?
        let status = SecItemCopyMatching(query as CFDictionary, &result)
        if status == errSecItemNotFound { return nil }
        guard status == errSecSuccess else { throw .keychain(status) }
        guard let data = result as? Data, let token = String(data: data, encoding: .utf8) else {
            throw .malformed
        }
        return token
    }

    public func save(_ token: String) throws(TokenStoreError) {
        let data = Data(token.utf8)
        let update = [kSecValueData as String: data]
        let status = SecItemUpdate(baseQuery as CFDictionary, update as CFDictionary)
        if status == errSecSuccess { return }
        guard status == errSecItemNotFound else { throw .keychain(status) }
        var add = baseQuery
        add[kSecValueData as String] = data
        add[kSecAttrAccessible as String] = kSecAttrAccessibleAfterFirstUnlockThisDeviceOnly
        if let accessGroup { add[kSecAttrAccessGroup as String] = accessGroup }
        var added = SecItemAdd(add as CFDictionary, nil)
        if added == errSecMissingEntitlement, accessGroup != nil {
            // A build without the App Group (a test host signed without it) keeps the token private.
            add.removeValue(forKey: kSecAttrAccessGroup as String)
            added = SecItemAdd(add as CFDictionary, nil)
        }
        guard added == errSecSuccess else { throw .keychain(added) }
    }

    public func clear() throws(TokenStoreError) {
        let status = SecItemDelete(baseQuery as CFDictionary)
        guard status == errSecSuccess || status == errSecItemNotFound else { throw .keychain(status) }
    }

    private var baseQuery: [String: Any] {
        [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrService as String: service,
            kSecAttrAccount as String: Self.account,
        ]
    }
}

/// For previews, fixture mode and tests.
public final class InMemoryTokenStore: TokenStore, @unchecked Sendable {
    private let lock = NSLock()
    private var token: String?

    public init(token: String? = nil) {
        self.token = token
    }

    public func load() throws(TokenStoreError) -> String? {
        lock.withLock { token }
    }

    public func save(_ token: String) throws(TokenStoreError) {
        lock.withLock { self.token = token }
    }

    public func clear() throws(TokenStoreError) {
        lock.withLock { token = nil }
    }
}
