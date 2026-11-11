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
    /// The item the check is about (the flask, the attunement item), when the data has its icon.
    var icon: String?

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
    let roles: [RoleLineContent]
}

/// One named player on a role line, the name in the class colour (`Brightwen (Late · 15 min)`).
struct RolePerson: Equatable {
    let name: String
    let className: String?
    /// The name as the line reads it: with a Late's minutes.
    let text: String
}

/// One role of Who is coming (S6, one tally model on both apps): `Healers 5 of 8 · 3 short` in tone
/// when the raid has a target, `Healers 6` without; then the names in class colour, and for Damage
/// the melee and ranged split.
struct RoleLineContent: Equatable, Identifiable {
    enum Tone: Equatable { case met, short, plain }

    let role: String
    let have: Int
    let want: Int?
    /// The names as one line (VoiceOver, and where no class is known).
    let names: String
    /// The names drawn in class colour; past five, the first three and `and N more`.
    var people: [RolePerson] = []
    var more = 0
    /// Damage: `5 melee, 6 ranged`.
    var split: String?

    var id: String { role }

    /// How many the target still wants; nil with no target or none short.
    var short: Int? { want.flatMap { $0 > have ? $0 - have : nil } }

    var tone: Tone {
        guard want != nil else { return .plain }
        return short == nil ? .met : .short
    }

    /// `5 of 8 · 3 short`, `2 of 2`, or `6` without a target.
    var tally: String {
        guard let want else { return "\(have)" }
        let count = Copy.raidTabOfWant(have: have, want: want)
        return short.map { Copy.raidTabTallyShort(tally: count, n: $0) } ?? count
    }

    /// `2/8` for the compact tally of Home's raid line.
    var compact: String? { want.map { Copy.raidTabTallyCompact(have: have, want: $0) } }

    /// Home's compact form (row 2): `5/8 · 3 short`, `2/2`, or `6` without a target.
    var compactTally: String {
        guard let compact else { return "\(have)" }
        return short.map { Copy.raidTabTallyShort(tally: compact, n: $0) } ?? compact
    }

    /// The card's meter, only against a target.
    var meter: FSTally.Item? { want.map { FSTally.Item(label: role, have: have, want: $0) } }

    /// The line without its target, `Healers 5`: under 60 the Guild tab shows who is coming, not what the raid
    /// still needs (spec 2026-11-10-guild-tab.md 2.1).
    var countOnly: RoleLineContent {
        RoleLineContent(role: role, have: have, want: nil, names: names, people: people, more: more, split: split)
    }
}

/// The Raid tab for a guild member.
struct RaidGuildContent: Equatable {
    /// Raid night leads while a raid is live.
    var night: RaidNightRootContent?
    /// The next raid's head; nil with no raid planned (and while raid night leads).
    var next: NextRaidContent?
    /// Who is coming to the next raid.
    var coming: [RoleLineContent] = []
    /// Before you pull (the next raid) or the Ready check (raid night).
    var ready: ReadyListContent?
    var needsYou: NeedsYouContent?
    var pages: GuildPagesContent
    var offlineStamp: String?
    var staleStamp: String?
    /// The officers' notes, two lines (the raid page has the rest).
    var notes: String?
    /// Who is coming's foot: `Bench 1 · Out 4 · 14 not answered`.
    var answers: String?
    /// G4: what the raid's cover misses, by name (`Missing Fort · WF 1 of 2`).
    var cover: RaidCoverContent?
    /// G1, officers: `Not synced this week: Kaelra, Thalrin, Brakka, Mirethil`, opening the roster.
    var notSynced: NotSyncedContent?
    /// The Guild tab's own sections (spec 2026-11-10-guild-tab.md), nil until the guild home loads.
    var tab: GuildTabContent?
    /// Under 60: `Raids open at 60 · you're 36` under the raid head (they sign up all the same).
    var raidsOpen: String?
    /// Your job in the next raid (spec 2.3a); empty when you hold no slot.
    var jobs: [RaidJobLine] = []
    /// An officer's read-only assignments; nil for everyone else.
    var assignments: [AssignmentEncounterLine]?
}

/// A gap in the cover (G4): nobody brings `name` (`have` nil), or `have` of the `need` the raid wants.
struct CoverGap: Equatable {
    /// The buff's short label (`Fort`), else its name.
    let name: String
    let icon: String?
    var have: Int?
    var need = 0

    /// `Fort`, or `WF 1 of 2`.
    var text: String { have.map { Copy.raidTabCoverShort(name: name, have: $0, need: need) } ?? name }
}

/// The Covered line (row 4; round 4, row 1): `Missing` once over the missing names, then the short
/// ones, then `and N more`, every name parted by the one ` · `; the view sets the names in wipe red
/// beside their spell icons, and a name without one closes up.
struct RaidCoverContent: Equatable {
    let missing: [CoverGap]
    let short: [CoverGap]
    let more: Int

    /// The line as one string (VoiceOver, tests): `Missing Kings · Might · WF 1 of 2 and 3 more`.
    var line: String {
        let separator = RaidFlowPresenter.separator
        let text = [
            missing.isEmpty ? nil : Copy.raidTabCoverMissing(names: missing.map(\.text).joined(separator: separator)),
            short.isEmpty ? nil : short.map(\.text).joined(separator: separator),
        ].compactMap(\.self).joined(separator: separator)
        return more > 0 ? Copy.raidTabAndMore(names: text, n: more) : text
    }
}

/// The officers' not-synced line (row 5) and whom it names, for the filtered roster.
struct NotSyncedContent: Equatable {
    let line: String
    let keys: [String]
}

/// A guild recruiting the player's spec, with Apply.
struct RecruitingRow: Equatable, Identifiable {
    let guildID: Int64
    let name: String
    /// `Alliance · Wed and Sat, 8:00 pm CST · needs damage`.
    let line: String
    /// `Applied Mon` once the player has an open application to the guild (spec 2.4); nil shows Apply.
    var applied: String?

    var id: Int64 { guildID }
}

struct RecruitingContent: Equatable {
    /// `Guilds recruiting a Fury Warrior`.
    let heading: String
    var rows: [RecruitingRow]
    /// Under 60: `Raids open at 60 · you're 36` (spec 2.4).
    var raidsOpen: String?
}

/// What the Raid tab's root draws (flows spec 4.2).
enum RaidRootContent: Equatable {
    case loading
    case failed
    case signedOut
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
    /// The Discord DM link, only for a member who lets officers DM them.
    var dm: URL?

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
