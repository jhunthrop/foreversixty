@testable import ForeverSixty
import DesignSystem
import Models
import Networking
import SnapshotTesting
import SwiftUI
import Testing
import UIKit

/// Every F10 screen and state on iPhone 17 (402 x 874), at the default size and at AX5, from the
/// contract fixtures. Recorded on the shared iOS 26.5 simulator like SimsSnapshotTests; the PNGs
/// double as the review captures in design/reviews/2026-10-31-mobile-wave3-reports.
@MainActor
@Suite("Reports snapshots")
struct ReportsSnapshotTests {
    private let specs = SpecCatalog.bundled()
    private let now = ReportsTestData.now
    private let clock = ReportsTestData.clock

    private func snap(_ view: some View, _ name: String, height: CGFloat = 874, large: Bool = false) {
        let traits = UITraitCollection { traits in
            traits.userInterfaceStyle = .dark
            traits.preferredContentSizeCategory = large ? .accessibilityExtraExtraExtraLarge : .large
        }
        let config = ViewImageConfig(safeArea: SimsSnapshotTests.safeArea, size: CGSize(width: SimsSnapshotTests.phone.width, height: height), traits: traits)
        let sized = view
            .environment(\.itemIconSource, RepoItemIconSource())
            .environment(\.dynamicTypeSize, large ? .accessibility5 : .large)
            .preferredColorScheme(.dark)
        assertSnapshot(
            of: sized,
            as: .image(precision: 0.99, perceptualPrecision: 0.98, layout: .device(config: config), traits: traits),
            testName: name
        )
    }

    /// The screen as pushed, so the back control shows (root titled `back`).
    private func pushed(back: String, trailing: String, @ViewBuilder _ content: @escaping () -> some View) -> some View {
        NavigationStack(path: .constant([1])) {
            Color.clear.navigationTitle(back).navigationDestination(for: Int.self) { _ in
                ScrollView { content() }
                    .background(FSColor.bg)
                    .reportsBar(title: back) { Button {} label: { Image(systemName: trailing) } }
            }
        }
    }

    private func list(_ phase: ReportsPhase, account: ReportsAccount? = nil) throws -> some View {
        let account = try account ?? ReportsTestData.account(premium: false)
        return pushed(back: Copy.tabHome, trailing: "arrow.clockwise") {
            ReportsListView(phase: phase, account: account, now: now, clock: clock)
        }
    }

    private func report(_ phase: ReportPhase, parses: ParsesState = .hidden, coach: CoachState = .hidden) -> some View {
        pushed(back: Copy.reportsBack, trailing: "arrow.up.right.square") {
            ReportView(phase: phase, parses: parses, coach: coach, now: now, clock: clock)
        }
    }

    private func loaded(_ report: Report, offline: Date? = nil) -> ReportPhase {
        .loaded(ReportContent(report: report, offlineSince: offline))
    }

    private func meters(_ row: FightRow, metric: MeterMetric = .damage, phase: MetersPhase) throws -> some View {
        let target = try #require(MetersTarget.of(row, reportID: ReportsTestData.reportID))
        return MetersView(target: target, metric: metric, phase: phase, showAll: false, specs: specs, build: ReportsTestData.build)
    }

    private var listContent: ReportsContent {
        get throws {
            let list = try ReportsTestData.list()
            return ReportsContent(rows: list.rows, total: list.total, page: list.page, offlineSince: nil)
        }
    }

    // MARK: F10.1 Your reports

    @Test func reports_list() throws {
        snap(try list(.loaded(try listContent)), "reports-list-ios")
    }

    @Test func reports_list_large() throws {
        snap(try list(.loaded(try listContent)), "reports-list-large-ios", height: 2400, large: true)
    }

    @Test func reports_empty() throws {
        snap(try list(.loaded(ReportsContent(rows: [], total: 0, page: 1, offlineSince: nil))), "reports-empty-ios")
    }

    @Test func reports_loading() throws {
        snap(try list(.loading), "reports-loading-ios")
    }

    @Test func reports_error() throws {
        snap(try list(.failed), "reports-error-ios")
    }

    @Test func reports_signedOut() throws {
        snap(try list(.signedOut, account: .signedOut), "reports-signed-out-ios")
    }

    // MARK: F10.2 A report

    @Test func report_kill() throws {
        snap(report(loaded(try ReportsTestData.report()), parses: .loaded(try ReportsTestData.me())), "report-kill-ios")
    }

