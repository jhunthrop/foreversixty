import Foundation
@testable import Auth
import Models
import Networking
import TestSupport
import Testing

/// A browser sheet that answers with a callback built from the start URL it was given.
final class FakeWebAuthenticator: WebAuthenticating, @unchecked Sendable {
    enum Answer {
        case echoState(code: String)
        case wrongState
        case error(String)
        case fail(AuthError)
    }

    private let answer: Answer
    private let lock = NSLock()
    private var opened: [URL] = []

    init(_ answer: Answer) {
        self.answer = answer
    }

    var startURLs: [URL] { lock.withLock { opened } }

    func authenticate(url: URL, callbackHost: String, callbackPath: String) async throws(AuthError) -> URL {
        lock.withLock { opened.append(url) }
        let state = URLComponents(url: url, resolvingAgainstBaseURL: false)?
            .queryItems?.first { $0.name == "state" }?.value ?? ""
        var callback = URLComponents()
        callback.scheme = "https"
        callback.host = callbackHost
        callback.path = callbackPath
        switch answer {
        case let .echoState(code):
            callback.queryItems = [.init(name: "code", value: code), .init(name: "state", value: state)]
        case .wrongState:
            callback.queryItems = [.init(name: "code", value: "c"), .init(name: "state", value: "forged")]
        case let .error(message):
            callback.queryItems = [.init(name: "error", value: message), .init(name: "state", value: state)]
        case let .fail(error):
            throw error
        }
        return callback.url!
    }
}

@Suite("App sign-in")
struct AppSignInTests {
    let base = URL(string: "https://api.example.test")!

    private func signIn(
        _ web: FakeWebAuthenticator,
        store: InMemoryTokenStore = InMemoryTokenStore(),
        transport: any HTTPTransport = FixtureTransport(directory: FixtureLocation.directory)
    ) -> AppSignIn {
        AppSignIn(client: APIClient(baseURL: base, transport: transport, tokens: store), web: web, store: store)
    }

