import Foundation
import Models

/// One block of the time-aware Home (flows spec 4.1), from `/v1/home` `sections` in the server's
/// order (`homesections.MomentOf`). Types the app does not know are skipped: the list is open.
enum HomeBlock: Equatable, Sendable {
    case head
    case upgrades(limit: Int)
    /// `card` on raid day (the countdown, the answers, who is coming); `line` on other days.
    case nextRaid(card: Bool)
    case raidLive
    case raidSummary
    case activity(HomeActivityQuery)
    case guildCTA
    case notice(text: String, href: String?)

    static let defaultUpgrades = 3
    static let defaultActivity = 3

    static func of(_ sections: [HomeSection]) -> [HomeBlock] {
        sections.compactMap { section in
            switch section.type {
            case "head": .head
            case "upgrades": .upgrades(limit: section.limit ?? defaultUpgrades)
            case "next_raid": .nextRaid(card: section.style == "card")
            case "raid_live": .raidLive
            case "raid_summary": .raidSummary
            case "activity":
                .activity(HomeActivityQuery(
                    limit: section.limit ?? defaultActivity, group: section.group, exclude: Set(section.excludeKinds ?? [])
                ))
            case "guild_cta": .guildCTA
            case "notice": section.text.map { .notice(text: $0, href: section.href) }
            default: nil
            }
        }
    }

    /// The Best run line rides with Next upgrades except when the Home carries a raid card: the raid
    /// is tonight (flows spec F5, appendix C).
    static func showsBestRun(_ blocks: [HomeBlock]) -> Bool {
        !blocks.contains(.nextRaid(card: true))
    }

    /// The raid the Home needs loaded: a next raid, or the live one.
    static func needsRaid(_ blocks: [HomeBlock]) -> Bool {
        blocks.contains { block in
            switch block {
            case .nextRaid, .raidLive: true
            default: false
            }
        }
    }
}

/// The New for you preview: how many rows, which group (raids on raid night), and the kinds a block
/// above already shows (the day after: Last night's loot and report), read past to fill the limit.
struct HomeActivityQuery: Equatable, Sendable {
    let limit: Int
    let group: String?
    let exclude: Set<String>

    func rows(_ rows: [ActivityRow]) -> [ActivityRow] {
        Array(rows.filter { !exclude.contains($0.kind) }.sorted { $0.occurredAt > $1.occurredAt }.prefix(limit))
    }
}
