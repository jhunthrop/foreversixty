@testable import ForeverSixty
import Foundation
import Models
import Networking
import Testing

/// The one F5 test list (spec §5.5), mirrored in Android's SimsLogicTest.
@Suite("Sims logic")
struct SimsLogicTests {
    static let now = RaidTime.date("2026-12-15T05:22:10Z")

    private func rows() throws -> [SimRow] { try RaidFixture.load(SimList.self, "sims.list").rows }

    private func row(_ id: String) throws -> SimRow {
        try #require(try rows().first { $0.simId == id })
    }

    @Test func states_followTheSite() throws {
        #expect(ServerRuns.state(try row("km3xibdrsiyr"), now: Self.now) == .doneUnseen)
        #expect(ServerRuns.state(try row("q8wd2c9vh4tz"), now: Self.now) == .running)
        #expect(ServerRuns.state(try row("b7nf5x1pa3le"), now: Self.now) == .failed)
    }

    @Test func running_turnsOverduePastOneAndAHalfEstimates() throws {
        let running = try row("q8wd2c9vh4tz") // created 05:20:40, estimate 90 s
        #expect(ServerRuns.state(running, now: RaidTime.date("2026-12-15T05:22:55Z")) == .running)
        #expect(ServerRuns.state(running, now: RaidTime.date("2026-12-15T05:22:56Z")) == .runningOverdue)
    }

    @Test func running_withoutEstimate_hasNoSuffix() {
        let row = SimRow(simId: "x", state: "running", createdAt: Self.now)
        #expect(ServerRuns.state(row, now: Self.now) == .runningNoEstimate)
        #expect(SimsCopy.stateSuffix(row, state: .runningNoEstimate, now: Self.now) == "")
    }

    @Test func estimate_roundsToOneUnit() {
        #expect(RunEstimate(seconds: 50) == .aboutAMinute)
        #expect(RunEstimate(seconds: 90) == .minutes(2))
        #expect(RunEstimate(seconds: 240) == .minutes(4))
        #expect(RunEstimate(seconds: 89 * 60) == .minutes(89))
        #expect(RunEstimate(seconds: 90 * 60) == .hours(2))
    }

    @Test func visible_isTheNewestFive() {
        let many = (0..<7).map { SimRow(simId: "r\($0)", createdAt: Self.now.addingTimeInterval(Double($0))) }
        let visible = ServerRuns.visible(many)
        #expect(visible.rows.map(\.simId) == ["r6", "r5", "r4", "r3", "r2"])
        #expect(visible.hasMore)
        #expect(!ServerRuns.visible(Array(many.prefix(5))).hasMore)
    }

    @Test func section_showsForPremiumOrOldServerRows() throws {
        #expect(ServerRuns.showsSection(entitled: true, serverRows: []))
        #expect(ServerRuns.showsSection(entitled: false, serverRows: try rows()))
        #expect(!ServerRuns.showsSection(entitled: false, serverRows: []))
    }

    @Test func poll_asksOnlyAboutRunsInFlight() throws {
        #expect(ServerRuns.inFlightIDs(try rows(), now: Self.now) == ["q8wd2c9vh4tz"])
    }

    @Test func bar_onlyForCombinations() throws {
        let plain = try RaidFixture.load(SimProgress.self, "sims.progress")
        #expect(ServerRuns.fraction(plain) == nil)
        let bulk = SimProgress(state: "running", iterationsDone: 0, stage: 1, combosDone: 40, combosTotal: 64)
        #expect(ServerRuns.fraction(bulk) == 0.625)
        #expect(ServerRuns.hasEnded(SimProgress(state: "done", iterationsDone: 1, stage: 0, combosDone: 0, combosTotal: 0)))
        #expect(!ServerRuns.hasEnded(plain))
    }

    @Test func rerun_402IsALapsedPlan() throws {
        #expect(RerunStatus.refused(try RaidFixture.error("sims.rerun.error-premium-required", status: 402)) == .lapsed)
        #expect(RerunStatus.refused(try RaidFixture.error("sims.rerun.error-not-found", status: 404)) == .failed)
        #expect(RerunStatus.refused(.transport("offline")) == .failed)
        #expect(RerunStatus.pending.isLocked && RerunStatus.started.isLocked && !RerunStatus.failed.isLocked)
    }

    @Test func failedLine_namesTheReasonWhenThereIsOne() throws {
        #expect(SimsCopy.failedLine(try row("b7nf5x1pa3le")) == "Failed: the engine stopped")
        #expect(SimsCopy.failedLine(SimRow(simId: "x", state: "error")) == "Failed")
    }
}
