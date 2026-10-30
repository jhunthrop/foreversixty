import Foundation

/// Everything an API call can fail with. Views switch on the case; the
/// server's own message is user-facing copy (the API writes it that way).
public enum APIError: Error, Sendable, Equatable {
    /// No response: offline, timed out, TLS. Carries the system description.
    case transport(String)
    /// 401 (`unauthorized` or `session_expired`): the device token is missing, revoked or
    /// expired. The client has already discarded it; show sign-in.
    case unauthorized(requestID: String?)
    /// The API answered `ok: false` with its own code and message.
    case server(status: Int, body: ErrorBody, requestID: String?)
    /// A body that is not the envelope, or a payload that does not match its model.
    case decoding(String)
    /// A non-2xx status without a readable envelope.
    case unexpectedStatus(Int)
    /// 426 `update_required`: this build is below the API's minimum. The app has already
    /// switched to its blocking update screen; `storeURL` is the gate's store page.
    case updateRequired(storeURL: URL?)
    /// Fixture mode has no file for this route.
    case missingFixture(String)

    /// The server's error code, when it sent one.
    public var code: String? {
        if case let .server(_, body, _) = self { return body.code }
        return nil
    }

    /// 403 `browser_session_required`: billing writes happen on foreversixty.gg in a browser.
    /// Send the player to the site; never retry from the app.
    public var needsBrowserSession: Bool {
        code == Self.browserSessionRequired
    }

    public static let browserSessionRequired = "browser_session_required"
}
