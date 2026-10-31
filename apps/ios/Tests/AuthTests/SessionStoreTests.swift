import Foundation
@testable import Auth
import Models
import Networking
import TestSupport
import Testing

/// F1's test list (one list, both platforms: Android's SessionManagerTest mirrors it).
@MainActor
@Suite("Session")
struct SessionStoreTests {
    let base = URL(string: "https://api.example.test")!

    private func store(
        transport: any HTTPTransport = FixtureTransport(directory: FixtureLocation.directory),
        token: String? = nil,
        web: any WebAuthenticating = FixtureWebAuthenticator(),
        deviceID: MemoryValue = MemoryValue()
    ) -> (SessionStore, SessionTokenSource) {
        let tokens = SessionTokenSource(InMemoryTokenStore(token: token))
        let client = APIClient(baseURL: base, transport: transport, tokens: tokens)
        let store = SessionStore(client: client, tokens: tokens, web: web, picked: MemoryValue(), deviceID: deviceID)
        return (store, tokens)
    }

    private func fixtures(_ variants: [String: FixtureVariant]) -> FixtureTransport {
        FixtureTransport(directory: FixtureLocation.directory, variants: variants)
    }

    private func fixture(_ file: String, status: Int) throws -> RecordingTransport {
        RecordingTransport(status: status, body: String(decoding: try FixtureLocation.data(file), as: UTF8.self))
    }

    @Test("starts signed out with no token and signed in with one")
    func startPhase() {
        #expect(store().0.phase == .signedOut(nil))
        #expect(store(token: "fsd_x").0.phase == .signedIn(nil))
    }

    @Test("sign-in exchanges the code, keeps the device id and imports the account")
    func signInImports() async throws {
        let deviceID = MemoryValue()
        let (session, tokens) = store(deviceID: deviceID)
        await session.signIn(deviceName: "iPhone")
        let me = try #require(session.phase.me)
        #expect(!me.characters.isEmpty)
        #expect(deviceID.load() == "dvc_8Qk2nXw4")
        #expect(tokens.bearerToken() != nil)
    }

    @Test("a spent code shows the expired line and keeps no token")
    func signInExpired() async {
        let (session, tokens) = store(transport: fixtures([
            "app.signin.exchange": .error("invalid-grant", status: 400),
        ]))
        await session.signIn(deviceName: "iPhone")
        #expect(session.phase == .signedOut(.signInFailed(.expired)))
        #expect(tokens.bearerToken() == nil)
    }

    @Test("cancelling the browser returns to signed out with no line")
    func signInCancelled() async {
        let (session, _) = store(web: FailingWeb(error: .cancelled))
        await session.signIn(deviceName: "iPhone")
        #expect(session.phase == .signedOut(nil))
    }

    @Test("sign-in failures map to the spec's three lines", arguments: [
        (AuthError.stateMismatch, SignInFailure?.some(.battleNet)),
        (.callbackRejected("access_denied"), .battleNet),
        (.exchange(.transport("offline")), .offline),
        (.exchange(.server(status: 400, body: ErrorBody(code: "invalid_grant", message: "x"), requestID: nil)), .expired),
        (.exchange(.server(status: 400, body: ErrorBody(code: "invalid", message: "x"), requestID: nil)), .battleNet),
        (.cancelled, nil),
    ])
    func failureLines(error: AuthError, line: SignInFailure?) {
        #expect(SignInFailure(error) == line)
    }

    @Test("401 unauthorized and 401 session_expired both clear the token and say so", arguments: [
        "error-unauthorized.json", "error-session-expired.json",
    ])
    func deadToken(file: String) async throws {
        let (session, tokens) = store(transport: try fixture(file, status: 401), token: "fsd_dead")
        await session.refresh()
        #expect(session.phase == .signedOut(.sessionEnded))
        #expect(tokens.bearerToken() == nil)
    }

    @Test("a 401 on any other request reaches the session through the token source")
    func deadTokenElsewhere() async throws {
        let (session, tokens) = store(transport: try fixture("error-session-expired.json", status: 401), token: "fsd_dead")
        let watcher = Task { await session.watchForDeadTokens() }
        defer { watcher.cancel() }
        _ = try? await session.devices()
        for _ in 0..<50 where session.phase.isSignedIn { await Task.yield() }
        #expect(session.phase == .signedOut(.sessionEnded))
        #expect(tokens.bearerToken() == nil)
    }

    @Test("sign out forgets the token and shows no line")
    func signOut() async {
        let (session, tokens) = store(token: "fsd_x")
        await session.signOut()
        #expect(session.phase == .signedOut(nil))
        #expect(tokens.bearerToken() == nil)
    }

    @Test("this device is listed first and cannot be revoked")
    func deviceRows() async throws {
        let (session, _) = store(token: "fsd_x", deviceID: MemoryValue("dvc_8Qk2nXw4"))
        let now = try #require(ISO8601DateFormatter().date(from: "2026-10-30T21:41:00Z"))
        let rows = try await session.devices(now: now)
        #expect(rows.first?.isThisDevice == true)
        #expect(rows.first?.canRevoke == false)
        #expect(rows.dropFirst().allSatisfy { $0.canRevoke })
        #expect(rows[1].meta == .lastUsed(.hours(2)))
    }

    @Test("revoke answers done, or failed for a device that is not yours")
    func revoke() async throws {
        #expect(await store(token: "fsd_x").0.revoke("dvc_3Hm7pLa9") == .done)
        let missing = store(transport: try fixture("devices.revoke.error-not-found.json", status: 404), token: "fsd_x")
        #expect(await missing.0.revoke("dvc_nope") == .failed)
    }

    @Test("deleting the account signs out with the date the data goes")
    func deleteAccount() async throws {
        let (session, tokens) = store(token: "fsd_x")
        #expect(await session.deleteAccount() == .done)
        let completeBy = try #require(ISO8601DateFormatter().date(from: "2026-11-29T21:41:00Z"))
        #expect(session.phase == .signedOut(.accountDeleted(completeBy: completeBy)))
        #expect(tokens.bearerToken() == nil)
    }

    @Test("the deletion sends the confirmation phrase")
    func deletionPhrase() throws {
        let body = try #require(try Endpoint.deleteAccount(.confirmed).body)
        #expect(String(decoding: body, as: UTF8.self) == #"{"confirm":"delete my account"}"#)
    }

    @Test("a refused deletion keeps the session")
    func deleteRefused() async throws {
        let (session, _) = store(transport: try fixture("me.delete.error-invalid.json", status: 400), token: "fsd_x")
        #expect(await session.deleteAccount() == .failed)
        #expect(session.phase.isSignedIn)
    }

    @Test("403 browser_session_required opens the site instead of failing")
    func browserOnly() async {
        let body = #"{"ok":false,"data":null,"error":{"code":"browser_session_required","message":"this can only be done on foreversixty.gg in a browser"},"request_id":"r"}"#
        let (session, _) = store(transport: RecordingTransport(status: 403, body: body), token: "fsd_x")
        #expect(await session.deleteAccount() == .needsBrowser)
        #expect(await session.revoke("dvc_3Hm7pLa9") == .needsBrowser)
    }
}

private struct FailingWeb: WebAuthenticating {
    let error: AuthError

    func authenticate(url _: URL, callbackHost _: String, callbackPath _: String) async throws(AuthError) -> URL {
        throw error
    }
}
