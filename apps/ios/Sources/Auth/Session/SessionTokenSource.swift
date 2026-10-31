import Foundation
import Networking

/// The token store the API client and sign-in share. It forwards to the real store and reports
/// every 401 discard on `discarded`, so `SessionStore` can show signed out the moment any
/// request learns the token is dead (`unauthorized` or `session_expired`).
public final class SessionTokenSource: TokenStore {
    public let discarded: AsyncStream<Void>
    private let base: any TokenStore
    private let continuation: AsyncStream<Void>.Continuation

    public init(_ base: any TokenStore) {
        self.base = base
        (discarded, continuation) = AsyncStream.makeStream(bufferingPolicy: .bufferingNewest(1))
    }

    public func load() throws(TokenStoreError) -> String? { try base.load() }

    public func save(_ token: String) throws(TokenStoreError) { try base.save(token) }

    public func clear() throws(TokenStoreError) { try base.clear() }

    public func discardBearerToken() {
        base.discardBearerToken()
        continuation.yield()
    }
}

/// Fixture mode's browser: answers at once with the callback the API would send, so the whole
/// sign-in (state check, exchange, import) runs against the contract fixtures, offline.
public struct FixtureWebAuthenticator: WebAuthenticating {
    public static let code = "fixture-code"

    public init() {}

    public func authenticate(url: URL, callbackHost: String, callbackPath: String) async throws(AuthError) -> URL {
        let state = URLComponents(url: url, resolvingAgainstBaseURL: false)?
            .queryItems?.first { $0.name == "state" }?.value
        var callback = URLComponents()
        callback.scheme = "https"
        callback.host = callbackHost
        callback.path = callbackPath
        callback.queryItems = [.init(name: "code", value: Self.code), .init(name: "state", value: state)]
        guard let answer = callback.url else { throw .webSession("no fixture callback") }
        return answer
    }
}
