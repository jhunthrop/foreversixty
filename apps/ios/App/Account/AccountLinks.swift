import Foundation

/// The site pages F1 opens in the browser (never an in-app web view).
enum AccountLinks {
    private static let site = "https://foreversixty.gg"

    /// Where a `403 browser_session_required` act is done.
    static let account = page("/account")
    /// F1.3's `Get the addon`.
    static let setup = page("/setup")
    /// You's Privacy and Terms rows (App Store 5.1.1(i)).
    static let privacy = page("/privacy")
    static let terms = page("/terms")

    private static func page(_ path: String) -> URL {
        guard let url = URL(string: site + path) else { preconditionFailure("bad site path \(path)") }
        return url
    }
}
