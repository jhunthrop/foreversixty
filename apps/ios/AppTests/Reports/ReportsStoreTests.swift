@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// A scripted reports service: each call answers from its closure; every call is recorded.
final class FakeReportsService: ReportsService, @unchecked Sendable {
    var listAnswer: (Int) throws(APIError) -> ReportList = { _ throws(APIError) in try RaidFixture.api(ReportList.self, "reports.list") }
    var reportAnswer: (String) throws(APIError) -> Report = { _ throws(APIError) in try RaidFixture.api(Report.self, "reports.get") }
    var meAnswer: (String) throws(APIError) -> ReportMe = { _ throws(APIError) in try RaidFixture.api(ReportMe.self, "reports.me") }
    var metersAnswer: (MeterMetric) throws(APIError) -> FightMeters = { _ throws(APIError) in try RaidFixture.api(FightMeters.self, "reports.meters") }
    var nightAnswer: () throws(APIError) -> CoachNight? = { () throws(APIError) in try RaidFixture.api(CoachNight.self, "coach.night") }
    var coachAnswer: () throws(APIError) -> CoachReview = { () throws(APIError) in try RaidFixture.api(CoachReview.self, "coach.fight") }
    private(set) var pages: [Int] = []
    private(set) var meCalls = 0
    private(set) var coachCalls: [String] = []
    private(set) var meterCalls: [(n: Int, metric: MeterMetric, scope: String?)] = []

    func list(page: Int) async throws(APIError) -> ReportList {
        pages.append(page)
        return try listAnswer(page)
    }

    func report(id: String) async throws(APIError) -> Report { try reportAnswer(id) }

    func me(id: String) async throws(APIError) -> ReportMe {
        meCalls += 1
        return try meAnswer(id)
    }

    func meters(id: String, n: Int, metric: MeterMetric, scope: String?) async throws(APIError) -> FightMeters {
        meterCalls.append((n, metric, scope))
        return try metersAnswer(metric)
    }

    func coachNight(id: String, player: String) async throws(APIError) -> CoachNight? {
        coachCalls.append("night:\(player)")
        return try nightAnswer()
    }

    func fightCoach(id: String, n: Int, player: String) async throws(APIError) -> CoachReview {
        coachCalls.append("fight:\(n)")
        return try coachAnswer()
    }
}

@MainActor
@Suite("Reports store")
struct ReportsStoreTests {
    private let now: @Sendable () -> Date = { ReportsTestData.now }

    private func listStore(_ service: FakeReportsService, cache: MemoryRaidCache = MemoryRaidCache(), signedIn: Bool = true) -> ReportsStore {
        ReportsStore(service: service, cache: cache, now: now, signedIn: signedIn)
    }

    private func content(_ store: ReportsStore) throws -> ReportsContent {
        guard case let .loaded(content) = store.phase else { throw APIError.missingFixture("not loaded: \(store.phase)") }
        return content
    }

    @Test func list_loadsTheFirstPage() async throws {
        let service = FakeReportsService()
        let store = listStore(service)
        await store.load()
        #expect(try content(store).rows.count == 4)
        #expect(service.pages == [1])
    }

    @Test func list_signedOut_makesNoCall() async {
        let service = FakeReportsService()
        let store = listStore(service, signedIn: false)
        await store.load()
        #expect(store.phase == .signedOut)
        #expect(service.pages.isEmpty)
    }

    @Test func list_offline_showsTheCacheUnderTheStrip() async throws {
        let cache = MemoryRaidCache()
        await listStore(FakeReportsService(), cache: cache).load()
        let offline = FakeReportsService()
        offline.listAnswer = { _ throws(APIError) in throw .transport("offline") }
        let store = listStore(offline, cache: cache)
        await store.load()
        #expect(try content(store).offlineSince == ReportsTestData.now)
    }

    @Test func list_error_withNothingCached_fails() async {
        let service = FakeReportsService()
        service.listAnswer = { _ throws(APIError) in throw .unexpectedStatus(500) }
        let store = listStore(service)
        await store.load()
        #expect(store.phase == .failed)
    }

    @Test func list_loadMore_appendsTheNextPage() async throws {
        let service = FakeReportsService()
        service.listAnswer = { page throws(APIError) in
            let list = try RaidFixture.api(ReportList.self, "reports.list")
            return ReportList(rows: page == 1 ? list.rows : [list.rows[0]], total: 5, page: page, perPage: 100)
        }
        let store = listStore(service)
        await store.load()
        #expect(try content(store).hasMorePages)
        await store.loadMore()
        #expect(try content(store).rows.count == 5)
        #expect(service.pages == [1, 2])
    }

    @Test func homeLine_fallsBackToTheCache() async {
        let cache = MemoryRaidCache()
        #expect(await ReportsStore.homeLine(service: FakeReportsService(), cache: cache, now: now) == "Your last raid: Blackrock Spire, 0 kills")
        let offline = FakeReportsService()
        offline.listAnswer = { _ throws(APIError) in throw .transport("offline") }
        #expect(await ReportsStore.homeLine(service: offline, cache: cache, now: now) == "Your last raid: Blackrock Spire, 0 kills")
        #expect(await ReportsStore.homeLine(service: offline, cache: MemoryRaidCache(), now: now) == nil)
    }

    // MARK: A report

