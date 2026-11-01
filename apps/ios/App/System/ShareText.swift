import Foundation
import Models

/// A share sheet's link and its message (spec §5.7 Share): only things with a working web page
/// share, each with the page whose Open Graph card the site makes.
struct ShareItem: Equatable, Sendable {
    let url: URL
    let text: String
}

/// The share messages and links, pure (spec §5.7 Share table).
enum ShareText {
    static let site = "https://foreversixty.gg"

    private static func url(_ path: String) -> URL {
        URL(string: site + path) ?? URL(fileURLWithPath: "/")
    }

    /// `/character/{region}/{ruleset}/{name}`, `Lightningbolltz, Elemental Shaman, level 60`.
    /// With no spec, class or level known yet, the name alone.
    static func character(key: String, name: String, spec: String?, className: String?, level: Int?) -> ShareItem {
        let classWord = className.map(SearchWords.className) ?? ""
        let text: String
        if let spec, let level, !classWord.isEmpty {
            text = Copy.shareCharacter(name: name, spec: spec, className: classWord, n: level)
        } else {
            text = name
        }
        return ShareItem(url: url("/character/\(key)"), text: text)
    }

    /// `/guides/{class}/{spec}`, `Elemental Shaman guide`.
    static func guide(_ guide: GuideFile) -> ShareItem {
        ShareItem(
            url: GuideCopy.shareURL(guide),
            text: Copy.shareGuide(spec: guide.specName, className: guide.className)
        )
    }

    /// `/tiers`, `/tiers/tank`, `/tiers/healer`; `DPS tier list`.
    static func tiers(_ role: TierRole) -> ShareItem {
        let path = role == .dps ? "/tiers" : "/tiers/\(role.rawValue)"
        return ShareItem(url: url(path), text: Copy.shareTiers(role: role.title))
    }

    /// `/reports/{id}`, `Onyxia's Lair · 15 Dec`. Guild and private reports have no share (nil).
    static func report(id: String, zone: String, visibility: String, createdAt: Date, zoneOfPhone: TimeZone = .current) -> ShareItem? {
        guard sharedVisibilities.contains(visibility) else { return nil }
        let date = WidgetClock(zone: zoneOfPhone).dayMonth(createdAt)
        return ShareItem(url: url("/reports/\(id)"), text: Copy.shareReport(zone: zone, date: date))
    }

    static let sharedVisibilities: Set<String> = ["public", "unlisted"]

    /// A loaded report's share, nil for a guild or private one.
    static func report(_ report: Report) -> ShareItem? {
        Self.report(id: report.id, zone: report.zone ?? report.title, visibility: report.visibility, createdAt: report.createdAt)
    }
}
