import Foundation
import Models

/// Where a history row goes (spec F16.3).
enum HistoryTarget: Equatable, Sendable {
    /// Your raid night (F4.22): you were there and the night has a planned raid.
    case night(raid: Int64)
    /// The report on the web (F10 is not built in the app yet).
    case report(URL)
}

/// One logged raid night, worded (spec F16.3).
struct HistoryNight: Equatable, Identifiable, Sendable {
    /// Ten or fewer is a dungeon: `5 players`, not raiders.
    static let dungeonSize = 10

    let id: String
    let zone: String
    /// The night's start: created_at minus its length.
    let start: Date
    let bosses: Int
    let kills: Int
    let wipes: Int
    let size: Int
    let minutes: Int
    let raidID: Int64?
    let youWereThere: Bool

    init(row: RaidNightRow, yourKeys: Set<String>) {
        id = row.id
        zone = row.zone
        start = row.createdAt.addingTimeInterval(-TimeInterval(row.durationMs) / 1000)
        bosses = Set(row.fights.map(\.name)).count
        kills = row.killCount
        wipes = row.wipeCount
        // The night's size is never below the kill's player count (wow-player verdict).
        size = max(row.raiders, row.fights.filter(\.kill).map(\.players).max() ?? 0)
        minutes = Int((Double(row.durationMs) / 60000).rounded())
        raidID = row.raidId
        youWereThere = row.present.contains { yourKeys.contains($0.characterKey) }
    }

    var isDungeon: Bool { size <= Self.dungeonSize }

    private var wipeWords: String { wipes == 0 ? Copy.historyNoWipes : Copy.historyWipes(n: wipes) }
    private var sizeWords: String { isDungeon ? Copy.historyPlayers(n: size) : Copy.historyRaiders(n: size) }

    /// `1 boss · 1 kill · no wipes`.
    var factsLine: String {
        [Copy.historyBosses(n: bosses), Copy.historyKills(n: kills), wipeWords].joined(separator: " · ")
    }

    /// `38 raiders · 1 h 47 m`, or `5 players · 1 h 38 m`.
    var sizeLine: String {
        let (hours, rest) = (minutes / 60, minutes % 60)
        let length = hours > 0 ? Copy.historyLength(h: hours, m: rest) : Copy.historyLengthMinutes(m: rest)
        return "\(sizeWords) · \(length)"
    }

    var target: HistoryTarget {
        if youWereThere, let raidID { return .night(raid: raidID) }
        return .report(RaidLinks.report(id))
    }

    /// Spec 6: `Onyxia's Lair, Thursday 17 December, 1 boss, 1 kill, no wipes, 38 raiders, 1 hour 47 minutes. You were there.`
    func accessibilityLabel(clock: RaidClock) -> String {
        let (hours, rest) = (minutes / 60, minutes % 60)
        let length = [hours > 0 ? Copy.historyA11yHours(n: hours) : nil, rest > 0 || hours == 0 ? Copy.historyA11yMinutes(n: rest) : nil]
            .compactMap { $0 }.joined(separator: " ")
        let words = [zone, clock.longDate(start), Copy.historyBosses(n: bosses), Copy.historyKills(n: kills), wipeWords, sizeWords, length]
            .joined(separator: ", ")
        return youWereThere ? "\(words). \(Copy.historyA11yThere)" : words
    }
}
