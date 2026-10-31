import AuthenticationServices
import Foundation

/// Opens a URL in the system browser sheet and returns the callback URL.
/// Injected into `AppSignIn` so tests never present UI.
public protocol WebAuthenticating: Sendable {
    func authenticate(url: URL, callbackHost: String, callbackPath: String) async throws(AuthError) -> URL
}

/// ASWebAuthenticationSession with an https (universal link) callback.
///
/// TODO(universal links): the https callback only completes once the app has the
/// Associated Domains entitlement `applinks:foreversixty.gg` and the site serves
/// apple-app-site-association for /app/*. Both wait on the Apple Team ID (lane O);
/// add `com.apple.developer.associated-domains` to project.yml's entitlements then.
@MainActor
public final class SystemWebAuthenticator: NSObject, WebAuthenticating {
    private let anchor: @MainActor () -> ASPresentationAnchor

    public init(anchor: @escaping @MainActor () -> ASPresentationAnchor) {
        self.anchor = anchor
    }

    public nonisolated func authenticate(
        url: URL,
        callbackHost: String,
        callbackPath: String
    ) async throws(AuthError) -> URL {
        let result: Result<URL, AuthError> = await withCheckedContinuation { continuation in
            Task { @MainActor in
                let session = ASWebAuthenticationSession(
                    url: url,
                    callback: .https(host: callbackHost, path: callbackPath)
                ) { callback, error in
                    continuation.resume(returning: Self.result(callback: callback, error: error))
                }
                session.presentationContextProvider = self
                session.prefersEphemeralWebBrowserSession = false
                if !session.start() {
                    continuation.resume(returning: .failure(.webSession("could not start")))
                }
            }
        }
        return try result.get()
    }

    private nonisolated static func result(callback: URL?, error: (any Error)?) -> Result<URL, AuthError> {
        if let callback { return .success(callback) }
        if let error = error as? ASWebAuthenticationSessionError, error.code == .canceledLogin {
            return .failure(.cancelled)
        }
        return .failure(.webSession(error?.localizedDescription ?? "no callback"))
    }
}

extension SystemWebAuthenticator: ASWebAuthenticationPresentationContextProviding {
    public nonisolated func presentationAnchor(for _: ASWebAuthenticationSession) -> ASPresentationAnchor {
        MainActor.assumeIsolated { anchor() }
    }
}
