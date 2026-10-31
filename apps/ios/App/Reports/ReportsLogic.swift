import Foundation
import Models

/// The API's words for a report's status and a fight's kind.
enum ReportStatus {
    static let live = "live"
    static let processing = "processing"
    static let failed = "failed"
}

enum FightKind {
    static let encounter = "encounter"
}

/// A list row's tag (spec F10.1): live, processing, failed, or none for a finished report.
enum ReportRowState: Equatable, Sendable {
    case live, processing, failed, complete

    init(status: String) {
        switch status {
        case ReportStatus.live: self = .live
        case ReportStatus.processing: self = .processing
        case ReportStatus.failed: self = .failed
        default: self = .complete
        }
    }
}

/// A report's visibility, with its glyph (spec F10.1).
enum ReportVisibility: Equatable, Sendable {
    case `public`, unlisted, guild, `private`

    init(_ word: String) {
        switch word {
        case "public": self = .public
        case "unlisted": self = .unlisted
        case "guild": self = .guild
        default: self = .private
        }
    }

    var systemImage: String {
        switch self {
        case .public: "globe"
        case .unlisted: "link"
        case .guild: "person.2"
        case .private: "lock"
        }
    }
}

/// The list's rules (spec F10.1): pure data in, data out. Every sentence is in `ReportsCopy`.
enum ReportsList {
    /// `This week` holds reports created within this window of now.
    static let thisWeek: TimeInterval = 7 * 24 * 60 * 60

    /// The rows split into `This week` and `Earlier`, each keeping the server's newest-first order.
    static func sections(_ rows: [ReportRow], now: Date) -> (thisWeek: [ReportRow], earlier: [ReportRow]) {
        let recent = rows.filter { now.timeIntervalSince($0.createdAt) < thisWeek }
        let older = rows.filter { now.timeIntervalSince($0.createdAt) >= thisWeek }
        return (recent, older)
    }

    /// Whole seconds since the companion last spoke (heartbeat, else the last upload), never negative.
    static func secondsSinceUpdate(_ row: ReportRow, now: Date) -> Int {
        let last = row.heartbeatAt ?? row.lastUploadAt ?? row.createdAt
        return max(0, Int(now.timeIntervalSince(last)))
    }

    /// The row's crest: the logging character's class from the account, else the current
    /// character's class, else nil (the neutral crest).
    static func crestClass(_ row: ReportRow, characters: [AccountCharacter], currentClass: String?) -> String? {
        let logger = row.loggingCharacter.flatMap { logger in characters.first { $0.key == logger.key } }
        return logger?.className ?? currentClass
    }

    /// The row's name: the report's title when it has one, the zone otherwise.
    static func name(_ row: ReportRow) -> String {
        row.title.isEmpty ? row.zone : row.title
    }
}

/// One row of the Fights card (spec F10.2).
enum FightRow: Equatable, Identifiable, Sendable {
    /// A boss pull, kill or wipe; opens its meters.
    case boss(FightEntry)
    /// A zone's trash, folded by the server into one entry; opens the combined meters.
    case trash(TrashGroup)
    /// A trash pull with no zone, never grouped.
    case loneTrash(FightEntry)
    /// The fight a live report is still recording; no meters yet.
    case inProgress(FightEntry)

    var id: String {
        switch self {
        case let .boss(fight), let .loneTrash(fight), let .inProgress(fight): "fight-\(fight.index)"
        case let .trash(group): "trash-\(group.key)"
        }
    }
}

/// The Fights card's rows from the report and its server-folded `trash_groups` (never grouped here).
enum FightRows {
    static func rows(_ report: Report) -> [FightRow] {
        let groups = report.trashGroups ?? []
        let grouped = Set(groups.flatMap(\.fights))
        var placed: [(index: Int, after: Int, row: FightRow)] = report.fights.compactMap { fight in
            if fight.kind != FightKind.encounter, grouped.contains(fight.index) { return nil }
            if fight.inProgress == true { return (fight.index, 0, .inProgress(fight)) }
            return (fight.index, 0, fight.kind == FightKind.encounter ? .boss(fight) : .loneTrash(fight))
        }
        placed += groups.map { (anchor(of: $0, in: report.fights), 1, .trash($0)) }
        return placed.sorted { ($0.index, $0.after) < ($1.index, $1.after) }.map(\.row)
    }

    /// A zone's trash sits after that zone's last boss pull; a zone without one keeps its first pull's place.
    static func anchor(of group: TrashGroup, in fights: [FightEntry]) -> Int {
        let bosses = fights.filter { $0.kind == FightKind.encounter && $0.zone == group.zone }.map(\.index)
        return bosses.max() ?? group.fights.min() ?? Int.max
    }

    /// The label's count: boss pulls only; trash is never counted.
    static func bossPulls(_ report: Report) -> Int {
        report.fights.filter { $0.kind == FightKind.encounter }.count
    }

