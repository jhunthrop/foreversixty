import Auth
import Foundation
import Testing

/// Runs in the hosted app test bundle: the Keychain refuses an unhosted
/// package test bundle on the simulator (errSecMissingEntitlement, -34018).
@Suite("Keychain token store")
struct KeychainTokenStoreTests {
    @Test("saves, overwrites and clears")
    func roundTrip() throws {
        let store = KeychainTokenStore(service: "gg.foreversixty.app.tests.\(UUID().uuidString)")
        defer { try? store.clear() }
        #expect(try store.load() == nil)
        try store.save("fsd_first")
        try store.save("fsd_second")
        #expect(try store.load() == "fsd_second")
        #expect(store.bearerToken() == "fsd_second")
        try store.clear()
        #expect(try store.load() == nil)
        try store.clear()
    }
}
