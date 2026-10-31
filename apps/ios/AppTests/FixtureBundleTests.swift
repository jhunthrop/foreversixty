@testable import ForeverSixty
import Foundation
import Networking
import Testing

/// Fixture mode reads the contract fixtures from the app bundle. A folder reference keeps its
/// on-disk name (`app/`), so every reader goes through `FixtureTransport.bundledFolder`.
@Suite("Fixture mode")
struct FixtureBundleTests {
    @Test("the app bundle carries the contract fixtures where fixture mode reads them")
    func bundledFixtures() {
        #expect(FixtureTransport.bundled() != nil)
        #expect(FixtureTransport.bundled()?.fileName(method: .get, path: "/v1/me") == "me.json")
        #expect(Bundle.main.url(forResource: "me", withExtension: "json", subdirectory: FixtureTransport.bundledFolder) != nil)
    }

    @Test("fixture sign-in seeds the device id from the bundled exchange")
    func fixtureDeviceID() {
        #expect(AccountStorage.fixtureDeviceID(.main) != nil)
    }
}
