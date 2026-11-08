import Foundation
import Models

/// The raid page filtered to Not answered (flows spec F3): the short role first, the mention that
/// names it, rows marked on the page, and the bench that can close the gap.
struct NotAnsweredPresenter {
    let raid: RaidView
    let guild: String
    /// Raiders marked In or Bench on this page (the row confirms with Undo).
    let marks: [String: SignupState]
    var clock = RaidClock()

    private static let roleOrder = ["tank", "healer", "dps"]

    func content(filter: NotAnsweredFilter) -> NotAnsweredContent {
        let short = shortRole
        let unanswered = unansweredRows(short: short)
        let open = unanswered.filter { $0.marked == nil }
        let facts = [guild, when] + (short.map { [Copy.homeOfficerShort(role: $0.lowerWord, n: $0.missing)] } ?? [])
        let listed = filter == .notAnswered ? unanswered : signedRows(filter.state)
        let bench = filter == .notAnswered ? benchRows(short: short, open: open) : []
        return NotAnsweredContent(
            title: raid.zoneName,
            icon: raid.zoneIcon ?? RaidIcons.generic,
            facts: facts,
            chips: chips(unanswered: open.count),
            filter: filter,
            mention: filter == .notAnswered && !open.isEmpty ? mention(open, short: short) : nil,
            listTitle: filter == .notAnswered ? Copy.raidNaTitle : filter.label,
            listMeta: filter == .notAnswered ? listMeta(open: open.count, short: short) : nil,
            rows: listed,
            actionable: filter == .notAnswered,
            benchTitle: bench.isEmpty ? nil : short.map { Copy.raidNaBench(role: $0.word) },
            bench: bench,
            benchNote: bench.isEmpty ? nil : benchNote(short: short, open: open, bench: bench)
        )
    }

    // MARK: Counts

    private var total: Int { raid.signups.count + raid.notSigned.count }

    private func count(_ state: SignupState) -> Int {
        raid.signups.filter { $0.state == state.rawValue && marks[$0.characterKey] == nil }.count
            + marks.values.filter { $0 == state }.count
    }

    private func chips(unanswered: Int) -> [NotAnsweredChip] {
        NotAnsweredFilter.allCases.map { filter in
            let n = switch filter {
            case .notAnswered: unanswered
            case .all: total
            default: filter.state.map(count) ?? 0
            }
            return NotAnsweredChip(filter: filter, title: Copy.raidNaFilter(label: filter.label, n: n))
        }
    }

    /// The role furthest short, less the raiders of it marked In here.
    var shortRole: RaidShortRole? {
        guard let short = RaidShortRole.of(raid) else { return nil }
        let marked = marks.filter { $0.value == .signedIn && role(of: $0.key) == short.role }.count
        return short.missing > marked ? RaidShortRole(role: short.role, missing: short.missing - marked) : nil
    }

    private var when: String {
        Copy.homeRaidWhen(
            date: clock.shortDate(raid.startsAt),
            time: clock.guildClock(raid.startsAt, guildZone: TimeZone(identifier: raid.guildTimezone) ?? clock.zone)
        )
    }

    private func role(of key: String) -> String? {
        raid.notSignedRaiders?.first { $0.characterKey == key }?.role ?? raid.signups.first { $0.characterKey == key }?.role
    }

    // MARK: Rows

    /// The unanswered, the short role first, then tanks, healers, damage, in the server's order.
    private func unansweredRows(short: RaidShortRole?) -> [NotAnsweredRow] {
        let raiders = raid.notSignedRaiders ?? []
        let rank = { (role: String?) -> Int in
            if let short, role == short.role { return -1 }
            return Self.roleOrder.firstIndex(of: role ?? "dps") ?? Self.roleOrder.count
        }
        return raiders.enumerated().sorted { a, b in
            (rank(a.element.role), a.offset) < (rank(b.element.role), b.offset)
        }.map { _, raider in
            let spec = raid.notSigned.first { $0.characterKey == raider.characterKey }?.specName
            return NotAnsweredRow(
                characterKey: raider.characterKey, name: raider.name, className: raider.className,
                sub: Copy.raidNaSub(kind: RaidFlowPresenter.kind(spec: spec, className: raider.className), line: Self.lastWeek(raider.lastWeek)),
                role: raider.role, marked: marks[raider.characterKey], dm: raider.discordDM
            )
        }
    }

