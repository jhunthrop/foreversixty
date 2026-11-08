@testable import ForeverSixty
import Auth
import DesignSystem
import SwiftUI
import Testing

/// The account chores against their mocks (design/parity.md "Native"): ios-devices,
/// ios-offline-data, ios-offline-data-empty, ios-offline-strip and ios-update-required.
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

    /// The mock's phone: Fury Warrior and Beast Mastery Hunter guides, a list, talents, the search
    /// index and the icons, 27 MB in all.
    private var saved: OfflineReport {
        let build = "1.60.1.70707"
        let savedAt = SystemStory.now.addingTimeInterval(-600)
        return OfflineReport.build(files: [
            SavedStaticFile(relativePath: "\(build)/guides/warrior-fury.json", bytes: 2_100_000, savedAt: savedAt),
            SavedStaticFile(relativePath: "\(build)/guides/hunter-beast-mastery.json", bytes: 2_100_000, savedAt: savedAt),
            SavedStaticFile(relativePath: "\(build)/bis/warrior-fury.json", bytes: 1_100_000, savedAt: savedAt),
            SavedStaticFile(relativePath: "\(build)/talents/warrior.json", bytes: 600_000, savedAt: savedAt),
            SavedStaticFile(relativePath: "\(build)/search.json", bytes: 2_300_000, savedAt: savedAt),
        ], iconBytes: 18_400_000, iconsUpdated: savedAt)
    }

    @Test("ios-offline-data")
    func offlineData() throws {
        try ParityCheck.assert(offline(saved), screen: "offline-data", size: Self.device)
    }

    @Test("ios-offline-data-empty")
    func offlineDataEmpty() throws {
        try ParityCheck.assert(offline(OfflineReport(rows: [])), screen: "offline-data-empty", size: Self.device)
    }

    @Test("ios-update-required")
    func updateRequired() throws {
        try ParityCheck.assert(UpdateRequiredView(storeURL: UpdateGate.fallbackStoreURL), screen: "update-required", size: Self.device)
    }

    @Test("ios-offline-strip")
    func offlineStrip() throws {
        try ParityCheck.assert(offline(saved, online: false), screen: "offline-strip", size: Self.device)
    }
}