    /// A wipe's boss health as a whole percent, when the log has it.
    static func bossPercent(_ fight: FightEntry) -> Int? {
        guard let pct = fight.bossHealthPct, pct >= 0 else { return nil }
        return Int(pct.rounded())
    }
}

/// The parse number's colour band (spec §4 Parse number): the rankings page's floors.
enum ParseTier: Equatable, Sendable, CaseIterable {
    case grey, green, blue, purple, orange, pink

    init(_ percentile: Int) {
        switch percentile {
        case 99...: self = .pink
        case 95...: self = .orange
        case 75...: self = .purple
        case 50...: self = .blue
        case 25...: self = .green
        default: self = .grey
        }
    }
}

/// The three meters and their API words (spec F10.3).
enum MeterMetric: String, CaseIterable, Equatable, Sendable, Identifiable {
    case damage, healing, taken

    var id: String { rawValue }

    /// The per-second unit the row prints (`report.unit.*`).
    var unit: String {
        switch self {
        case .damage: "dps"
        case .healing: "hps"
        case .taken: "dtps"
        }
    }

    func value(_ player: MeterPlayer) -> Double {
        switch self {
        case .damage: player.dps
        case .healing: player.hps
        case .taken: player.dtps
        }
    }
}

/// The meter rows' rules: the cut, the bar, the server's order kept.
enum MeterRows {
    /// Rows shown before `Show all`, at the default sizes and at the accessibility sizes (spec §9).
    static let defaultVisible = 11
    static let largeVisible = 4
    static let trashScope = "trash"

    static func visible(_ players: [MeterPlayer], limit: Int, showAll: Bool) -> [MeterPlayer] {
        showAll ? players : Array(players.prefix(limit))
    }

    static func hasMore(_ players: [MeterPlayer], limit: Int, showAll: Bool) -> Bool {
        !showAll && players.count > limit
    }

    /// The bar's length: the value over the top row's value, clamped to 0...1.
    static func fraction(_ value: Double, top: Double) -> Double {
        guard top > 0 else { return 0 }
        return min(1, max(0, value / top))
    }
}

/// What the Your parses card shows for one row (spec F10.2).
struct ParseRow: Equatable, Identifiable, Sendable {
    let character: ReportMeCharacter
    let fight: ReportMeFight

    var id: String { "\(character.key)-\(fight.index)" }
    var tier: ParseTier? { fight.percentile.map(ParseTier.init) }
}

enum ParseRows {
    /// One row per character fight, in the order the server gives them.
    static func rows(_ me: ReportMe) -> [ParseRow] {
        me.characters.flatMap { character in character.fights.map { ParseRow(character: character, fight: $0) } }
    }
}

/// One section of the coach card (spec F10.4); each line prints as given.
struct CoachSection: Equatable, Sendable, Identifiable {
    enum Kind: Equatable, Sendable { case fixFirst, practice, then, doingWell, couldntCheck }
    struct Line: Equatable, Sendable {
        let title: String?
        let body: String
    }

    let kind: Kind
    let lines: [Line]
    var id: Kind { kind }
}

/// The coach card's state: hidden (free, no line, or still asking), the coach's own busy words,
/// the review, or its error.
enum CoachState: Equatable, Sendable {
    case hidden, writing, preparing, ready(CoachReview), failed
}

enum Coach {
    static let writing = "writing"
    static let preparing = "preparing"

    /// The character to ask about: only with `log_coach`, the first in /me with a kill.
    static func player(logCoach: Bool, me: ReportMe?) -> String? {
        guard logCoach, let me else { return nil }
        return me.characters.first { !$0.fights.isEmpty }?.key
    }

    static func state(_ review: CoachReview) -> CoachState {
        switch review.status {
        case writing: .writing
        case preparing: .preparing
        default: .ready(review)
        }
    }

    /// The sections present, in the card's order.
    static func sections(_ doc: CoachReviewDoc) -> [CoachSection] {
        var sections: [CoachSection] = []
        if let fix = doc.fixFirst {
            sections.append(CoachSection(kind: .fixFirst, lines: [.init(title: fix.title, body: fix.why)]))
            if !fix.drillText.isEmpty {
                sections.append(CoachSection(kind: .practice, lines: [.init(title: nil, body: fix.drillText)]))
            }
        }
        let groups: [(CoachSection.Kind, [CoachSection.Line])] = [
            (.then, doc.thenFix.map { .init(title: $0.title, body: $0.why) }),
            (.doingWell, doc.doingWell.map { .init(title: nil, body: $0.text) }),
            (.couldntCheck, doc.couldntCheck.map { .init(title: nil, body: $0) }),
        ]
        sections += groups.filter { !$0.1.isEmpty }.map { CoachSection(kind: $0.0, lines: $0.1) }
        return sections
    }
}

/// A live report re-reads itself while the screen shows (spec F10.2 Live).
enum ReportsLive {
    static let pollInterval: Duration = .seconds(20)

    static func polls(_ report: Report?) -> Bool {
        report?.status == ReportStatus.live
    }
}