    private func reportStore(_ service: FakeReportsService, premium: Bool = true, signedIn: Bool = true, cache: MemoryRaidCache = MemoryRaidCache()) throws -> ReportStore {
        let account = signedIn ? try ReportsTestData.account(premium: premium) : .signedOut
        return ReportStore(id: ReportsTestData.reportID, account: account, service: service, cache: cache, now: now)
    }

    @Test func report_premium_loadsParsesAndTheCoach() async throws {
        let service = FakeReportsService()
        let store = try reportStore(service)
        await store.load()
        #expect(store.report?.id == ReportsTestData.reportID)
        #expect(store.parses == .loaded(try ReportsTestData.me()))
        #expect(store.coach == .ready(try ReportsTestData.review()))
        #expect(service.coachCalls == ["night:us/pvp/obnoxious-yell", "fight:9"])
    }

    @Test func report_free_makesNoCoachCall() async throws {
        let service = FakeReportsService()
        let store = try reportStore(service, premium: false)
        await store.load()
        #expect(store.coach == .hidden)
        #expect(service.coachCalls.isEmpty)
    }

    @Test func report_signedOut_hasNoParses() async throws {
        let service = FakeReportsService()
        let store = try reportStore(service, signedIn: false)
        await store.load()
        #expect(store.parses == .hidden)
        #expect(service.meCalls == 0)
        #expect(service.coachCalls.isEmpty)
    }

    @Test func report_noCoachLine_drawsNoCard() async throws {
        let service = FakeReportsService()
        service.nightAnswer = { () throws(APIError) in nil }
        let store = try reportStore(service)
        await store.load()
        #expect(store.coach == .hidden)
        #expect(service.coachCalls == ["night:us/pvp/obnoxious-yell"])
    }

    @Test func report_coachError_failsInTheCardOnly() async throws {
        let service = FakeReportsService()
        service.coachAnswer = { () throws(APIError) in throw .unexpectedStatus(500) }
        let store = try reportStore(service)
        await store.load()
        #expect(store.coach == .failed)
        #expect(store.report != nil)
    }

    @Test func report_parsesError_failsInTheCardOnly() async throws {
        let service = FakeReportsService()
        service.meAnswer = { _ throws(APIError) in throw .unexpectedStatus(500) }
        let store = try reportStore(service)
        await store.load()
        #expect(store.parses == .failed)
        #expect(store.report != nil)
        service.meAnswer = { _ throws(APIError) in try RaidFixture.api(ReportMe.self, "reports.me") }
        await store.retryParses()
        #expect(store.parses == .loaded(try ReportsTestData.me()))
    }

    @Test func report_404_isNotYours() async throws {
        let service = FakeReportsService()
        let refusal = try RaidFixture.error("reports.get.error-not-found", status: 404)
        service.reportAnswer = { _ throws(APIError) in throw refusal }
        let store = try reportStore(service)
        await store.load()
        #expect(store.phase == .notYours)
        #expect(service.meCalls == 0)
    }

    @Test func report_offline_showsTheCache() async throws {
        let cache = MemoryRaidCache()
        await (try reportStore(FakeReportsService(), cache: cache)).load()
        let offline = FakeReportsService()
        offline.reportAnswer = { _ throws(APIError) in throw .transport("offline") }
        offline.meAnswer = { _ throws(APIError) in throw .transport("offline") }
        let store = try reportStore(offline, cache: cache)
        await store.load()
        guard case let .loaded(content) = store.phase else { Issue.record("not loaded"); return }
        #expect(content.offlineSince == ReportsTestData.now)
        #expect(store.parses == .loaded(try ReportsTestData.me()))
    }

    // MARK: Meters

    @Test func meters_loadEachSegmentOnce() async throws {
        let service = FakeReportsService()
        let target = try #require(MetersTarget.of(try ReportsTestData.fightRow(9), reportID: ReportsTestData.reportID))
        let store = MetersStore(target: target, service: service)
        await store.load()
        #expect(store.phase == .loaded(try ReportsTestData.meters()))
        await store.select(.healing)
        await store.select(.damage)
        #expect(service.meterCalls.map(\.metric) == [.damage, .healing])
        #expect(service.meterCalls.allSatisfy { $0.n == 9 && $0.scope == nil })
    }

    @Test func meters_trashAsksTheFold() async throws {
        let service = FakeReportsService()
        let target = try #require(MetersTarget.of(try ReportsTestData.trashRow(), reportID: ReportsTestData.reportID))
        await MetersStore(target: target, service: service).load()
        #expect(service.meterCalls.first?.n == 1)
        #expect(service.meterCalls.first?.scope == "trash")
    }

    @Test func meters_flaggedAndFailed() async throws {
        let target = try #require(MetersTarget.of(try ReportsTestData.fightRow(9), reportID: ReportsTestData.reportID))
        let flagged = FakeReportsService()
        flagged.metersAnswer = { _ throws(APIError) in
            throw .server(status: 404, body: ErrorBody(code: "flagged", message: "flagged"), requestID: nil)
        }
        let flaggedStore = MetersStore(target: target, service: flagged)
        await flaggedStore.load()
        #expect(flaggedStore.phase == .flagged)
        let failing = FakeReportsService()
        failing.metersAnswer = { _ throws(APIError) in throw .unexpectedStatus(500) }
        let failedStore = MetersStore(target: target, service: failing)
        await failedStore.load()
        #expect(failedStore.phase == .failed)
    }
}
