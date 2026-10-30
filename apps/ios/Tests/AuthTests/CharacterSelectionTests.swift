import Foundation
@testable import Auth
import Models
import Networking
import TestSupport
import Testing

@Suite("Character selection and ages")
struct CharacterSelectionTests {
    private func me() throws -> Me {
        let data = try FixtureLocation.data("me.json")
        return try WireFormat.decoder().decode(Envelope<Me>.self, from: data).data!
    }

    @Test("the picked character wins, then the main, then the first")
    func order() throws {
        let me = try me()
        let other = try #require(me.characters.first { $0.key != me.mainCharacterKey })
        #expect(CharacterSelection.current(in: me, picked: other.key)?.key == other.key)
        #expect(CharacterSelection.current(in: me, picked: "us/pvp/gone")?.key == me.mainCharacterKey)
        #expect(CharacterSelection.current(in: me, picked: nil)?.key == me.mainCharacterKey)
    }

    @Test("ages read in one unit", arguments: [
        (30.0, RelativeAge.now), (300, .minutes(5)), (7200, .hours(2)), (3 * 86400, .days(3)), (-10, .now),
    ])
    func ages(seconds: Double, age: RelativeAge) {
        let now = Date(timeIntervalSince1970: 1_000_000)
        #expect(RelativeAge(from: now.addingTimeInterval(-seconds), to: now) == age)
    }
}
