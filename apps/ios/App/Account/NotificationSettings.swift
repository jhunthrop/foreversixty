import DesignSystem
import Models
import Networking
import Observation
import SwiftUI

/// `GET` and `PATCH /v1/me/notifications`, behind a protocol for tests.
protocol NotificationSettingsService: Sendable {
    func load() async throws(APIError) -> NotificationSettings
    func save(_ patch: NotificationsPatch) async throws(APIError) -> NotificationSettings
}

struct LiveNotificationSettingsService: NotificationSettingsService {
    let client: APIClient

    func load() async throws(APIError) -> NotificationSettings {
        try await client.send(.getNotificationSettings)
    }

    func save(_ patch: NotificationsPatch) async throws(APIError) -> NotificationSettings {
        try await client.send(.updateNotificationSettings(patch))
    }

}

/// The notification groups in the order the settings and Activity's filters list them (flows
/// spec F4), with their words. The API names each event's group (`events.GroupOf`).
enum NotificationGroup: String, CaseIterable, Identifiable {
    case raids, gear, logs, ladder, tiers, officer

    var id: String { rawValue }

    var title: String {
        switch self {
        case .raids: Copy.pushGroupRaids
        case .gear: Copy.pushGroupGear
        case .logs: Copy.pushGroupLogs
        case .ladder: Copy.pushGroupLadder
        case .tiers: Copy.pushGroupTiers
        case .officer: Copy.pushGroupOfficer
        }
    }

    var line: String {
        switch self {
        case .raids: Copy.pushGroupRaidsLine
        case .gear: Copy.pushGroupGearLine
        case .logs: Copy.pushGroupLogsLine
        case .ladder: Copy.pushGroupLadderLine
        case .tiers: Copy.pushGroupTiersLine
        case .officer: Copy.pushGroupOfficerLine
        }
    }
}

/// The settings as the screen shows them: one switch per group the player gets events of (Raids
/// and Guild ladder for a guild member, Officer for an officer, as the API's event list says).
struct NotificationGroupsSummary: Equatable {
    struct Row: Equatable, Identifiable {
        let group: NotificationGroup
        let on: Bool

        var id: NotificationGroup { group }
    }

    let rows: [Row]
    /// The raid-day reminder, `15:00` in the player's zone; the API default until it sends one.
    let reminder: String

    static let defaultReminder = "15:00"

    init(_ settings: NotificationSettings) {
        let byGroup = Dictionary(grouping: settings.events) { $0.group.flatMap(NotificationGroup.init(rawValue:)) }
        rows = NotificationGroup.allCases.compactMap { group in
            guard let events = byGroup[group], !events.isEmpty else { return nil }
            return Row(group: group, on: events.allSatisfy(\.channels.push))
        }
        reminder = settings.raidReminder ?? Self.defaultReminder
    }

    /// The PATCH that turns a group's push on or off: every kind in it.
    static func patch(_ group: NotificationGroup, on: Bool, in settings: NotificationSettings) -> NotificationsPatch {
        let kinds = settings.events.filter { $0.group == group.rawValue }.map(\.kind)
        return NotificationsPatch(events: kinds.map { NotificationChoice(kind: $0, channel: "push", enabled: on) })
    }
}

@MainActor
@Observable
final class NotificationSettingsModel {
    enum Phase: Equatable {
        case loading
        case loaded(NotificationSettings)
        case failed
    }

    private(set) var phase: Phase = .loading
    private(set) var saveFailed = false
    private let service: any NotificationSettingsService

    init(service: any NotificationSettingsService) {
        self.service = service
    }

    func load() async {
        do {
            phase = .loaded(try await service.load())
        } catch {
            if case .loaded = phase { return }
            phase = .failed
        }
    }

    /// The raid-day reminder hour (`15:00`), drawn at once; a refused save puts it back.
    func setReminder(_ hhmm: String) async {
        guard case let .loaded(settings) = phase else { return }
        saveFailed = false
        phase = .loaded(NotificationSettings(
            timeZone: settings.timeZone, quietHours: settings.quietHours, officerDms: settings.officerDms, raidReminder: hhmm,
            events: settings.events
        ))
        do {
            phase = .loaded(try await service.save(NotificationsPatch(raidReminder: hhmm)))
        } catch {
            phase = .loaded(settings)
            saveFailed = true
        }
    }

    /// Officers may DM the player on Discord (`officer_dms`, off by default).
    var officerDMs: Bool {
        if case let .loaded(settings) = phase { return settings.officerDms ?? false }
        return false
    }

    /// Let officers DM me on Discord, drawn at once; a refused save puts it back with the line.
    func setOfficerDMs(_ on: Bool) async {
        guard case let .loaded(settings) = phase else { return }
        saveFailed = false
        phase = .loaded(settings.with(officerDMs: on))
        do {
            phase = .loaded(try await service.save(NotificationsPatch(officerDms: on)))
        } catch {
            phase = .loaded(settings)
            saveFailed = true
        }
    }

    /// Flips a group at once; a refused save puts it back with the line.
    func set(_ group: NotificationGroup, on: Bool) async {
        guard case let .loaded(settings) = phase else { return }
        saveFailed = false
        let patch = NotificationGroupsSummary.patch(group, on: on, in: settings)
        phase = .loaded(Self.applying(patch, to: settings))
        do {
            phase = .loaded(try await service.save(patch))
        } catch {
            phase = .loaded(settings)
            saveFailed = true
        }
    }

