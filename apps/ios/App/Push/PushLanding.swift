import Foundation
import Routing

/// Where a tapped notification lands (flows spec F4, E4): every event on its answer, never a list
/// the player must search. Read from the APNs payload's `kind`, `link` and `data`; pure, so the
/// table is tested without a notification.
enum PushLanding {
    /// The event kinds the API sends (api/internal/events/catalogue.go), by their landing.
    enum Kind {
        static let raidStarting = "raid.starting"
        static let raidSummary = "raid.summary"
        static let lootAwarded = "loot.awarded"
        static let gear: Set<String> = ["bis.changed", "band.entered", "game.milestone"]
        static let characterLinked = "character.linked"
        static let account: Set<String> = ["billing.failed", "premium.seat"]
    }

    /// The destination for a notification's userInfo, or nil when it names nothing the app opens
    /// (the app then simply comes forward on its current screen).
    static func destination(userInfo: [AnyHashable: Any]) -> IntentDestination? {
        let data = (userInfo["data"] as? [String: Any] ?? [:]).compactMapValues { $0 as? String }
        let link = (userInfo["link"] as? String).flatMap(URL.init(string:))
        return destination(kind: userInfo["kind"] as? String, link: link, data: data)
    }

    static func destination(kind: String?, link: URL?, data: [String: String]) -> IntentDestination? {
        if let kind, let byKind = destination(kind: kind, data: data) { return byKind }
        return link.flatMap(Self.link(from:)).map(IntentDestination.link)
    }

    /// The kinds whose answer the payload's data names better than its web link.
    private static func destination(kind: String, data: [String: String]) -> IntentDestination? {
        let guild = data["guild_id"].flatMap(Int64.init)
        let raid = data["raid_id"].flatMap(Int64.init)
        switch kind {
        case Kind.raidSummary:
            return data["report_id"].map { .link(.report(id: $0)) }
        case Kind.raidStarting:
            guard let guild, let raid else { return nil }
            return .link(.raidNight(guild: guild, raid: raid))
        case _ where kind == Kind.lootAwarded || Kind.gear.contains(kind):
            // The item sheet and the Gear tab replace this with the navigation step (spec 4.3).
            return .tab(.gear)
        case Kind.characterLinked:
            return data["character_key"].map { .character(key: $0) }
        case _ where Kind.account.contains(kind):
            return .link(.plan)
        default:
            guard let guild, let raid else { return nil }
            return .link(.raid(guild: guild, raid: raid))
        }
    }

    /// The screen a notification's link names (an app link or one of the site pages the app opens).
    static func link(from url: URL) -> DeepLink? {
        DeepLink(url: url)
    }
}