    static func lastWeek(_ state: String?) -> String {
        guard let state else { return Copy.raidNaNew }
        guard let answer = SignupState(rawValue: state) else { return Copy.raidNaMissed }
        return Copy.raidNaLastWeek(state: answer.label)
    }

    /// The sign-ups in a state (All: everyone), for the other chips.
    private func signedRows(_ state: SignupState?) -> [NotAnsweredRow] {
        let signed = raid.signups.filter { state == nil || $0.state == state?.rawValue }.map { signup in
            NotAnsweredRow(
                characterKey: signup.characterKey, name: signup.name, className: signup.className,
                sub: RaidFlowPresenter.kind(spec: signup.specName, className: signup.className),
                role: signup.role, marked: nil
            )
        }
        return state == nil ? signed + unansweredRows(short: shortRole) : signed
    }

    /// The short role's bench, when the unanswered of it cannot close the gap.
    private func benchRows(short: RaidShortRole?, open: [NotAnsweredRow]) -> [NotAnsweredRow] {
        guard let short else { return [] }
        let openOfRole = open.filter { $0.role == short.role }.count
        guard openOfRole < short.missing else { return [] }
        return raid.signups.filter { $0.state == SignupState.bench.rawValue && $0.role == short.role }.map { signup in
            NotAnsweredRow(
                characterKey: signup.characterKey, name: signup.name, className: signup.className,
                sub: Copy.raidNaSub(kind: RaidFlowPresenter.kind(spec: signup.specName, className: signup.className), line: Copy.raidNaOnBench),
                role: signup.role, marked: marks[signup.characterKey]
            )
        }
    }

    private func benchNote(short: RaidShortRole?, open: [NotAnsweredRow], bench: [NotAnsweredRow]) -> String? {
        guard let short else { return nil }
        let names = open.filter { $0.role == short.role }.map(\.name) + bench.filter { $0.marked == nil }.map(\.name)
        guard names.count >= short.missing else { return nil }
        return Copy.raidNaFills(names: RaidFlowPresenter.list(names), role: short.lowerWord)
    }

    // MARK: The mention

    private func listMeta(open: Int, short: RaidShortRole?) -> String {
        short.map { Copy.raidNaMeta(n: open, total: total, role: $0.lowerWord) } ?? Copy.raidNaMetaPlain(n: open, total: total)
    }

    private func mention(_ open: [NotAnsweredRow], short: RaidShortRole?) -> MentionContent {
        let ofRole = short.map { short in open.filter { $0.role == short.role } } ?? []
        let named = ofRole.isEmpty ? open : ofRole
        let names = named.map { "@\($0.name)" }.joined(separator: " ")
        let preview = if let short, !ofRole.isEmpty {
            Copy.raidNaPreview(names: names, zone: raid.zoneName, when: when, role: short.lowerWord)
        } else {
            Copy.raidNaPreviewAll(names: names, zone: raid.zoneName, when: when)
        }
        return MentionContent(
            role: short.flatMap { ofRole.isEmpty ? nil : Copy.raidNaMentionRole(people: Self.people($0.role, n: ofRole.count)) },
            all: Copy.raidNaMentionAll(n: open.count),
            preview: preview
        )
    }

    /// `2 healers`.
    static func people(_ role: String, n: Int) -> String {
        switch role {
        case "tank": Copy.raidsPartTanks(n: n)
        case "healer": Copy.raidsPartHealers(n: n)
        default: Copy.raidsPartDps(n: n)
        }
    }
}
