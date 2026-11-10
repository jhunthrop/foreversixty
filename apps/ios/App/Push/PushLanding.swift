import Foundation
import Routing

/// Where a tapped notification lands (flows spec F4, E4): every event on its answer, never a list
/// the player must search. Read from the APNs payload's `kind`, `link` and `data`; pure, so the
/// table is tested without a notification.
enum PushLanding {
    /// The event kinds the API sends (api/internal/events/catalogue.go), by their landing.
    enum Kind {
        static let raidStarting = "raid.starting"
        /// An officer's: the raid is short, so the raid page opens filtered to Not answered (F3).
        static let raidShort = "raid.short"
        static let raidSummary = "raid.summary"
        static let lootAwarded = "loot.awarded"
        static let gear: Set<String> = ["bis.changed", "band.entered", "game.milestone"]
        static let characterLinked = "character.linked"
        static let account: Set<String> = ["billing.failed", "premium.seat"]
        /// An officer's: a new guild application; the tap opens its review.
        static let guildApplicant = "guild.applicant"
    }

    /// The application a `guild.applicant` payload names (its data's guild_id and application_id).
    static func applicant(data: [String: String]) -> (guild: Int64, application: Int64)? {
        guard let guild = data["guild_id"].flatMap(Int64.init),
              let application = data["application_id"].flatMap(Int64.init) else { return nil }
        return (guild, application)
    }

    /// The payload's data strings, as APNs delivers them in userInfo.
    static func data(userInfo: [AnyHashable: Any]) -> [String: String] {
        (userInfo["data"] as? [String: Any] ?? [:]).compactMapValues { $0 as? String }
    }

    /// The destination for a notification's userInfo, or nil when it names nothing the app opens
    /// (the app then simply comes forward on its current screen).
    static func destination(userInfo: [AnyHashable: Any]) -> IntentDestination? {
        let data = data(userInfo: userInfo)
        let link = (userInfo["link"] as? String).flatMap(URL.init(string:))
        return destination(kind: userInfo["kind"] as? String, link: link, data: data)
    }

    static func destination(kind: String?, link: URL?, data: [String: String]) -> IntentDestination? {
        let linked = link.flatMap(Self.link(from:))
        if let kind, let byKind = destination(kind: kind, data: data, linked: linked) { return byKind }
        return linked.map(IntentDestination.link)
    }

    /// The kinds whose answer the payload's data names better than its web link.
    private static func destination(kind: String, data: [String: String], linked: DeepLink?) -> IntentDestination? {
        let guild = data["guild_id"].flatMap(Int64.init)
        let raid = data["raid_id"].flatMap(Int64.init)
        switch kind {
        case Kind.raidSummary:
            return data["report_id"].map { .link(.report(id: $0)) }
        case Kind.raidStarting:
            guard let guild, let raid else { return nil }
            return .link(.raidNight(guild: guild, raid: raid))
        case Kind.raidShort:
            guard let guild, let raid else { return nil }
            return .raid(.notAnswered(guild: guild, raid: raid))
        case _ where kind == Kind.lootAwarded || Kind.gear.contains(kind):
            // The item sheet and the Gear tab replace this with the navigation step (spec 4.3).
            return .tab(.gear)
        case Kind.guildApplicant:
            return applicant(data: data).map { .link(.guildApplicant(guild: $0.guild, application: $0.application)) }
        case Kind.characterLinked:
            return data["character_key"].map { .character(key: $0) }
        case _ where Kind.account.contains(kind):
            return .link(.plan)
        default:
            guard let guild, let raid else { return nil }
            // The ids win; the server's app link for the same raid still says what to press (`state=`, `spec=ask`).
            guard case let .raid(linkGuild, linkRaid, spec, state, _)? = linked, linkGuild == guild, linkRaid == raid else {
                return .link(.raid(guild: guild, raid: raid))
            }
            return .link(.raid(guild: guild, raid: raid, spec: spec, state: state))
        }
    }

    /// The screen a notification's link names (an app link or one of the site pages the app opens).
    static func link(from url: URL) -> DeepLink? {
        DeepLink(url: url)
    }
}
