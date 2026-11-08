@testable import ForeverSixty
import Auth
import DesignSystem
import SwiftUI
import Testing

/// The account chores against their mocks (design/parity.md "Native"): ios-devices,
/// ios-offline-data, ios-offline-data-empty and ios-offline-strip.
@MainActor
@Suite("Account parity")
struct AccountParityTests {
    static let device = CGSize(width: 393, height: 852)

    @Test("ios-devices")
    func devices() throws {
        let view = DevicesView(state: .loaded(try AccountSnapshotTests.deviceRows()), failedRevoke: nil, onRevoke: { _ in }, onRetry: {})
        try ParityCheck.assert(view, screen: "devices", size: Self.device)
    }

    private var chicago: Calendar {
        var calendar = Calendar(identifier: .gregorian)
        calendar.timeZone = SystemStory.zone
        return calendar
    }

    private func offline(_ report: OfflineReport, online: Bool = true) -> some View {
        OfflineDataView(
            report: report, update: .idle, online: online, checkedAt: SystemStory.now, words: OfflineWords(calendar: chicago),
            now: SystemStory.now, keepClassesReady: .constant(true), wifiOnly: .constant(true), updateNow: {}, remove: {}
        )
    }

    private var saved: OfflineReport {
        SystemStory.offlineReport(savedAt: SystemStory.now.addingTimeInterval(-600))
    }

    @Test("ios-offline-data")
    func offlineData() throws {
        try ParityCheck.assert(offline(saved), screen: "offline-data", size: Self.device)
    }

    @Test("ios-offline-data-empty")
    func offlineDataEmpty() throws {
        try ParityCheck.assert(offline(OfflineReport(rows: [])), screen: "offline-data-empty", size: Self.device)
    }

    @Test("ios-offline-strip")
    func offlineStrip() throws {
        try ParityCheck.assert(offline(saved, online: false), screen: "offline-strip", size: Self.device)
    }
}