    /// The settings with a PATCH's push choices applied, drawn before the server answers.
    nonisolated static func applying(_ patch: NotificationsPatch, to settings: NotificationSettings) -> NotificationSettings {
        let choices = Dictionary((patch.events ?? []).map { ($0.kind, $0.enabled) }) { first, _ in first }
        let events = settings.events.map { event in
            guard let on = choices[event.kind] else { return event }
            return event.withPush(on)
        }
        return NotificationSettings(
            timeZone: settings.timeZone, quietHours: settings.quietHours, officerDms: settings.officerDms,
            raidReminder: settings.raidReminder, events: events
        )
    }
}

private extension NotificationSettings {
    func with(officerDMs on: Bool) -> NotificationSettings {
        NotificationSettings(timeZone: timeZone, quietHours: quietHours, officerDms: on, raidReminder: raidReminder, events: events)
    }
}

private extension NotificationEvent {
    func withPush(_ on: Bool) -> NotificationEvent {
        NotificationEvent(
            kind: kind, group: group, who: who, officer: officer, quietExempt: quietExempt, timeSensitive: timeSensitive,
            channels: NotificationChannels(push: on, discord: channels.discord)
        )
    }
}

/// Notifications (mock ios-notification-settings), pushed from the menu: Send me, a switch per
/// group, and When, the raid-day reminder.
struct NotificationSettingsScreen: View {
    @State private var model: NotificationSettingsModel

    init(service: any NotificationSettingsService) {
        _model = State(initialValue: NotificationSettingsModel(service: service))
    }

    var body: some View {
        NotificationSettingsView(
            phase: model.phase, saveFailed: model.saveFailed, officerDMs: model.officerDMs,
            set: { group, on in Task { await model.set(group, on: on) } },
            setOfficerDMs: { on in Task { await model.setOfficerDMs(on) } },
            setReminder: { hhmm in Task { await model.setReminder(hhmm) } },
            retry: { Task { await model.load() } }
        )
        .task { await model.load() }
    }
}

struct NotificationSettingsView: View {
    let phase: NotificationSettingsModel.Phase
    var saveFailed = false
    var officerDMs = false
    var set: (NotificationGroup, Bool) -> Void = { _, _ in }
    var setOfficerDMs: (Bool) -> Void = { _ in }
    var setReminder: (String) -> Void = { _ in }
    var retry: () -> Void = {}

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                PageHeader(title: Copy.menuNotifications)
                switch phase {
                case .loading:
                    ProgressView().frame(maxWidth: .infinity).padding(.top, FSSpace.space6)
                case .failed:
                    FSButton(Copy.commonTryAgain, action: retry).padding(.top, FSSpace.space6)
                case let .loaded(settings):
                    loaded(NotificationGroupsSummary(settings))
                }
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .padding(.bottom, FSSpace.space6)
            .parityContent()
        }
        .screenGround()
        .navigationBarTitleDisplayMode(.inline)
    }

    @ViewBuilder
    private func loaded(_ summary: NotificationGroupsSummary) -> some View {
        FSSection(Copy.notificationsSendMe) {
            VStack(spacing: 0) {
                ForEach(summary.rows) { row in
                    SwitchRow(title: row.group.title, sub: row.group.line, on: row.on) { set(row.group, $0) }
                }
                if saveFailed {
                    SmallText(Copy.commonFailed, color: FSColor.wipe).padding(.top, FSSpace.space2)
                }
            }
        }
        .padding(.top, FSSpace.space7)
        FSSection(Copy.notificationsWhen) {
            FSRow(
                title: Copy.notificationsReminder,
                sub: Copy.notificationsReminderLine(time: ClockText.time(hhmm: summary.reminder)), inset: 0,
                trailPlacement: .below
            ) {
                EmptyView()
            } trail: {
                ReminderPicker(hhmm: summary.reminder, save: setReminder)
            }
        }
        .padding(.top, FSSpace.space7)
        // A guild member's: officers may DM them about a raid they have not answered (F3).
        if summary.rows.contains(where: { $0.group == .raids }) {
            FSSection(Copy.notificationsDiscord) {
                SwitchRow(title: Copy.notificationsOfficerDms, sub: Copy.notificationsOfficerDmsLine, on: officerDMs, set: setOfficerDMs)
            }
            .padding(.top, FSSpace.space7)
        }
    }
}

/// A setting's row: its title and line, the switch on the title row.
struct SwitchRow: View {
    let title: String
    let sub: String
    let on: Bool
    let set: (Bool) -> Void

    var body: some View {
        FSRow(title: title, sub: sub, inset: 0) {
            EmptyView()
        } trail: {
            Toggle(title, isOn: Binding(get: { on }, set: set))
                .labelsHidden()
                .tint(FSColor.gold)
                .frame(minHeight: FSTarget.min)
        }
    }
}

/// The reminder row's Change: a menu of the afternoon and evening hours, the current
/// one checked.
struct ReminderPicker: View {
    let hhmm: String
    let save: (String) -> Void

    /// What the link adds above and below itself to reach the 44 pt target.
    static let targetPad: CGFloat = 12

    /// The hours offered: noon to 8 pm, when a raid-day nudge is useful.
    static let hours = Array(12 ... 20)

    var body: some View {
        Menu {
            ForEach(Self.hours, id: \.self) { hour in
                let value = String(format: "%02d:00", hour)
                Button { save(value) } label: {
                    if value == hhmm {
                        Label(ClockText.time(hhmm: value), systemImage: "checkmark")
                    } else {
                        Text(ClockText.time(hhmm: value))
                    }
                }
            }
        } label: {
            // A link's look, padded to the 44 pt target without taking the room (web ActionLink).
            Text(Copy.notificationsChange)
                .type(FSType.item)
                .foregroundStyle(FSColor.gold)
                .padding(.vertical, ReminderPicker.targetPad)
                .contentShape(.rect)
                .padding(.vertical, -ReminderPicker.targetPad)
        }
    }
}
