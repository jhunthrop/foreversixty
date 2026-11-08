import DesignSystem
import Foundation

/// One of the guild's pages under the raid answers (flows spec 4.2): Roster, Calendar, Loot
/// history, Guild sims.
enum GuildPage: String, CaseIterable, Identifiable, Sendable {
    case roster, calendar, loot, sims

    var id: String { rawValue }

    var title: String {
        switch self {
        case .roster: Copy.raidsRosterRow
        case .calendar: Copy.raidPagesCalendar
        case .loot: Copy.raidPagesLoot
        case .sims: Copy.raidPagesSims
        }
    }
}

struct GuildPageRow: Equatable, Identifiable {
    let page: GuildPage
    /// `38 members, 31 at 60`; nil while the page's figure is unknown.
    let sub: String?

    var id: String { page.id }
}

/// The guild's pages, titled with the guild's name (the guild head is gone: appendix C).
struct GuildPagesContent: Equatable {
    let guild: String
    let rows: [GuildPageRow]
}

/// An officer's open items (F3): the badge's number and its two rows.
struct NeedsYouContent: Equatable {
    /// `6 open`.
    let meta: String
    let items: OfficerContent
}

/// One ready-check line (F2): its mark, what it is, and where to get a miss.
struct ReadyRowContent: Equatable, Identifiable {
    enum Mark: String, Equatable {
        case ok, warn, miss
    }

    let key: String
    let mark: Mark
    let title: String
    let line: String
    let fix: String?
    let link: URL?

    var id: String { key }
}

struct ReadyListContent: Equatable {
    let rows: [ReadyRowContent]
    /// `From your last log-out, 9 h ago`.
    let stamp: String?
}

/// Raid night on the Raid tab (F2): the live card, then who is in by role with names.
struct RaidNightRootContent: Equatable {
    let live: RaidLiveContent
    let roles: [HomeLine]
}

/// The Raid tab for a guild member.
struct RaidGuildContent: Equatable {
    /// Raid night leads while a raid is live.
    var night: RaidNightRootContent?
    /// The next raid's card; nil with no raid planned (and while raid night leads).
    var next: NextRaidContent?
    var ready: ReadyListContent?
    var needsYou: NeedsYouContent?
    var pages: GuildPagesContent
    var offlineStamp: String?
    var staleStamp: String?
}

/// A guild recruiting the player's spec, with Apply.
struct RecruitingRow: Equatable, Identifiable {
    let guildID: Int64
    let name: String
    /// `Alliance · Wed and Sat, 8:00 pm CST · needs damage`.
    let line: String

    var id: Int64 { guildID }
}

struct RecruitingContent: Equatable {
    /// `Guilds recruiting a Fury Warrior`.
    let heading: String
    let rows: [RecruitingRow]
}

/// What the Raid tab's root draws (flows spec 4.2).
enum RaidRootContent: Equatable {
    case loading
    case failed
    case signedOut
    /// Under 60: the levels to go.
    case leveler(toGo: String)
    case noGuild(RecruitingContent)
    /// In a guild, waiting for an officer to approve the character.
    case waiting
    case guild(RaidGuildContent)
}

/// One raider on Not answered (F3).
struct NotAnsweredRow: Equatable, Identifiable {
    let characterKey: String
    let name: String
    let className: String
    /// `Holy Priest · was In last week`.
    let sub: String
    let role: String?
    /// Marked In or Bench on this page, with Undo while the hold runs.
    var marked: SignupState?

    var id: String { characterKey }
}

/// A filter chip: Not answered (preselected), In, Late, Bench, Out, All, each with its count.
enum NotAnsweredFilter: String, CaseIterable, Identifiable, Sendable {
    case notAnswered, signedIn = "in", late, bench, out, all

    var id: String { rawValue }

    var label: String {
        switch self {
        case .notAnswered: Copy.raidNaTitle
        case .signedIn: SignupState.signedIn.label
        case .late: SignupState.late.label
        case .bench: SignupState.bench.label
        case .out: SignupState.out.label
        case .all: Copy.raidNaAll
        }
    }

    /// The sign-up state the chip lists; nil for Not answered and All.
    var state: SignupState? {
        switch self {
        case .signedIn: .signedIn
        case .late: .late
        case .bench: .bench
        case .out: .out
        case .notAnswered, .all: nil
        }
    }
}

struct NotAnsweredChip: Equatable, Identifiable {
    let filter: NotAnsweredFilter
    let title: String

    var id: String { filter.id }
}

/// The mention: the short role's button, Mention all, the message it posts.
struct MentionContent: Equatable {
    /// `Mention the 2 healers in Discord`; nil with no short role among the unanswered.
    let role: String?
    /// `Mention all 4`.
    let all: String
    let preview: String
}

/// The raid page filtered to Not answered (F3, mock ios-raid-unanswered).
struct NotAnsweredContent: Equatable {
    let title: String
    let icon: String
    /// `Olympus XXVII`, `Wed 6 Jan, 8:00 pm CST`, `healers short by 3`.
    let facts: [String]
    let chips: [NotAnsweredChip]
    let filter: NotAnsweredFilter
    let mention: MentionContent?
    /// `Not answered`, or the filter's name.
    let listTitle: String
    /// `4 of 31, healers first`.
    let listMeta: String?
    let rows: [NotAnsweredRow]
    /// Only Not answered's rows carry Mark In, Bench, Discord DM.
    let actionable: Bool
    /// `Healers on the bench`, with Mark In.
    let benchTitle: String?
    let bench: [NotAnsweredRow]
    /// `If Mirethil, Selka and Aurelith are In, healers are full.`
    let benchNote: String?
}

/// One figure on the applicant: `74 Forever Score`, `612 HPS Sim`.
struct ApplicantFigure: Equatable, Identifiable {
    let label: String
    let value: String
    let unit: String?

    var id: String { label }
}

/// The applicant (F3, mocks ios-raid-applicant, -accepted).
struct ApplicantContent: Equatable {
    let name: String
    let className: String?
    /// `Holy Priest · 60`, the faction and race when known, `applied 2 min ago`.
    let facts: [String]
    let figures: [ApplicantFigure]
    let ready: [HomeLine]
    let history: [HomeLine]
    let note: String?
    /// After Accept: `Accepted into Olympus XXVII. Invite Mirelle in game:`.
    let accepted: String?
    let invite: String
    let after: String
}

/// The raid page on raid night (mocks ios-raid-night-live, -signup).
struct RaidNightPageContent: Equatable {
    let title: String
    let icon: String
    let facts: [String]
    let eyebrow: String
    let clock: [FSClockPart]
    /// `Wed 6 Jan, 8:00 pm CST · 18 in, 3 late`.
    let aside: String
    /// `You are In`, or `Not signed up`.
    let tag: String
    let answered: Bool
    let selected: SignupState?
    let lateOpen: Bool
    let roles: [HomeLine]
    let ready: ReadyListContent?
}
