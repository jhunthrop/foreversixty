@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// Fixture mode answers a sign-up write with the tapped answer in place (the fixture app's answers move, as the server's do).
@Suite("Fixture sign-up echo")
struct SignupEchoTests {
    private func put(_ body: String) async throws -> RaidView {
        let transport = try #require(FixtureTransport.bundled())
        var request = URLRequest(url: URL(string: "https://api.foreversixty.gg/v1/guilds/2/raids/31/signups/us/pvp/gorvash")!)
        request.httpMethod = "PUT"
        request.httpBody = Data(body.utf8)
        let (data, _) = try await transport.send(request)
        let answer = try WireFormat.decoder().decode(Envelope<SignupAnswer>.self, from: data).data
        return try #require(answer).raid
    }

    @Test("Late with its minutes lands on the character's sign-up and viewer row")
    func late() async throws {
        let raid = try await put(#"{"state":"late","late_minutes":30}"#)
        #expect(raid.signups.first { $0.characterKey == "us/pvp/gorvash" }?.state == "late")
        #expect(raid.signups.first { $0.characterKey == "us/pvp/gorvash" }?.lateMinutes == 30)
        #expect(raid.viewer.characters.first { $0.characterKey == "us/pvp/gorvash" }?.state == "late")
    }

    @Test("Out lands too")
    func out() async throws {
        let raid = try await put(#"{"state":"out"}"#)
        #expect(raid.viewer.characters.first?.state == "out")
    }
}
