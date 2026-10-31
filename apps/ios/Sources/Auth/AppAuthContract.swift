import Foundation
import Models
import Networking

/// The app sign-in contract (lane A1; docs/contracts/app-v1.md "Sign in"):
///
/// 1. The app opens `GET /v1/auth/app/start?platform=ios&code_challenge=…&code_challenge_method=S256&state=…`
///    in ASWebAuthenticationSession. The API answers a confirm page; the player taps
///    "Continue to Battle.net" and signs in there.
/// 2. The API sends the browser to exactly `https://foreversixty.gg/app/auth?code=…&state=…`
///    (no `/callback` suffix). The server picks that landing; any `redirect_uri` is ignored.
/// 3. The app posts the one-time code with its verifier and a device name to
///    `POST /v1/auth/app/exchange` and gets 201 `{token, device, me}`.
///
/// Afterwards a 401 (`unauthorized` or `session_expired`) means the token is dead: the client
/// discards it and the app signs in again. Billing writes answer 403 `browser_session_required`.
public enum AppAuthContract {
    public static let platform = "ios"
    public static let startPath = "/v1/auth/app/start"
    public static let callbackHost = "foreversixty.gg"
    public static let callbackPath = "/app/auth"

    public static func startURL(baseURL: URL, pkce: PKCE, state: String) -> URL {
        let start = baseURL.appending(path: startPath)
        return start.appending(queryItems: [
            URLQueryItem(name: "platform", value: platform),
            URLQueryItem(name: "code_challenge", value: pkce.challenge),
            URLQueryItem(name: "code_challenge_method", value: PKCE.method),
            URLQueryItem(name: "state", value: state),
        ])
    }

    /// The one-time code from a callback URL, after checking its state.
    public static func code(fromCallback url: URL, expectedState: String) throws(AuthError) -> String {
        let items = URLComponents(url: url, resolvingAgainstBaseURL: false)?.queryItems ?? []
        let value = { (name: String) in items.first { $0.name == name }?.value }
        guard url.host == callbackHost, url.path == callbackPath else {
            throw .callbackRejected("unexpected callback \(url.host ?? "")\(url.path)")
        }
        guard value("state") == expectedState else { throw .stateMismatch }
        if let error = value("error") { throw .callbackRejected(error) }
        guard let code = value("code"), !code.isEmpty else { throw .callbackRejected("no code") }
        return code
    }
}
