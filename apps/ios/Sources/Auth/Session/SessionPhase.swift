import Foundation
import Models
import Networking

/// Where the app's account stands. The views draw one screen per phase (spec §5.1).
public enum SessionPhase: Equatable, Sendable {
    /// F1.1, with the line the last ending left (F1.4's toast, the 401 line, a deletion).
    case signedOut(SignedOutNotice?)
    /// F1.2: the system browser is open; the button shows its spinner.
    case signingIn
    /// F1.3: the code is exchanged; the characters are on their way.
    case importing
    /// Signed in. `me` is nil until the first `GET /v1/me` of this launch answers.
    case signedIn(Me?)

    public var isSignedIn: Bool {
        if case .signedIn = self { return true }
        return false
    }

    public var me: Me? {
        if case let .signedIn(me) = self { return me }
        return nil
    }
}

/// The one line F1.1 shows above its buttons, held until the next tap.
public enum SignedOutNotice: Equatable, Sendable {
    case signInFailed(SignInFailure)
    /// A 401 (`unauthorized` or `session_expired`) ended the session without asking.
    case sessionEnded
    /// `DELETE /v1/me` was accepted; the data is gone by this date.
    case accountDeleted(completeBy: Date)
}

/// The known sign-in failures (spec F1.4); each has one sentence in the copy catalog.
public enum SignInFailure: Equatable, Sendable {
    /// Battle.net or the callback refused, the state did not match, or the exchange failed.
    case battleNet
    /// The one-time code was spent or expired (`400 invalid_grant`).
    case expired
    /// No connection: signing in needs the internet.
    case offline

    /// The failure an `AuthError` shows, or nil when the player cancelled (back to F1.1, no line).
    public init?(_ error: AuthError) {
        switch error {
        case .cancelled:
            return nil
        case let .exchange(apiError):
            self = Self.exchange(apiError)
        case .webSession, .callbackRejected, .stateMismatch, .randomUnavailable, .storage:
            self = .battleNet
        }
    }

    private static let invalidGrant = "invalid_grant"

    private static func exchange(_ error: APIError) -> SignInFailure {
        switch error {
        case .transport: .offline
        case let .server(_, body, _) where body.code == invalidGrant: .expired
        default: .battleNet
        }
    }
}
