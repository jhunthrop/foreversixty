import Foundation
import Models

/// One roster row: a sign-up, or a member not signed up (state nil).
struct RosterRow: Identifiable, Equatable, Sendable {
    let characterKey: String
    let name: String
    let className: String
    let spec: String
    let specName: String
    let specIcon: String?
    let role: String
    let state: SignupState?
    let note: String
    let own: Bool

    var id: String { characterKey }

    init(_ signup: RaidSignup) {
        characterKey = signup.characterKey
        name = signup.name
        className = signup.className
        spec = signup.spec
        specName = signup.specName ?? ""
        specIcon = signup.specIcon
        role = signup.role
        state = SignupState(rawValue: signup.state)
        note = signup.note
        own = signup.own
    }

    init(_ member: RaidNotSigned, own: Bool) {
        characterKey = member.characterKey
        name = member.name
        className = member.className
        spec = member.spec
        specName = member.specName ?? ""
        specIcon = member.specIcon
        role = member.role ?? "dps"
        state = nil
        note = ""
        self.own = own
    }

    /// The same row in another place (an officer's optimistic move).
    func moved(to state: SignupState?) -> RosterRow {
        RosterRow(copying: self, state: state)
    }

    private init(copying row: RosterRow, state: SignupState?) {
        characterKey = row.characterKey
        name = row.name
        className = row.className
        spec = row.spec
        specName = row.specName
        specIcon = row.specIcon
        role = row.role
        self.state = state
        note = row.note
        own = row.own
    }
}

/// The roster's groups (spec F4.2 item 8; the site's raid-roster.ts): In by role, then Late,
/// Bench, Out and Not signed up. Within a group: class in the game's order, then name.
struct RosterGroups: Equatable, Sendable {
    enum Kind: String, CaseIterable, Sendable {
        case tanks, healers, dps, late, bench, out

        var label: String {
            switch self {
            case .tanks: Copy.raidGroupTanks
            case .healers: Copy.raidGroupHealers
            case .dps: Copy.raidGroupDps
            case .late: Copy.raidGroupLate
            case .bench: Copy.raidGroupBench
            case .out: Copy.raidGroupOut
            }
        }

        /// `No tanks yet.` for an empty role; nil for Late, Bench and Out, which are left out.
        var emptyLine: String? {
            switch self {
            case .tanks: Copy.raidEmptyRoleTank
            case .healers: Copy.raidEmptyRoleHealer
            case .dps: Copy.raidEmptyRoleDps
            case .late, .bench, .out: nil
            }
        }
    }

    struct Group: Equatable, Sendable, Identifiable {
        let kind: Kind
        let rows: [RosterRow]
        var id: Kind { kind }
    }

    let groups: [Group]
    let notSigned: [RosterRow]

    init(raid: RaidView, overrides: [String: SignupState?] = [:]) {
        let ownKeys = Set(raid.viewer.characters.map(\.characterKey))
        let base = raid.signups.map(RosterRow.init) + raid.notSigned.map { RosterRow($0, own: ownKeys.contains($0.characterKey)) }
        let rows = base.map { row in overrides[row.characterKey].map { row.moved(to: $0) } ?? row }
        let sorted = rows.sorted(by: Self.order)
        func inState(_ state: SignupState, role: String? = nil) -> [RosterRow] {
            sorted.filter { $0.state == state && (role == nil || $0.role == role) }
        }
        let all: [Group] = [
            Group(kind: .tanks, rows: inState(.signedIn, role: "tank")),
            Group(kind: .healers, rows: inState(.signedIn, role: "healer")),
            Group(kind: .dps, rows: inState(.signedIn, role: "dps")),
            Group(kind: .late, rows: inState(.late)),
            Group(kind: .bench, rows: inState(.bench)),
            Group(kind: .out, rows: inState(.out)),
        ]
        groups = all.filter { !$0.rows.isEmpty || $0.kind.emptyLine != nil }
        notSigned = sorted.filter { $0.state == nil }
    }

    static func order(_ a: RosterRow, _ b: RosterRow) -> Bool {
        let classA = GameClass.rank(a.className), classB = GameClass.rank(b.className)
        if classA != classB { return classA < classB }
        return a.name.localizedCompare(b.name) == .orderedAscending
    }
}

/// The classes strip (spec F4.2 item 6): In plus Late per class, in the game's order. A count of
/// the view's own rows.
struct ClassStrip: Equatable, Sendable {
    struct Entry: Equatable, Sendable, Identifiable {
        let gameClass: GameClass
        let count: Int
        var id: GameClass { gameClass }
    }

    let entries: [Entry]

    init(raid: RaidView) {
        let coming = raid.signups.filter { SignupState(rawValue: $0.state)?.isComing == true }
        entries = GameClass.allCases.map { cls in
            Entry(gameClass: cls, count: coming.filter { $0.className.lowercased() == cls.rawValue }.count)
        }
    }

    /// `Classes coming: Warrior 4, Hunter 2, …, no Paladins`.
    var accessibilityLabel: String {
        let present = entries.filter { $0.count > 0 }.map { Copy.raidsA11yClassCount(className: $0.gameClass.name, n: $0.count) }
        let missing = entries.filter { $0.count == 0 }.map(\.gameClass.none)
        return Copy.raidsA11yClasses(list: (present + missing).joined(separator: ", "))
    }
}

/// The officer's Copy the list (the site's copyList): the title and guild time, the count and its
/// breakdown, then each non-empty group's names.
enum RosterCopy {
    static func text(raid: RaidView, clock: RaidClock, now: Date) -> String {
        let facts = RaidFacts(raid: raid, now: now)
        let guildZone = TimeZone(identifier: raid.guildTimezone) ?? clock.zone
        let header = [
            raid.title,
            Copy.raidsWhenDate(date: clock.shortDate(raid.startsAt), clock: clock.guildClock(raid.startsAt, guildZone: guildZone)),
            "<t:\(Int(raid.startsAt.timeIntervalSince1970)):F>",
        ].joined(separator: " · ")
        let size = raid.size.map { Copy.raidsCopySize(n: facts.signedUp, size: $0) } ?? Copy.raidCountNoSize(n: facts.signedUp)
        let roster = RosterGroups(raid: raid)
        let groups = roster.groups.filter { !$0.rows.isEmpty }.map {
            Copy.raidsCopyGroup(label: $0.kind.label, names: $0.rows.map(\.name).joined(separator: ", "))
        }
        return ([header, Copy.raidsCopyCount(size: size, breakdown: facts.breakdownParts.joined(separator: ", "))] + groups).joined(separator: "\n")
    }
}
