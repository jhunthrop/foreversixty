import Foundation
@testable import Networking
import TestSupport
import Testing

/// Records every gate answer the transport reports.
final class RecordingSink: UpdateRequiredSink, @unchecked Sendable {
    private let lock = NSLock()
    private var seen: [URL?] = []
    var answers: [URL?] { lock.withLock { seen } }
    func updateRequired(storeURL: URL?) { lock.withLock { seen.append(storeURL) } }
}

@Suite("Client identity and the update gate")
struct ClientIdentityTests {
    let identity = ClientIdentity(version: "1.2.0", build: "42", installID: "0b6f6b8e-0000-4000-8000-000000000001")
    let base = URL(string: "https://api.example.test")!
    static let gateBody = #"{"ok":false,"data":null,"error":{"code":"update_required","message":"Update the app.","fields":{"store_url":"https://apps.apple.com/app/id1"}},"request_id":"r"}"#

    private func send(_ transport: RecordingTransport, sink: RecordingSink? = nil) async -> APIError? {
        let client = APIClient(
            baseURL: base,
            transport: IdentifiedTransport(transport, identity: identity, sink: sink),
            tokens: FixedToken(token: nil)
        )
        do {
            _ = try await client.send(.publicPing)
            return nil
        } catch {
            return error
        }
    }

    @Test("X-App reads ios/<version> (<build>)")
    func appHeaderValue() {
        #expect(identity.appHeaderValue == "ios/1.2.0 (42)")
    }

    @Test("every request carries X-App and X-App-Install")
    func sendsHeaders() async throws {
        let transport = RecordingTransport(body: #"{"ok":true,"data":{"value":1},"error":null,"request_id":"r"}"#)
        _ = await send(transport)
        let request = try #require(transport.requests.first)
        #expect(request.value(forHTTPHeaderField: "X-App") == "ios/1.2.0 (42)")
        #expect(request.value(forHTTPHeaderField: "X-App-Install") == identity.installID)
    }

    @Test("a 426 update_required is its own error and reaches the gate with the store page")
    func gateAnswer() async {
        let sink = RecordingSink()
        let error = await send(RecordingTransport(status: 426, body: Self.gateBody), sink: sink)
        let store = URL(string: "https://apps.apple.com/app/id1")!
        #expect(error == .updateRequired(storeURL: store))
        #expect(sink.answers == [store])
    }

    @Test("a gate answer with a non-https store link keeps the gate but drops the link")
    func gateWithoutHTTPS() async {
        let body = Self.gateBody.replacingOccurrences(of: "https://apps.apple.com/app/id1", with: "javascript:alert(1)")
        let sink = RecordingSink()
        let error = await send(RecordingTransport(status: 426, body: body), sink: sink)
        #expect(error == .updateRequired(storeURL: nil))
        #expect(sink.answers == [nil])
    }

    @Test("another 426, or update_required on another status, is not the gate")
    func notTheGate() async {
        let sink = RecordingSink()
        let other = #"{"ok":false,"data":null,"error":{"code":"upgrade","message":"No."},"request_id":"r"}"#
        _ = await send(RecordingTransport(status: 426, body: other), sink: sink)
        _ = await send(RecordingTransport(status: 400, body: Self.gateBody), sink: sink)
        #expect(sink.answers.isEmpty)
    }

    @Test("the install id is made once, then reused; a fresh store makes a new one")
    func installID() throws {
        let suite = "install-id-\(UUID().uuidString)"
        let defaults = try #require(UserDefaults(suiteName: suite))
        defer { defaults.removePersistentDomain(forName: suite) }
        let first = InstallID.current(in: defaults)
        #expect(UUID(uuidString: first) != nil)
        #expect(InstallID.current(in: defaults) == first)
        defaults.removePersistentDomain(forName: suite)
        #expect(InstallID.current(in: defaults) != first)
    }
}
