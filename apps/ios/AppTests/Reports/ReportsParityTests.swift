@testable import ForeverSixty
import DesignSystem
import Models
import SwiftUI
import Testing

/// Your reports and a report that cannot show, against their mocks (design/parity.md "Native"):
/// ios-reports, ios-reports-empty, ios-reports-failed and ios-report-error.
@MainActor
@Suite("Reports parity")
struct ReportsParityTests {
    static let device = CGSize(width: 393, height: 852)

    private func list(_ phase: ReportsPhase) -> some View {
        ScrollView {
            ReportsListView(phase: phase, now: ReportsTestData.now, clock: ReportsTestData.clock)
        }
        .background(FSColor.bg)
    }

    /// The mock's one finished report (the Onyxia night).
    private var oneReport: ReportsContent {
        get throws {
            let row = try #require(try ReportsTestData.list().rows.first { ReportRowState(status: $0.status) == .complete })
            return ReportsContent(rows: [row], total: 1, page: 1, offlineSince: nil)
        }
    }

    @Test("ios-reports")
    func reports() throws {
        try ParityCheck.assert(list(.loaded(try oneReport)), screen: "reports", size: Self.device)
    }

    @Test("ios-reports-empty")
    func empty() throws {
        try ParityCheck.assert(
            list(.loaded(ReportsContent(rows: [], total: 0, page: 1, offlineSince: nil))), screen: "reports-empty", size: Self.device
        )
    }

    @Test("ios-reports-failed")
    func failed() throws {
        try ParityCheck.assert(list(.failed), screen: "reports-failed", size: Self.device)
    }

    @Test("ios-report-error")
    func reportError() throws {
        let view = ScrollView {
            ReportView(phase: .failed, parses: .hidden, coach: .hidden, now: ReportsTestData.now, clock: ReportsTestData.clock)
        }
        .background(FSColor.bg)
        try ParityCheck.assert(view, screen: "report-error", size: Self.device)
    }
}