    @Test("start URL carries platform, S256 challenge and state")
    func startURL() throws {
        let pkce = PKCE(verifier: "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk")
        let url = AppAuthContract.startURL(baseURL: base, pkce: pkce, state: "s1")
        let components = try #require(URLComponents(url: url, resolvingAgainstBaseURL: false))
        let query = Dictionary(uniqueKeysWithValues: (components.queryItems ?? []).map { ($0.name, $0.value ?? "") })
        #expect(components.path == "/v1/auth/app/start")
        #expect(query == [
            "platform": "ios",
            "code_challenge": "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM",
            "code_challenge_method": "S256",
            "state": "s1",
        ])
    }

    @Test("a full sign-in stores the exchanged device token")
    func happyPath() async throws {
        let store = InMemoryTokenStore()
        let session = try await signIn(FakeWebAuthenticator(.echoState(code: "one-time")), store: store)
            .signIn(deviceName: "Test iPhone")
        #expect(session.token.hasPrefix("fsd_"))
        #expect(session.device.platform == "ios")
        #expect(session.me.user.id > 0)
        #expect(try store.load() == session.token)
    }

    @Test("the exchange posts the code with the verifier that matches the challenge")
    func exchangeBody() async throws {
        let transport = RecordingTransport(
            status: 201,
            body: String(decoding: try FixtureLocation.data("app.signin.exchange.json"), as: UTF8.self)
        )
        let web = FakeWebAuthenticator(.echoState(code: "one-time"))
        _ = try await signIn(web, transport: transport).signIn(deviceName: "Test iPhone")

        let request = try #require(transport.requests.first)
        #expect(request.url?.path == "/v1/auth/app/exchange")
        #expect(request.httpMethod == "POST")
        #expect(request.value(forHTTPHeaderField: "Authorization") == nil)
        let body = try JSONSerialization.jsonObject(with: try #require(request.httpBody)) as? [String: String]
        #expect(Set(body?.keys.map(\.self) ?? []) == ["code", "code_verifier", "device_name"])
        #expect(body?["code"] == "one-time")
        #expect(body?["device_name"] == "Test iPhone")

        let start = try #require(web.startURLs.first)
        let challenge = URLComponents(url: start, resolvingAgainstBaseURL: false)?
            .queryItems?.first { $0.name == "code_challenge" }?.value
        #expect(challenge == PKCE.challenge(for: try #require(body?["code_verifier"])))
    }

    @Test("a callback with a different state is refused and nothing is stored")
    func stateMismatch() async {
        let store = InMemoryTokenStore()
        await #expect(throws: AuthError.stateMismatch) {
            try await signIn(FakeWebAuthenticator(.wrongState), store: store).signIn(deviceName: "x")
        }
        #expect((try? store.load()) == nil)
    }

    @Test("an error callback is rejected")
    func errorCallback() async {
        await #expect(throws: AuthError.callbackRejected("access_denied")) {
            try await signIn(FakeWebAuthenticator(.error("access_denied"))).signIn(deviceName: "x")
        }
    }

    @Test("cancelling the sheet surfaces cancelled")
    func cancelled() async {
        await #expect(throws: AuthError.cancelled) {
            try await signIn(FakeWebAuthenticator(.fail(.cancelled))).signIn(deviceName: "x")
        }
    }

    @Test("a spent or mismatched code surfaces invalid_grant from lane C's fixture")
    func exchangeRefused() async {
        let transport = FixtureTransport(
            directory: FixtureLocation.directory,
            variants: ["app.signin.exchange": .error("invalid-grant", status: 400)]
        )
        do {
            _ = try await signIn(FakeWebAuthenticator(.echoState(code: "c")), transport: transport).signIn(deviceName: "x")
            Issue.record("expected the exchange to be refused")
        } catch {
            guard case let .exchange(apiError) = error else {
                Issue.record("unexpected \(error)")
                return
            }
            #expect(apiError.code == "invalid_grant")
        }
    }

    @Test("the callback lands on exactly https://foreversixty.gg/app/auth")
    func callbackPath() throws {
        let url = URL(string: "https://foreversixty.gg/app/auth?code=c&state=s")!
        #expect(try AppAuthContract.code(fromCallback: url, expectedState: "s") == "c")
    }

    @Test("a callback on another path or host is rejected", arguments: [
        "https://foreversixty.gg/app/auth/callback?code=c&state=s",
        "https://evil.example.test/app/auth?code=c&state=s",
    ])
    func wrongCallback(raw: String) {
        let url = URL(string: raw)!
        #expect(throws: AuthError.callbackRejected("unexpected callback \(url.host ?? "")\(url.path)")) {
            try AppAuthContract.code(fromCallback: url, expectedState: "s")
        }
    }

    @Test("sign out revokes the device on the server, then clears the token")
    func signOut() async throws {
        let store = InMemoryTokenStore(token: "fsd_x")
        let transport = RecordingTransport(body: String(decoding: try FixtureLocation.data("app.logout.json"), as: UTF8.self))
        try await signIn(FakeWebAuthenticator(.wrongState), store: store, transport: transport).signOut()
        let request = try #require(transport.requests.first)
        #expect(request.url?.path == "/v1/auth/app/logout")
        #expect(request.httpMethod == "POST")
        #expect(request.value(forHTTPHeaderField: "Authorization") == "Bearer fsd_x")
        #expect(try store.load() == nil)
    }

    @Test("sign out clears the token even when the server cannot be reached")
    func signOutOffline() async throws {
        let store = InMemoryTokenStore(token: "fsd_x")
        let transport = RecordingTransport(status: 500, body: "")
        try await signIn(FakeWebAuthenticator(.wrongState), store: store, transport: transport).signOut()
        #expect(try store.load() == nil)
    }

    @Test("a 401 on any call discards the stored token", arguments: ["unauthorized", "session_expired"])
    func unauthorizedDiscardsToken(code: String) async throws {
        let store = InMemoryTokenStore(token: "fsd_dead")
        let body = #"{"ok":false,"data":null,"error":{"code":"\#(code)","message":"Sign in again."},"request_id":"r"}"#
        let client = APIClient(baseURL: base, transport: RecordingTransport(status: 401, body: body), tokens: store)
        await #expect(throws: APIError.unauthorized(requestID: "r")) {
            try await client.send(.getMe)
        }
        #expect(try store.load() == nil)
    }

    @Test("billing writes answer browser_session_required, which the app sends to the site")
    func browserSessionRequired() {
        let error = APIError.server(
            status: 403,
            body: ErrorBody(code: "browser_session_required", message: "This can only be done on foreversixty.gg."),
            requestID: "r"
        )
        #expect(error.needsBrowserSession)
    }
}
