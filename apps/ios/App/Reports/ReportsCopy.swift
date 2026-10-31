import Foundation
import Models

/// A report's dates as the viewer reads them: `15 Dec`, `Tuesday`, `15 December 2026`. The zone
/// is injected so a test pins it.
struct ReportsClock: Sendable {
    let zone: TimeZone

    init(zone: TimeZone = .current) {
        self.zone = zone
    }

    private func format(_ template: String, _ date: Date) -> String {
        let formatter = DateFormatter()
        formatter.locale = Locale(identifier: "en_GB")
        formatter.timeZone = zone
        formatter.dateFormat = template
        return formatter.string(from: date)
    }

    /// `15 Dec`, the spec's day-month order on every locale.
    func shortDate(_ date: Date) -> String { format("d MMM", date) }
    /// `Tuesday`.
    func weekday(_ date: Date) -> String { format("EEEE", date) }
    /// `15 December 2026`.
    func longDate(_ date: Date) -> String { format("d MMMM yyyy", date) }
    /// `Tuesday 15 December`, the spoken row date.
    func spokenDate(_ date: Date) -> String { format("EEEE d MMMM", date) }
}

/// Every F10 sentence built from data (spec §5.2); the words themselves live in the copy catalog.
enum ReportsCopy {
    private static let millisPerSecond: Int64 = 1000
    private static let secondsPerMinute: Int64 = 60

    /// `4:28`.
    static func duration(_ ms: Int64) -> String {
        let seconds = max(0, ms) / millisPerSecond
        return String(format: "%lld:%02lld", seconds / secondsPerMinute, seconds % secondsPerMinute)
    }

    /// `4 minutes 28 seconds`, from the platform's duration formatter.
    static func spokenDuration(_ ms: Int64) -> String {
        let formatter = DateComponentsFormatter()
        formatter.unitsStyle = .full
        formatter.allowedUnits = [.hour, .minute, .second]
        formatter.calendar?.locale = Locale(identifier: "en_US")
        let words = formatter.string(from: TimeInterval(max(0, ms) / millisPerSecond)) ?? ""
        return words.replacingOccurrences(of: ", ", with: " ")
    }

    /// `91st`.
    static func ordinal(_ n: Int) -> String {
        let formatter = NumberFormatter()
        formatter.numberStyle = .ordinal
        formatter.locale = Locale(identifier: "en_US")
        return formatter.string(from: NSNumber(value: n)) ?? "\(n)"
    }

    /// A per-second figure rounded to a whole number, grouped: `712`, `1,043`.
    static func figure(_ value: Double) -> String {
        Int(value.rounded()).formatted(.number.grouping(.automatic).locale(Locale(identifier: "en_US")))
    }

    /// `DPS`, `HPS` or `DTPS`.
    static func unit(_ unit: String) -> String {
        switch unit {
        case "hps": Copy.reportUnitHps
        case "dtps": Copy.reportUnitDtps
        default: Copy.reportUnitDps
        }
    }

    /// `damage per second` and the like, for VoiceOver.
    static func spokenUnit(_ unit: String) -> String {
        switch unit {
        case "hps": Copy.reportA11yPerSecondHps
        case "dtps": Copy.reportA11yPerSecondDtps
        default: Copy.reportA11yPerSecondDps
        }
    }

    static func visibility(_ visibility: ReportVisibility) -> String {
        switch visibility {
        case .public: Copy.reportsVisibilityPublic
        case .unlisted: Copy.reportsVisibilityUnlisted
        case .guild: Copy.reportsVisibilityGuild
        case .private: Copy.reportsVisibilityPrivate
        }
    }

    // MARK: F10.1 list

    /// `14 Dec · 1 boss · 1 kill`, `Updated 20 s ago` (live) or `29 Nov · could not be read`.
    static func rowMeta(_ row: ReportRow, now: Date, clock: ReportsClock) -> String {
        let date = clock.shortDate(row.createdAt)
        switch ReportRowState(status: row.status) {
        case .failed:
            return Copy.reportsUnreadable(date: date)
        case .live:
            return Copy.reportsUpdated(n: ReportsList.secondsSinceUpdate(row, now: now))
        case .processing, .complete:
            return Copy.reportsRowMeta(date: date, bosses: Copy.reportsBosses(n: row.bossCount), kills: Copy.reportsKills(n: row.killCount))
        }
    }

