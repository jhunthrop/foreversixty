@testable import ForeverSixty
import Foundation
import Models
import Networking
import SwiftUI
import Testing

/// The mock's member: Raids, Gear and sims and Logs on; Guild ladder and Tiers off; no officer rows.
enum NotificationTestData {
    static func event(_ kind: String, _ group: String, push: Bool, officer: Bool = false) -> NotificationEvent {
        NotificationEvent(
            kind: kind, group: group, who: "player", officer: officer, quietExempt: false, timeSensitive: false,
            channels: NotificationChannels(push: push, discord: false)
        )
    }

    static let member = NotificationSettings(
        timeZone: "America/Chicago", quietHours: QuietHours(enabled: false, start: "22:00", end: "08:00"), raidReminder: "15:00",
        events: [
            event("raid.posted", "raids", push: true), event("raid.starting", "raids", push: true),
            event("loot.awarded", "gear", push: true), event("sim.finished", "gear", push: true),
            event("report.ready", "logs", push: true),
            event("guildsim.passed", "ladder", push: false),
            event("tier.moved", "tiers", push: false),
            event("billing.failed", "account", push: true),
        ]
    )
}

final class FakeNotificationSettingsService: NotificationSettingsService, @unchecked Sendable {
    var settings = NotificationTestData.member
    var patches: [NotificationsPatch] = []
    var fails = false

    func load() async throws(APIError) -> NotificationSettings { settings }

    func save(_ patch: NotificationsPatch) async throws(APIError) -> NotificationSettings {
        if fails { throw .transport("offline") }
        patches.append(patch)
        settings = NotificationSettingsModel.applying(patch, to: settings)
        return settings
    }
}

@Suite("Notification settings")
struct NotificationGroupsTests {
    @Test("one switch per group the player gets, in the settings' order; account and officer stay off the list for a member")
    func rows() {
        let summary = NotificationGroupsSummary(NotificationTestData.member)
        #expect(summary.rows.map(\.group) == [.raids, .gear, .logs, .ladder, .tiers])
        #expect(summary.rows.map(\.on) == [true, true, true, false, false])
        #expect(summary.reminder == "15:00")
    }

    @Test("a group's switch writes every kind in it, push only")
    func patch() {
        let patch = NotificationGroupsSummary.patch(.gear, on: false, in: NotificationTestData.member)
        #expect(patch.events?.map(\.kind) == ["loot.awarded", "sim.finished"])
        #expect(patch.events?.allSatisfy { $0.channel == "push" && !$0.enabled } == true)
    }

    @Test("the reminder reads as the phone's clock, lower case")
    func reminder() {
        #expect(ClockText.time(hhmm: "15:00", locale: Locale(identifier: "en_US")) == "3:00 pm")
        #expect(ClockText.time(hhmm: "later", locale: Locale(identifier: "en_US")) == "later")
    }
}

@MainActor
@Suite("Notification settings model")
struct NotificationSettingsModelTests {
    @Test("a switch draws at once and saves; a refused save puts it back with the line")
    func toggle() async {
        let service = FakeNotificationSettingsService()
        let model = NotificationSettingsModel(service: service)
        await model.load()
        await model.set(.ladder, on: true)
        #expect(service.patches.count == 1)
        guard case let .loaded(settings) = model.phase else { Issue.record("not loaded"); return }
        #expect(NotificationGroupsSummary(settings).rows.first { $0.group == .ladder }?.on == true)

        service.fails = true
        await model.set(.raids, on: false)
        #expect(model.saveFailed)
        guard case let .loaded(after) = model.phase else { Issue.record("not loaded"); return }
        #expect(NotificationGroupsSummary(after).rows.first { $0.group == .raids }?.on == true)
    }
}

@MainActor
@Suite("Notification settings parity")
struct NotificationSettingsParityTests {
    @Test("ios-notification-settings")
    func parity() async throws {
        let view = NotificationSettingsView(phase: .loaded(NotificationTestData.member))
        try await ParityCheck.assert(view, screen: "notification-settings", size: CGSize(width: 393, height: 852))
    }
}

@Suite("Site links")
struct SiteLinkTests {
    @Test("a guild's path name is its name in lower case with hyphens")
    func slug() {
        #expect(GuildSlug.of("Olympus XXVII") == "olympus-xxvii")
        #expect(GuildSlug.of(" Night Watch ") == "night-watch")
    }
}
