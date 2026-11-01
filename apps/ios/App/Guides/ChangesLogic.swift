import Foundation
import Models

/// A key date as F9.5 draws it.
struct KeyDateRow: Identifiable, Equatable, Sendable {
    /// The data's display date, `Dec 9`.
    let key: String
    let label: String
    let note: String?
    /// The gold sentence while the date is in progress.
    let running: String?
    let iso: String

    var id: String { "\(iso)-\(label)" }
}

enum KeyDatesLogic {
    /// Past dates dropped, the rest sorted by `iso`; `running` only while today is inside the
    /// date's span (dates.json is compared by `iso`/`endIso`, never by `key`).
    static func rows(_ dates: [KeyDate], now: Date, zone: TimeZone = .current) -> [KeyDateRow] {
        let today = GuideDates.iso(now, zone: zone)
        return dates
            .filter { ($0.endIso ?? $0.iso) >= today }
            .sorted { $0.iso < $1.iso }
            .map { date in
                let inProgress = date.iso <= today && today <= (date.endIso ?? date.iso)
                return KeyDateRow(key: date.key, label: date.value, note: date.note, running: inProgress ? date.running : nil, iso: date.iso)
            }
    }
}

/// A changelog row (spec F9.5): `10/28`, the title, the tag.
struct ChangelogRow: Identifiable, Equatable, Sendable {
    let entry: ChangelogEntry
    let date: String
    let tag: String

    var id: String { entry.id }
    var title: String { entry.title }
}

enum ChangelogLogic {
    static let siteKind = "site"

    /// Newest first, as the file has them.
    static func rows(_ changelog: Changelog) -> [ChangelogRow] {
        changelog.entries.map { ChangelogRow(entry: $0, date: GuideDates.monthDay($0.date), tag: tag($0.kind)) }
    }

    /// `Site` for a change to this site; every other kind is game news.
    static func tag(_ kind: String) -> String {
        kind == siteKind ? Copy.changesTagSite : Copy.changesTagGame
    }
}

/// Whether the new-data strip shows over an open screen (spec §5.7).
enum NewDataState: Equatable, Sendable {
    case none
    /// A newer build is ready; `Update` applies it.
    case ready(String)
    /// The update failed; `Update` tries again.
    case failed(String)
}

enum NewDataCheck {
    /// A newer build than the one on screen; nil (offline) never shows the strip.
    static func state(shown: String, latest: String?) -> NewDataState {
        guard let latest, latest != shown else { return .none }
        return .ready(latest)
    }
}

/// Keeping the player's classes ready offline (spec F9 Sync).
enum GuidesSync {
    /// Foreground checks run at most hourly.
    static let interval: TimeInterval = 3600

    static func isDue(last: Date?, now: Date) -> Bool {
        guard let last else { return true }
        return now.timeIntervalSince(last) >= interval
    }

    /// The guides of the classes the player has characters in, in index order.
    static func specs(in index: GuideIndex, classes: Set<String>) -> [String] {
        index.classes.filter { classes.contains($0.classSlug) }.flatMap { $0.specs.map(\.spec) }
    }
}