    static func rowA11y(_ row: ReportRow, now: Date, clock: ReportsClock) -> String {
        let base = Copy.reportsA11yRow(
            zone: ReportsList.name(row), date: clock.spokenDate(row.createdAt),
            bosses: Copy.reportsBosses(n: row.bossCount), kills: Copy.reportsKills(n: row.killCount),
            visibility: visibility(ReportVisibility(row.visibility))
        )
        switch ReportRowState(status: row.status) {
        case .live: return "\(base), \(Copy.reportsA11yLive(n: "\(ReportsList.secondsSinceUpdate(row, now: now))"))"
        case .processing: return "\(base), \(Copy.reportsStateProcessing)"
        case .failed: return "\(base), \(Copy.reportsStateFailed)"
        case .complete: return base
        }
    }

    // MARK: F10.2 report

    /// `Tuesday 15 December 2026 · logged by Obnoxious Yell`.
    static func reportMeta(_ report: Report, clock: ReportsClock) -> String {
        Copy.reportMeta(
            weekday: clock.weekday(report.createdAt), date: clock.longDate(report.createdAt),
            character: report.loggingCharacter?.name ?? ""
        )
    }

    static func bossPulls(_ n: Int) -> String {
        n == 0 ? Copy.reportNoBossPulls : Copy.reportBossPulls(n: n)
    }

    /// The fight row's title.
    static func title(_ row: FightRow) -> String {
        switch row {
        case let .boss(fight), let .inProgress(fight): fight.name
        case let .trash(group): Copy.reportTrashTitle(zone: group.zone)
        case .loneTrash: Copy.reportTrash
        }
    }

    /// `Kill · 4:28 · 40 players · 3 deaths`, `Wipe · 1:12 · boss at 61%`, `6 pulls · 7:40`.
    static func meta(_ row: FightRow) -> String {
        switch row {
        case let .boss(fight): bossMeta(fight)
        case let .trash(group): Copy.reportTrashMeta(pulls: Copy.reportPulls(n: group.fights.count), time: duration(group.durationMs))
        case let .loneTrash(fight): Copy.reportTrashMeta(pulls: Copy.reportPulls(n: 1), time: duration(fight.durationMs))
        case .inProgress: Copy.reportInProgress
        }
    }

    static func bossMeta(_ fight: FightEntry) -> String {
        let time = duration(fight.durationMs)
        if fight.kill {
            return Copy.reportKillMeta(time: time, players: Copy.reportPlayers(n: fight.players.count), deaths: Copy.reportDeaths(n: fight.deaths ?? 0))
        }
        guard let percent = FightRows.bossPercent(fight) else { return Copy.reportWipeMetaPlain(time: time) }
        return Copy.reportWipeMeta(time: time, percent: "\(percent)%")
    }

    static func fightA11y(_ row: FightRow) -> String {
        switch row {
        case let .boss(fight):
            let time = spokenDuration(fight.durationMs)
            if fight.kill {
                return Copy.reportA11yKill(boss: fight.name, time: time, players: Copy.reportPlayers(n: fight.players.count), deaths: Copy.reportDeaths(n: fight.deaths ?? 0))
            }
            guard let percent = FightRows.bossPercent(fight) else { return Copy.reportA11yWipePlain(boss: fight.name, time: time) }
            return Copy.reportA11yWipe(boss: fight.name, time: time, n: percent)
        case let .trash(group):
            return Copy.reportA11yTrash(zone: group.zone, pulls: Copy.reportPulls(n: group.fights.count), time: spokenDuration(group.durationMs))
        case let .loneTrash(fight):
            return "\(Copy.reportTrash), \(Copy.reportPulls(n: 1)), \(spokenDuration(fight.durationMs))"
        case let .inProgress(fight):
            return "\(fight.name), \(Copy.reportInProgress)"
        }
    }

