import Foundation
import Models
import TestSupport
@testable import Networking
import Testing

struct Pong: Decodable, Sendable, Equatable {
    let value: Int
}

extension Endpoint where Response == Pong {
    static let ping = Endpoint(method: .get, path: "/v1/ping", query: [URLQueryItem(name: "q", value: "1")])
    static let publicPing = Endpoint(method: .get, path: "/v1/ping", authenticated: false)
}

@Suite("APIClient")
struct APIClientTests {
    let base = URL(string: "https://api.example.test")!

    private func client(_ transport: RecordingTransport, token: String? = "fsd_abc") -> APIClient {
        APIClient(baseURL: base, transport: transport, tokens: FixedToken(token: token))
    }

    @Test("unwraps data from an ok envelope")
    func unwrapsData() async throws {
        let transport = RecordingTransport(body: #"{"ok":true,"data":{"value":7},"error":null,"request_id":"r1"}"#)
        let pong = try await client(transport).send(.ping)
        #expect(pong == Pong(value: 7))
    }

    @Test("builds the URL, method and Accept header")
    func buildsRequest() async throws {
        let transport = RecordingTransport(body: #"{"ok":true,"data":{"value":1},"error":null,"request_id":"r"}"#)
        _ = try await client(transport).send(.ping)
        let request = try #require(transport.requests.first)
        #expect(request.url?.absoluteString == "https://api.example.test/v1/ping?q=1")
        #expect(request.httpMethod == "GET")
        #expect(request.value(forHTTPHeaderField: "Accept") == "application/json")
    }

    @Test("sends the device bearer token")
    func sendsBearer() async throws {
        let transport = RecordingTransport(body: #"{"ok":true,"data":{"value":1},"error":null,"request_id":"r"}"#)
        _ = try await client(transport).send(.ping)
        #expect(transport.requests.first?.value(forHTTPHeaderField: "Authorization") == "Bearer fsd_abc")
    }

    @Test("an unauthenticated endpoint never carries the token")
    func omitsBearerWhenUnauthenticated() async throws {
        let transport = RecordingTransport(body: #"{"ok":true,"data":{"value":1},"error":null,"request_id":"r"}"#)
        _ = try await client(transport).send(.publicPing)
        #expect(transport.requests.first?.value(forHTTPHeaderField: "Authorization") == nil)
    }

    @Test("an ok:false envelope is a typed server error")
    func serverError() async {
        let body = #"{"ok":false,"data":null,"error":{"code":"invalid","message":"Bad name.","fields":{"name":"too long"}},"request_id":"r9"}"#
        let transport = RecordingTransport(status: 422, body: body)
        await #expect(throws: APIError.server(
            status: 422,
            body: ErrorBody(code: "invalid", message: "Bad name.", fields: ["name": "too long"]),
            requestID: "r9"
        )) {
            try await client(transport).send(.ping)
        }
    }

    @Test("401 is unauthorized with the request id")
    func unauthorized() async {
        let body = #"{"ok":false,"data":null,"error":{"code":"unauthorized","message":"Sign in."},"request_id":"r4"}"#
        let transport = RecordingTransport(status: 401, body: body)
        await #expect(throws: APIError.unauthorized(requestID: "r4")) {
            try await client(transport, token: nil).send(.ping)
        }
    }

    @Test("a non-envelope error body keeps its status")
    func unexpectedStatus() async {
        let transport = RecordingTransport(status: 502, body: "<html>bad gateway</html>")
        await #expect(throws: APIError.unexpectedStatus(502)) {
            try await client(transport).send(.ping)
        }
    }

    @Test("a payload that does not match its model is a decoding error")
    func decodingError() async {
        let transport = RecordingTransport(body: #"{"ok":true,"data":{"value":"seven"},"error":null,"request_id":"r"}"#)
        await #expect {
            try await client(transport).send(.ping)
        } throws: { error in
            guard case .decoding = error as? APIError else { return false }
            return true
        }
    }

    @Test("an empty 204 body is Empty")
    func emptyBody() async throws {
        let transport = RecordingTransport(status: 204, body: "")
        let empty = try await client(transport).send(Endpoint<Empty>(method: .delete, path: "/v1/x"))
        #expect(empty == Empty())
    }

    @Test("a JSON body is snake_case with a Content-Type")
    func jsonBody() async throws {
        struct Body: Encodable { let deviceName: String }
        let transport = RecordingTransport(body: #"{"ok":true,"data":{"value":1},"error":null,"request_id":"r"}"#)
        let endpoint = try Endpoint<Pong>.json(.post, "/v1/ping", body: Body(deviceName: "Phone"))
        _ = try await client(transport).send(endpoint)
        let request = try #require(transport.requests.first)
        #expect(request.value(forHTTPHeaderField: "Content-Type") == "application/json")
        #expect(request.httpBody.flatMap { String(data: $0, encoding: .utf8) } == #"{"device_name":"Phone"}"#)
    }

    @Test("fixture mode answers /v1/me from the contract fixture")
    func fixtureMe() async throws {
        let client = APIClient(baseURL: base, transport: FixtureTransport(directory: FixtureLocation.directory))
        let me = try await client.send(.me)
        #expect(!me.characters.isEmpty)
    }

    @Test("fixture mode fails fast on a route with no fixture")
    func fixtureMissing() async {
        let client = APIClient(baseURL: base, transport: FixtureTransport(directory: FixtureLocation.directory))
        await #expect(throws: APIError.missingFixture("GET /v1/ping")) {
            try await client.send(.ping)
        }
    }

    @Test("every catalogued fixture file exists")
    func catalogFilesExist() {
        for route in FixtureCatalog.app {
            let url = FixtureLocation.directory.appending(path: route.file)
            #expect(FileManager.default.fileExists(atPath: url.path), "missing \(route.file)")
        }
    }
}
