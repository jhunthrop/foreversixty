import Foundation
@testable import Networking
import Testing

@Suite("Envelope decoding")
struct EnvelopeTests {
    struct Payload: Decodable, Sendable, Equatable {
        let itemLevel: Int
        let capturedAt: Date
    }

    @Test("decodes snake_case keys and request_id")
    func snakeCase() throws {
        let json = #"{"ok":true,"data":{"item_level":40,"captured_at":"2026-11-04T21:03:00Z"},"error":null,"request_id":"abc"}"#
        let envelope = try WireFormat.decoder().decode(Envelope<Payload>.self, from: Data(json.utf8))
        #expect(envelope.ok)
        #expect(envelope.requestId == "abc")
        #expect(envelope.data?.itemLevel == 40)
        #expect(envelope.error == nil)
    }

    @Test("decodes an error envelope with fields")
    func errorBody() throws {
        let json = #"{"ok":false,"data":null,"error":{"code":"x","message":"y","fields":{"a":"b"}},"request_id":"r"}"#
        let envelope = try WireFormat.decoder().decode(Envelope<Payload>.self, from: Data(json.utf8))
        #expect(envelope.data == nil)
        #expect(envelope.error == ErrorBody(code: "x", message: "y", fields: ["a": "b"]))
    }

    @Test("parses Go timestamps with and without fractional seconds", arguments: [
        "2026-11-04T21:03:00Z",
        "2026-11-04T21:03:00.123456789Z",
        "2026-11-04T21:03:00.5+00:00",
    ])
    func timestamps(text: String) throws {
        let date = try #require(WireFormat.parseDate(text))
        let whole = try #require(WireFormat.parseDate("2026-11-04T21:03:00Z"))
        #expect(abs(date.timeIntervalSince(whole)) < 1)
    }

    @Test("rejects a non-timestamp")
    func badTimestamp() {
        #expect(WireFormat.parseDate("yesterday") == nil)
    }
}