    /// `Kill 4:28 · Obnoxious Yell, Fury`.
    static func parseMeta(_ row: ParseRow) -> String {
        Copy.reportParseMeta(time: duration(row.fight.durationMs), name: row.character.name, spec: row.character.specName)
    }

    static func parseA11y(_ row: ParseRow) -> String {
        let fight = row.fight
        let time = spokenDuration(fight.durationMs)
        let value = figure(fight.value)
        let unit = spokenUnit(fight.unit)
        var words = fight.percentile.map {
            Copy.reportA11yParse(boss: fight.encounter, time: time, value: value, unit: unit, ordinal: ordinal($0))
        } ?? Copy.reportA11yParseNoRank(boss: fight.encounter, time: time, value: value, unit: unit)
        if fight.newBest { words += ", \(Copy.reportNewBest)" }
        return words
    }

    // MARK: F10.3 meters

    static func segment(_ metric: MeterMetric) -> String {
        switch metric {
        case .damage: Copy.metersDamage
        case .healing: Copy.metersHealing
        case .taken: Copy.metersTaken
        }
    }

    static func metersTitle(_ metric: MeterMetric) -> String {
        switch metric {
        case .damage: Copy.metersTitleDamage
        case .healing: Copy.metersTitleHealing
        case .taken: Copy.metersTitleTaken
        }
    }

    /// The sheet's head meta: `Kill · 4:28 · 40 players`, `Wipe · 1:12 · 40 players`, or trash's `6 pulls · 7:40`.
    static func metersHead(_ target: MetersTarget, meters: FightMeters?) -> String {
        let durationMs = meters?.durationMs ?? target.durationMs
        let time = duration(durationMs)
        switch target.kind {
        case .trash, .loneTrash:
            return Copy.metersTrashHead(pulls: Copy.reportPulls(n: target.pulls), time: time)
        case .boss:
            // The pull's roster, as the Fights row counts it: a healing meter lists only healers.
            let players = Copy.reportPlayers(n: target.players)
            return (meters?.kill ?? target.kill) ? Copy.metersKillHead(time: time, players: players) : Copy.metersWipeHead(time: time, players: players)
        }
    }

    /// `Fury Warrior`.
    static func specClass(_ player: MeterPlayer) -> String {
        let className = player.className.prefix(1).uppercased() + player.className.dropFirst()
        return player.specName.isEmpty ? className : "\(player.specName) \(className)"
    }

    static func meterA11y(rank: Int, player: MeterPlayer, metric: MeterMetric) -> String {
        let row = Copy.metersA11yRow(n: rank, name: player.name, spec: specClass(player), value: figure(metric.value(player)))
        return player.mine ? Copy.metersA11yMine(row: row) : row
    }

    // MARK: F10.4 coach

    static func coachSection(_ kind: CoachSection.Kind) -> String {
        switch kind {
        case .fixFirst: Copy.coachFixFirst
        case .practice: Copy.coachPractice
        case .then: Copy.coachThen
        case .doingWell: Copy.coachDoingWell
        case .couldntCheck: Copy.coachCouldntCheck
        }
    }

    /// `Reviewed 15 Dec · Based on: {labels}.` The labels are sentences; the copy's own full stop
    /// replaces the last one's, so the footer never ends in two.
    static func coachFooter(_ review: CoachReview, clock: ReportsClock) -> String {
        var labels = review.basisLabels.joined(separator: " ")
        if labels.hasSuffix(".") { labels.removeLast() }
        return Copy.coachFooter(date: clock.shortDate(review.reviewedAt), labels: labels)
    }

    // MARK: Home

    /// Home's Logs row sub: the newest report's zone and kills, or `No reports yet.`
    static func homeLine(_ list: ReportList) -> String {
        guard let newest = list.rows.first else { return Copy.homeLogsNone }
        return Copy.homeLogsLast(zone: newest.zone, kills: Copy.reportsKills(n: newest.killCount))
    }
}
