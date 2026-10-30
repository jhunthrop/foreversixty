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
    public static let exchangePath = "/v1/auth/app/exchange"
    public static let logoutPath = "/v1/auth/app/logout"
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

/// `POST /v1/auth/app/exchange` request body. `device_name` is optional on the server (it
/// defaults to "iOS"); the app always sends one so /account names the phone.
public struct AppExchangeRequest: Encodable, Sendable, Equatable {
    public let code: String
    public let codeVerifier: String
    public let deviceName: String
}

/// `POST /v1/auth/app/exchange` data: the device token (shown only here), the device row it
/// created and the `GET /v1/me` body, so the first screen needs no second call.
public struct AppSession: Decodable, Sendable, Equatable {
    public let token: String
    public let device: Device
    public let me: Me
}

/// `POST /v1/auth/app/logout` data.
public struct SignedOut: Decodable, Sendable, Equatable {
    public let signedOut: Bool
}

public extension Endpoint where Response == AppSession {
    static func appExchange(_ body: AppExchangeRequest) throws(APIError) -> Endpoint {
        try .json(.post, AppAuthContract.exchangePath, body: body, authenticated: false)
    }
}

public extension Endpoint where Response == SignedOut {
    /// Revokes this device's token on the server.
    static var appLogout: Endpoint { Endpoint(method: .post, path: AppAuthContract.logoutPath) }
}
