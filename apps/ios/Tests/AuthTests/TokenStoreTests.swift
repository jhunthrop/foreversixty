import Foundation
import Auth
import Testing

/// The Keychain store is tested in AppTests, inside the hosted app.
@Suite("Token stores")
struct TokenStoreTests {
    @Test("in-memory store round-trips and clears")
    func inMemory() throws {
        let store = InMemoryTokenStore()
        #expect(try store.load() == nil)
        try store.save("fsd_a")
        #expect(store.bearerToken() == "fsd_a")
        try store.clear()
        #expect(try store.load() == nil)
    }
}
