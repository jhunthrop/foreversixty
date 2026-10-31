import Foundation
import Models

/// The role chips (spec F16.1). They read `members[].role`, never a spec word.
enum RosterRoleFilter: String, CaseIterable, Identifiable, Sendable {
    case all, tanks, healers, dps

    var id: String { rawValue }

    var label: String {
        switch self {
        case .all: Copy.rosterChipAll
        case .tanks: Copy.rosterChipTanks
        case .healers: Copy.rosterChipHealers
        case .dps: Copy.rosterChipDps
        }
    }

    func admits(_ role: String?) -> Bool {
        switch self {
        case .all: true
        case .tanks: role == "tank"
        case .healers: role == "healer"
        case .dps: role == "dps"
        }
    }
}

/// `Sort: {by}` (spec F16.1).
enum RosterSort: String, CaseIterable, Identifiable, Sendable {
    case rank, className, name, attendance, level

    var id: String { rawValue }

    var label: String {
        switch self {
        case .rank: Copy.rosterSortRank
        case .className: Copy.rosterSortClassName
        case .name: Copy.rosterSortName
        case .attendance: Copy.rosterSortAttendance
        case .level: Copy.rosterSortLevel
        }
    }
}

/// What a row shows on its right.
enum RosterStanding: Equatable, Sendable {
    case attendance(present: Int, nights: Int)
    /// The guild has no logged nights yet.
    case noNights
    case waiting
    /// In game, no account: no attendance, no page.
    case noAccount
    /// A member without a roster row: nothing to show.
    case unknown

    init(present: Int, nights: Int) {
        self = nights > 0 ? .attendance(present: min(present, nights), nights: nights) : .noNights
    }
}

/// The attendance bar: one segment per night, at most eight (spec F16.1).
struct AttendanceBar: Equatable, Sendable {
    static let maxSegments = 8
    let segments: Int
    let filled: Int

    init(present: Int, nights: Int) {
        segments = min(max(nights, 0), Self.maxSegments)
        filled = min(max(present, 0), segments)
    }
}

/// One roster row, worded.
struct RosterEntry: Equatable, Identifiable, Sendable {
    let id: String
    let name: String
    let className: String?
    let specName: String?
    let specIcon: String?
    let role: String?
    let level: Int?
    let rank: String
    let characterKey: String?
    let supporter: Bool
    let standing: RosterStanding

    init(member: GuildMember, row: GuildRosterRow?) {
        id = member.characterKey ?? "name:\(member.name)"
        name = member.name
        className = member.className
        specName = member.specName
        specIcon = member.specIcon
        role = member.role
        level = member.level
        rank = member.rank
        characterKey = member.characterKey
        supporter = row?.supporter ?? false
        standing = switch member.state {
        case "waiting": .waiting
        case "no_account": .noAccount
        default: member.characterKey == nil ? .noAccount : Self.standing(row)
        }
    }

    private static func standing(_ row: GuildRosterRow?) -> RosterStanding {
        guard let attendance = row?.attendance else { return .unknown }
        return RosterStanding(present: attendance.present ?? 0, nights: attendance.nights ?? 0)
    }

    static func rankWord(_ rank: String) -> String {
        switch rank {
        case "leader": Copy.rosterRankLeader
        case "officer": Copy.rosterRankOfficer
        default: Copy.rosterRankMember
        }
    }

    /// Leader, officer, member.
    static func rankOrder(_ rank: String) -> Int {
        switch rank {
        case "leader": 0
        case "officer": 1
        default: 2
        }
    }

    var rankWord: String { Self.rankWord(rank) }

    /// The full spec name (`Holy Priest`), else the class name.
    var specWords: String {
        specName ?? className.flatMap(GameClass.init(rawValue:))?.name ?? className?.capitalized ?? ""
    }

