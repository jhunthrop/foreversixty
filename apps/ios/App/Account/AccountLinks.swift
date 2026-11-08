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
    /// F8's `See plans`, link regions only (decision P).
    static let premium = page("/premium")
    /// F8.4's `Manage billing on foreversixty.gg`, link regions only.
    static let billing = page("/account#plan")

    /// The menu's Help: the Discord server, from the site's own links file (web/src/data/links.json,
    /// bundled by project.yml), else the site.
    static let help: URL = {
        struct Links: Decodable { let discordInvite: URL }
        let links = Bundle.main.url(forResource: "links", withExtension: "json")
            .flatMap { try? Data(contentsOf: $0) }
            .flatMap { try? JSONDecoder().decode(Links.self, from: $0) }
        return links?.discordInvite ?? page("/")
    }()

    private static func page(_ path: String) -> URL {
        guard let url = URL(string: site + path) else { preconditionFailure("bad site path \(path)") }
        return url
    }
}
