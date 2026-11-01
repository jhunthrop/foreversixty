import Auth
import Foundation
import Networking
import Observation

/// The calendar link's secure store, per account (iOS Keychain live). Cleared at sign-out.
protocol CalendarLinkStore: Sendable {
    func load(account: String) -> String?
    func save(_ url: String, account: String)
    func clear()
}

/// One Keychain item through the app's Keychain helper, holding `{account}\n{url}`: a link
/// stored for another account never shows.
struct KeychainCalendarLinkStore: CalendarLinkStore {
    static let service = "gg.foreversixty.app.calendar-link"
    private let keychain = KeychainTokenStore(service: Self.service)

    func load(account: String) -> String? {
        guard let value = try? keychain.load() else { return nil }
        let parts = value.split(separator: "\n", maxSplits: 1).map(String.init)
        guard parts.count == 2, parts[0] == account else { return nil }
        return parts[1]
    }

    func save(_ url: String, account: String) { try? keychain.save("\(account)\n\(url)") }

    func clear() { try? keychain.clear() }
}

/// What the Calendar link row shows (spec F16.4 part 2).
enum CalendarLinkPhase: Equatable, Sendable {
    case loading
    case ready(String)
    /// The live link was shown on another phone or browser and none is stored here: only Reset.
    case elsewhere
    /// It did not load and nothing is stored: `calendar.linkError` and Try again.
    case failed
}

/// The link's state machine: `GET /v1/me/calendar-feed` shows the url once; this keeps it in the
/// secure store. The url is never logged.
@MainActor
@Observable
final class CalendarLinkModel {
    private(set) var phase: CalendarLinkPhase = .loading
    /// The read could not connect: the buttons draw disabled at 50% (writes never queue).
    private(set) var offline = false
    private(set) var resetting = false
    /// `calendar.linkError` or `calendar.resetError`.
    private(set) var message: String?

    private let service: any GuildExtrasService
    private let store: any CalendarLinkStore
    private let account: String

    init(service: any GuildExtrasService, store: any CalendarLinkStore, account: String) {
        self.service = service
        self.store = store
        self.account = account
    }

    var url: String? {
        if case let .ready(url) = phase { return url }
        return nil
    }

    func load() async {
        message = nil
        do {
            let feed = try await service.calendarFeed()
            offline = false
            accept(feed.url)
        } catch {
            offline = SignupRefusal.isOffline(error)
            if let stored = store.load(account: account) {
                phase = .ready(stored)
                if !offline { message = Copy.calendarLinkError }
            } else {
                phase = .failed
                message = Copy.calendarLinkError
            }
        }
    }

    private func accept(_ shown: String?) {
        if let shown {
            store.save(shown, account: account)
            phase = .ready(shown)
        } else if let stored = store.load(account: account) {
            phase = .ready(stored)
        } else {
            phase = .elsewhere
        }
    }

    /// After the confirmation: a new link, the old one stops working.
    func reset() async {
        guard !resetting else { return }
        resetting = true
        defer { resetting = false }
        message = nil
        do {
            let feed = try await service.resetCalendarFeed()
            offline = false
            accept(feed.url)
        } catch {
            message = Copy.calendarResetError
        }
    }

    /// `webcal://` in place of `https://`, for Add to Calendar.
    static func webcal(_ url: String) -> URL? {
        guard var components = URLComponents(string: url), components.scheme == "https" else { return nil }
        components.scheme = "webcal"
        return components.url
    }
}
