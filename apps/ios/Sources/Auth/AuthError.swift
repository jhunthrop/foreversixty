import Foundation
import Networking

public enum AuthError: Error, Equatable {
    /// The player closed the sign-in sheet.
    case cancelled
    /// The browser session failed before it reached the callback.
    case webSession(String)
    /// The callback carried no code, or an `error=` from the API.
    case callbackRejected(String)
    /// The callback's state does not match the one this sign-in sent: a forged or stale redirect.
    case stateMismatch
    case randomUnavailable
    case exchange(APIError)
    case storage(TokenStoreError)
}
