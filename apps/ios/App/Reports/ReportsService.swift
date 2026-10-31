import Foundation
import Models
import Networking

/// The report calls F10 makes. Live goes through `APIClient`; tests pass a fake.
protocol ReportsService: Sendable {
    func list(page: Int) async throws(APIError) -> ReportList
    func report(id: String) async throws(APIError) -> Report
    func me(id: String) async throws(APIError) -> ReportMe
    func meters(id: String, n: Int, metric: MeterMetric, scope: String?) async throws(APIError) -> FightMeters
    /// The night's coach line, or nil when the coach has none for this character.
    func coachNight(id: String, player: String) async throws(APIError) -> CoachNight?
    func fightCoach(id: String, n: Int, player: String) async throws(APIError) -> CoachReview
}

struct LiveReportsService: ReportsService {
    let client: APIClient

    /// What `APIClient` throws for `ok: true` with `data: null`: the coach's "no line".
    static let emptyEnvelope = "envelope ok with no data"
    /// The server's fold: each zone's trash as one entry.
    static let trashFold = "trash"

    func list(page: Int) async throws(APIError) -> ReportList {
        try await client.send(.listMyReports(page: page))
    }

    func report(id: String) async throws(APIError) -> Report {
        try await client.send(.getReport(id: id, fold: Self.trashFold))
    }

    func me(id: String) async throws(APIError) -> ReportMe { try await client.send(.getReportMe(id: id)) }

    func meters(id: String, n: Int, metric: MeterMetric, scope: String?) async throws(APIError) -> FightMeters {
        try await client.send(.getFightMeters(id: id, n: n, metric: metric.rawValue, scope: scope))
    }

    func coachNight(id: String, player: String) async throws(APIError) -> CoachNight? {
        do {
            return try await client.send(.getCoachNight(id: id, player: player))
        } catch .decoding(Self.emptyEnvelope) {
            return nil
        }
    }

    func fightCoach(id: String, n: Int, player: String) async throws(APIError) -> CoachReview {
        try await client.send(.getFightCoach(id: id, n: n, player: player))
    }
}