    /// `{Spec} · {level} · {rank}`.
    var meta: String {
        level.map { Copy.rosterMeta(spec: specWords, level: $0, rank: rankWord) }
            ?? Copy.rosterMetaNoLevel(spec: specWords, rank: rankWord)
    }

    var bar: AttendanceBar? {
        if case let .attendance(present, nights) = standing { return AttendanceBar(present: present, nights: nights) }
        return nil
    }

    /// `6 of 8` under the bar.
    var attendanceCaption: String? {
        if case let .attendance(present, nights) = standing { return Copy.rosterAttendance(present: present, nights: nights) }
        return nil
    }

    /// The character's page on the site; nil for a name no account has claimed.
    var pageURL: URL? {
        guard standing != .noAccount, let key = characterKey, (try? CharacterKey(key)) != nil else { return nil }
        return StaticSite.url.appending(path: "character/\(key)")
    }

    /// Spec 6: `Kraggor, Protection Warrior, level 60, Guild Master, attended 8 of 8 recent raid nights`.
    var accessibilityLabel: String {
        let tail: String? = switch standing {
        case let .attendance(present, nights): Copy.rosterA11yAttended(present: present, nights: nights)
        case .noNights: Copy.rosterNoNights
        case .waiting: Copy.rosterA11yWaiting
        case .noAccount: Copy.rosterNoAccountTag
        case .unknown: nil
        }
        return [name, specWords, level.map { Copy.rosterA11yLevel(level: $0) }, rankWord, tail]
            .compactMap { $0 }.filter { !$0.isEmpty }.joined(separator: ", ")
    }

    /// Present nights over nights, for the attendance sort; nil sorts last.
    fileprivate var attendanceRatio: Double? {
        if case let .attendance(present, nights) = standing { return Double(present) / Double(nights) }
        return nil
    }
}

/// The roster read from `GET /v1/guilds/{id}/home`: members joined to roster rows by character key.
struct RosterList: Equatable, Sendable {
    let entries: [RosterEntry]
    let counts: GuildMemberCounts

    init(home: GuildHome) {
        let rows = Dictionary(home.roster.map { ($0.characterKey, $0) }, uniquingKeysWith: { first, _ in first })
        entries = home.members.map { member in RosterEntry(member: member, row: member.characterKey.flatMap { rows[$0] }) }
        counts = home.counts
    }

    func visible(role: RosterRoleFilter, query: String, sort: RosterSort) -> [RosterEntry] {
        let needle = query.trimmingCharacters(in: .whitespacesAndNewlines)
        return entries
            .filter { role.admits($0.role) }
            .filter { needle.isEmpty || $0.name.range(of: needle, options: [.caseInsensitive, .diacriticInsensitive]) != nil }
            .sorted { Self.precedes($0, $1, by: sort) }
    }

    private static func byName(_ lhs: RosterEntry, _ rhs: RosterEntry) -> Bool {
        lhs.name.localizedCaseInsensitiveCompare(rhs.name) == .orderedAscending
    }

    private static func precedes(_ lhs: RosterEntry, _ rhs: RosterEntry, by sort: RosterSort) -> Bool {
        switch sort {
        case .name:
            return byName(lhs, rhs)
        case .rank:
            let (left, right) = (RosterEntry.rankOrder(lhs.rank), RosterEntry.rankOrder(rhs.rank))
            return left != right ? left < right : byName(lhs, rhs)
        case .className:
            let (left, right) = (GameClass.rank(lhs.className ?? ""), GameClass.rank(rhs.className ?? ""))
            return left != right ? left < right : byName(lhs, rhs)
        case .level:
            let (left, right) = (lhs.level ?? 0, rhs.level ?? 0)
            return left != right ? left > right : byName(lhs, rhs)
        case .attendance:
            switch (lhs.attendanceRatio, rhs.attendanceRatio) {
            case let (left?, right?) where left != right: return left > right
            case (.some, nil): return true
            case (nil, .some): return false
            default: return byName(lhs, rhs)
            }
        }
    }
}
