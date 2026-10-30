import Foundation
import Models
import Networking

/// The app sign-in contract (lane A1, docs/superpowers/plans/2026-10-30-mobile-apps-lanes.md):
///
/// 1. The app opens `GET /v1/auth/app/start?platform=ios&code_challenge=…&code_challenge_method=S256&state=…`
///    in ASWebAuthenticationSession. The API runs Battle.net OAuth.
/// 2. The API redirects to the app link `https://foreversixty.gg/app/auth/callback?code=…&state=…`
///    (or `?error=…&state=…`).
/// 3. The app posts the one-time code with its verifier to `POST /v1/auth/app/exchange`
///    and gets a device token (`fsd_…`).
///
/// Every name here is the one place to change if lane A1's final contract differs.
public enum AppAuthContract {
    public static let platform = "ios"
    public static let startPath = "/v1/auth/app/start"
    public static let exchangePath = "/v1/auth/app/exchange"
    public static let callbackHost = "foreversixty.gg"
    public static let callbackPath = "/app/auth/callback"

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

/// `POST /v1/auth/app/exchange` request body.
public struct AppExchangeRequest: Encodable, Sendable, Equatable {
    public let code: String
    public let codeVerifier: String
    public let platform: String
    public let deviceName: String
}

/// `POST /v1/auth/app/exchange` data: the device token and the device row it created.
public struct AppSession: Decodable, Sendable, Equatable {
    public let token: String
    public let device: Device
}

public extension Endpoint where Response == AppSession {
    static func appExchange(_ body: AppExchangeRequest) throws(APIError) -> Endpoint {
        try .json(.post, AppAuthContract.exchangePath, body: body, authenticated: false)
    }
}
