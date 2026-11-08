import DesignSystem
import Foundation
import Models

/// One head on every Home (flows spec 4.1): the class crest, the name in the class colour, the
/// fact line. No figures.
struct HomeHeadContent: Equatable {
    let name: String
    let className: String?
    let facts: [String]
}

/// A row title and its line.
struct HomeLine: Equatable {
    let title: String
    let sub: String
}

/// One Next upgrades row: the item, its slot and source, its gain.
struct HomeUpgradeItem: Equatable, Identifiable {
    let slot: String
    let look: HomeItemLook
    let source: String
    let gain: HomeGain?
    let wornID: Int?
    let sourceLine: String

    var id: String { slot }
}

struct HomeGain: Equatable {
    let figure: String
    let unit: String
}

struct HomeUpgradesContent: Equatable {
    let rows: [HomeUpgradeItem]
    /// `Best run: Blackrock Spire, Blackcrow +1.9 DPS, 8%`; nil on raid day or with no dungeon pick.
    let bestRun: String?
    /// The answer's sentence when there are no rows (all slots match, no gear yet).
    let note: String?
}

/// The next raid as Home's card (raid day) or line (other days), and the Raid tab's card.
struct NextRaidContent: Equatable {
    let guildID: Int64
    let raidID: Int64
    let characterKey: String
    /// `Olympus XXVII · Onyxia's Lair tonight`.
    let eyebrow: String
    let clock: [FSClockPart]
    /// `Wed 6 Jan, 8:00 pm CST` (no guild stores a sign-up cutoff, so there is no `answer by`).
    let when: String
    /// The line's title and sub (other days).
    let line: HomeLine
    let icon: String?
    /// `You have not answered, as Thandrel, Damage`.
    let answer: String
    let selected: SignupState?
    /// Late's minutes are open under the answers.
    let lateOpen: Bool
    let tally: [FSTally.Item]
    let ready: HomeLine?
    let officer: OfficerContent?
}

/// The officer's open items (F3): the unanswered with the short role named, the applicants.
struct OfficerContent: Equatable {
    let unanswered: HomeLine?
    let applicants: HomeLine?
}

struct RaidLiveContent: Equatable {
    let eyebrow: String
    let line: String
}

struct LastNightLoot: Equatable, Identifiable {
    let look: HomeItemLook
    let line: String

    var id: Int { look.id }
}

struct LastNightContent: Equatable {
    let loot: [LastNightLoot]
    let report: HomeLine?
    let reportIcon: String?
    let reportID: String
    let simAgain: HomeLine?
}

struct NewForYouRow: Equatable, Identifiable {
    let row: ActivityRow
    let line: String

    var id: Int64 { row.id }
}

struct NewForYouContent: Equatable {
    let rows: [NewForYouRow]
    let unread: Int
}

/// A block as drawn, in the server's order.
enum HomeBlockContent: Equatable, Identifiable {
    case upgrades(HomeUpgradesContent)
    case raidCard(NextRaidContent)
    case raidLine(NextRaidContent)
    case raidLive(RaidLiveContent)
    case lastNight(LastNightContent)
    case newForYou(NewForYouContent)
    case findGuild(String)
    case notice(String)

    var id: String {
        switch self {
        case .upgrades: "upgrades"
        case .raidCard: "raid-card"
        case .raidLine: "raid-line"
        case .raidLive: "raid-live"
        case .lastNight: "last-night"
        case .newForYou: "new-for-you"
        case .findGuild: "find-guild"
        case .notice: "notice"
        }
    }
}

/// What Home shows.
enum HomeFlowPhase: Equatable {
    /// The head, and the blocks at their final size.
    case loading
    case failed
    case loaded([HomeBlockContent])
}

struct HomeFlowContent: Equatable {
    let head: HomeHeadContent
    let phase: HomeFlowPhase
    /// Set while the last answers show because Forever Sixty did not answer (lane F15).
    var staleStamp: String?
    /// Set while the last answers show because the connection is gone (E5).
    var offlineStamp: String?
}
