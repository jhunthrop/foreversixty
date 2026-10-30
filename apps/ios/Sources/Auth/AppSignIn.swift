import Foundation
import Models
import Networking

/// Battle.net sign-in for the app: PKCE, the system browser, the code
/// exchange, and the device token into the store.
public struct AppSignIn: Sendable {
    private let client: APIClient
    private let web: any WebAuthenticating
    private let store: any TokenStore
    private let random: any RandomBytes

    public init(
        client: APIClient,
        web: any WebAuthenticating,
        store: any TokenStore,
        random: any RandomBytes = SystemRandomBytes()
    ) {
        self.client = client
        self.web = web
        self.store = store
        self.random = random
    }

    /// Runs the whole flow and returns the new session; the token is already stored.
    public func signIn(deviceName: String) async throws(AuthError) -> AppSession {
        let pkce = try PKCE.generate(random: random)
        let state = try random.bytes(16).base64URLEncoded()
        let start = AppAuthContract.startURL(baseURL: client.baseURL, pkce: pkce, state: state)
        let callback = try await web.authenticate(
            url: start,
            callbackHost: AppAuthContract.callbackHost,
            callbackPath: AppAuthContract.callbackPath
        )
        let code = try AppAuthContract.code(fromCallback: callback, expectedState: state)
        let session = try await exchange(code: code, verifier: pkce.verifier, deviceName: deviceName)
        do throws(TokenStoreError) {
            try store.save(session.token)
        } catch {
            throw .storage(error)
        }
        return session
    }

    /// Revokes the device on the server (`POST /v1/auth/app/logout`), then forgets the token.
    /// The local token goes even when the server cannot be reached: the player asked to sign
    /// out, and the site's device list can still revoke a row left behind.
    public func signOut() async throws(AuthError) {
        if store.bearerToken() != nil {
            _ = try? await client.send(.appLogout)
        }
        do throws(TokenStoreError) {
            try store.clear()
        } catch {
            throw .storage(error)
        }
    }

    private func exchange(code: String, verifier: String, deviceName: String) async throws(AuthError) -> AppSession {
        let body = AppExchangeRequest(code: code, codeVerifier: verifier, deviceName: deviceName)
        do throws(APIError) {
            return try await client.send(.appExchange(body))
        } catch {
            throw .exchange(error)
        }
    }
}
