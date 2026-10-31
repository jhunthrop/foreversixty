@testable import ForeverSixty
import Foundation
import Models
import Networking

/// F10's test data: the contract fixtures (one Onyxia night), and variants derived from them with
/// the arithmetic kept. Every helper throws, so a missing file fails only its own test.
enum ReportsTestData {
    static let build = "1.60.1.70658"
    /// Twenty seconds after the live report's heartbeat (reports.list).
    static let now = RaidTime.date("2026-12-16T00:32:00Z")
    /// The guild's evening in US Central: the Onyxia night is Monday 14 December there.
    static let clock = ReportsClock(zone: RaidTime.chicago)
    static let reportID = "ewlrx3enxbeb"

    struct Failure: Error, CustomStringConvertible { let description: String }

    static func list() throws -> ReportList { try RaidFixture.load(ReportList.self, "reports.list") }
    static func report() throws -> Report { try RaidFixture.load(Report.self, "reports.get") }
    static func me() throws -> ReportMe { try RaidFixture.load(ReportMe.self, "reports.me") }
    static func meters() throws -> FightMeters { try RaidFixture.load(FightMeters.self, "reports.meters") }
    static func review() throws -> CoachReview { try RaidFixture.load(CoachReview.self, "coach.fight") }
    static func account(premium: Bool) throws -> ReportsAccount {
        let me = try RaidFixture.load(Me.self, premium ? "me.premium" : "me")
        return ReportsAccount.of(me, current: me.characters.first)
    }

    /// The fixture JSON's `data`, to derive a variant from.
    static func data(_ name: String) throws -> [String: Any] {
        let root = try JSONSerialization.jsonObject(with: RaidFixture.data(name)) as? [String: Any]
        guard let data = root?["data"] as? [String: Any] else { throw Failure(description: "\(name) has no data") }
        return data
    }

    static func decode<Value: Decodable>(_ type: Value.Type, _ object: Any) throws -> Value {
        try WireFormat.decoder().decode(type, from: JSONSerialization.data(withJSONObject: object))
    }

    /// reports.get with a second zone: two Molten Core trash pulls folded into one group whose
    /// length is their sum.
    static func twoZones() throws -> Report {
        var data = try data("reports.get")
        var fights = data["fights"] as? [[String: Any]] ?? []
        guard let template = fights.first else { throw Failure(description: "reports.get has no fights") }
        let added: [(Int, Int)] = [(10, 70000), (11, 85000)]
        for (index, ms) in added {
            var fight = template
            fight["index"] = index
            fight["zone"] = "Molten Core"
            fight["duration_ms"] = ms
            fights.append(fight)
        }
        data["fights"] = fights
        var groups = data["trash_groups"] as? [[String: Any]] ?? []
        groups.append(["key": 2, "zone": "Molten Core", "fights": added.map(\.0), "duration_ms": added.map(\.1).reduce(0, +)])
        data["trash_groups"] = groups
        return try decode(Report.self, data)
    }

    /// reports.get without fight 9 (the kill): only the two wipes and the trash.
    static func onlyWipes() throws -> Report {
        var data = try data("reports.get")
        data["fights"] = (data["fights"] as? [[String: Any]] ?? []).filter { ($0["index"] as? Int) != 9 }
        return try decode(Report.self, data)
    }

    /// reports.me without its fight 9 row: no kills of yours.
    static func meWithoutKill() throws -> ReportMe {
        try RaidFixture.load(ReportMe.self, "reports.me.empty")
    }

    /// The live Blackrock Spire run (reports.get.live): four trash pulls, no boss.
    static func liveReport() throws -> Report { try RaidFixture.load(Report.self, "reports.get.live") }

    /// A processing or failed report: the list's row of that status (Molten Core, Stratholme) as a
    /// report with no fights yet.
    static func report(status: String) throws -> Report {
        guard let row = try list().rows.first(where: { $0.status == status }) else { throw Failure(description: "no \(status) row") }
        var data = try data("reports.get")
        data["id"] = row.id
        data["title"] = row.title
        data["zone"] = row.zone
        data["visibility"] = row.visibility
        data["status"] = status
        data["created_at"] = ISO8601DateFormatter().string(from: row.createdAt)
        data["trash_groups"] = [] as [Any]
        data["fights"] = [] as [Any]
        return try decode(Report.self, data)
    }

    /// The healing meter: the fixture's players sorted by hps, those with none dropped.
    static func healing() throws -> FightMeters {
        let meters = try meters()
        return FightMeters(
            fightIndex: meters.fightIndex, scope: meters.scope, metric: MeterMetric.healing.rawValue, kill: meters.kill,
            durationMs: meters.durationMs, players: meters.players.filter { $0.hps > 0 }.sorted { $0.hps > $1.hps }
        )
    }

    /// Damage taken: the fixture's players sorted by dtps.
    static func taken() throws -> FightMeters {
        let meters = try meters()
        return FightMeters(
            fightIndex: meters.fightIndex, scope: meters.scope, metric: MeterMetric.taken.rawValue, kill: meters.kill,
            durationMs: meters.durationMs, players: meters.players.sorted { $0.dtps > $1.dtps }
        )
    }

    /// The Onyxia trash group folded by the server (six pulls, 7:40).
    static func trashMeters() throws -> FightMeters { try RaidFixture.load(FightMeters.self, "reports.meters.trash") }

    static func fightRow(_ index: Int) throws -> FightRow {
        guard let fight = try report().fights.first(where: { $0.index == index }) else { throw Failure(description: "no fight \(index)") }
        return .boss(fight)
    }

    static func trashRow() throws -> FightRow {
        guard let group = try report().trashGroups?.first else { throw Failure(description: "no trash group") }
        return .trash(group)
    }
}
