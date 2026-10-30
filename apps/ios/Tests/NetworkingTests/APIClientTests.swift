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

    /// Sample values for a route's `{name}` segments; the CDN fixtures are for one spec and class.
    private static let sampleParams = ["spec": "warrior-fury", "class": "warrior"]

    private static func samplePath(_ pattern: String) -> String {
        sampleParams.merging(["build": "1.60.1.70658"]) { first, _ in first }
            .reduce(pattern) { $0.replacingOccurrences(of: "{\($1.key)}", with: $1.value) }
            .replacing(/\{[a-z]+\}/, with: "1")
    }

    @Test("every catalogued route reads a fixture lane C wrote", arguments: FixtureCatalog.app)
    func catalogFilesExist(route: FixtureRoute) throws {
        let transport = FixtureTransport(directory: FixtureLocation.directory)
        let path = Self.samplePath(route.pattern)
        guard route.fixture != nil else {
            #expect(transport.fileName(method: route.method, path: path) == nil)
            return
        }
        let file = try #require(transport.fileName(method: route.method, path: path), "\(route.pattern) unmatched")
        let url = FixtureLocation.directory.appending(path: file)
        #expect(FileManager.default.fileExists(atPath: url.path), "missing \(file)")
    }

    @Test("a success variant reads <endpoint>.<variant>.json")
    func namedVariant() async throws {
        let transport = FixtureTransport(directory: FixtureLocation.directory, variants: ["me": .named("empty")])
        let me = try await APIClient(baseURL: base, transport: transport).send(.me)
        #expect(me.characters.isEmpty)
    }

    @Test("an error variant answers its status with lane C's error body")
    func errorVariant() async {
        let transport = FixtureTransport(
            directory: FixtureLocation.directory,
            variants: ["character.card": .error("not-found", status: 404)]
        )
        let client = APIClient(baseURL: base, transport: transport)
        do {
            _ = try await client.send(Endpoint<Pong>(method: .get, path: "/v1/characters/us/pvp/nobody"))
            Issue.record("expected a 404")
        } catch {
            #expect(error.code == "not_found")
        }
    }

    @Test("the shared unauthorized fixture is a 401 on any route")
    func unauthorizedVariant() async {
        let transport = FixtureTransport(directory: FixtureLocation.directory, variants: ["me": .unauthorized])
        await #expect(throws: APIError.unauthorized(requestID: "9f2c41d07ab3e655")) {
            try await APIClient(baseURL: base, transport: transport).send(.me)
        }
    }

    @Test("a 204 route answers an empty body")
    func noContentRoute() async throws {
        let client = APIClient(baseURL: base, transport: FixtureTransport(directory: FixtureLocation.directory))
        let empty = try await client.send(Endpoint<Empty>(method: .post, path: "/v1/sims/7/seen"))
        #expect(empty == Empty())
    }

    @Test("CDN fixtures are bare JSON, read through the static client")
    func cdnBare() async throws {
        let site = URL(string: "https://site.example.test")!
        let client = StaticDataClient(baseURL: site, transport: FixtureTransport(directory: FixtureLocation.directory))
        let active = try await client.activeBuild()
        #expect(!active.build.isEmpty)
        let classes = try await client.get([ClassInfo].self, path: "/data/\(active.build)/classes.json")
        #expect(classes.allSatisfy { !$0.slug.isEmpty })
    }
}