    @Test func report_kill_large() throws {
        snap(report(loaded(try ReportsTestData.report()), parses: .loaded(try ReportsTestData.me())), "report-kill-large-ios", height: 3000, large: true)
    }

    @Test func report_onlyWipes() throws {
        snap(report(loaded(try ReportsTestData.onlyWipes()), parses: .loaded(try ReportsTestData.meWithoutKill())), "report-wipes-ios")
    }

    @Test func report_twoZones() throws {
        snap(report(loaded(try ReportsTestData.twoZones()), parses: .loaded(try ReportsTestData.me())), "report-two-zones-ios", height: 1000)
    }

    @Test func report_offline() throws {
        let since = RaidTime.date("2026-12-15T21:41:00Z")
        snap(report(loaded(try ReportsTestData.report(), offline: since), parses: .loaded(try ReportsTestData.me())), "report-offline-ios")
    }

    @Test func report_parsesFailed() throws {
        snap(report(loaded(try ReportsTestData.report()), parses: .failed), "report-parses-failed-ios")
    }

    // MARK: F10.4 Coach

    @Test func report_premium() throws {
        let coach = CoachState.ready(try ReportsTestData.review())
        snap(report(loaded(try ReportsTestData.report()), parses: .loaded(try ReportsTestData.me()), coach: coach), "report-premium-ios", height: 2100)
    }

    @Test func report_free() throws {
        snap(report(loaded(try ReportsTestData.report()), parses: .loaded(try ReportsTestData.me())), "report-free-ios", height: 2100)
    }

    // MARK: F10.5 States

    @Test func report_live() throws {
        snap(report(loaded(try ReportsTestData.report(status: ReportStatus.live)), parses: .loaded(try ReportsTestData.me())), "report-live-ios")
    }

    @Test func report_processing() throws {
        snap(report(loaded(try ReportsTestData.report(status: ReportStatus.processing))), "report-processing-ios")
    }

    @Test func report_failed() throws {
        snap(report(loaded(try ReportsTestData.report(status: ReportStatus.failed))), "report-failed-ios")
    }

    @Test func report_notYours() {
        snap(report(.notYours), "report-not-yours-ios")
    }

    @Test func report_loading() {
        snap(report(.loading), "report-loading-ios")
    }

    // MARK: F10.3 A fight

    @Test func meters_damage() throws {
        snap(try meters(try ReportsTestData.fightRow(9), phase: .loaded(try ReportsTestData.meters())), "meters-damage-ios")
    }

    @Test func meters_damage_large() throws {
        snap(try meters(try ReportsTestData.fightRow(9), phase: .loaded(try ReportsTestData.meters())), "meters-damage-large-ios", height: 2000, large: true)
    }

    @Test func meters_healing() throws {
        snap(try meters(try ReportsTestData.fightRow(9), metric: .healing, phase: .loaded(try ReportsTestData.healing())), "meters-healing-ios")
    }

    @Test func meters_taken() throws {
        snap(try meters(try ReportsTestData.fightRow(9), metric: .taken, phase: .loaded(try ReportsTestData.taken())), "meters-taken-ios")
    }

    @Test func meters_trash() throws {
        snap(try meters(try ReportsTestData.trashRow(), phase: .loaded(try ReportsTestData.trashMeters())), "meters-trash-ios")
    }

    @Test func meters_loading() throws {
        snap(try meters(try ReportsTestData.fightRow(9), phase: .loading), "meters-loading-ios")
    }

    @Test func meters_flagged() throws {
        snap(try meters(try ReportsTestData.fightRow(9), phase: .flagged), "meters-flagged-ios")
    }

    @Test func meters_failed() throws {
        snap(try meters(try ReportsTestData.fightRow(9), phase: .failed), "meters-failed-ios")
    }

    // MARK: Home

    @Test func home_logsRow() throws {
        var presenter = HomeFixtures.presenter()
        presenter.logsLine = ReportsCopy.homeLine(try ReportsTestData.list())
        let phase = HomePhase.loaded(try HomeFixtures.home(), offlineSince: nil)
        let view = HomeCaptureFrame {
            HomeAnswerView(phase: phase, presenter: presenter, build: ReportsTestData.build)
        }
        snap(view, "home-logs-ios", height: 2100)
    }
}
